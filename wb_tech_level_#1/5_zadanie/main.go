package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// функция, которая последовательно отправляет данные в канал каждую секунду
func Write(dataChan chan<- int, second int, wg *sync.WaitGroup) {

	defer wg.Done()

	//отправляем данные в канал раз в секунду
	for i := second; i >= 0; i-- {
		if i != 0 {
			dataChan <- i
			time.Sleep(1 * time.Second) //иммитация задержки
		} else {
			fmt.Println("Программа завершена!")
			close(dataChan) //при достижении 0, закрываем канал
		}
	}
}

// функция, которая последовательно считывает данные из канала
func Read(dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	//после закрытия канала dataChan функция прекращает работу, отправляя сигнал в waitgroup
	for data := range dataChan {
		fmt.Printf("Секунд до завершения работы: %d\n", data)
	}
}

func main() {

	//просим ввести значение продолжительности работы программы в stdin
	var Second int
	for {
		fmt.Println("Введите продолжительность работы программы: ")
		fmt.Fscan(os.Stdin, &Second)
		if Second < 0 {
			fmt.Println("Продолжительность не может быть отрицательной")
			continue
		} else {
			break
		}
	}

	//waitgroup для синхронизации работы горутин
	var wg sync.WaitGroup

	//иннициализируем канал для передачи данных
	dataChan := make(chan int)

	//запускаем горутину для записи данных в канал
	wg.Add(1)
	go Write(dataChan, Second, &wg)

	//запускаем горутину для чтения данных из канала
	wg.Add(1)
	go Read(dataChan, &wg)

	//ожидание завершения работы всех горутин
	wg.Wait()
}
