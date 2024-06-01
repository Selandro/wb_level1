package main

import (
	"fmt"
)

// OldFinaleCost - старая структура
type OldFinaleCost struct {
	Cost     float64
	Discount int
}

// NewCost - метод для расчета финальной цены в старой структуре
func (f *OldFinaleCost) NewCost() float64 {
	return f.Cost - (f.Cost * float64(f.Discount) / 100)
}

// NewInterface - новый интерфейс
type NewInterface interface {
	NewCost() float64
}

// Adapter - адаптер для OldFinaleCost, реализующий интерфейс NewInterface
type Adapter struct {
	OldFinaleCost *OldFinaleCost
}

// Area - метод адаптера для расчета площади
func (a *Adapter) NewCost() float64 {
	return a.OldFinaleCost.NewCost()
}

func main() {
	// Создаем экземпляр OldFinaleCost
	newCostWithAdapter := &OldFinaleCost{Cost: 400, Discount: 15}

	// Используем адаптер для совместимости с новым интерфейсом
	var shape NewInterface = &Adapter{OldFinaleCost: newCostWithAdapter}

	// Теперь мы можем использовать метод Cost из нового интерфейса
	fmt.Printf("New cost: %.2f\n", shape.NewCost())
}
