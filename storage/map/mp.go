package mp

import (
	"errors"
	"fmt"
	"reflect"
)

var ErrMismatchType = errors.New("mismatched type: the type of the provided value does not match the type of items already in the storage")

type Mp struct {
	HashTable map[int64]interface{}
	len       int64
}

func (myMap *Mp) isEmpty() bool {
	if myMap.len <= 0 {
		return true
	}
	return false
}

func NewMap() *Mp {
	hashTable := make(map[int64]interface{}, 0)
	return &Mp{HashTable: hashTable}
}

func (myMap *Mp) Clear() {
	myMap.HashTable = map[int64]interface{}{}
	myMap.len = 0
}

func (myMap *Mp) Len() int64 {

	return myMap.len
}

func (myMap *Mp) GetAll() (result []interface{}, ok bool) {

	if myMap.isEmpty() {
		return result, false
	}
	var i int64 = 0
	result = make([]interface{}, myMap.len, myMap.len)
	for ; i < myMap.len; i++ {
		result[i] = myMap.HashTable[i]
	}
	return result, true
}

func (myMap *Mp) GetAllByValue(value interface{}) (ids []int64, ok bool) {

	if myMap.isEmpty() {
		return
	}
	var i int64 = 0
	for ; i < myMap.len; i++ {
		if myMap.HashTable[i] == value {
			ids = append(ids, i)
		}
	}
	if ids != nil {
		return ids, true
	}
	return
}

func (myMap *Mp) Print() {

	if myMap.isEmpty() {
		fmt.Println("Map is empty")
		return
	}
	var i int64 = 0
	for ; i < myMap.len-1; i++ {
		fmt.Print(myMap.HashTable[i], "[", i, "], ")
	}
	fmt.Print(myMap.HashTable[i], "[", i, "]")
	fmt.Println("")
}

func (myMap *Mp) GetByIndex(id int64) (value interface{}, ok bool) {

	if id >= myMap.len {
		return 0, false
	}
	return myMap.HashTable[id], true
}

func (myMap *Mp) GetByValue(value interface{}) (id int64, ok bool) {
	if myMap.isEmpty() {
		return 0, false
	}
	var i int64 = 0
	for ; i < myMap.len; i++ {
		if myMap.HashTable[i] == value {
			return i, true
		}
	}
	return 0, false
}

func (myMap *Mp) Add(data interface{}) (int64, error) {
	if myMap.isEmpty() {
		myMap.HashTable[0] = data
		myMap.len++
		return 0, nil
	}

	if reflect.TypeOf(data) != reflect.TypeOf(myMap.HashTable[0]) {
		return 0, ErrMismatchType
	}

	myMap.HashTable[myMap.len] = data
	myMap.len++
	return myMap.len - 1, nil
}

func (myMap *Mp) RemoveByIndex(id int64) {

	if myMap.isEmpty() {
		fmt.Println("Map is empty")
		return
	}

	if id < 0 || id > myMap.len {
		return
	}

	if id == myMap.len-1 {
		delete(myMap.HashTable, id)
		myMap.len--
		return
	}

	id += 1
	for ; id < myMap.len; id++ {
		id_ := id - 1
		myMap.HashTable[id_] = myMap.HashTable[id]
	}
	delete(myMap.HashTable, myMap.len-1)
	myMap.len--
	return

}

func (myMap *Mp) Count(value interface{}) int {
	counter := 0
	var i int64 = 0
	for ; i < myMap.len; i++ {
		if myMap.HashTable[i] == value {
			counter++
		}
	}
	return counter
}

func (myMap *Mp) RemoveAllByValue(value interface{}) {
	amount := myMap.Count(value)
	for ; amount != 0; amount-- {
		myMap.RemoveByValue(value)
	}
}

func (myMap *Mp) RemoveByValue(value interface{}) {

	if myMap.isEmpty() {
		return
	}

	var i int64 = 0
	for ; i < myMap.len; i++ {
		if myMap.HashTable[i] == value {
			myMap.RemoveByIndex(i)
			return
		}
	}
}
