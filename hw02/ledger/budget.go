package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Budget задаёт фиксированный лимит категории без разделения по периодам.
type Budget struct {
	Category string `json:"category"`
	Limit    int    `json:"limit"` // Лимит в рублях, как и Transaction.Amount.
}

var budgets = make(map[string]Budget)

func SetBudget(b Budget) {
	budgets[b.Category] = b
}

func LoadBudgets(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("read budgets: %w", err)
	}

	var loaded []Budget
	if err := json.Unmarshal(data, &loaded); err != nil {
		return fmt.Errorf("parse budgets: %w", err)
	}
	if loaded == nil {
		return errors.New("budgets must be a JSON array")
	}

	for _, b := range loaded {
		SetBudget(b)
	}
	return nil
}
