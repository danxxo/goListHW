package mp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMap(t *testing.T) {
	myMap := NewMap()
	assert.Equal(t, int64(0), myMap.Len())
	assert.IsType(t, map[int64]int64{}, myMap.HashTable)
	assert.Equal(t, 0, len(myMap.HashTable))
	myMap.Clear()
	assert.Equal(t, myMap.HashTable, map[int64]int64{})
	assert.Equal(t, myMap.Len(), int64(0))
}

func TestGetAll(t *testing.T) {
	myMap := NewMap()
	var i int64 = 0
	expected := []int64{}
	for ; i < 100; i++ {
		myMap.Add(i)
		expected = append(expected, i)
	}
	given, ok := myMap.GetAll()
	assert.Equal(t, given, expected)
	assert.Equal(t, ok, true)
	myMap.Add(10)
	myMap.Add(10)
	expected = []int64{10, 100, 101}
	given, ok = myMap.GetAllByValue(10)
	assert.Equal(t, given, expected)
	assert.Equal(t, ok, true)
	myMap.Clear()
	_, okGetAll := myMap.GetAll()
	_, okGetAllByValue := myMap.GetAllByValue(100)
	assert.Equal(t, okGetAll, false)
	assert.Equal(t, okGetAllByValue, false)
}

func TestGetByIndex(t *testing.T) {
	myMap := NewMap()
	myMap.Add(1)
	expectedID, ok := myMap.GetByIndex(0)
	assert.Equal(t, expectedID, int64(1))
	assert.Equal(t, ok, true)
	for i := 0; i < 10; i++ {
		myMap.Add(10)
	}
	myMap.Add(200)
	expectedID, ok = myMap.GetByIndex(11)
	assert.Equal(t, expectedID, int64(200))
	assert.Equal(t, ok, true)

	myMap.Clear()
	expectedID, ok = myMap.GetByIndex(10)
	assert.Equal(t, expectedID, int64(0))
	assert.Equal(t, ok, false)
}

func TestRemoveByValue(t *testing.T) {
	myMap := NewMap()
	var i int64 = 0
	for ; i < 10; i++ {
		myMap.Add(i)
	}
	myMap.Add(0)

	ok := myMap.RemoveByValue(0)
	assert.Equal(t, myMap.HashTable[0], int64(1))
	assert.Equal(t, ok, true)
	myMap.RemoveAllByValue(0)
	_, ok = myMap.GetByValue(0)
	assert.Equal(t, ok, false)
}

func TestRemoveByIndex(t *testing.T) {
	myMap := NewMap()
	var i int64 = 0
	for ; i < 10; i++ {
		myMap.Add(i)
	}
	myMap.RemoveByIndex(0)
	myMap.RemoveByIndex(myMap.len - 1)
	myMap.RemoveByIndex(4)
	expect, _ := myMap.GetAll()
	assert.Equal(t, expect, []int64{1, 2, 3, 4, 6, 7, 8})
}
