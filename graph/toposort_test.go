package main

import (
	"testing"
)

// Test toposort on various acyclic and cyclic graphs.
func TestKahnTopoSort(t *testing.T) {
	// Test a simple directed acyclic graph
	graph := NewEmptyGraph(true)
	graph.AddNodes(3)
	graph.ConnectNodes(1, 2, 1.0)
	graph.ConnectNodes(2, 3, 1.0)

	order, err := graph.KahnTopoSort()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedOrder := []int{1, 2, 3}
	for i, v := range order {
		if v != expectedOrder[i] {
			t.Errorf("Expected order %v, got %v", expectedOrder, order)
		}
	}

	// Test a more complex acyclic graph
	graph2 := NewEmptyGraph(true)
	graph2.AddNodes(6)
	graph2.ConnectNodes(1, 2, 1.0)
	graph2.ConnectNodes(1, 3, 1.0)
	graph2.ConnectNodes(2, 4, 1.0)
	graph2.ConnectNodes(3, 4, 1.0)
	graph2.ConnectNodes(4, 5, 1.0)
	graph2.ConnectNodes(5, 6, 1.0)

	order2, err := graph2.KahnTopoSort()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedOrder2 := []int{1, 2, 3, 4, 5, 6}
	if !slicesEqual(order2, expectedOrder2) {
		t.Errorf("Expected order %v, got %v", expectedOrder2, order2)
	}

	// Test another complex acyclic graph
	graph3 := NewEmptyGraph(true)
	graph3.AddNodes(7)
	graph3.ConnectNodes(1, 2, 1.0)
	graph3.ConnectNodes(1, 3, 1.0)
	graph3.ConnectNodes(2, 4, 1.0)
	graph3.ConnectNodes(3, 4, 1.0)
	graph3.ConnectNodes(3, 5, 1.0)
	graph3.ConnectNodes(4, 6, 1.0)
	graph3.ConnectNodes(5, 6, 1.0)
	graph3.ConnectNodes(6, 7, 1.0)

	order3, err := graph3.KahnTopoSort()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedOrder3 := []int{1, 2, 3, 4, 5, 6, 7}
	if !slicesEqual(order3, expectedOrder3) {
		t.Errorf("Expected order %v, got %v", expectedOrder3, order3)
	}

	// Test a graph with a cycle
	graphWithCycle := NewEmptyGraph(true)
	graphWithCycle.AddNodes(3)
	graphWithCycle.ConnectNodes(1, 2, 1.0)
	graphWithCycle.ConnectNodes(2, 3, 1.0)
	graphWithCycle.ConnectNodes(3, 1, 1.0)

	_, err = graphWithCycle.KahnTopoSort()
	if err == nil {
		t.Errorf("Expected error due to cycle, got nil")
	}
}

// Helper function to check if two slices are equal.
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
