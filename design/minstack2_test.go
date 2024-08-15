package design

import (
	"fmt"
	"testing"
)

func TestMinStack2(t *testing.T) {
	ms := NewMinStack2()
	ms.Push(3)
	ms.Push(5)
	fmt.Println(ms.Min()) // 输出 3
	ms.Push(2)
	ms.Push(1)
	fmt.Println(ms.Min()) // 输出 1
	ms.Pop()
	fmt.Println(ms.Min()) // 输出 2
	ms.Pop()
	fmt.Println(ms.Min()) // 输出 3
}
