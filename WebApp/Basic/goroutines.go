package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func go1(){
	fmt.Println("go 1")
	wg.Done()

}
func go2(){
	fmt.Println("go 2")
	wg.Done()
}
var wg sync.WaitGroup
func main(){

	numGo:=runtime.NumGoroutine()
	fmt.Println(numGo)
	time.Sleep(time.Second)
	//synchronized goroutines
   /*logic 1

   go1()
   go2()

   logic 2
   */
wg.Add(2)
fmt.Println("...begin")
go go1()
go go2()
wg.Wait()
fmt.Println("...end")

}
