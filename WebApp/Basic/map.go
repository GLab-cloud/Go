package main

import (
	"fmt"
)

func main() {
	var myMap = make(map[string]int)
	var myMap1 map[string]int
	fmt.Println(myMap)
	if myMap == nil {
		fmt.Println("mymap==nil")
	}
	if myMap1 == nil {
		fmt.Println("mymap1==nil")
	}
	//assign map values
	myMap2 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 5,
		"key4": 6,
	}
	fmt.Println(myMap2)
	myMap2["key5"] = 50
	myMap2["key6"] = 60
	fmt.Println(myMap2)
	delete(myMap2, "key1")
	fmt.Println(myMap2)
}
