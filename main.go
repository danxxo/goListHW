package main

import (
	"fmt"

	list "storage.com/storage/list"
	mp "storage.com/storage/map"
)

func main() {

	l := list.Newlist()
	l.Add(1)

	myMap := mp.NewMap()
	myMap.Add(-10)
	myMap.Add(100)
	myMap.Add(-10)

	myMap.Print()
	fmt.Print("")
	fmt.Println(myMap.GetAll())
	fmt.Println(myMap.GetByIndex(0))
	fmt.Println(myMap.GetByValue(-10))

	fmt.Println(myMap.Len())
	fmt.Println(myMap.GetAllByValue(-100))

}
