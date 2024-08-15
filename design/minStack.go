package design

type MinStack struct {
	stack    []int
	minStack []int
}

// 构造函数，初始化两个栈
func NewMinStack() *MinStack {
	return &MinStack{}
}

// Push 方法，添加元素到栈中，同时更新 minStack 以保存当前最小值
func (ms *MinStack) Push(number int) {
	ms.stack = append(ms.stack, number)
	if len(ms.minStack) == 0 {
		ms.minStack = append(ms.minStack, number)
	} else {
		min := ms.minStack[len(ms.minStack)-1]
		if number < min {
			ms.minStack = append(ms.minStack, number)
		} else {
			ms.minStack = append(ms.minStack, min)
		}
	}
}

// Pop 方法，移除并返回栈顶元素，同时更新 minStack
func (ms *MinStack) Pop() int {
	if len(ms.stack) == 0 {
		panic("Stack is empty")
	}
	ms.minStack = ms.minStack[:len(ms.minStack)-1]
	val := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]
	return val
}

// Min 方法，返回当前栈中的最小值
func (ms *MinStack) Min() int {
	if len(ms.minStack) == 0 {
		panic("Stack is empty")
	}
	return ms.minStack[len(ms.minStack)-1]
}
