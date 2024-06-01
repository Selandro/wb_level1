package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Способ 1: Использование канала для остановки горутины
func workerWithChannel(done chan bool) {
	fmt.Println("Рабочая горутина начала выполнение... (канал)")
	time.Sleep(2 * time.Second)
	fmt.Println("Рабочая горутина завершила выполнение... (канал)")
	done <- true
}

// Способ 2: Использование контекста для остановки горутины
func workerWithContext(ctx context.Context) {
	fmt.Println("Рабочая горутина начала выполнение... (контекст)")
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Рабочая горутина завершила выполнение... (контекст)")
	case <-ctx.Done():
		fmt.Println("Рабочая горутина была остановлена... (контекст)")
	}
}

// Способ 3: Использование сигналов ОС для остановки горутины
func workerWithSignal() {
	fmt.Println("Рабочая горутина начала выполнение... (сигнал)")
	time.Sleep(2 * time.Second)
	fmt.Println("Рабочая горутина завершила выполнение... (сигнал)")
}

func main() {
	// Способ 1: Использование канала
	done := make(chan bool)
	go workerWithChannel(done)
	<-done

	// Способ 2: Использование контекста
	ctx, cancel := context.WithCancel(context.Background())
	go workerWithContext(ctx)
	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(2 * time.Second) // Ждем завершения горутины

	// Способ 3: Использование сигналов ОС
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go workerWithSignal()
	<-c
	fmt.Println("Главная горутина завершила выполнение...")
}
