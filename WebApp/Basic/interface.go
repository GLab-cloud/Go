package main

import "fmt"

//Interface
type Animal interface {
	speak()
}
type Movement interface {
	move()
}
type Dog struct{}

type MoveAnimal interface {
	Animal
	Movement
}

func (d Dog) speak() {
	fmt.Println("Woau Woau...")
}
func (d Dog) move() {
	fmt.Println("dog move with 4 legs")
}

//empty interface -generic
func goout(i interface{}) {
	fmt.Println(i)
}

type data struct {
	index int
}

func main() {
	var animal Animal
	animal = Dog{}
	animal.speak()
	dog := Dog{}
	dog.move()
	dog.speak()
	nextAnimal := dog
	nextAnimal.move()
	nextAnimal.speak()
	goout(10)
	goout(10.17)
	d := data{index: 123}
	goout(d)
}
