package main

import "fmt"

func main() {
	//дан исходный массив строк
	array := []string{"cat", "cat", "dog", "cat", "tree"}

	//создаем мапу для хранения уникальных значений из массива
	collection := make(map[string]struct{})

	//заполняем мапу данными. т.к. каждый ключ обязан быть являются уникальным
	//все повторяющиеся ключи будут автоматически удалены
	for _, data := range array {
		collection[data] = struct{}{}
	}

	//создаем множество животных
	animals := []string{}

	//заполняем его уникальными ключами из мапы
	for data := range collection {

		animals = append(animals, data)

	}

	//выводим множество на экран
	fmt.Println(animals)
}
