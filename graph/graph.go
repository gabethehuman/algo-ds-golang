package main

import "fmt"

// Edge definition.
// We keep information about both From and To, because it is useful in some algorithms.
// However, there's no need for From attribute, as it can be inferred directly from the adjacency list.
// Eg. when we see map[0: [{..., To:1, Weight:0}]], we know it's an edge from 0 to 1.
type Edge struct {
	From   int
	To     int
	Weight int
}

// Get a new Edge with specified attributes. Should be used to create new edges
// instead of raw Edge objects, to avoid issues with default struct values.
func newEdge(from int, to int, weight int) Edge {
	return Edge{From: from, To: to, Weight: weight}
}

// Graph consists of nodes and adjacency list (a map in this case).
// Nodes are represented as consecutive ints, eg. graph with 3 nodes will always have nodes 1, 2, 3.
// Node enumeration starts from 1, if Nodes==0 it means that the graph is empty.
type Graph struct {
	Nodes         int
	AdjacencyList map[int][]Edge // A map from integers to a slice of Edges.
	Directed      bool
}

// Get a new empty graph.
func NewEmptyGraph(directed bool) Graph {
	return Graph{
		Nodes:         0, // Empty graph has 0 nodes.
		AdjacencyList: make(map[int][]Edge),
		Directed:      directed,
	}
}

// Add numNodes number of nodes to the graph. Panic if numNodes less than one.
func (g *Graph) AddNodes(numNodes int) {
	if numNodes < 1 {
		panic("AddNodes: numNodes should be greater than 0.")
	}
	for i := g.Nodes + 1; i <= g.Nodes+numNodes; i++ {
		g.AdjacencyList[i] = []Edge{}
	}
	g.Nodes = g.Nodes + numNodes
}

// Get the adjacency matrix of a graph.
func (g *Graph) AdjacencyMatrix() [][]int {
	matrix := make([][]int, g.Nodes)
	for i := range matrix {
		matrix[i] = make([]int, g.Nodes)
	}

	for key, value := range g.AdjacencyList {
		for _, edge := range value {
			matrix[key-1][edge.To-1] = edge.Weight // Shift indices by 1, because the nodes start from 1, not 0.
		}
	}
	return matrix
}

// Check if there is an edge between nodes `from` and `to`.
// This is achieved by iterating over the adjacency list for the `from` node.
func (g *Graph) edgeExists(from int, to int) bool {
	for _, node := range g.AdjacencyList[from] {
		if node.To == to {
			return true
		}
	}
	return false
}

// Connect node `from` with node `to` by appending entry to the adjacency list.
// For undirected graphs it makes a two way connection, for directed ones only one way.
// If an edge already exists between the nodes, panic.
func (g *Graph) ConnectNodes(from int, to int, weight int) {
	if from < 1 || from > g.Nodes || to < 1 || to > g.Nodes {
		panic(fmt.Sprintf("ConnectNodes: from and to node should be in range [1, %v]", g.Nodes))
	}
	if from == to {
		panic(fmt.Sprintf("ConnectNodes: cannot make an edge between nodes %v and %v", from, to))
	}
	if g.edgeExists(from, to) {
		panic(fmt.Sprintf("ConnectNodes: edge between %v and %v already exists", from, to))
	}
	if weight == 0 {
		panic("Weight of the connection should be non-zero")
	}

	g.AdjacencyList[from] = append(g.AdjacencyList[from], newEdge(from, to, weight))
	if !g.Directed {
		// For undirected graph make the connection both ways.
		g.AdjacencyList[to] = append(g.AdjacencyList[to], newEdge(to, from, weight))
	}
}

func main() {
}
