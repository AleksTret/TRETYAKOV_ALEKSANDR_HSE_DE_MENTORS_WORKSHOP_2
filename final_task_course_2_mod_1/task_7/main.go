package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Сумма подходящих чисел (чётные положительные)")
	fmt.Print("Сколько чисел вы хотите ввести? ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Ошибка: количество чисел должно быть больше нуля.")
		return
	}

	sum := 0

	for i := 1; i <= n; i++ {
		var num int
		fmt.Printf("Введите число %d: ", i)
		fmt.Scan(&num)

		if num > 0 && num%2 == 0 {
			sum += num
		}
	}

	fmt.Printf("\nСумма подходящих чисел: %d\n", sum)
}