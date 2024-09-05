/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// Keep track of nodes we have made copies of
	copies := make(map[int]*Node)

	// Keep track of nodes we have visited
	visited := make(map[int]bool)

	// BFS: Search node, then search its neighbors
	queue := []*Node{node}

	getCopy := func(n *Node) *Node {
		var cpy, exists = copies[n.Val]
		if exists {
			return cpy
		}
		c := Node{Val: n.Val}
		copies[n.Val] = &c
		return copies[n.Val]
	}

	for len(queue) > 0 {
		// dequeue
		n := queue[0]
		queue = queue[1:]

		_, been_visited := visited[n.Val]
		if been_visited {
			continue
		}

		visited[n.Val] = true

		cpy := getCopy(n)

		for _, neighbor := range n.Neighbors {
			cpy.Neighbors = append(cpy.Neighbors, getCopy(neighbor))
			queue = append(queue, neighbor)
		}
	}
	return getCopy(node)
}
