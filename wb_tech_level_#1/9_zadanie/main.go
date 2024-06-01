package main

import (
	"fmt"
	"sync"
	"time"
)

// функция записи данных из массива в первый канал
func Write(array []int, ch1 chan<- int, wg *sync.WaitGroup) {

	defer close(ch1)
	defer wg.Done()

	for i := 0; i < len(array); i++ {
		time.Sleep(1 * time.Second)
		ch1 <- array[i]
	}

}

// функция преобразования данных из первого канала и записью результата во второй
func Change(ch1 <-chan int, ch2 chan<- int, wg *sync.WaitGroup) {
	defer close(ch2)
	defer wg.Done()
	for data := range ch1 { // Читаем данные из канала ch1 до его закрытия
		ch2 <- data * 2 // Умножаем данные на 2 и отправляем в канал ch2

	}
}

// функция чтения данных из 2 канала и вывода в stdout
func Read(ch2 <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch2 { // Читаем данные из канала ch2 до его закрытия
		fmt.Println(data) // Выводим данные в stdout

	}

}

func main() {

	var wg sync.WaitGroup

	//создаем массив чисел
	array := [10]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	//создаем первый канал для передачи данных
	ch1 := make(chan int)
	//создаем второй канал для преобразованных данных
	ch2 := make(chan int)

	//увеличиваем счетчик waitgroup на число, соответствующее
	//колличеству запускаемых параллельно горутин
	wg.Add(3)

	//запускаем функцию записи чисел из массива в первый канал
	go Write(array[:], ch1, &wg)

	//запускаем функцию для преобразования данных из первого канала
	//и записи из во второй
	go Change(ch1, ch2, &wg)

	//запускаем функцию чтения преобразованных данных из второго канала
	//и вывода их в stdout
	go Read(ch2, &wg)

	wg.Wait()
	fmt.Println("Работа конвейера завершена!")

}
