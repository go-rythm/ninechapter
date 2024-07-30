package ds

type listNode struct {
	num  int
	next *listNode
}

type DataStream struct {
	dummy, tail *listNode
	num2Prev    map[int]*listNode
	duplicates  map[int]bool
}

func NewDataStream() *DataStream {
	node := &listNode{}
	ds := &DataStream{
		dummy:      node,
		tail:       node,
		num2Prev:   make(map[int]*listNode),
		duplicates: make(map[int]bool),
	}
	return ds
}

func (ds *DataStream) Add(num int) {
	if _, ok := ds.duplicates[num]; ok {
		return
	}

	if _, ok := ds.num2Prev[num]; !ok {
		ds.push(num)
		return
	}

	ds.duplicates[num] = true
	ds.remove(num)
}

func (ds *DataStream) push(num int) {
	ds.num2Prev[num] = ds.tail
	newNode := &listNode{
		num: num,
	}
	ds.tail.next = newNode
	ds.tail = newNode
}

func (ds *DataStream) remove(num int) {
	prev := ds.num2Prev[num]
	delete(ds.num2Prev, num)
	prev.next = prev.next.next

	if prev.next != nil {
		ds.num2Prev[prev.next.num] = prev
	} else {
		ds.tail = prev
	}
}

func (ds *DataStream) FirstUnique() int {
	if ds.dummy.next != nil {
		return ds.dummy.next.num
	}
	return -1
}
