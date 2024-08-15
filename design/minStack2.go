package design

type MinStack2 struct {
	stack     []int
	MinStack2 []int
}

// 构造函数，初始化两个栈
func NewMinStack2() *MinStack2 {
	return &MinStack2{}
}

// Push 方法，添加元素到栈中，同时更新 MinStack2 以保存当前最小值
func (ms *MinStack2) Push(number int) {
	ms.stack = append(ms.stack, number)
	// 仅在 MinStack2 为空或新值小于等于当前最小值时推入 MinStack2
	if len(ms.MinStack2) == 0 || number <= ms.MinStack2[len(ms.MinStack2)-1] {
		ms.MinStack2 = append(ms.MinStack2, number)
	}
}

// Pop 方法，移除并返回栈顶元素，同时更新 MinStack2
func (ms *MinStack2) Pop() int {
	if len(ms.stack) == 0 {
		panic("Stack is empty")
	}
	val := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]
	// 仅当栈顶元素等于 MinStack2 栈顶元素时，移除 MinStack2 栈顶
	if val == ms.MinStack2[len(ms.MinStack2)-1] {
		ms.MinStack2 = ms.MinStack2[:len(ms.MinStack2)-1]
	}
	return val
}

// Min 方法，返回当前栈中的最小值
func (ms *MinStack2) Min() int {
	if len(ms.MinStack2) == 0 {
		panic("Stack is empty")
	}
	return ms.MinStack2[len(ms.MinStack2)-1]
}
