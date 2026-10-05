package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Ledger service started")

	SetBudget(Budget{Category: "Еда", Limit: 4000})
	SetBudget(Budget{Category: "Транспорт", Limit: 2000})

	file, err := os.Open("budgets.json")
	if err != nil {
		log.Fatal(err)
	}
	err = LoadBudgets(bufio.NewReader(file))
	file.Close()
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	examples := []Transaction{
		{Amount: 450, Category: "Еда", Description: "Обед", Date: now},
		{Amount: 70, Category: "Транспорт", Description: "Поездка на метро", Date: now},
		{Amount: 1200, Category: "Образование", Description: "Книга по Go", Date: now},
		{Amount: 4550, Category: "Еда", Description: "Продукты: достигаем лимита", Date: now},
		{Amount: 1, Category: "Еда", Description: "Покупка сверх лимита", Date: now},
	}
	for _, tx := range examples {
		if err := AddTransaction(tx); err != nil {
			fmt.Printf("Отказ: %s — %v\n", tx.Description, err)
		}
	}

	fmt.Println("Сохранённые транзакции:")
	for _, tx := range ListTransactions() {
		fmt.Printf("ID: %d | Сумма: %d руб. | Категория: %s | Описание: %s | Дата: %s\n",
			tx.ID, tx.Amount, tx.Category, tx.Description, tx.Date.Format("2006-01-02"))
	}
}
