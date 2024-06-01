package main

import (
	"fmt"
)

func main() {

	//пример исходной строки
	a := "!syug BW olleH"

	//создаем массив рун из данной строки
	array := []rune{}
	for _, e := range a {

		array = append(array, e)

	}
	fmt.Println(array)

	//заполняем слайс символами в обратном порядке
	rsvArray := make([]rune, len(array))
	for i, val := range array {
		rsvArray[len(array)-1-i] = val
	}

	//преобразуем слайс в строку
	rsvA := string(rsvArray)

	//выводим результат
	fmt.Println(rsvA)

}
