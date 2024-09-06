package main

import "fmt"

type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
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

func main() {
	test_cases := []string{"()", "()[]{}", "(])", "([])", "", "([]", ")]"}
	for i, test_case := range test_cases {
		fmt.Println("*** TEST # ", i)
		fmt.Println("input: ", test_case)
		fmt.Println("output: ", isValid(test_case))
		fmt.Print(("\n"))
	}
}

func isValid(s string) bool {
	parens := NewStack[rune]()

	for _, c := range s {
		parens.Push(c)
	}

	for range 10 {
		r, has_value := parens.Pop()
		if has_value {
			fmt.Println(r)
			if r == '(' {
				fmt.Println("This is an open parenthesis!")
			}
		} else {
			fmt.Println("Stack is empty!")
		}
	}

	return false
}
