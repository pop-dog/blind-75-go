package graph

type Node[T any] struct {
	Value    T
	incoming []*Edge[T]
	Edges    []*Edge[T]
	Graph    *Graph[T]
}

type Edge[T any] struct {
	parent *Node[T]
	Node   *Node[T]
	Weight int
}

type Graph[T any] struct {
	Nodes []*Node[T]
}

func (g *Graph[T]) New() *Graph[T] {
	return &Graph[T]{}
}

func (g *Graph[T]) AddNode(v T) *Node[T] {
	n := &Node[T]{Value: v, Graph: g}
	return n
}

func (n *Node[T]) CreateEdge(n2 *Node[T], w int) *Edge[T] {
	e := &Edge[T]{Weight: w, Node: n2, parent: n}
	n.Edges = append(n.Edges, e)
	return e
}

func (n *Node[T]) deleteEdge(e *Edge[T]) {
	if e == nil || e.parent != n {
		return
	}
	// Find index of edge on node
	index := -1
	for i := range n.Edges {
		if n.Edges[i] == e {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}
	if index == 0 {
		n.Edges = n.Edges[1:]
		return
	}
	if index == len(n.Edges)-1 {
		n.Edges = n.Edges[:index]
		return
	}
	lEdges := n.Edges[:index]
	rEdges := n.Edges[index+1:]
	n.Edges = append(lEdges, rEdges...)
}

func (n *Node[T]) DeleteEdge(e *Edge[T]) {

}

func (g *Graph[T]) RemoveNode(n *Node[T]) T {
	if n != nil && n.Graph == g {
		//
	}
	return n.Value
}
