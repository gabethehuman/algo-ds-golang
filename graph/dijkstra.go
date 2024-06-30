package main

import (
	"fmt"
	"math"
)

// Dijkstra's algorithm using min heap priority queue.
// Calculate minimum distance from a source node to every other node.
// Return a map of shortest distances to each node, and also a map of predecessor nodes
// on the shortest path from the source.
func (g *Graph) Dijkstra(source int) (map[int]int, map[int]int) {
	if source < 1 || source > g.Nodes {
		panic(fmt.Sprintf("Dijkstra: source should be in range [1, %v], got %v", g.Nodes, source))
	}

	// Initialize priorities to 'Inf'.
	priorities := make([]int, g.Nodes)
	for i := range priorities {
		priorities[i] = math.MaxInt
	}

	// Priority of the source is 0 (priority is like a distance to the node).
	priorities[source-1] = 0

	// Gather all nodes to a list.
	values := make([]int, g.Nodes)
	for i := range g.Nodes {
		values[i] = i + 1 // Shifted by 1, because `range` goes from 0, but we enumerate nodes from 1.
	}

	// Initialize priority queue using values (node identifiers) and priorities.
	prioQueue := NewHeap(priorities, values)

	prev := make(map[int]int) // For a predecessor of each node.
	dist := make(map[int]int) // For distances to each node.

	// Initialize `dist` and `prev` with 0 for source, undefined/int for other nodes.
	for i := range g.Nodes {
		if i+1 != source {
			prev[i+1] = -1
			dist[i+1] = math.MaxInt
		} else {
			prev[i+1] = 0
			dist[i+1] = 0
		}
	}

	// The steps above are not that important and might seem convoluted because of the indices.
	// What's important is that at this point in the function we have:
	// * `prioQueue` initialized with values [1, ..., g.Nodes] and priorities:
	//               0 for source, Inf for all the other nodes.
	// * `dist` with keys [1, ..., g.Nodes] and corresponding values: 0 for source, Inf otherwise.
	// * `prev` with the same keys and corresponding values: 0 for source, -1 otherwise.

	// Main loop of Dijkstra's algorithm.
	for prioQueue.Len() > 0 {
		_, currNode, _ := prioQueue.PopMin()             // Extract the best node.
		for _, edge := range g.AdjacencyList[currNode] { // Go through all neighbors of currNode.
			alt := dist[currNode] + edge.Weight // Calculate the alternative path distance.
			// If the alt path is shorter than the previously known shortest path to `edge.To`, update the path.
			if alt < dist[edge.To] {
				prev[edge.To] = currNode //  Update the predecessor of `edge.To` to be `currNode`.
				dist[edge.To] = alt      // Update the shortest distance to `edge.To`.
				prioQueue.DecreasePrio(edge.To, alt)
			}
		}
	}

	return dist, prev
}
