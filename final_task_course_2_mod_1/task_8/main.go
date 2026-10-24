package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var num int

	fmt.Println("Проверка простого числа")
	fmt.Print("Введите целое число: ")
	fmt.Scan(&num)

	if isPrime(num) {
		fmt.Printf("Число %d является простым.\n", num)
	} else {
		fmt.Printf("Число %d не является простым.\n", num)
	}
}