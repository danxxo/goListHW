package main

import (
	"fmt"

	"lesson1/storage/list"
)

func main() {
	var l = list.Newlist()
	l.Add(10)
	l.Add(12)
	l.Add(100)
	l.Add(1000)
	l.Add(10)
	l.Add(10)

	l.Add(10)
	l.Add(10)
	l.Add(10)
	l.Add(10)
	fmt.Print(l.GetAllByValue(10))

	l.Print()

	l.RemoveAllByValue(10)
	l.Print()

	fmt.Print(l.GetAllByValue(10))

}
