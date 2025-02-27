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

	//создадим переменную для записи в нее результата сложения всех квадратов чисел
	res := 0

	//создадим мьютекс для обеспечения синхронизации доступа к res
	var res_mu sync.Mutex

	//запуск горутин для расчета квадратов чисел
	for _, value := range A {

		wg.Add(1)

		//передаем значение value в анонимную функцию, для гарантирования
		//правильного захвата value на момент вызова горутины
		go func(value int) {

			//вызываем блокировку res для записи в нее данных с последующей разблокировкой
			res_mu.Lock()
			defer res_mu.Unlock()
			defer wg.Done()
			res = res + (value * value)
		}(value)
	}

	//ожидаем завершения всех вызванных горутин
	wg.Wait()

	//выводим результат
	fmt.Println(res)
}
