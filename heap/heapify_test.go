package heap

import "testing"

func TestHeapify(t *testing.T) {
	arr := []int{3, 2, 1, 5, 4, 6}
	Heapify(arr)
	t.Log(arr)
}
