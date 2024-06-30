package main

import (
	"testing"
)

// Test graph creation, number of nodes, length of adj list.
func TestGraphCreation(t *testing.T) {
	// Test creating an empty graph
	graph := NewEmptyGraph(true)
	if graph.Nodes != 0 {
		t.Errorf("Expected graph to have 0 nodes, got %d", graph.Nodes)
	}
	if len(graph.AdjacencyList) != 0 {
		t.Errorf("Expected adjacency list to be empty, got %d elements", len(graph.AdjacencyList))
	}
	if !graph.Directed {
		t.Errorf("Expected graph to be directed")
	}

	// Test adding nodes
	graph.AddNodes(3)
	if graph.Nodes != 3 {
		t.Errorf("Expected graph to have 3 nodes, got %d", graph.Nodes)
	}
	if len(graph.AdjacencyList) != 3 {
		t.Errorf("Expected adjacency list to have 3 elements, got %d", len(graph.AdjacencyList))
	}
}

// Test edges in a directed graph.
func TestGraphConnection(t *testing.T) {
	graph := NewEmptyGraph(true)
	graph.AddNodes(3)

	// Test connecting nodes
	graph.ConnectNodes(1, 2, 1.0)
	if !graph.edgeExists(1, 2) {
		t.Errorf("Expected an edge from 1 to 2 to exist")
	}
	if graph.edgeExists(2, 1) {
		t.Errorf("Did not expect an edge from 2 to 1 to exist in a directed graph")
	}

	// Test connection panic on duplicate edges
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected ConnectNodes to panic on duplicate edge")
		}
	}()
	graph.ConnectNodes(1, 2, 1.0)
}

// Test edges in an undirected graph.
func TestUndirectedGraph(t *testing.T) {
	graph := NewEmptyGraph(false)
	graph.AddNodes(3)

	graph.ConnectNodes(1, 2, 1.0)
	if !graph.edgeExists(1, 2) || !graph.edgeExists(2, 1) {
		t.Errorf("Expected undirected edges between 1 and 2")
	}
}
