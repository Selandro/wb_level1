package main

import (
	"fmt"
	"reflect"
)

// В языке Go нет прямого способа определить тип переменной в рантайме,
// когда она представлена как interface{}. Это связано с тем,
// что interface{} является пустым интерфейсом,
// который может содержать значения любого типа. Информация
// о конкретном типе переменной теряется при присвоении её interface{}.

// Но можно проверить, является ли переменная значением определенного
// типа, используя механизм утверждения типа (type assertion) в Go.
func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Переменная является int:", v)
	case string:
		fmt.Println("Переменная является string:", v)
	case bool:
		fmt.Println("Переменная является bool:", v)
	case chan int:
		fmt.Println("Переменная является каналом для int:", v)
	default:
		fmt.Printf("Неизвестный тип: %v\n", reflect.TypeOf(i))
	}
}

func main() {
	var a interface{}

	// помещаем данные типа int
	a = 111
	checkType(a)

	// помещаем данные типа string
	a = "awd"
	checkType(a)

	// помещаем данные типа bool
	a = true
	checkType(a)

	// помещаем данные типа chan
	ch := make(chan int)
	a = ch
	checkType(a)

	// помещаем данные неизвестного типа
	a = struct{}{}
	checkType(a)
}
