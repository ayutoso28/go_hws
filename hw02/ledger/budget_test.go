package main

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestAddTransactionBudget(t *testing.T) {
	tests := []struct {
		name     string
		category string
		amount   int
		wantErr  bool
	}{
		{name: "below limit", category: "Еда", amount: 2000},
		{name: "exact limit", category: "Еда", amount: 4000},
		{name: "cumulative excess", category: "Еда", amount: 4001, wantErr: true},
		{name: "single excess", category: "Еда", amount: 6000, wantErr: true},
		{name: "no budget", category: "Образование", amount: 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetLedger(t)
			SetBudget(Budget{Category: "Еда", Limit: 5000})
			for _, tx := range []Transaction{
				{Category: "Еда", Amount: 1000},
				{Category: "Транспорт", Amount: 9000},
			} {
				if err := AddTransaction(tx); err != nil {
					t.Fatal(err)
				}
			}
			before := ListTransactions()
			err := AddTransaction(Transaction{Category: tt.category, Amount: tt.amount})
			if (err != nil) != tt.wantErr {
				t.Fatalf("AddTransaction error = %v, want error = %v", err, tt.wantErr)
			}
			if tt.wantErr {
				if !reflect.DeepEqual(ListTransactions(), before) {
					t.Fatal("rejected transaction changed storage")
				}
				if err := AddTransaction(Transaction{Category: "Еда", Amount: 1}); err != nil {
					t.Fatal(err)
				}
			}
			got := ListTransactions()
			if len(got) != 3 || got[2].ID != 3 {
				t.Fatalf("expected three transactions with consecutive IDs, got %v", got)
			}
		})
	}
}

func TestSetBudgetUpdatesLimit(t *testing.T) {
	resetLedger(t)
	SetBudget(Budget{Category: "Еда", Limit: 100})
	if err := AddTransaction(Transaction{Category: "Еда", Amount: 100}); err != nil {
		t.Fatal(err)
	}
	if err := AddTransaction(Transaction{Category: "Еда", Amount: 1}); err == nil {
		t.Fatal("expected rejection before raising limit")
	}
	SetBudget(Budget{Category: "Еда", Limit: 200})
	if err := AddTransaction(Transaction{Category: "Еда", Amount: 100}); err != nil {
		t.Fatalf("raised limit was not applied: %v", err)
	}
	SetBudget(Budget{Category: "Еда", Limit: 50})
	if err := AddTransaction(Transaction{Category: "Еда", Amount: 1}); err == nil {
		t.Fatal("lowered limit was not applied")
	}
	if len(ListTransactions()) != 2 {
		t.Fatal("updating a budget must preserve existing transactions")
	}
}

func TestLoadBudgets(t *testing.T) {
	resetLedger(t)
	SetBudget(Budget{Category: "Еда", Limit: 100})
	SetBudget(Budget{Category: "Транспорт", Limit: 2000})
	input := `[{"category":"Еда","limit":5000},{"category":"Образование","limit":3000}]`
	if err := LoadBudgets(strings.NewReader(input)); err != nil {
		t.Fatal(err)
	}
	want := map[string]Budget{
		"Еда":         {Category: "Еда", Limit: 5000},
		"Транспорт":   {Category: "Транспорт", Limit: 2000},
		"Образование": {Category: "Образование", Limit: 3000},
	}
	if !reflect.DeepEqual(budgets, want) {
		t.Fatalf("budgets = %v, want %v", budgets, want)
	}
	if err := AddTransaction(Transaction{Category: "Еда", Amount: 5001}); err == nil {
		t.Fatal("loaded budget must apply to new transactions")
	}
	if err := LoadBudgets(strings.NewReader(`[]`)); err != nil {
		t.Fatalf("empty array: %v", err)
	}
	if !reflect.DeepEqual(budgets, want) {
		t.Fatal("empty array must preserve existing budgets")
	}
}

func TestLoadBudgetsInvalidJSON(t *testing.T) {
	for _, input := range []string{
		``, `null`, `{}`, `[`,
		`[{"category":"Еда","limit":"5000"}]`,
		`[{"category":"Еда","limit":5000},{"category":"Транспорт","limit":"bad"}]`,
		`[{"category":"Еда","limit":5000}] trailing`,
		`[] []`,
	} {
		t.Run(input, func(t *testing.T) {
			resetLedger(t)
			SetBudget(Budget{Category: "Еда", Limit: 100})
			if err := LoadBudgets(strings.NewReader(input)); err == nil {
				t.Fatal("expected JSON error")
			}
			want := map[string]Budget{"Еда": {Category: "Еда", Limit: 100}}
			if !reflect.DeepEqual(budgets, want) {
				t.Fatalf("invalid input changed budgets: %v", budgets)
			}
		})
	}
}

type failingReader struct{}

func (failingReader) Read(p []byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}

func TestLoadBudgetsReadError(t *testing.T) {
	resetLedger(t)
	SetBudget(Budget{Category: "Еда", Limit: 100})
	err := LoadBudgets(io.MultiReader(
		strings.NewReader(`[{"category":"Еда","limit":5000}]`),
		failingReader{},
	))
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Fatalf("error = %v, want wrapped read error", err)
	}
	if budgets["Еда"].Limit != 100 {
		t.Fatal("read error must not change budgets")
	}
}
