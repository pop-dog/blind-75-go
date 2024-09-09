package main

import (
	"fmt"
	"strconv"

	"github.com/pop-dog/blind-75-go/dstructs/list"
)

func main() {
	l1 := list.New[int]()
	l2 := list.New[int]()

	for i := range 10 {
		l1.PushBack(i)
		l2.PushBack(i)
	}

	l1.PushFront(12)
	l1.MoveAfter(l1.Front(), l1.Back())

	mergedList := mergeTwoLists(l1, l2)
	fmt.Println("Result: ", listToString(mergedList))
	fmt.Println("Total length: ", mergedList.Len())

	mergedList.PushBackList(l1)
	fmt.Println("Pushed l1 to back of merged: ", listToString(mergedList))

	mergedList.Remove(mergedList.Front())
	mergedList.Remove(mergedList.Back())

	fmt.Println("mergedList: ", listToString(mergedList))

	doubled := mergedList.Map(func(el *list.Element[int], args ...any) int {
		return el.Value * 2
	})
	fmt.Println("doubled: ", listToString(doubled))
}

func mergeTwoLists(l1 *list.List[int], l2 *list.List[int]) *list.List[int] {
	// Allocate a list to hold the merged lists
	nm := list.New[int]()

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
		if p2.Value < p1.Value {
			nm.PushBack(p2.Value)
			p2 = p2.Next()
			continue
		}
		nm.PushBack(p1.Value)
		p1 = p1.Next()
	}

	return nm
}

func listToString(list *list.List[int]) string {
	el := list.Front()
	var printVal string
	for el != nil {
		if el != list.Front() {
			printVal += ", "
		}
		printVal += strconv.Itoa(el.Value)
		el = el.Next()
	}
	return printVal
}
