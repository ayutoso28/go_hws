package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	username := os.Getenv("USER")
	if runtime.GOOS == "windows" {
		username = os.Getenv("USERNAME")
	}
	if username == "" {
		fmt.Println("Имя пользователя: переменная окружения не задана")
	} else {
		fmt.Println("Имя пользователя:", username)
	}

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Аргументы CLI: не переданы")
	} else {
		fmt.Println("Аргументы CLI:")
		for i, arg := range args {
			fmt.Printf("  %d: %s\n", i+1, arg)
		}
	}

	fmt.Println("Версия Go:", runtime.Version())
}
