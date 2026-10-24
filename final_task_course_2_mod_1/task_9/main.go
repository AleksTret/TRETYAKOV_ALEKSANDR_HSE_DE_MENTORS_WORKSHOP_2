package main

import (
	"fmt"
)

func main() {
	balance := 1000.0
	var choice int

	fmt.Println("Мини-банкомат")
	fmt.Println("Начальный баланс: 1000.00 руб.")

	for {
		fmt.Println("\nДоступные операции:")
		fmt.Println("1 - Проверить баланс")
		fmt.Println("2 - Пополнить счёт")
		fmt.Println("3 - Снять средства")
		fmt.Println("0 - Выход")
		fmt.Print("Выберите операцию: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Текущий баланс: %.2f руб.\n", balance)

		case 2:
			var amount float64
			fmt.Print("Введите сумму для пополнения: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Ошибка: сумма пополнения должна быть больше нуля.")
			} else {
				balance += amount
				fmt.Printf("Счёт пополнен на %.2f руб. Новый баланс: %.2f руб.\n", amount, balance)
			}

		case 3:
			var amount float64
			fmt.Print("Введите сумму для снятия: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Ошибка: сумма снятия должна быть больше нуля.")
			} else if amount > balance {
				fmt.Printf("Ошибка: недостаточно средств. Доступно: %.2f руб.\n", balance)
			} else {
				balance -= amount
				fmt.Printf("Снято %.2f руб. Новый баланс: %.2f руб.\n", amount, balance)
			}

		case 0:
			fmt.Println("До свидания!")
			return

		default:
			fmt.Println("Ошибка: неверный пункт меню.")
		}
	}
}