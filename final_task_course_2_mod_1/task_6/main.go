package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Подсчёт положительных, отрицательных и нулей")
	fmt.Print("Сколько чисел вы хотите ввести? ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Ошибка: количество чисел должно быть больше нуля.")
		return
	}

	positiveCount := 0
	negativeCount := 0
	zeroCount := 0

	var num float64

	for i := 1; i <= n; i++ {
		fmt.Printf("Введите число %d: ", i)
		fmt.Scan(&num)

		if num > 0 {
			positiveCount++
		} else if num < 0 {
			negativeCount++
		} else {
			zeroCount++
		}
	}

	fmt.Printf("\nРезультаты анализа %d чисел:\n", n)
	fmt.Printf("- Положительных: %d\n", positiveCount)
	fmt.Printf("- Отрицательных: %d\n", negativeCount)
	fmt.Printf("- Нулей:         %d\n", zeroCount)
}