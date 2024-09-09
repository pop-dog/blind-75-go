package queue

type Queue[T any] []T

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(value T) {
	*q = append(*q, value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if len(*q) == 0 {
		return zero, false
	}
	value := (*q)[0]
	*q = (*q)[1:]
	return value, true
}

func (q *Queue[T]) Peek() (T, bool) {
	var zero T
	if len(*q) == 0 {
		return zero, false
	}
	return (*q)[0], true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Clear() {
	*q = []T{}
}

func (q *Queue[T]) Len() int {
	return len(*q)
}
