package main

import (
	"fmt"
	"math"
)

// Структура Point для представления точки в двумерном пространстве
type Point struct {
	x float64
	y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Метод для вычисления расстояния между двумя точками
func (p1 *Point) DistanceTo(p2 *Point) float64 {
	deltaX := p2.x - p1.x
	deltaY := p2.y - p1.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func main() {
	// Создаем две точки
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 1000)

	// Вычисляем расстояние между точками
	distance := point1.DistanceTo(point2)

	// Выводим результат
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
