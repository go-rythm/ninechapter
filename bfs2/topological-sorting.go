package bfs2

// https://www.lintcode.com/problem/127/

type DirectedGraphNode struct {
	Label     int
	Neighbors []*DirectedGraphNode
}

/**
 * @param graph: A list of Directed graph node
 * @return: Any topological order for the given graph.
 */
func TopSort(graph []*DirectedGraphNode) []*DirectedGraphNode {
	order := make([]*DirectedGraphNode, 0, len(graph))
	indegree := make(map[*DirectedGraphNode]int, len(graph))
	for _, node := range graph {
		for _, neighbor := range node.Neighbors {
			indegree[neighbor]++
		}
	}

	queue := []*DirectedGraphNode{}
	for _, node := range graph {
		if _, ok := indegree[node]; !ok {
			queue = append(queue, node)
			order = append(order, node)
		}
	}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		for _, neighbor := range head.Neighbors {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
				order = append(order, neighbor)
			}
		}
	}

	return order
}
