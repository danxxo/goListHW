package main

import (
	list "storage.com/storage/list"
	mp "storage.com/storage/map"
)

func main() {

	l := list.Newlist()
	l.Add(1)

	myMap := mp.NewMap()
	var i int64 = 0
	for ; i < 7; i++ {
		myMap.Add(i)
	}

	i = 0
	for ; i < 7; i++ {
		myMap.Print()
		myMap.RemoveByIndex(0)
	}

}
