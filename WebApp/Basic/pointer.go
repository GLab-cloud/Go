package main
import "fmt"

func main(){
  a:=100
  var pointer *int
  pointer=&a
  fmt.Println(pointer)
  fmt.Println(*pointer)

  fmt.Println()
  fmt.Printf("%T",pointer)
  fmt.Println()
  fmt.Println(a)
  *pointer=99
  fmt.Println(a)
  //pointer->array

  array:=[3]int{1,2,3}
  var ptn *[3]int
  ptn=&array
  fmt.Println(ptn)
  fmt.Println()
  fmt.Printf("%T",ptn)
  fmt.Println()

   c:=45
   var pointer2 *int=&c
   applyPointer(pointer2)
   fmt.Println(c)
}
func applyPointer(pointer *int){
	*pointer=7777
}