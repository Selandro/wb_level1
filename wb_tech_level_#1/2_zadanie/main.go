package main

import (
	"fmt"
	"sync"
)

func main() {

	//иннициаллизируем массив из 5 элементов
	A := [5]int{2, 4, 6, 8, 10}

	//waitgroup для синхронизации горутин
	var wg sync.WaitGroup

	//запуск горутин для расчета квадратов чисел
	for _, value := range A {

		wg.Add(1)

		//передаем значение value в анонимную функцию, для гарантирования
		//правильного захвата value на момент вызова горутины
		go func(value int) {
			defer wg.Done()
			fmt.Println(value * value)
		}(value)
	}

	//ожидаем завершения всех вызванных горутин
	wg.Wait()
}
