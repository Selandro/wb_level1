package main

import (
	"fmt"
)

func main() {

	//исходный массив данных
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 38.1, 10.0, 13.1, -11.0, -14.5, 25.7}

	//создаем карту для хранения данных
	sortedTemp := make(map[int][]float64)

	//группировка по температуре с шагом 10
	for _, temp := range temperatures {
		// Округляем температуру до кратного 10, результат будет являться ключом
		key := int(temp/10) * 10
		sortedTemp[key] = append(sortedTemp[key], temp)
	}

	//выводим результат
	for key, values := range sortedTemp {
		fmt.Printf("%d: %v\n", key, values)
	}
}
