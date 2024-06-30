# Algorithms and data structures in golang

The repository contains various data structures along with associated algorithms, written from scratch in Go. This collection is not exhaustive and does not necessarily emphasize the most crucial algorithms. It's mostly what I found interesting at the time of implementing, and it served as an outlet for me to learn some Go.

Each directory is self-contained. For simplicity, all required data structures from this repository are replicated locally within each directory.

The implementations aim to follow Go's best practices, but as part of the learning process, some code may not align with idiomatic Go. For instance, `panic` is occasionally used instead of error-by-value for simplicity, and it's not consistent throughout the repository. These and other choices may differ from typical approaches used by experienced Go developers working on bigger projects.

## Content

* [Binary search tree](binary_search_tree): A BST with in-order, pre-order, post-order, and level-order traversals, rebalancing, and other methods.
* [Double-ended queue](deque): A ring buffer implementation of a deque.
* [Graph](graph): Various graph algorithms such as Dijkstra's shortest path, Kruskal's minimum spanning tree, topological sorting, depth-first search (DFS), breadth-first search (BFS), and Edmonds-Karp's maximum flow.
* [Hashmap](hashmap): A hashmap with linear probing for collision resolution.
* [Heap](heap): A binary min heap, array based.
* [Doubly linked list](linked_list): A very simple linked list.
* [Merge sort](merge_sort): Regular and parallel versions of merge sort. The parallel implementation utilizes goroutines to explore concurrent programming in Go.
* [Prefix tree](prefix_tree): A simple trie.
* [Union find](union_find): A disjoint set data structure with path compression and union by size.


## Tests

Each data structure and algorithm has their own tests. Some tests were written manually, other were generated with the help from LLMs.

To run all the tests enter the directory and run `go test` as such:
```bash
cd algorithms-golang
go test ./...
```

## Requirements and Go version

There are no external requirements. All the algorithms were implemented using Go version `1.22.1`.