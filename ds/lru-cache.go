package ds

type linkedNode struct {
	key   int
	value int
	next  *linkedNode
}

type LRUCache struct {
	keyToPrev map[int]*linkedNode
	dummy     *linkedNode
	tail      *linkedNode
	capacity  int
	size      int
}

func NewLRUCache(capacity int) *LRUCache {
	dummy := &linkedNode{}
	c := &LRUCache{
		keyToPrev: make(map[int]*linkedNode),
		dummy:     dummy,
		tail:      dummy,
		capacity:  capacity,
	}
	return c
}

func (c *LRUCache) moveToTail(key int) {
	prev := c.keyToPrev[key]
	cur := prev.next

	if cur == c.tail {
		return
	}

	prev.next = prev.next.next
	c.tail.next = cur
	cur.next = nil

	if prev.next != nil {
		c.keyToPrev[prev.next.key] = prev
	}
	c.keyToPrev[cur.key] = c.tail

	c.tail = cur
}

func (c *LRUCache) Get(key int) int {
	if _, ok := c.keyToPrev[key]; !ok {
		return -1
	}
	c.moveToTail(key)
	return c.tail.value
}

func (c *LRUCache) Set(key, value int) {
	if c.Get(key) != -1 {
		c.tail.value = value
		return
	}

	if c.size < c.capacity {
		c.size++
		cur := &linkedNode{
			key:   key,
			value: value,
			next:  nil,
		}
		c.tail.next = cur
		c.keyToPrev[key] = c.tail
		c.tail = cur
		return
	}

	first := c.dummy.next
	delete(c.keyToPrev, first.key)

	first.key = key
	first.value = value
	c.keyToPrev[key] = c.dummy

	c.moveToTail(key)
}
