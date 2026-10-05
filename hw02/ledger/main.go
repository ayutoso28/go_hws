package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Ledger service started")

	now := time.Now()
	examples := []Transaction{
		{Amount: 450, Category: "Еда", Description: "Обед", Date: now},
		{Amount: 70, Category: "Транспорт", Description: "Поездка на метро", Date: now},
		{Amount: 1200, Category: "Образование", Description: "Книга по Go", Date: now},
	}
	for _, tx := range examples {
		if err := AddTransaction(tx); err != nil {
			log.Fatal(err)
		}
	}

	for _, tx := range ListTransactions() {
		fmt.Printf("ID: %d | Сумма: %d руб. | Категория: %s | Описание: %s | Дата: %s\n",
			tx.ID, tx.Amount, tx.Category, tx.Description, tx.Date.Format("2006-01-02"))
	}
}
