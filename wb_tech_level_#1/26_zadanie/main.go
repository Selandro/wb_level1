package main

import (
	"fmt"
	"strings"
)

// функция проверки на одинаковые символы
func CheckArray(array string) bool {
	var res bool = true

	//приводим все символы в нижний регистр
	array = strings.ToLower(array)

	for i := 0; i < len(array); i++ {

		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				res = false
				break

			}
		}
		if !res {
			break
		}
	}
	return res
}

func main() {

	//вызываем функции для разных строк
	arrayOne := "abcd"
	fmt.Println(CheckArray(arrayOne)) //вывод: true

	arrayTwo := "abCdefAaf"
	fmt.Println(CheckArray(arrayTwo)) //вывод: false

	arrayThree := "aabcd"
	fmt.Println(CheckArray(arrayThree)) //вывод: false

}
