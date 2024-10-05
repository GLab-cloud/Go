package main

import "fmt"

func main(){
//close channel
queue1:=make(chan int,10)
go func ()  {
 for i:=0;i<10;i++{
	if i>5 {
		close(queue1)
		//queue1<-100
        break
	} else{
		queue1<-i
	}
 }

}()
for value:=range queue1{
  fmt.Println(value)
}
}