package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Статистика последовательности")
	fmt.Print("Сколько чисел вы хотите ввести? ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Ошибка: количество чисел должно быть больше нуля.")
		return
	}

	var num float64

	fmt.Print("Введите число 1: ")
	fmt.Scan(&num)

	sum := num
	minVal := num
	maxVal := num

	for i := 2; i <= n; i++ {
		fmt.Printf("Введите число %d: ", i)
		fmt.Scan(&num)

		sum += num
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}

	average := sum / float64(n)

	fmt.Printf("\nСтатистика по %d числам:\n", n)
	fmt.Printf("- Количество элементов: %d\n", n)
	fmt.Printf("- Сумма:                %.2f\n", sum)
	fmt.Printf("- Среднее значение:     %.2f\n", average)
	fmt.Printf("- Минимум:              %.2f\n", minVal)
	fmt.Printf("- Максимум:             %.2f\n", maxVal)
}