package main

import "fmt"

// Устанавливает i-й бит в определенное значение (1 или 0) в числе n.
func changeBit(n int64, i int, bitValue int) int64 {
	// Создаем маску, устанавливающуюся в 1 в позиции i-го бита.
	mask := int64(1 << i)

	// Если bitValue равно 1, устанавливаем i-й бит в 1, иначе в 0.
	if bitValue == 1 {
		return n | mask // Устанавливаем i-й бит в 1.
	}
	return n &^ mask // Устанавливаем i-й бит в 0.
}

func main() {
	// Исходное число
	var num int64 = 256

	// Устанавливаем 3 бит в 0
	num = changeBit(num, 3, 0)
	//число не поменялось, т.к. 3 бит и так был равен 0
	fmt.Println(num)

	// Устанавливаем 3 бит в 0
	num = changeBit(num, 3, 1)
	fmt.Println(num)

	// Устанавливаем 3 бит обратно в 1
	num = changeBit(num, 3, 0)
	fmt.Println(num)

}
