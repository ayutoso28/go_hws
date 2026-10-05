package main

import (
	"errors"
	"fmt"
	"time"
)

type Transaction struct {
	ID          int
	Amount      int // Сумма в рублях.
	Category    string
	Description string
	Date        time.Time
}

var transactions = make([]Transaction, 0)

func AddTransaction(tx Transaction) error {
	if tx.Amount == 0 {
		return errors.New("сумма транзакции не должна быть равна нулю")
	}

	if budget, ok := budgets[tx.Category]; ok {
		total := tx.Amount
		for _, saved := range transactions {
			if saved.Category == tx.Category {
				total += saved.Amount
			}
		}
		if total > budget.Limit {
			return fmt.Errorf("budget exceeded: категория %q, сумма %d руб., лимит %d руб.",
				tx.Category, total, budget.Limit)
		}
	}

	tx.ID = len(transactions) + 1
	transactions = append(transactions, tx)
	return nil
}

func ListTransactions() []Transaction {
	result := make([]Transaction, len(transactions))
	copy(result, transactions)
	return result
}
