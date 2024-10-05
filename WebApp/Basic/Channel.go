package main

import "fmt"

func main(){
	//unbuferred channels
	ch:=make(chan int)
	//go - annonymous func
	go func(){
        ch<-100
		fmt.Println("Sent") //deadlock until another channel - main go take value
	}()
	fmt.Println(<-ch)//main go take value, //deadlock until another channel - go sent value to channel
	fmt.Println("Done")
	//buferred channels
	ch1:=make(chan int,2)
	ch1<-1 //not deadlock until another channel - main go take value
	ch1<-2 //not deadlock
	//ch1<-3 //deadlock
	close(ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)//deadlock, if not close(ch1)
	fmt.Println(<-ch1)


	//select
	queue:=make(chan int)
	done:=make(chan bool)
	fmt.Println("select")

	go func(){
		for i:=0;i<10;i++{
			queue<-i
		}
		done<-true
	}()
	for {
		select {
		case v:=<-queue:
			fmt.Println(v)
		case <-done:
			fmt.Println("Done")
			return
		}
	}
  

}