package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Поиск максимума и минимума")
	fmt.Print("Сколько чисел вы хотите ввести? ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Ошибка: количество чисел должно быть больше нуля.")
		return
	}

	var num float64

	fmt.Print("Введите число 1: ")
	fmt.Scan(&num)
	maxVal := num
	minVal := num

	for i := 2; i <= n; i++ {
		fmt.Printf("Введите число %d: ", i)
		fmt.Scan(&num)

		if num > maxVal {
			maxVal = num
		}

		if num < minVal {
			minVal = num
		}
	}

	fmt.Printf("\nРезультаты анализа %d чисел:\n", n)
	fmt.Printf("- Максимальное значение: %.2f\n", maxVal)
	fmt.Printf("- Минимальное значение:  %.2f\n", minVal)
}