package main

import "fmt"

type Human struct {
	// Родительская структура

	name string
	age  int
}

func (Human) sayHello() {
	fmt.Println("Hello")
}

func (h Human) greet() {
	fmt.Printf("I am %v and I am %v years old\n", h.name, h.age)
}

type Action struct {
	// Структура, которую мы композируем (типа наследуемся от родителя)

	human Human
}

func (a Action) act() {
	a.human.sayHello()
	a.human.greet()
}

func main() {
	human := Human{name: "Billy", age: 20}
	action := Action{human: human}
	action.act()
}
