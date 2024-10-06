package main

import (
	"fmt"
	//"os/user"
)

    //class
	type User struct{
		Email string
		Age int
	}
	//constructor
	func New() User{
		return User{Email:"trung@gmail.com"}
	}
	//getter
	func (user *User) getEmail() string{
		return user.Email
	}
	//getter
	func (user *User) getAge() int{
		return user.Age
	}
	//setter
	func (user *User) setAge(age int){
		user.Age=age
	}

func main(){
u:= New()
fmt.Println(u.getEmail())
u.setAge(18)
fmt.Println(u.getAge())

}