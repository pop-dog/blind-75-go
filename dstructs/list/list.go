package list

import "errors"

type Element[T any] struct {
	Value T
	next  *Element[T]
	prev  *Element[T]
}

type List[T any] struct {
	front *Element[T]
	back  *Element[T]
	count int
}

func (el *Element[T]) Next() *Element[T] {
	return el.next
}

func (el *Element[T]) Prev() *Element[T] {
	return el.prev
}

func New[T any]() *List[T] {
	return &List[T]{count: 0}
}

func (l *List[T]) Back() *Element[T] {
	return l.back
}

func (l *List[T]) Front() *Element[T] {
	return l.front
}

func (l *List[T]) Init() *List[T] {
	return New[T]()
}

// Validate that mark is an element of the list
func (l *List[T]) contains(mark *Element[T]) bool {
	if mark == nil {
		return false
	}
	for el := l.front; el != nil; el = el.next {
		if el == mark {
			return true // Found it!!
		}
	}
	// We never found the element; it's not in the list!
	return false
}

// Shared internal function to insert a value after an element
func (l *List[T]) insertAfter(v T, mark *Element[T]) *Element[T] {
	// Allocate a new element and assign it the value
	newNode := &Element[T]{}
	newNode.Value = v

	oldNext := mark.next
	// Insert newNode between mark and its old "next"
	mark.next = newNode
	newNode.prev = mark
	newNode.next = oldNext
	if oldNext == nil {
		// mark was the last element
		l.back = newNode
	} else {
		oldNext.prev = newNode
	}
	l.count++
	return newNode
}

func (l *List[T]) InsertAfter(v T, mark *Element[T]) *Element[T] {
	if !l.contains(mark) {
		return nil
	}
	newNode := l.insertAfter(v, mark)
	return newNode
}

// Shared internal function to insert a value before an element
func (l *List[T]) insertBefore(v T, mark *Element[T]) *Element[T] {
	// Allocate a new element and assign it the value
	newNode := &Element[T]{}
	newNode.Value = v

	oldPrev := mark.prev
	// Insert newNode between mark and its old "prev"
	mark.prev = newNode
	newNode.next = mark
	if oldPrev == nil {
		// newNode is now the front of the list
		l.front = newNode
	} else {
		oldPrev.next = newNode
	}
	l.count++
	return newNode
}

func (l *List[T]) InsertBefore(v T, mark *Element[T]) *Element[T] {
	if !l.contains(mark) {
		return nil
	}
	newNode := l.insertBefore(v, mark)
	return newNode
}

func (l *List[T]) Len() int {
	return l.count
}

func (l *List[T]) Remove(e *Element[T]) (T, error) {
	var zero T
	if !l.contains(e) {
		return zero, errors.New("Element is not in list")
	}
	l.remove(e)
	l.count--
	return e.Value, nil
}

// Shared internal function to remove an element
func (l *List[T]) remove(e *Element[T]) {
	oldPrev := e.prev
	oldNext := e.next

	// check if the front or back pointers need to update
	if oldPrev == nil {
		l.front = e.next
	} else {
		oldPrev.next = oldNext
	}
	if oldNext == nil {
		l.back = e.prev
	} else {
		oldNext.prev = oldPrev
	}
	l.count--
}

func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	if !l.contains(e) || !l.contains(mark) || e == mark {
		return
	}
	// Remove the element from its current position, then re-insert after target
	l.remove(e)
	l.insertAfter(e.Value, mark)
}

func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	if !l.contains(e) || !l.contains(mark) || e == mark {
		return
	}
	// Remove the element from its current position, then re-insert before target
	l.remove(e)
	l.insertBefore(e.Value, mark)
}

func (l *List[T]) MoveToBack(e *Element[T]) {
	// Alias for MoveAfter(list.back)
	l.MoveAfter(e, l.back)
}

func (l *List[T]) MoveToFront(e *Element[T]) {
	// Alias for MoveBefore(list.front)
	l.MoveBefore(e, l.front)
}

func (l *List[T]) init(v T) *Element[T] {
	// Used if the list currently has no elements.
	// Create an element with the given value and make it the front/back of the list
	el := &Element[T]{Value: v}
	l.front = el
	l.back = el
	l.count = 1
	return el
}

func (l *List[T]) PushBack(v T) *Element[T] {
	if l.back == nil {
		return l.init(v)
	} else {
		return l.insertAfter(v, l.back)
	}
}

func (l *List[T]) PushBackList(other *List[T]) {
	if l == nil || other == nil {
		return
	}
	// Go through other and append each el to the back of l
	for el := other.front; el != nil; el = el.next {
		l.PushBack(el.Value)
	}
}

func (l *List[T]) PushFront(v T) *Element[T] {
	if l.front == nil {
		return l.init(v)
	} else {
		return l.insertBefore(v, l.front)
	}
}

func (l *List[T]) PushFrontList(other *List[T]) {
	if l == nil || other == nil {
		return
	}
	// Go through other and prepend each el to the front of l
	for el := other.back; el != nil; el = el.prev {
		l.PushFront(el.Value)
	}
}
