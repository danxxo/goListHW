package main

import (
	"fmt"

	list "storage.com/storage/list"
	mp "storage.com/storage/map"
)

func main() {

	type Storage interface {
		Clear()
		Len()
		GetAll()
		GetAllByValue(int64) ([]int64, bool)
		Print()
		GetByIndex(int64) (int64, bool)
		GetByValue(int64) (int64, bool)
		Add(int64)
		RemoveAllByValue(int64) bool
		RemoveByIndex(int64) bool
		RemoveByValue(int64) bool
	}

	myMap := mp.NewMap()
	myMap.Add(1)

	myList := list.Newlist()
	myList.Add(1)
	myMap.Clear()
	exp, ok := myMap.GetAll()
	fmt.Println(exp, ok)
}
