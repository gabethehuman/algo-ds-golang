package main

import (
	"fmt"
	"sort"
)

// Find a minimum spanning tree for an undirected graph with weighted edges.
func (g *Graph) KruskalMST() []Edge {
	if g.Directed {
		panic("KruskalMST: cannot be applied to directed graphs.")
	}

	// Gather all unique edges by using a map, and representing edges as `from-to` strings,
	// where `from`` is always smaller than `to``.
	uniqueEdges := map[string]Edge{}
	for from, edgeList := range g.AdjacencyList {
		for _, edge := range edgeList {
			// Ensure from is less than to to avoid duplicates.
			from, to := from, edge.To
			if from > to {
				from, to = to, from
			}
			key := fmt.Sprintf("%d-%d", from, to)
			if _, exists := uniqueEdges[key]; !exists {
				uniqueEdges[key] = Edge{From: from, To: to, Weight: edge.Weight}
			}
		}
	}

	// Extract values from the uniqueEdges map to a slice,
	// these are all the edges in the graph without duplicates.
	allEdges := make([]Edge, 0, len(uniqueEdges))
	for _, edge := range uniqueEdges {
		allEdges = append(allEdges, edge)
	}

	// Sort all the edges by weight.
	sort.Slice(allEdges, func(i, j int) bool {
		return allEdges[i].Weight < allEdges[j].Weight
	})

	uf := NewUnionFind(g.Nodes) // For now each node in the graph is its own disjoint set.

	// Process each edge starting with the smallest weight.
	// If the edge connects two distinct components, add it to the MST.
	// If we were to connect two edges from the same component, this would create a cycle,
	// which we don't want.
	minSpanTree := []Edge{}
	cost := 0
	for _, edge := range allEdges {
		// Adjust node indices for Union-Find, grahp enumerates nodes from 1 to N, UnionFind from 0 to N-1.
		from, to := edge.From-1, edge.To-1
		if uf.Find(from) != uf.Find(to) {
			minSpanTree = append(minSpanTree, edge) // Append original Edge, so no off by 1 index issues.
			cost += edge.Weight
			uf.Union(from, to)
		}
	}
	return minSpanTree
}
