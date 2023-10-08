package main

import (
	"fmt"

	Storage "storage.com/storage"
)

func main() {

	var myStorage Storage.Storage

	myStorage = Storage.NewList()
	myStorage.Add(10)
	for i := 0; i < 10; i++ {
		myStorage.Add(i)
		myStorage.Add(i)
	}
	myStorage.RemoveByIndex(10)
	myStorage.RemoveByValue(4)
	myStorage.RemoveAllByValue(2)
	myStorage.RemoveAllByValue(11)
	myStorage.Print()
	fmt.Println(myStorage.Add("wsdefrtg"))
	fmt.Println(myStorage.GetAll())
	fmt.Println(myStorage.GetAllByValue(9))
	fmt.Println(myStorage.GetByIndex(10))
}
