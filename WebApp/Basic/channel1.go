package main

import "fmt"

func main(){
//close channel
queue1:=make(chan int,10)
go func ()  {
 for i:=0;i<10;i++{
  queue1<-i
 }
 close(queue1)

}()
for value:=range queue1{
  fmt.Println(value)
}
}