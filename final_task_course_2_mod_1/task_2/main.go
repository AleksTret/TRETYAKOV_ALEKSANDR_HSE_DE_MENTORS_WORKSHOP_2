package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Println("Анализ числа")
	fmt.Print("Введите целое число: ")
	fmt.Scan(&num)

	fmt.Printf("\nАнализ числа %d:\n", num)

	// 1. Определение знака
	if num > 0 {
		fmt.Println("- Знак: положительное")
	} else if num < 0 {
		fmt.Println("- Знак: отрицательное")
	} else {
		fmt.Println("- Знак: ноль")
	}

	// 2. Четность
	if num == 0 {
		fmt.Println("- Чётность: ноль является чётным числом")
	} else if num%2 == 0 {
		fmt.Println("- Чётность: чётное")
	} else {
		fmt.Println("- Чётность: нечётное")
	}

	// 3. Делимость на 3
	if num%3 == 0 {
		fmt.Println("- Делимость на 3: делится без остатка")
	} else {
		fmt.Println("- Делимость на 3: не делится (остаток", num%3, ")")
	}

	// 4. Делимость на 5
	if num%5 == 0 {
		fmt.Println("- Делимость на 5: делится без остатка")
	} else {
		fmt.Println("- Делимость на 5: не делится (остаток", num%5, ")")
	}
}