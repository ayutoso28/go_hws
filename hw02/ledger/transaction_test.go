package main

import (
	"testing"
	"time"
)

func TestTransactions(t *testing.T) {
	previous := transactions
	transactions = make([]Transaction, 0)
	t.Cleanup(func() { transactions = previous })

	if got := ListTransactions(); len(got) != 0 {
		t.Fatalf("initial transactions = %v, want empty storage", got)
	}
	if err := AddTransaction(Transaction{Amount: 0}); err == nil {
		t.Fatal("zero amount must return an error")
	}
	if len(ListTransactions()) != 0 {
		t.Fatal("invalid transaction must not change storage")
	}

	date := time.Date(2026, time.October, 5, 12, 0, 0, 0, time.UTC)
	first := Transaction{ID: 99, Amount: 450, Category: "Еда", Description: "Обед", Date: date}
	second := Transaction{Amount: 70, Category: "Транспорт", Description: "Метро", Date: date}
	for _, tx := range []Transaction{first, second} {
		if err := AddTransaction(tx); err != nil {
			t.Fatalf("AddTransaction: %v", err)
		}
	}

	got := ListTransactions()
	first.ID, second.ID = 1, 2
	if len(got) != 2 {
		t.Fatalf("transaction count = %d, want 2", len(got))
	}
	if got[0] != first || got[1] != second {
		t.Fatalf("transactions = %v, want %v", got, []Transaction{first, second})
	}

	got[0].Amount = 999
	if ListTransactions()[0].Amount != first.Amount {
		t.Fatal("changing the returned slice must not change storage")
	}
}
