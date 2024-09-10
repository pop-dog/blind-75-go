package graph

type Node[T any] struct {
	Value    T
	incoming []*Edge[T]
	Edges    []*Edge[T]
}

type Edge[T any] struct {
	parent *Node[T]
	Node   *Node[T]
	Weight int
}

type Graph[T any] struct {
	Nodes []*Node[T]
}

func New[T any]() *Graph[T] {
	return &Graph[T]{}
}

func NewFromAdjacency[T any](values []T, edges [][]int) *Graph[T] {
	// Create a fresh graph
	g := &Graph[T]{}
	if len(values) != len(edges) {
		return &Graph[T]{}
	}
	for _, v := range values {
		g.AddNode(v)
	}
	for i := range edges {
		for _, j := range edges[i] {
			if j < 0 || j >= len(g.Nodes) {
				return &Graph[T]{}
			}
			g.Nodes[i].CreateEdge(g.Nodes[j])
		}
	}
	return g
}

func (g *Graph[T]) AddNode(v T) *Node[T] {
	n := &Node[T]{Value: v}
	g.Nodes = append(g.Nodes, n)
	return n
}

func (n *Node[T]) createEdge(n2 *Node[T], w int) *Edge[T] {
	e := &Edge[T]{Weight: w, Node: n2, parent: n}
	n.Edges = append(n.Edges, e)
	n2.incoming = append(n2.incoming, e)
	return e
}

func (n *Node[T]) CreateEdge(n2 *Node[T]) *Edge[T] {
	return n.createEdge(n2, 1)
}

func (n *Node[T]) CreateWeightedEdge(n2 *Node[T], w int) *Edge[T] {
	return n.createEdge(n2, w)
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
		panic("attempt to delete edge that is not on node")
	}
	n.Edges = append(n.Edges[:index], n.Edges[index+1:]...)
}

func (n *Node[T]) DeleteEdge(e *Edge[T]) {
	n.deleteEdge(e)
}

func (g *Graph[T]) RemoveNode(n *Node[T]) T {
	if n != nil {
		// Remove the node from the graph
		index := -1
		for i, node := range g.Nodes {
			if node == n {
				index = i
				break
			}
		}
		if index == -1 {
			panic("removed node must be part of the graph")
		}
		g.Nodes = append(g.Nodes[:index], g.Nodes[index+1:]...)
		// Delete all edges that point to the node
		for _, e := range n.incoming {
			e.parent.deleteEdge(e)
		}
		// Delete all edges from the node
		for _, e := range n.Edges {
			n.deleteEdge(e)
		}
	}
	return n.Value
}
