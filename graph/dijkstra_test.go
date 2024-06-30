package main

import (
	"reflect"
	"testing"
)

// Test with an invalid source node
func TestDijkstraInvalidSource(t *testing.T) {
	g := NewEmptyGraph(false)
	g.AddNodes(5)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Dijkstra with invalid source did not panic")
		}
	}()

	g.Dijkstra(6) // This should panic
}

// Test with a simple undirected graph
func TestDijkstraSimpleUndirected(t *testing.T) {
	g := NewEmptyGraph(false)
	g.AddNodes(5)

	g.ConnectNodes(1, 2, 1)
	g.ConnectNodes(2, 3, 2)
	g.ConnectNodes(3, 4, 3)
	g.ConnectNodes(4, 5, 4)

	dist, _ := g.Dijkstra(1)
	expected := map[int]int{1: 0, 2: 1, 3: 3, 4: 6, 5: 10}

	if !reflect.DeepEqual(dist, expected) {
		t.Errorf("Dijkstra (simple undirected) = %v, want %v", dist, expected)
	}
}

// Test with a single node graph
func TestDijkstraSingleNode(t *testing.T) {
	g := NewEmptyGraph(false)
	g.AddNodes(1)

	dist, _ := g.Dijkstra(1)
	expected := map[int]int{1: 0}

	if !reflect.DeepEqual(dist, expected) {
		t.Errorf("Dijkstra (single node) = %v, want %v", dist, expected)
	}
}

// Test Dijkstra's algorithm on the predefined graph
func TestDijkstraMainExampleDir(t *testing.T) {
	g := NewEmptyGraph(true)
	g.AddNodes(8)
	g.ConnectNodes(1, 2, 4)  // AB
	g.ConnectNodes(1, 3, 2)  // AC
	g.ConnectNodes(1, 6, 7)  // AG
	g.ConnectNodes(2, 4, 2)  // BD
	g.ConnectNodes(4, 7, 6)  // DH
	g.ConnectNodes(4, 6, 5)  // DG
	g.ConnectNodes(3, 6, 3)  // CG
	g.ConnectNodes(3, 5, 8)  // CF
	g.ConnectNodes(6, 8, 4)  // GJ
	g.ConnectNodes(5, 8, 3)  // FJ
	g.ConnectNodes(7, 8, 2)  // HJ
	dist, _ := g.Dijkstra(1) // Focus on distances, ignore predecessors
	expected := map[int]int{1: 0, 2: 4, 3: 2, 4: 6, 5: 10, 6: 5, 7: 12, 8: 9}
	for k, v := range expected {
		if dist[k] != v {
			t.Errorf("Expected shortest distance to node %d to be %d, got %d", k, v, dist[k])
		}
	}
}

// Test Dijkstra's algorithm on the predefined graph
func TestDijkstraMainExampleUndir(t *testing.T) {
	g := NewEmptyGraph(false)
	g.AddNodes(8)
	g.ConnectNodes(1, 2, 4)  // AB
	g.ConnectNodes(1, 3, 2)  // AC
	g.ConnectNodes(1, 6, 7)  // AG
	g.ConnectNodes(2, 4, 2)  // BD
	g.ConnectNodes(4, 7, 6)  // DH
	g.ConnectNodes(4, 6, 5)  // DG
	g.ConnectNodes(3, 6, 3)  // CG
	g.ConnectNodes(3, 5, 8)  // CF
	g.ConnectNodes(6, 8, 4)  // GJ
	g.ConnectNodes(5, 8, 3)  // FJ
	g.ConnectNodes(7, 8, 2)  // HJ
	dist, _ := g.Dijkstra(1) // Focus on distances, ignore predecessors
	expected := map[int]int{1: 0, 2: 4, 3: 2, 4: 6, 5: 10, 6: 5, 7: 11, 8: 9}
	for k, v := range expected {
		if dist[k] != v {
			t.Errorf("Expected shortest distance to node %d to be %d, got %d", k, v, dist[k])
		}
	}
}

// Test Dijkstra's algorithm on the predefined graph
func TestDijkstraSimpleDir(t *testing.T) {
	g := NewEmptyGraph(true)
	g.AddNodes(4)
	g.ConnectNodes(1, 2, 5)
	g.ConnectNodes(2, 3, 5)
	g.ConnectNodes(3, 4, 5)
	g.ConnectNodes(1, 4, 5)
	g.ConnectNodes(1, 3, 1)

	dist, _ := g.Dijkstra(1) // Focus on distances, ignore predecessors
	expected := map[int]int{1: 0, 2: 5, 3: 1, 4: 5}
	for k, v := range expected {
		if dist[k] != v {
			t.Errorf("Expected shortest distance to node %d to be %d, got %d", k, v, dist[k])
		}
	}
}
