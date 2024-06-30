package main

import "fmt"

func (g *Graph) inDegree() map[int]int {
	if !g.Directed {
		panic("inDegree: cannot be applied to undirected graphs.")
	}

	// Initialize node degree map.
	inDegree := make(map[int]int)
	for n := range g.Nodes {
		inDegree[n+1] = 0 // n+1 because we enumaret graph nodes from 1
	}

	// Count degrees of all nodes.
	for _, edgeList := range g.AdjacencyList {
		for _, edge := range edgeList {
			inDegree[edge.To] += 1
		}
	}

	return inDegree
}

// Remove a specified Edge element from a slice.
// This implementation seems very inefficient, but is not the point of the exercise.
// Because golang doesn't have a nice remove(idx) function, it's easier to iterate over
// all elements and append them selectively, than to remove them via index and maneuvers
// such as s = append(s[:idx], s[idx+1:]...). The latter looks weird, and there's more edge cases,
// such as when len(s)==0 or len(s)==1, etc.
func remove(s []Edge, e Edge) []Edge {
	// Create a new slice to hold the result.
	var result []Edge
	for _, item := range s {
		// Check if the current item is the one to be removed.
		if item != e {
			// If not, add it to the result slice.
			result = append(result, item)
		}
	}
	// Return the result slice.
	return result
}

// Topological ordering of a graph using Kahn's algorithm. Only possible for directed graphs.
// This function works with a copy of the original graph, as it removes edges during the procedure.
// If it's not possible to topologically sort a graph, because it has cycles, return an error.
func (g Graph) KahnTopoSort() ([]int, error) {
	if !g.Directed {
		panic("KahnTopoSort: cannot be applied to undirected graphs.")
	}

	result := []int{}
	nodesToProcess := NewQueue()
	inDegree := g.inDegree() // Map that counts incoming nodes for each node.

	// Gather all nodes with 0 incoming edges, so called sources of the graph.
	for node, degree := range inDegree {
		if degree == 0 {
			nodesToProcess.Enqueue(node)
		}
	}

	// Get a node from the queue and append it to the result slice.
	// For this node, remove all outgoing edges. If after removing such an edge,
	// the destination node has no more incoming edges, enqueue it and proceed.
	for nodesToProcess.Length() > 0 {
		node := nodesToProcess.Dequeue()
		result = append(result, node)

		// Gather all outgoing edges from the current node and save them to another variable.
		// This is because we will be modifying g.AdjacencyList[node] later, so it seems cleaner
		// iterate over edgesToConsider than over g.AdjacencyList[node] and modify it in the same loop.
		edgesToConsider := []Edge{}
		edgesToConsider = append(edgesToConsider, g.AdjacencyList[node]...)

		for _, e := range edgesToConsider {
			g.AdjacencyList[node] = remove(g.AdjacencyList[node], e)
			inDegree[e.To] -= 1
			if inDegree[e.To] == 0 {
				nodesToProcess.Enqueue(e.To)
			}
		}
	}

	// At the end, if there's no more edges left, the sort is complete.
	// If there are edges left, it means there is a cycle in the graph, and it cannot be
	// topologically sorted, return an error.
	for _, adj := range g.AdjacencyList {
		if len(adj) > 0 {
			return nil, fmt.Errorf("KahnTopoSort: topological sort failed, cycle detected")
		}
	}
	return result, nil
}
