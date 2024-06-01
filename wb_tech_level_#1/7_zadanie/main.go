package main

import (
	"fmt"
	"sync"
)

// функция записи данных в мапу
func WriteOne(a map[int]int, mu *sync.Mutex, wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 0; i < 10; i++ {
		mu.Lock()
		a[i] = i
		mu.Unlock()
	}
}

// функция записи данных в мапу
func WriteTwo(a map[int]int, mu *sync.Mutex, wg *sync.WaitGroup) {

	defer wg.Done()
	for j := 10; j < 20; j++ {
		mu.Lock()
		a[j] = j
		mu.Unlock()
	}
}

func main() {

	//мьютекс для синхронизации доступа к мапе
	var mu sync.Mutex

	//waitgroup для синхронизации работы горутин
	var wg sync.WaitGroup

	//иннициализация мапы
	a := make(map[int]int)

	//увеличение счетчика работающих горутин
	wg.Add(1)

	//параллельный запуск первой функции для записи
	go WriteOne(a, &mu, &wg)

	//увеличение счетчика работающих горутин
	wg.Add(1)

	//параллельный запуск второй функции для записи
	go WriteTwo(a, &mu, &wg)

	//ожидание завершения работы горутин
	wg.Wait()

	//вывод результата
	fmt.Println(a)
}
