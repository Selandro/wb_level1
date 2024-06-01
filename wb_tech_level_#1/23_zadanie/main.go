package main

import "fmt"

// Удаление элемента из среза
func delElem(slice *[]int, index int) {
	if index < len(*slice) {
		*slice = append((*slice)[:index], (*slice)[index+1:]...)
	}
}

func main() {
	// Исходный срез
	slice := []int{1, 2, 3, 4, 5}

	// Индекс элемента, который нужно удалить
	index := 4

	// Вызываем функцию удаления элемента из среза
	delElem(&slice, index)

	// Вывод среза после удаления элемента
	fmt.Println(slice)
}
