package container

type Stack[T any] struct {
	data []*T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		data: make([]*T, 0),
	}
}

// Push 将元素推入栈
func (t *Stack[T]) Push(e *T) {
	if t.data != nil {
		t.data = append(t.data, e)
	}
}

// Pop 从栈中弹出并返回元素
func (t *Stack[T]) Pop() *T {
	if !t.IsEmpty() {
		e := t.data[len(t.data)-1]
		t.data = t.data[:len(t.data)-1]
		return e

	}
	return nil
}

// Peek 返回栈顶元素但不弹出
func (t *Stack[T]) Peek() *T {
	if t.IsEmpty() {
		return nil
	}
	return t.data[len(t.data)-1]
}

// IsEmpty 检查栈是否为空
func (t *Stack[T]) IsEmpty() bool {
	return len(t.data) == 0
}

// Size 返回栈的大小
func (t *Stack[T]) Size() int {
	return len(t.data)
}
