package main
import(
	"fmt"
)

func Index[T comparable](s[]T,x T) int{
	for i,v:=range s{
		//v,i are type T, which has the comparable
		if v==x{
		return i
		}
	}
	return -1
}
func main (){
	slice_int:=[]int{10,20,-10,8}
	fmt.Println(Index(slice_int,20))
	slice_string:=[]string{"foo","bar","go","c++"}
	fmt.Println(Index(slice_string,"c++"))

}