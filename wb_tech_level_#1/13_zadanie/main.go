package main

import "fmt"

//замена местами двух чисел, без использования доп переменной.
//функция использует указатели
func change(c *int, d *int) {
	*c = *c + *d
	*d = *c - *d
	*c = *c - *d
}

func main() {
	a := 1
	b := 2
	c := 7
	d := 19
	fmt.Printf("До замены: a=%d, b=%d\n", a, b)
	fmt.Printf("До замены: c=%d, d=%d\n", c, d)

	//меняем значения переменных a и b с помощью побитовых операций
	//используем операцию исключающего ИЛИ (XOR)
	a = a ^ b
	b = a ^ b
	a = a ^ b

	//замена
	change(&c, &d)

	fmt.Printf("После замены: a=%d, b=%d\n", a, b)
	fmt.Printf("После замены: c=%d, d=%d\n", c, d)

}
