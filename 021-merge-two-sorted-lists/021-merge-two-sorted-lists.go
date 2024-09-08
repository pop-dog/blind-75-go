package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	l1 := list.New()
	l2 := list.New()

	for i := range 5 {
		l1.PushBack(i)
		l2.PushBack(i)
	}

	fmt.Println("Result: ", listToString(mergeTwoLists(l1, l2)))
}

func mergeTwoLists(l1 *list.List, l2 *list.List) *list.List {
	// Allocate a list to hold the merged lists
	nm := list.New()

	// Get pointers to the front of the lists
	p1 := l1.Front()
	p2 := l2.Front()

	// Iterate using two points until both lists are exhausted
	for p1 != nil || p2 != nil {
		// If either pointer reached the end
		// just add the next el from the other list
		if p1 == nil {
			nm.PushBack(p2.Value)
			p2 = p2.Next()
			continue
		}
		if p2 == nil {
			nm.PushBack(p1.Value)
			p1 = p1.Next()
			continue
		}

		// Both lists have a value, use the smaller one
		if p2.Value.(int) < p1.Value.(int) {
			nm.PushBack(p2.Value)
			p2 = p2.Next()
			continue
		}
		nm.PushBack(p1.Value)
		p1 = p1.Next()
	}

	return nm
}

func listToString(list *list.List) string {
	el := list.Front()
	var printVal string
	for el != nil {
		if el != list.Front() {
			printVal += ", "
		}
		printVal += strconv.Itoa(el.Value.(int))
		el = el.Next()
	}
	return printVal
}
