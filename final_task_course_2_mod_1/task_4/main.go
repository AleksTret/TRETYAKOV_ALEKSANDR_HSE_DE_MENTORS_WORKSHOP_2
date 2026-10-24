package main

import (
	"fmt"
)

func printMultiplicationTable(num int) {
	fmt.Printf("\nТаблица умножения для числа %d:\n", num)
	fmt.Println("-------------------------")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %-2d = %-3d\n", num, i, num*i)
	}
	fmt.Println("-------------------------")
}

func main() {
	var n int

	fmt.Println("Таблица умножения")
	fmt.Print("Введите число для построения таблицы: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Ошибка: число должно быть не равно нулю.")
		return
	}

	printMultiplicationTable(n)
}