package main

import (
	"fmt"
	"testing"
)

// Test KruskalMST on a few connected, undirected graphs with weighted edges.
func TestKruskalMST(t *testing.T) {
	tests := []struct {
		name        string
		directed    bool
		numNodes    int
		edges       [][3]int // [from, to, weight]
		expectedMST []Edge
	}{
		{
			name:     "Simple undirected graph",
			directed: false,
			numNodes: 4,
			edges: [][3]int{
				{1, 2, 1},
				{1, 3, 4},
				{2, 3, 2},
				{2, 4, 5},
				{3, 4, 3},
			},
			expectedMST: []Edge{
				{From: 1, To: 2, Weight: 1},
				{From: 2, To: 3, Weight: 2},
				{From: 3, To: 4, Weight: 3},
			},
		},
		{
			name:     "More complex",
			directed: false,
			numNodes: 5,
			edges: [][3]int{
				{1, 2, 1},
				{1, 3, 2},
				{1, 4, 3},
				{2, 3, 4},
				{2, 5, 5},
				{3, 4, 6},
				{4, 5, 7},
			},
			expectedMST: []Edge{
				{From: 1, To: 2, Weight: 1},
				{From: 1, To: 3, Weight: 2},
				{From: 1, To: 4, Weight: 3},
				{From: 2, To: 5, Weight: 5},
			},
		},
		{
			name:     "Most complex",
			directed: false,
			numNodes: 9,
			edges: [][3]int{
				{1, 2, 10},
				{1, 3, 9},
				{1, 4, 6},
				{1, 5, 12},
				{2, 5, 8},
				{3, 4, 7},
				{3, 6, 5},
				{4, 5, 8},
				{4, 7, 7},
				{4, 6, 8},
				{5, 7, 4},
				{5, 9, 13},
				{6, 7, 14},
				{6, 8, 6},
				{7, 9, 8},
				{7, 8, 8},
				{8, 9, 10},
			},
			expectedMST: []Edge{
				{From: 5, To: 7, Weight: 4},
				{From: 3, To: 6, Weight: 5},
				{From: 1, To: 4, Weight: 6},
				{From: 6, To: 8, Weight: 6},
				{From: 3, To: 4, Weight: 7},
				{From: 4, To: 7, Weight: 7},
				{From: 7, To: 9, Weight: 8},
				{From: 2, To: 5, Weight: 8},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := NewEmptyGraph(tt.directed)
			graph.AddNodes(tt.numNodes)
			for _, edge := range tt.edges {
				graph.ConnectNodes(edge[0], edge[1], edge[2])
			}
			mst := graph.KruskalMST()

			if !compareEdges(mst, tt.expectedMST) {
				t.Errorf("Expected MST %v, but got %v", tt.expectedMST, mst)
			}
		})
	}
}

// Helper function to compare edges. Compares two slices of edges in a permutation-invariant manner.
// It first checks if the slices have the same length. If not, it returns false.
// Then, it sorts both slices to ensure that the order of elements is consistent.
// Finally, it iterates through the sorted slices and compares corresponding elements.
// This way, the function verifies that both slices contain the same edges,
// regardless of their initial order in the slices.
func compareEdges(got, expected []Edge) bool {
	if len(got) != len(expected) {
		return false
	}

	edgeMap := make(map[string]int)
	for _, edge := range got {
		key := fmt.Sprintf("%d-%d", edge.From, edge.To)
		if edge.From > edge.To {
			key = fmt.Sprintf("%d-%d", edge.To, edge.From)
		}
		edgeMap[key] = edge.Weight
	}

	for _, edge := range expected {
		key := fmt.Sprintf("%d-%d", edge.From, edge.To)
		if edge.From > edge.To {
			key = fmt.Sprintf("%d-%d", edge.To, edge.From)
		}
		if w, exists := edgeMap[key]; !exists || w != edge.Weight {
			return false
		}
	}

	return true
}
