package mp

type mp struct {
	hashTable map[int64]int64
}

func NewMap() *mp {
	hashTable := make(map[int64]int64, 0)
	return &mp{hashTable: hashTable}
}

func (myMap *mp) Clear() {}

func (myMap *mp) Len() int {
	return len(myMap.hashTable)
}

func (myMap *mp) GetAll() (result []int64, ok bool) {}

func (myMap *mp) GetAllByValue(value int64) (ids []int64, ok bool) {}

func (myMap *mp) Print() {}

func (myMap *mp) GetByIndex(id int64) (value int64, ok bool) {}

func (myMap *mp) GetByValue(value int64) (id int64, ok bool) {}

func (myMap *mp) updateIndexes() {}

func (myMap *mp) Add(data int64) {}

func (myMap *mp) RemoveByIndex(id int64) (ok bool) {}

func (myMap *mp) RemoveAllByValue(value int64) (ok bool) {}

func (myMap *mp) RemoveByValue(nodeData int64) (ok bool) {}
