package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var operator string

	fmt.Println("Калькулятор")
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	fmt.Println("Результат:")

	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно!")
		} else {
			fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
		}
	default:
		fmt.Println("Ошибка: неизвестный оператор. Доступны: +, -, *, /")
	}
}