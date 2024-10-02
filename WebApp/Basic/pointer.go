package main
import "fmt"

func main(){
  a:=100
  var pointer *int
  pointer=&a
  fmt.Println(pointer)
  fmt.Println(*pointer)

  fmt.Println()
  fmt.Printf("T",pointer)

}