package list

import (
	"fmt"
)

type List struct {
	len int64
	fn  *node
}

func Newlist() *List {
	return &List{len: 0, fn: nil}
}

func (l *List) Clear() {
	l.fn = nil
}

func (l *List) Len() int64 {
	return l.len
}

func (l *List) GetAll() (values []int64, ok bool) {
	curr := l.fn

	if curr == nil {
		return
	}

	for curr != nil {
		values = append(values, curr.value)
		curr = curr.next
	}
	return values, true
}

func (l *List) GetAllByValue(value int64) (ids []int64, ok bool) {
	valuesArr, _ := l.GetAll()
	if valuesArr != nil {
		for key, val := range valuesArr {
			if val == value {
				ids = append(ids, int64(key))
				ok = true
			}
		}
	}
	return
}

func CreateListFromSlice(inputSlice []int64) *List {
	var myList *List = Newlist()
	for _, val := range inputSlice {
		myList.Add(val)
	}
	return myList
}

func (l *List) Print() {
	var curr = l.fn
	for curr != nil {
		fmt.Print(curr.value, "[", curr.index, "]", "->")
		curr = curr.next
	}
	fmt.Println("nil")
}

func (l *List) getNodeByIndex(id int64) *node {
	if id < 0 || id > l.len-1 {
		return nil
	}
	var curr = l.fn
	for curr != nil {
		if curr.index == id {
			return curr
		}
		curr = curr.next
	}
	return nil
}

func (l *List) GetByIndex(id int64) (value int64, ok bool) {

	curr := l.getNodeByIndex(id)
	if curr == nil {
		return 0, false
	}
	return curr.value, true
}

func (l *List) GetByValue(value int64) (id int64, ok bool) {
	curr := l.fn
	for curr != nil {
		if curr.value == value {
			return curr.index, true
		}
		curr = curr.next
	}
	return 0, false
}

func (l *List) updateIndexes() {
	curr := l.fn
	var id int64 = 0
	for curr != nil {
		curr.index = id
		id++
		curr = curr.next
	}
}

func (l *List) Add(value int64) {
	newNode := &node{value: value, next: nil, index: l.len}

	if l.fn == nil {
		l.len++
		l.fn = newNode
		return
	}

	curr := l.fn
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
	l.len++
}

func (l *List) RemoveByIndex(id int64) (ok bool) {
	switch {
	case id < 0 || id > l.len-1:
		return false
	case id == l.len-1:
		if id == 0 {
			l.fn = nil
			l.updateIndexes()
			return true
		} else {
			curr := l.getNodeByIndex(id - 1)
			curr.next = nil
			l.updateIndexes()
			return true
		}
	case id == 0:
		l.fn = l.getNodeByIndex(1)
		l.updateIndexes()
		return true

	default:
		l.getNodeByIndex(id - 1).next = l.getNodeByIndex(id + 1)
		l.updateIndexes()
		return true
	}
}

func (l *List) RemoveAllByValue(value int64) (ok bool) {
	_, isFinded := l.GetByValue(value)
	for isFinded {
		l.RemoveByValue(value)
		_, isFinded = l.GetByValue(value)
		ok = true
	}
	return
}

func (l *List) RemoveByValue(value int64) (ok bool) {
	curr := l.fn
	if curr == nil {
		return false
	}

	if curr.value == value {
		if curr.next == nil {
			l.fn = nil
			fmt.Println("List is empty")
			l.len--
			l.updateIndexes()
			return true
		}
		l.fn = curr.next
		l.len--
		l.updateIndexes()
		return true
	}

	for curr.next != nil {
		if curr.next.value == value {
			if curr.next.next != nil {
				curr.next = curr.next.next
				l.len--
				l.updateIndexes()
				return true
			} else {
				curr.next = nil
				l.len--
				l.updateIndexes()
				return true
			}

		} else {
			curr = curr.next
		}
	}
	return false
}

// Optional, not Tested

// func (l *List) ReveseList() {
// 	var newSlice = GetSLiceFromList(l)
// 	l = Newlist()
// 	var i = len(newSlice) - 1
// 	for ; i >= 0; i-- {
// 		l.Add(newSlice[i])
// 	}
// }

// func (l *List) GetReversedList() (output *List) {
// 	var newSlice = GetSLiceFromList(l)
// 	output = Newlist()
// 	var i = len(newSlice) - 1
// 	for ; i >= 0; i-- {
// 		output.Add(newSlice[i])
// 	}
// 	return output
// }
