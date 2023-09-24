package main

import (
	list "storage.com/storage/list"
	mp "storage.com/storage/map"
)

func main() {

	type Storage interface {
		Clear()
		Len() int64
		GetAll() ([]int64, bool)
		GetAllByValue(int64) ([]int64, bool)
		Print()
		GetByIndex(int64) (int64, bool)
		GetByValue(int64) (int64, bool)
		Add(int64)
		RemoveAllByValue(int64) bool
		RemoveByIndex(int64) bool
		RemoveByValue(int64) bool
	}

	var myStorage Storage

	// list
	myStorage = list.Newlist()
	myStorage.Add(3)
	myStorage.Add(1)
	myStorage.Add(2)
	myStorage.Add(3)
	myStorage.RemoveAllByValue(3)
	myStorage.Print()

	// map
	myStorage = mp.NewMap()
	myStorage.Add(3)
	myStorage.Add(1)
	myStorage.Add(2)
	myStorage.Add(3)
	myStorage.RemoveAllByValue(3)
	myStorage.Print()
}
