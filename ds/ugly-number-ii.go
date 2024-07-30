package ds

import "container/heap"

/**
 * @param n: An integer
 * @return: return a  integer as description.
 */
func NthUglyNumber(n int) int {
	h := &IntHeap{1}
	heap.Init(h)
	visited := make(map[int]bool)
	var val int
	// for range n {
	for i := 0; i < n; i++ {
		val = heap.Pop(h).(int)
		for _, factor := range []int{2, 3, 5} {
			cand := val * factor
			if !visited[cand] {
				visited[cand] = true
				heap.Push(h, cand)
			}
		}
	}
	return val
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NthUglyNumber2(n int) int {
	uglys := []int{1}

	p2, p3, p5 := 0, 0, 0

	for i := 1; i < n; i++ {
		lastNumber := uglys[i-1]
		for uglys[p2]*2 <= lastNumber {
			p2++
		}
		for uglys[p3]*3 <= lastNumber {
			p3++
		}
		for uglys[p5]*5 <= lastNumber {
			p5++
		}

		nextUgly := min(uglys[p2]*2, uglys[p3]*3, uglys[p5]*5)
		uglys = append(uglys, nextUgly)
	}

	return uglys[n-1]
}
