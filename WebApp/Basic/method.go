package main
import "fmt"

type Student struct {
	name string
}
//define method
func (s Student) getName() string {
return s.name
}
func main(){
student:=Student{name: "Batman"}
name:=student.getName()
fmt.Println(name)
}