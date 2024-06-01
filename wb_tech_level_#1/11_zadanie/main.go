package main

import "fmt"

// Функция для нахождения пересечения двух множеств
//итерируемся по каждому элементу массива сравнивая его
//со всеми элементами из другого массива
func intersection(set1, set2 []int) []int {
	res := []int{}
	for i := 0; i < len(set1); i++ {

		for j := 0; j < len(set2); j++ {
			if set1[i] == set2[j] {
				res = append(res, set1[i])
			}
		}
	}
	//возвращаем массив пересечений множеств
	return res
}

func main() {
	set1 := []int{1, 2, 3, 4, 5, 6, -10, 33, 1000}
	set2 := []int{4, 5, 7, 8, 9, 10, -10, 11, 1000}

	result := intersection(set1, set2)
	fmt.Println("Пересечение множеств:", result)
}
