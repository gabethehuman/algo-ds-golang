package main

import "fmt"

// Iterative depth first search starting from a `node`, this is done in pre-order fashion.
func (g *Graph) IterDFS(node int) []int {
	if node < 1 || node > g.Nodes {
		panic(fmt.Sprintf("DFS: node should be in range [1, %v], got %v", g.Nodes, node))
	}
	seen := NewSet()
	stack := NewStack()
	result := []int{}

	stack.Push(node)
	for stack.Length() > 0 {
		v := stack.Pop()
		if !seen.Contains(v) {
			seen.Add(v)
			result = append(result, v) // Because we append it here, not after the loop, it's pre-order.
			for _, edge := range g.AdjacencyList[v] {
				stack.Push(edge.To)
			}
		}
	}
	return result
}

// Recursive depth first search starting from a `node`, , this is done in pre-order fashion.
func (g *Graph) RecDFS(node int) []int {
	if node < 1 || node > g.Nodes {
		panic(fmt.Sprintf("DFS: node should be in range [1, %v], got %v", g.Nodes, node))
	}

	seen := NewSet()
	result := []int{}

	var dfs func(v int)
	dfs = func(v int) {
		seen.Add(v)
		result = append(result, v) // Because we append it here, not after the loop, it's pre-order.
		for _, edge := range g.AdjacencyList[v] {
			if !seen.Contains(edge.To) {
				dfs(edge.To)
			}
		}
	}

	dfs(node)
	return result
}

// Breadth first search starting from a node `node`.
func (g *Graph) BFS(node int) []int {
	if node < 1 || node > g.Nodes {
		panic(fmt.Sprintf("BFS: node should be in range [1, %v], got %v", g.Nodes, node))
	}

	queue := NewQueue()
	seen := NewSet()
	result := []int{}

	seen.Add(node)
	queue.Enqueue(node)

	for queue.Length() > 0 {
		v := queue.Dequeue()
		result = append(result, v)
		for _, edge := range g.AdjacencyList[v] {
			if !seen.Contains(edge.To) {
				seen.Add(edge.To)
				queue.Enqueue(edge.To)
			}
		}
	}
	return result
}
