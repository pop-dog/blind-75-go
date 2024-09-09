package stack

type Stack[T any] []T

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(*s) == 0 {
		return zero, false
	}
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return value, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(*s) == 0 {
		return zero, false
	}
	index := len(*s) - 1
	value := (*s)[index]
	return value, true
}

func (s *Stack[T]) Clear() {
	*s = []T{}
}

func (s *Stack[T]) Len() int {
	return len(*s)
}
