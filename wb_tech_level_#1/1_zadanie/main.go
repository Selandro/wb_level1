package main

import "fmt"

//создаем структуру Human
type Human struct {
	Id   int
	Name string
	Age  int
}

//создаем структуру Action с вложенной структурой Human
type Action struct {
	Human
	Group string
}

//создаем метод структуры Human для получения Id
func (h Human) humanId() int {
	return h.Id
}
func (h Human) humanName() string {
	return h.Name
}

func main() {

	//инициализируем экземпляр структуры Action с инициализированной структурой Human
	Vanya := Action{
		Human: Human{
			Id:   122,
			Name: "Vanya",
			Age:  15,
		},
		Group: "A",
	}

	//вызываем методы структуры Human, через экземпляр Action
	Age := Vanya.humanId()
	Name := Vanya.humanName()
	//вывод результатов
	fmt.Println(Age)
	fmt.Println(Name)

}
