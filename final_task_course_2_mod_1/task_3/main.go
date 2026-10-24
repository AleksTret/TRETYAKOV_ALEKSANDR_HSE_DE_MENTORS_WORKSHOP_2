package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Сумма и среднее от 1 до N")
	fmt.Print("Введите число N (целое, больше 0): ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Ошибка: N должно быть больше нуля.")
		return
	}

	sum := 0 

	for i := 1; i <= n; i++ {
		sum += i
	}

	average := float64(sum) / float64(n)

	fmt.Printf("\nДля чисел от 1 до %d:\n", n)
	fmt.Printf("- Сумма: %d\n", sum)
	fmt.Printf("- Среднее значение: %.2f\n", average)
}