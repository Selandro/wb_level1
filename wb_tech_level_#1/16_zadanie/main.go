package main

import "fmt"

func quickSort(ar []int) []int {

	//если длина массива меньше двух, возвращаем массив
	if len(ar) < 2 {
		return ar
	}

	//в качестве pivot будем использовать середину массива
	pivotIndex := len(ar) / 2
	pivot := ar[pivotIndex]

	var low, up []int

	// Разделяем элементы относительно опорного элемента
	for i, num := range ar {
		// если число является pivot, пропускаем его,
		//чтобы он не оказался в одном из подмассивов
		if i == pivotIndex {
			continue
		}

		//если число меньше pivot, отправляем его в массив с меньшими числами,
		//если больше, то в массив с большими
		if num <= pivot {
			low = append(low, num)
		} else {
			up = append(up, num)
		}
	}

	//рекурсивно сортируем и объединяем
	sortedLow := quickSort(low)
	sortedUp := quickSort(up)

	sortedArray := append(sortedLow, pivot)
	sortedArray = append(sortedArray, sortedUp...)
	return sortedArray
}

func main() {

	//массив данных для сортировки
	array := []int{4, 1, 2, 53, 23, 123, 62, -12, -11, 10, 12, 11, 0, 100, -1000}

	//вызываем функцию быстрой сортировки массива
	fmt.Println(quickSort(array))
}
