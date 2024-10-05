package main

import (
	"fmt"
	"runtime"
	"time"
)

func go1(){
	fmt.Println("go 1")
}
func main(){
	go go1()
	numGo:=runtime.NumGoroutine()
	fmt.Println(numGo)
	time.Sleep(time.Second)
}
