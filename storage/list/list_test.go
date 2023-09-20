package list

import (
	"testing"
)

func TestNewList(t *testing.T) {
	l := Newlist()
	if l.len != 0 {
		t.Errorf("Test failed, newList length != 0")
	}
	if l.fn != nil {
		t.Errorf("Test failed, firstNode of newList is not nil")
	}
}

func TestAdd(t *testing.T) {
	l := Newlist()
	var i int64 = 0
	var testingArray []int64
	for ; i < 100; i++ {
		l.Add(i)
		testingArray = append(testingArray, i)
	}
	if l.Len() != int64(len(testingArray)) {
		t.Errorf("Error, Len must be 100.")
	}

	listToArray, ok := l.GetAll()
	if !ok {
		t.Errorf("GetAll returned false as second arg")
	}

	for i := 0; i < int(l.len); i++ {
		if testingArray[i] != listToArray[i] {
			t.Errorf("same values are not equal in testingArray and listArray")
		}
	}
}

func TestGet(t *testing.T) {
	l := Newlist()
	l.Add(10) // [0]
	l.Add(20)
	l.Add(90)
	l.Add(10) // [3]
	l.Add(10) // [4]
	l.Add(15)
	l.Add(10) // [6]

	var testingArray = [4]int64{0, 3, 4, 6}
	idsArray, ok := l.GetAllByValue(10)

	if !ok {
		t.Errorf("Error, GetAllByValues returned false")
	}

	for i := 0; i < 4; i++ {
		if testingArray[i] != idsArray[i] {
			t.Errorf("testing n real lists are not equal")
		}
	}

	l.RemoveAllByValue(10)
	if l.len != 3 {
		t.Errorf("Error, removing by value failed in length")
	}

	id20 := 0
	testingId20, _ := l.GetByValue(20)
	if id20 != int(testingId20) {
		t.Errorf("wrong id of 20 after removing")
	}

	val90 := 90
	testingVal90, _ := l.GetByIndex(1)
	if val90 != int(testingVal90) {
		t.Errorf("Eror in GetByIndex")
	}
}

func TestRemove(t *testing.T) {
	l := Newlist()
	l.Add(1)
	l.Add(2)

	var value int64 = 20
	for i := 0; i < 20; i++ {
		l.Add(value)
	}

	isAllRemovingFuncWorked := l.RemoveByIndex(l.len - 1)
	isAllRemovingFuncWorked = isAllRemovingFuncWorked && l.RemoveAllByValue(20)
	isAllRemovingFuncWorked = isAllRemovingFuncWorked && l.RemoveByValue(2)
	isAllRemovingFuncWorked = isAllRemovingFuncWorked && l.RemoveByIndex(0)

	if !isAllRemovingFuncWorked || l.fn != nil {
		t.Errorf("Removing failed")
	}
}

func TestClear(t *testing.T) {
	l := Newlist()
	l.Add(10)
	l.Clear()
	if l.fn != nil {
		t.Errorf("Clear is not working")
	}
}
