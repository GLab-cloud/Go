package main

import "fmt"

//Interface
type Animal interface {
	speak()
}
type Dog struct{}

func (d Dog) speak() {
	fmt.Println("Woau Woau...")
}
func main() {
	var animal Animal
	animal = Dog{}
	animal.speak()
}
