package mp

import "fmt"

type mp struct {
	HashTable map[int64]int64
	len       int64
}

func (myMap *mp) isEmpty() bool {
	if myMap.len <= 0 {
		return true
	}
	return false
}

func NewMap() *mp {
	hashTable := make(map[int64]int64, 0)
	return &mp{HashTable: hashTable}
}

func (myMap *mp) Clear() {
	myMap.HashTable = map[int64]int64{}
	myMap.len = 0
}

func (myMap *mp) Len() int64 {
	return myMap.len
}

func (myMap *mp) GetAll() (result []int64, ok bool) {
	if myMap.isEmpty() {
		return result, false
	}
	var i int64 = 0
	result = make([]int64, myMap.len, myMap.len)
	for ; i < myMap.len; i++ {
		result[i] = myMap.HashTable[i]
	}
	return result, true
}

func (myMap *mp) GetAllByValue(value int64) (ids []int64, ok bool) {
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

func (myMap *mp) Print() {
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

func (myMap *mp) GetByIndex(id int64) (value int64, ok bool) {
	if id >= myMap.len {
		return 0, false
	}
	return myMap.HashTable[id], true
}

func (myMap *mp) GetByValue(value int64) (id int64, ok bool) {
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

func (myMap *mp) Add(data int64) {
	myMap.HashTable[myMap.len] = data
	myMap.len++
}

func (myMap *mp) RemoveByIndex(id int64) (ok bool) {

	if myMap.isEmpty() {
		fmt.Println("Map is empty")
		return false
	}

	if id < 0 || id > myMap.len {
		return false
	}

	if id == myMap.len-1 {
		delete(myMap.HashTable, id)
		myMap.len--
		return true
	}

	id += 1
	for ; id < myMap.len; id++ {
		id_ := id - 1
		myMap.HashTable[id_] = myMap.HashTable[id]
	}
	delete(myMap.HashTable, myMap.len-1)
	myMap.len--
	return true

}

func (myMap *mp) RemoveAllByValue(value int64) (ok bool) {
	if myMap.isEmpty() {
		return false
	}
	//var i int64 = 0
	finded := myMap.RemoveByValue(value)
	for finded {
		finded = myMap.RemoveByValue(value)
		ok = true
	}
	return

}

func (myMap *mp) RemoveByValue(value int64) (ok bool) {
	if myMap.isEmpty() {
		return false
	}

	var i int64 = 0
	for ; i < myMap.len; i++ {
		if myMap.HashTable[i] == value {
			myMap.RemoveByIndex(i)
			return true
		}
	}
	return false
}
