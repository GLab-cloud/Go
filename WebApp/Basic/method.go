package main
import "fmt"

type Student struct {
	name string
}
//define method
//func(t type) methodName (params) returns{body code}
//t type
//1. Value receiver
//2. Pointer receiver

func (s Student) getName() string {
return s.name
}
//1. Value receiver
func (s Student) changeName(newName string) {
	fmt.Printf("p2=%p",&s)
	s.name=newName
}
//2. Pointer receiver
func (s *Student) changeName2(newName string) {
	fmt.Printf("p3=%p",s)
	fmt.Printf("p4=%p",*s)

	s.name=newName
}
func main(){
student:=Student{name: "Batman"}
fmt.Printf("p1=%p",&student)

name:=student.getName()
fmt.Println(name)
student.changeName("Superman")
fmt.Println(student.name)
student.changeName2("Superman")
fmt.Println(student.name)
}