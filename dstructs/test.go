package main

import (
	"fmt"

	"github.com/pop-dog/blind-75-go/dstructs/graph"
	"github.com/pop-dog/blind-75-go/dstructs/queue"
)

func main() {
	testGraph()
	testQueue()
}

func testGraph() {
	g := graph.NewFromAdjacency(
		[]int{1, 2, 3, 4},
		[][]int{
			{1, 3},
			{0, 2},
			{1, 3},
			{0, 2},
		},
	)
	printGraph(g)
	g.RemoveNode(g.Nodes[2])
	printGraph(g)
}

func printGraph(g *graph.Graph[int]) {
	fmt.Println("Graph g has ", len(g.Nodes), " nodes")
	for i, n := range g.Nodes {
		if i > 0 {
			fmt.Print("\n")
		}
		fmt.Print("[ ", n.Value, " -> ")
		for j, e := range n.Edges {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Print(e.Node.Value)
		}
		fmt.Print(" ]")
	}
	fmt.Print("\n")
}

func testQueue() {
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
