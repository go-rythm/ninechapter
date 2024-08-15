package design

type Queue []int

type Stack struct {
	queue1 Queue
	queue2 Queue
}

// 构造函数，初始化两个队列
func NewStack() *Stack {
	return &Stack{}
}

// 将 queue1 中除最后一个元素外的其他元素全部移动到 queue2
func (s *Stack) moveItems() {
	for len(s.queue1) > 1 {
		s.queue2 = append(s.queue2, s.queue1[0])
		s.queue1 = s.queue1[1:]
	}
}

// 交换 queue1 和 queue2 的指针
func (s *Stack) swapQueues() {
	s.queue1, s.queue2 = s.queue2, s.queue1
}

// Push 方法，向栈中压入一个新元素
func (s *Stack) Push(value int) {
	s.queue1 = append(s.queue1, value)
}

// Top 方法，返回栈顶元素
func (s *Stack) Top() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	s.moveItems()
	item := s.queue1[0]
	s.queue1 = s.queue1[1:]
	s.swapQueues()
	s.queue1 = append(s.queue1, item)
	return item
}

// Pop 方法，移除并返回栈顶元素
func (s *Stack) Pop() {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	s.moveItems()
	s.queue1 = s.queue1[1:]
	s.swapQueues()
}

// IsEmpty 方法，检查栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.queue1) == 0
}
