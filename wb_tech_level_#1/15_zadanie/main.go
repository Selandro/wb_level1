package main

import "fmt"

//в изначальном примере может возникнуть проблема с управлением паматью,
//justString будет содержать ссылку на оригинальную строку, и она не будет
//высвобождена из памяти

//Такой подход устраняет потенциальные проблемы с управлением памятью,
//поскольку теперь justString содержит независимую копию первых 100
//символов, а не ссылается на оригинальную большую строку.
//создадим функцию, которая возвращает строку указанного размера

var justString string

func createHugeString(size int) string {

	hugeString := make([]byte, size)
	for i := range hugeString {
		hugeString[i] = 'a'
	}

	//выведем результат для наглядности
	fmt.Printf("first string: %s\n", string(hugeString))
	return string(hugeString)
}

func someFunc() {

	//создадим строку большого размера с помощью функции createHugeString
	//указав, что нужно сдвинуть 1 влево на 10 битов
	v := createHugeString(1 << 10)
	// с помощью string() мы создаем новую строку из первых 100 элементов
	justString = string(v[:100])
}

func main() {

	someFunc()
	// Печатаем результат для проверки
	fmt.Printf("new string: %s", justString)
}
