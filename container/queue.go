package container

type Queue[T any] struct {
	data []*T
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		data: make([]*T, 0),
	}
}

// Enqueue 入队
func (q *Queue[T]) Enqueue(item *T) {
	q.data = append(q.data, item)
}

// Dequeue 出队
func (q *Queue[T]) Dequeue() *T {
	if q.IsEmpty() {
		return nil
	}
	item := q.data[0]
	q.data = q.data[1:]
	return item
}

// IsEmpty 队列是否为空
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// Size 队列大小
func (q *Queue[T]) Size() int {
	return len(q.data)
}
