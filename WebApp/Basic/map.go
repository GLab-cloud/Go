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
	//map is reference type
	myMap3 := myMap2
	myMap3["key5"] = 50000
	myMap3["key6"] = 60000

	fmt.Println(myMap3)
	fmt.Println(myMap2)
	key := "key5"
	value, found := myMap3[key]
	if found {
		fmt.Println(value)
	} else {
		fmt.Println("not found value with this key")
	}
	addItem(1, 100, 200, 300, 400)
	var list2 = []int{1, 2, 3, 4}
	addItem(100, list2...)
}

// variadic function
func addItem(item int, list ...int) {
	list = append(list, item)
	fmt.Println(list)
}
