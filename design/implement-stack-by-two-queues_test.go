package design

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Top()) // 输出 2
	stack.Pop()
	fmt.Println(stack.Top())     // 输出 1
	fmt.Println(stack.IsEmpty()) // 输出 false
	stack.Pop()
	fmt.Println(stack.IsEmpty()) // 输出 true
}
