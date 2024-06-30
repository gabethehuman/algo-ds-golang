package main

import (
	"fmt"
	"math"
)

type ResidualEdge struct {
	From int           // Source node.
	To   int           // Target node.
	Cap  int           // Capacity of the edge.
	Flow int           // Current flow through the edge.
	Rev  *ResidualEdge // Pointer to the reverse edge.
}

type ResidualGraph struct {
	Nodes         int // Number of nodes, the nodes are indexed from 1 to `Nodes`.
	AdjacencyList map[int][]ResidualEdge
}

// Construct a ResidualGraph based on a regular graph. It is to make the Edmonds-Karp algorithms more clear.
// New graph is built based on the weights of the original graph, the weights become the capacities.
func constructResidualGraph(g *Graph) *ResidualGraph {
	rg := &ResidualGraph{
		Nodes:         g.Nodes,
		AdjacencyList: make(map[int][]ResidualEdge),
	}

	// Initialize the adjacency list for each node.
	for i := 1; i <= g.Nodes; i++ {
		rg.AdjacencyList[i] = []ResidualEdge{}
	}

	// Create forward and reverse edges for each edge in the original graph.
	for _, edges := range g.AdjacencyList {
		for _, edge := range edges {
			// Create the forward edge
			forward := ResidualEdge{
				From: edge.From,
				To:   edge.To,
				Cap:  edge.Weight, // Assuming original weight is the capacity.
				Flow: 0,           // Initial flow is zero.
				Rev:  nil,         // Reverse edge, to be linked later.
			}
			// Create the reverse edge
			reverse := ResidualEdge{
				From: edge.To,
				To:   edge.From,
				Cap:  0, // Reverse edge initially has zero capacity.
				Flow: 0,
				Rev:  &forward,
			}
			forward.Rev = &reverse

			// Add to the adjacency lists.
			rg.AdjacencyList[edge.From] = append(rg.AdjacencyList[edge.From], forward)
			rg.AdjacencyList[edge.To] = append(rg.AdjacencyList[edge.To], reverse)
		}
	}

	return rg
}

// Floyd-Fulkerson method, Edmonds-Karp implementation.
// Augmenting paths are found using BFS. This implementation utilizes an auxiliary structures,
// ResidualEdge and ResidualGraph, which is a graph with reverse edges added.
// While there exists an augmenting path from source to sink it updates the flow along
// the path by the minimum residual capacity along this path.
func (g *Graph) EdmondsKarp(source int, sink int) int {
	if source < 1 || sink < 1 || source > g.Nodes || sink > g.Nodes {
		panic(fmt.Sprintf("EdmondsKarp: source and sink should be in range [1, %v], got %v", g.Nodes, source))
	}
	res := constructResidualGraph(g)

	flow := 0

	for {
		// Run BFS to find an augmenting path.
		queue := NewQueue()
		queue.Enqueue(source)               // A queue for BFS.
		path := make(map[int]*ResidualEdge) // A map to store the path taken in BFS.

		for i := 1; i <= g.Nodes; i++ {
			path[i] = nil
		}

		// A regular BFS loop.
		for queue.Length() > 0 && path[sink] == nil {
			curr := queue.Dequeue()
			for i := range res.AdjacencyList[curr] {
				edge := &res.AdjacencyList[curr][i]
				if path[edge.To] == nil && edge.To != source && edge.Cap > edge.Flow {
					path[edge.To] = edge
					queue.Enqueue(edge.To)
				}
			}
		}

		// If no augmenting path is found, we're done.
		if path[sink] == nil {
			break
		}

		// Find the minimum residual capacity along the path.
		minFlow := math.MaxInt32
		for v := sink; v != source; v = path[v].From { // Go backwards, from sink to source.
			residualCap := path[v].Cap - path[v].Flow
			if residualCap < minFlow {
				minFlow = residualCap
			}
		}

		// Update the flow along the path.
		for v := sink; v != source; v = path[v].From {
			edge := path[v]
			edge.Flow += minFlow     // Increase the flow of the forward edge by the minimum flow found.
			edge.Rev.Flow -= minFlow // Decrease the flow of the reverse edge by the same amount.
		}

		// Update the total flow by adding the flow pushed through this augmenting path.
		// The maximum flow is the sum of these augmenting path flows.
		flow += minFlow
	}

	return flow
}
