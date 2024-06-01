package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

//при применении атомарных операций, нет необходимости использовать mutex
//т.к. атомарные операции безопасны в конкурентной среде

// структура счетчика
type counter struct {
	count int64
}

// функция увеличения счетчика
func (c *counter) Count() {
	atomic.AddInt64(&c.count, 1)
}

func main() {
	var Counter counter

	//конкурентно запускаем первую анонимную функцию для увеличения счетчика
	go func() {

		for {
			Counter.Count()
			fmt.Printf("Первая горутина увеличила счетчик: %d\n", (Counter.count))
			time.Sleep(1 * time.Second)

		}

	}()

	//конкурентно запускаем вторую анонимную функцию для увеличения счетчика
	go func() {

		for {
			Counter.Count()
			fmt.Printf("Вторая горутина увеличила счетчик: %d\n", (Counter.count))
			time.Sleep(2 * time.Second)

		}

	}()

	//конкурентно запускаем третью анонимную функцию для увеличения счетчика
	go func() {

		for {
			Counter.Count()
			fmt.Printf("Третья горутина увеличила счетчик: %d\n", (Counter.count))
			time.Sleep(3 * time.Second)

		}

	}()

	time.Sleep(10 * time.Second)

	fmt.Printf("Итоговое значение счетчика: %d", atomic.LoadInt64(&Counter.count))
}
