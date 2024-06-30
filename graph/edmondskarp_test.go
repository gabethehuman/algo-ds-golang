package main

import (
	"testing"
)

func TestEdmondsKarp1(t *testing.T) {
	g := NewEmptyGraph(true)
	g.AddNodes(6)

	g.ConnectNodes(1, 2, 11)
	g.ConnectNodes(1, 3, 12)
	g.ConnectNodes(3, 2, 1)
	g.ConnectNodes(3, 5, 11)
	g.ConnectNodes(2, 4, 12)
	g.ConnectNodes(5, 4, 7)
	g.ConnectNodes(4, 6, 19)
	g.ConnectNodes(5, 6, 4)

	maxFlow := g.EdmondsKarp(1, 6)

	// Check if the max flow is correct
	expectedMaxFlow := 23
	if maxFlow != expectedMaxFlow {
		t.Errorf("EdmondsKarp() = %d; want %d", maxFlow, expectedMaxFlow)
	}
}

func TestEdmondsKarp2(t *testing.T) {
	g := NewEmptyGraph(true)
	g.AddNodes(6)

	g.AddNodes(6)
	g.ConnectNodes(1, 2, 10)
	g.ConnectNodes(1, 3, 10)
	g.ConnectNodes(3, 5, 15)
	g.ConnectNodes(5, 2, 6)
	g.ConnectNodes(2, 4, 25)
	g.ConnectNodes(4, 6, 10)
	g.ConnectNodes(5, 6, 10)

	maxFlow := g.EdmondsKarp(1, 6)

	// Check if the max flow is correct
	expectedMaxFlow := 20
	if maxFlow != expectedMaxFlow {
		t.Errorf("EdmondsKarp() = %d; want %d", maxFlow, expectedMaxFlow)
	}
}

func TestEdmondsKarp3(t *testing.T) {
	g := NewEmptyGraph(true)
	g.AddNodes(5)

	g.ConnectNodes(1, 2, 10)
	g.ConnectNodes(2, 3, 5)
	g.ConnectNodes(3, 4, 10)
	g.ConnectNodes(4, 5, 10)
	g.ConnectNodes(1, 3, 5)
	g.ConnectNodes(3, 2, 4)

	maxFlow := g.EdmondsKarp(1, 5)

	expectedMaxFlow := 10
	if maxFlow != expectedMaxFlow {
		t.Errorf("EdmondsKarp() for 5 nodes = %d; want %d", maxFlow, expectedMaxFlow)
	}
}

func TestEdmondsKarp4(t *testing.T) {
	g := NewEmptyGraph(true)
	g.AddNodes(4)

	g.ConnectNodes(1, 2, 5)
	g.ConnectNodes(1, 3, 7)
	g.ConnectNodes(2, 3, 3)
	g.ConnectNodes(2, 4, 4)
	g.ConnectNodes(3, 4, 8)

	maxFlow := g.EdmondsKarp(1, 4)

	expectedMaxFlow := 12
	if maxFlow != expectedMaxFlow {
		t.Errorf("EdmondsKarp() for 4 nodes = %d; want %d", maxFlow, expectedMaxFlow)
	}
}
