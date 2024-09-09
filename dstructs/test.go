package main

import (
	"fmt"

	"github.com/pop-dog/blind-75-go/dstructs/queue"
)

func main() {
	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.Dequeue())
	q.Clear()
	q.Enqueue(3)
	q.Enqueue(5)
	q.Enqueue(2)
	for !q.IsEmpty() {
		fmt.Println(q.Dequeue())
	}
}
