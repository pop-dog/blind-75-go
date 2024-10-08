package main

import (
	"fmt"
	"strings"

	"github.com/pop-dog/blind-75-go/dstructs/stack"
)

func main() {
	test_cases := []string{"()", "()[]{}", "(])", "([])", "", "([]", ")]"}
	for i, test_case := range test_cases {
		fmt.Println("*** TEST # ", i)
		fmt.Println("input: ", test_case)
		fmt.Println("output: ", isValid(test_case))
		fmt.Print(("\n"))
	}
}

func isParenthesis(r rune) bool {
	return strings.ContainsRune("()[]{}", r)
}

func isOpenParenthesis(r rune) bool {
	return strings.ContainsRune("([{", r)
}

func isCloseParenthesis(r rune) bool {
	return strings.ContainsRune(")]}", r)
}

func isValidPair(l rune, r rune) bool {
	return l == '(' && r == ')' ||
		l == '[' && r == ']' ||
		l == '{' && r == '}'
}

func isValid(s string) bool {
	parens := stack.New[rune]()

	for _, r := range s {
		// Ignore anything that is not a parenthesis
		if !isParenthesis(r) {
			continue
		}
		// 1. If we have an "open" parenthesis, push to the stack
		if isOpenParenthesis(r) {
			parens.Push(r)
			continue
		}
		// 2. If we have a "closed" parenthesis, pop from stack and compare
		if isCloseParenthesis(r) {
			if parens.IsEmpty() {
				// There is no open paren! Invalid!
				return false
			}
			l := parens.Pop()
			if !isValidPair(l, r) {
				// Parentheses do not match! INVALID!
				return false
			}
		}
	}

	// Final check: Do we hae any leftover "open" parentheses? If so, we have orphans :(
	return parens.IsEmpty() // If empty, we GOOD :)
}
