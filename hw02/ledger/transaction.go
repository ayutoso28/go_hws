package main

import (
	"errors"
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

	tx.ID = len(transactions) + 1
	transactions = append(transactions, tx)
	return nil
}

func ListTransactions() []Transaction {
	result := make([]Transaction, len(transactions))
	copy(result, transactions)
	return result
}
