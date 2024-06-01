package main

import (
	"fmt"
	"strings"
)

func main() {

	//пример исходной строки
	a := "guys! WB Hello"

	//разбеляем строку на отдельные слова
	words := strings.Fields(a)

	//мапа для хранения отдельных слов
	mapWords := make(map[int]string)

	//заполняем мапу словами, где ключом является индекс
	for i, word := range words {
		mapWords[i] = word
	}

	//массив для хранения слов в обратном порядке
	b := []string{}

	//заполняем массив словами
	for j := len(mapWords) - 1; j >= 0; j-- {

		b = append(b, mapWords[j])
	}

	//преобразуем массив в строку
	res := strings.Join(b, " ")

	//выводим результат
	fmt.Println(res)

}
