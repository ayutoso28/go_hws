# Домашняя работа 3 — бюджеты

Решение развивает сервис [Ledger из домашней работы 2](../hw02/ledger).

- [budget.go](../hw02/ledger/budget.go) — структура `Budget`, хранилище, `SetBudget` и `LoadBudgets`.
- [transaction.go](../hw02/ledger/transaction.go) — проверка лимита при добавлении транзакции.
- [main.go](../hw02/ledger/main.go) и [budgets.json](../hw02/ledger/budgets.json) — загрузка бюджетов и примеры успешных и отклонённых транзакций.
- [budget_test.go](../hw02/ledger/budget_test.go) — тесты бюджетов и загрузки JSON.

Запуск и проверка из корня репозитория:

```sh
cd hw02/ledger
go run .
go test ./...
```

Ожидаемый результат: четыре сохранённые транзакции и отказ `budget exceeded` для покупки сверх лимита.
