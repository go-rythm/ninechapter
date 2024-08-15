package design

import (
	"testing"
)

func TestMyQueue(t *testing.T) {
	queue := NewMyQueue()
	queue.Push(1)
	queue.Push(2)
	println(queue.Top()) // 输出 1
	println(queue.Pop()) // 输出 1
	println(queue.Pop()) // 输出 2
}
