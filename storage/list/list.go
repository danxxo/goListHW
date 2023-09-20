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

func (l *List) GetAll() (result []int64, ok bool) {
	cur := l.fn

	if cur == nil {
		return
	}

	for cur != nil {
		result = append(result, cur.value)
		cur = cur.next
	}
	return result, true
}

func (l *List) GetAllByValue(value int64) (ids []int64, ok bool) {
	arr, _ := l.GetAll()
	if arr != nil {
		for key, val := range arr {
			if val == value {
				ids = append(ids, int64(key))
				ok = true
			}
		}
	}
	return
}

func CreateListFromSlice(in_slice []int64) *List {
	var myList *List = Newlist()
	for _, val := range in_slice {
		myList.Add(val)
	}
	return myList
}

func (l *List) Print() {
	var cur = l.fn
	for cur != nil {
		fmt.Print(cur.value, "[", cur.index, "]", "->")
		cur = cur.next
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
	var index int64 = 0
	for curr != nil {
		curr.index = index
		index++
		curr = curr.next
	}
}

func (l *List) Add(data int64) {
	newNode := &node{value: data, next: nil, index: l.len}

	if l.fn == nil {
		l.len++
		l.fn = newNode
		return
	}

	cur := l.fn
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = newNode
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
	_, exist := l.GetByValue(value)
	for exist {
		l.RemoveByValue(value)
		_, exist = l.GetByValue(value)
		ok = true
	}
	return
}

func (l *List) RemoveByValue(nodeData int64) (ok bool) {
	cur := l.fn
	if cur == nil {
		return false
	}

	if cur.value == nodeData {
		if cur.next == nil {
			l.fn = nil
			fmt.Println("List is empty")
			l.len--
			l.updateIndexes()
			return true
		}
		l.fn = cur.next
		l.len--
		l.updateIndexes()
		return true
	}

	for cur.next != nil {

		if cur.next.value == nodeData {
			if cur.next.next != nil {
				cur.next = cur.next.next
				l.len--
				l.updateIndexes()
				return true
			} else {
				cur.next = nil
				l.len--
				l.updateIndexes()
				return true
			}

		} else {
			cur = cur.next
		}
	}
	return false
}

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
