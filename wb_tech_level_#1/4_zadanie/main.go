package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Generate генерирует случайные данные и отправляет их в канал
func Generate(dataChan chan<- int, doneChan <-chan struct{}) {
	for {
		select {
		case <-doneChan:
			close(dataChan) // Закрываем канал при получении сигнала завершения
			return
		default:
			data := rand.Intn(200) - 100 // Генерируем случайное число от -100 до 100
			dataChan <- data

		}
	}
}

// worker читает данные из канала и выводит их в stdout
func worker(id int, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range dataChan {
		fmt.Printf("worker %d прочитал данные: %d\n", id+1, data)
	}
}

func main() {
	var numWorkers int
	for {
		fmt.Println("Введите колличество worker'ов:")
		fmt.Fscan(os.Stdin, &numWorkers)
		if numWorkers <= 0 {
			fmt.Println("колличество worker'ов не может быть меньше одного")
			continue
		} else {
			break
		}
	}

	//создаем канал для передачи данных
	dataChan := make(chan int)

	// Для сигнализирования о завершении работы будем использовать канал типа struct{},
	// Такой подход имеет ряд преимуществ, которые заключаются в простоте и ясности использования,
	// он достаточно надежен , поскольку каждый поток может проверить канал,
	// и он достаточно эффектифен, поскольку каналы реализованы с использованием низкоуровневых
	// механизмов синхронизации, что делает их эффективным инструментом для обмена данными между горутинами.

	//создаем канал типа struct{}, который не содержит данных,
	//но используется для сигнализирования о прекращении работы для всех горутин
	doneChan := make(chan struct{})
	var wg sync.WaitGroup

	// Запускаем поток записи данных
	go Generate(dataChan, doneChan)

	// Запускаем воркеров
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChan, &wg)
	}

	// Канал для получения системных сигналов и перенаправления их в канал sigChan
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Ожидание сигнала прерывания
	<-sigChan
	fmt.Println("Получен сигнал о завершении работы")

	// Отправка сигнала завершения, после получения сигнала о закрытие канала doneChan,
	// функция Generate закроет канал dataChan, что послужит сигналом для worker о завершении работы
	close(doneChan)

	// Ожидание завершения всех воркеров
	wg.Wait()
	fmt.Println("все worker завершили работу")
}
