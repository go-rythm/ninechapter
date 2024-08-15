package design

type MyQueue struct {
	stack1 []int
	stack2 []int
}

// 构造函数，初始化两个栈
func NewMyQueue() *MyQueue {
	return &MyQueue{}
}

// 将 stack2 的所有元素移动到 stack1 中
func (q *MyQueue) stack2ToStack1() {
	for len(q.stack2) > 0 {
		// 从 stack2 弹出元素并压入 stack1
		q.stack1 = append(q.stack1, q.stack2[len(q.stack2)-1])
		q.stack2 = q.stack2[:len(q.stack2)-1]
	}
}

// Push 方法，添加元素到队列末尾
func (q *MyQueue) Push(element int) {
	q.stack2 = append(q.stack2, element)
}

// Pop 方法，移除并返回队列头部元素
func (q *MyQueue) Pop() int {
	if len(q.stack1) == 0 {
		q.stack2ToStack1()
	}
	if len(q.stack1) == 0 {
		panic("Queue is empty")
	}
	val := q.stack1[len(q.stack1)-1]
	q.stack1 = q.stack1[:len(q.stack1)-1]
	return val
}

// Top 方法，返回队列头部元素
func (q *MyQueue) Top() int {
	if len(q.stack1) == 0 {
		q.stack2ToStack1()
	}
	if len(q.stack1) == 0 {
		panic("Queue is empty")
	}
	return q.stack1[len(q.stack1)-1]
}
