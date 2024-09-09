package queue

type Queue[T any] []T

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(value T) {
	*q = append(*q, value)
}

func (q *Queue[T]) Dequeue() T {
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

func (q *Queue[T]) Peek() T {
	return (*q)[0]
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
