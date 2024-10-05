package main

import "fmt"

func main(){
	//unbuferred channels
	ch:=make(chan int)
	//go - annonymous func
	go func(){
        ch<-100
		fmt.Println("Sent") //locked until another channel - main go take value
	}()
	fmt.Println(<-ch)//another channel - main go take value
	fmt.Println("Done")
}