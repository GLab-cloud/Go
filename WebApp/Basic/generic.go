package main

import (
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
func SumIntsOrFloats[K comparable, V int64|float64] (m map[K]V) V{
	var s V
	for _,v:=range m{
		s+=v
	}
	return s
}
func main (){
	slice_int:=[]int{10,20,-10,8}
	fmt.Println(Index(slice_int,20))
	slice_string:=[]string{"foo","bar","go","c++"}
	fmt.Println(Index(slice_string,"c++"))
	ints:=map[string]int64{
		"first":3,
		"second":8,
	}
	floats:=map[string]float64{
		"first":3.6,
		"second":8.5,
	}
	fmt.Println(SumIntsOrFloats(ints))
	fmt.Println(SumIntsOrFloats(floats))

    for _,v:=range ints{
		fmt.Println(v)
	}
	for _,v:=range floats{
		fmt.Println(v)
	}
	var slice_int1[]int=slice_int[1:4]
	fmt.Println(slice_int1)
    a:=slice_int[:2]
	b:=slice_int[1:3]
	fmt.Println(a, b)
    a[0]=3
	b[0]=0
	fmt.Println(slice_int)
}
