package main

import (
	"fmt"
	"sync"
	"time"
)

// структура счетчика
type counter struct {
	count int
	mu    sync.Mutex
}

// функция увеличения счетчика
func (c *counter) Count() {
	c.mu.Lock()

	defer c.mu.Unlock()
	c.count++

}

func main() {
	//создаем экземпляр счетчика
	Counter := counter{count: 0}

	//конкурентно запускаем первую анонимную функцию для увеличения счетчика
	go func() {
		for {
			Counter.Count()
			fmt.Printf("Первая горутина увеличила счетчик:%d\n", Counter.count)
			time.Sleep(1 * time.Second)

		}
	}()

	//конкурентно запускаем вторую анонимную функцию для увеличения счетчика
	go func() {
		for {
			Counter.Count()
			fmt.Printf("Вторая горутина увеличила счетчик:%d\n", Counter.count)
			time.Sleep(2 * time.Second)

		}
	}()

	//конкурентно запускаем третью анонимную функцию для увеличения счетчика
	go func() {
		for {
			Counter.Count()
			fmt.Printf("Третья горутина увеличила счетчик:%d\n", Counter.count)
			time.Sleep(3 * time.Second)

		}
	}()

	time.Sleep(10 * time.Second)

	fmt.Printf("Итоговое значение счетчика: %d", Counter.count)
}
