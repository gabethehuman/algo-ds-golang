## Graph algorithms

### Features
* Variants: directed and undirected graphs are supported.
* Graph functions:
    - `AdjacencyMatrix`: builds a adjacency matrix from the adjacency list of a graph.
    - `Dijkstra`: finds a shortest distance from a source node to every other node in a graph. Also provides predecessors of each node on the shortest path. A min binary heap is used for efficiency.
    - `KruskalMST`: builds a minimum spanning tree of a graph using Kruskal's algorithm. `UnionFind` structure is used to check if adding a node to the current MST would create a cycle.
    - `KahnTopoSort`: topological sorting of a graph using Kahn's algorithm.
    - `IterDFS`, `RecDFS`, `BFS`: traversal methods.
    - `EdmondsKarp`: computes the maximal flow.

### Limitations
* Only integer values for weights.
* More well known algorithms could be implemented, such as Bellman–Ford and Floyd–Warshall for shortest paths, algorithms related to cliques, connected components, matching, etc.

### Usage
```golang
g := NewEmptyGraph(false) // Directed graph.
g.AddNodes(5)
g.ConnectNodes(1, 2, 5) // Connect 1 to 2 with weight 5

// Adjacency matrix.
matrix := g.AdjacencyMatrix()

// Traversal methods.
nodes := g.BFS(1)    // BFS starting from node 1.
nodes = g.RecDFS(1)  // Recursive implementation of DFS.
nodes = g.IterDFS(1) // Iterative implementation of DFS.

// Topological sorting.
toposort, err := g.KahnTopoSort()
if err != nil {
    fmt.Println("Cycle detected, topological sort impossible.")
}

// Kruskal's minimum spanning tree.
mst := g.KruskalMST()

// Maximum flow, Edmonds-Karp implementation of Floyd-Fulkerson method.
maxFlow := g.EdmondsKarp(1, 2)

// Dijkstras shortest path algorithm.
distsances, previous := g.Dijkstra(1)
```
