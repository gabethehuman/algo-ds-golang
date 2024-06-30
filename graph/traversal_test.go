package main

import (
	"testing"
)

// Test iter and rec implementations of DFS. Note, that they differ in the order, as one
// implementation explores left to right, the other one right to left.
func TestDFS(t *testing.T) {
	graph := NewEmptyGraph(true)
	graph.AddNodes(5)
	graph.ConnectNodes(1, 2, 1.0)
	graph.ConnectNodes(1, 3, 1.0)
	graph.ConnectNodes(2, 4, 1.0)
	graph.ConnectNodes(3, 5, 1.0)

	// Iter and rec version have different order.
	expectedIterOrder := []int{1, 3, 5, 2, 4}
	expectedRecOrder := []int{1, 2, 4, 3, 5}
	iterOrder := graph.IterDFS(1)
	recOrder := graph.RecDFS(1)
	if !slicesEqual(iterOrder, expectedIterOrder) {
		t.Errorf("IterDFS: Expected order %v, got %v", expectedIterOrder, iterOrder)
	}
	if !slicesEqual(recOrder, expectedRecOrder) {
		t.Errorf("RecDFS: Expected order %v, got %v", expectedRecOrder, recOrder)
	}

	// Add another test with different graph structure
	graph2 := NewEmptyGraph(true)
	graph2.AddNodes(4)
	graph2.ConnectNodes(1, 2, 1.0)
	graph2.ConnectNodes(1, 3, 1.0)
	graph2.ConnectNodes(3, 4, 1.0)

	expectedIterOrder2 := []int{1, 3, 4, 2}
	expectedRecOrder2 := []int{1, 2, 3, 4}
	iterOrder2 := graph2.IterDFS(1)
	recOrder2 := graph2.RecDFS(1)
	if !slicesEqual(iterOrder2, expectedIterOrder2) {
		t.Errorf("IterDFS: Expected order %v, got %v", expectedIterOrder2, iterOrder2)
	}
	if !slicesEqual(recOrder2, expectedRecOrder2) {
		t.Errorf("RecDFS: Expected order %v, got %v", expectedRecOrder2, recOrder2)
	}

	// Add another test with a more complex graph structure
	graph3 := NewEmptyGraph(true)
	graph3.AddNodes(6)
	graph3.ConnectNodes(1, 2, 1.0)
	graph3.ConnectNodes(1, 3, 1.0)
	graph3.ConnectNodes(2, 4, 1.0)
	graph3.ConnectNodes(2, 5, 1.0)
	graph3.ConnectNodes(3, 6, 1.0)

	expectedIterOrder3 := []int{1, 3, 6, 2, 5, 4}
	expectedRecOrder3 := []int{1, 2, 4, 5, 3, 6}
	iterOrder3 := graph3.IterDFS(1)
	recOrder3 := graph3.RecDFS(1)
	if !slicesEqual(iterOrder3, expectedIterOrder3) {
		t.Errorf("IterDFS: Expected order %v, got %v", expectedIterOrder3, iterOrder3)
	}
	if !slicesEqual(recOrder3, expectedRecOrder3) {
		t.Errorf("RecDFS: Expected order %v, got %v", expectedRecOrder3, recOrder3)
	}
}

func TestBFS(t *testing.T) {
	graph := NewEmptyGraph(true)
	graph.AddNodes(5)
	graph.ConnectNodes(1, 2, 1.0)
	graph.ConnectNodes(1, 3, 1.0)
	graph.ConnectNodes(2, 4, 1.0)
	graph.ConnectNodes(3, 5, 1.0)

	expectedOrder := []int{1, 2, 3, 4, 5}
	order := graph.BFS(1)
	if !slicesEqual(order, expectedOrder) {
		t.Errorf("Expected order %v, got %v", expectedOrder, order)
	}

	// Add another test with different graph structure
	graph2 := NewEmptyGraph(true)
	graph2.AddNodes(4)
	graph2.ConnectNodes(1, 2, 1.0)
	graph2.ConnectNodes(1, 3, 1.0)
	graph2.ConnectNodes(3, 4, 1.0)

	expectedOrder2 := []int{1, 2, 3, 4}
	order2 := graph2.BFS(1)
	if !slicesEqual(order2, expectedOrder2) {
		t.Errorf("Expected order %v, got %v", expectedOrder2, order2)
	}

	// Add another test with a more complex graph structure
	graph3 := NewEmptyGraph(true)
	graph3.AddNodes(6)
	graph3.ConnectNodes(1, 2, 1.0)
	graph3.ConnectNodes(1, 3, 1.0)
	graph3.ConnectNodes(2, 4, 1.0)
	graph3.ConnectNodes(2, 5, 1.0)
	graph3.ConnectNodes(3, 6, 1.0)

	expectedOrder3 := []int{1, 2, 3, 4, 5, 6}
	order3 := graph3.BFS(1)
	if !slicesEqual(order3, expectedOrder3) {
		t.Errorf("Expected order %v, got %v", expectedOrder3, order3)
	}
}
