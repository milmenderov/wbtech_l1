package main

import (
	"fmt"
)

type Human struct {
	Name string
	Age  int
}

func (h Human) Describe() {
	fmt.Printf("Name: %s, Age: %d\n", h.Name, h.Age)
}

func (h *Human) SetName(newName string) {
	h.Name = newName
}

type Action struct {
	Human  // Встраивание структуры Human
	Action string
}

// Дополнительный метод для структуры Action
func (a Action) Perform() {
	fmt.Printf("%s is performing: %s\n", a.Name, a.Action)
}

func main() {
	// Создание объекта Action
	act := Action{
		Human:  Human{Name: "John", Age: 30},
		Action: "running",
	}

	// Вызываем метод Describe, унаследованный от Human
	act.Describe()

	// Изменяем имя, используя метод SetName от Human
	act.SetName("Mike")

	// Снова вызываем Describe для отображения изменений
	act.Describe()
	// Выполняем действие, используя метод Perform от Action
	act.Perform()
}

func sum(a, b int) int {
	return a + b
}
