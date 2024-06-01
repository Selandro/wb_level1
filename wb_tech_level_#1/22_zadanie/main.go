package main

import (
	"fmt"
)

func main() {
	//Задаем значения переменных a и b
	//используя битовый сдвиг
	a := 111 << 20 // a > 2^20
	b := 222 << 20 // b > 2^20

	// Выполняем операции
	sum := a + b
	sub := a - b
	mul := a * b
	div := float64(a) / float64(b)

	// Выводим результаты
	fmt.Printf("результат сложения: %d\n", sum)
	fmt.Printf("результат вычитания: %d\n", sub)
	fmt.Printf("результат умножения: %d\n", mul)
	fmt.Printf("результат деления: %f\n", div)
}
