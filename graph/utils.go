package main

import (
	"fmt"
	"math"
	"strings"
)

// ============================
// Set implementation.
// Used in: traversal.go in DFS algorithm.
// ============================

// Set for integers, wrapper around map with empty struct values.
type Set struct {
	elements map[int]struct{}
}

// NewSet creates a new Set and initializes the map.
func NewSet() *Set {
	return &Set{elements: make(map[int]struct{})}
}

// Add an element to the set.
func (s *Set) Add(value int) {
	s.elements[value] = struct{}{}
}

// Delete an element from the set.
func (s *Set) Delete(value int) {
	delete(s.elements, value)
}

// Check if the set cointains a value.
func (s *Set) Contains(value int) bool {
	_, exists := s.elements[value]
	return exists
}

func (s *Set) Length() int {
	return len(s.elements)
}

// ============================
// Stack implementation.
// Used in: traversal.go in DFS algorithm.
// ============================
type Stack struct {
	elements []int
}

// Get a new, empty stack.
func NewStack() *Stack {
	return &Stack{}
}

// Push an element to the stack.
func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

// Pop and return the last element of the stack.
func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		panic("Pop: stack is empty")
	}
	index := len(s.elements) - 1
	popped := s.elements[index]
	s.elements = s.elements[:index]
	return popped
}

func (s *Stack) Length() int {
	return len(s.elements)
}

// Return the last element of the stack without popping it.
func (s *Stack) Peek() int {
	if len(s.elements) == 0 {
		panic("Peek: stack is empty")
	}
	return s.elements[len(s.elements)-1]
}

// ============================
// Queue implementation.
// Used in: traversal.go in BFS algorithm.
// ============================
type Queue struct {
	elements []int
}

func NewQueue() *Queue {
	return &Queue{}
}

// Add an element to the queue.
func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

// Remove and return the first element of the queue.
func (q *Queue) Dequeue() int {
	if len(q.elements) == 0 {
		panic("Dequeue: queue is empty")
	}

	// Memory usage might not be optimal here. The slice queue[1:] does not free
	// the memory used by the first element q[0], but it does not matter for this use case.
	element := q.elements[0]
	q.elements = q.elements[1:]

	return element
}

func (q *Queue) Length() int {
	return len(q.elements)
}

// ============================
// UnionFind implementation.
// Used in: kruskal.go in Kruskal's algorithm.
// ============================

// Union find (disjoint set) data structure.
type UnionFind struct {
	parents []int // An array, where parents[i] is the parent of the i-th element.
	sizes   []int // To keep track of the sizes.
	numSets int   // Number of disjoint sets.
}

func NewUnionFind(num int) UnionFind {
	parents := make([]int, num)
	sizes := make([]int, num)

	for i := 0; i < num; i++ {
		parents[i] = i // At the beginning, each element is its own parent.
		sizes[i] = 1
	}
	return UnionFind{parents: parents, sizes: sizes, numSets: num}
}

// Add a new set with a single element.
func (u *UnionFind) NewSet() {
	u.parents = append(u.parents, len(u.parents))
	u.sizes = append(u.sizes, 1)
	u.numSets += 1
}

// Find the representative of the set `k` belongs to. Apply path compression along the way.
// Uses additional space for recursion stack, but it's simpler than iterative version.
func (u *UnionFind) Find(k int) int {
	if u.parents[k] != k {
		u.parents[k] = u.Find(u.parents[k]) // Path compression.
	}
	return u.parents[k]
}

// Iterative version of `Find`, doesn't require additional space for the stack, but it requires
// two passes from `k` to the root (or representative) of the set.
func (u *UnionFind) FindIterative(k int) int {
	root := k
	// Find the root of the tree.
	for root != u.parents[root] {
		root = u.parents[root]
	}
	// Path compression: go once again node by node, make every node point directly to the root.
	for k != root {
		parent := u.parents[k]
		u.parents[k] = root
		k = parent
	}
	return root
}

// Connect two sets which `k` and `l` belong to. The bigger set always becomes the parent.
// This is so called union-by-size, as opposed to union-by-rank.
func (u *UnionFind) Union(k int, l int) {
	root1 := u.Find(k)
	root2 := u.Find(l)

	if root1 == root2 {
		return
	}

	if u.sizes[root1] >= u.sizes[root2] {
		u.parents[root2] = root1
		u.sizes[root1] += u.sizes[root2]
	} else {
		u.parents[root1] = root2
		u.sizes[root2] += u.sizes[root1]
	}
	u.numSets -= 1
}

// ============================
// Priority queue (based on mean heap) implementation.
// Used in: dijkstra.go in Dijkstra's algorithm,
//          prim.go in Prim's algorithm
// ============================

// Min heap, which can be used as a priority queue. All the heap operations work on `priorities`,
// but we can associate `values` with them. An auxiliary hash map (indexMap) is used for quick access.
type Heap struct {
	Priorities []int
	Values     []int
	IndexMap   map[int]int // Maps value to its index in the heap to avoid searching for elements all the time
}

// Get the index of the left child.
func left(i int) int {
	return 2*i + 1
}

// Get the index of the right child.
func right(i int) int {
	return 2*i + 2
}

// Get the index of the parent.
func parent(i int) int {
	return (i - 1) / 2
}

// Ensure that the subtree rooted at 'index' satisfies the min-heap property.
func (h *Heap) heapify(index int) {
	smallest := index
	leftIndex := left(index)
	rightIndex := right(index)

	// Check if left child exists and is smaller than the current node.
	if leftIndex < len(h.Priorities) && h.Priorities[leftIndex] < h.Priorities[smallest] {
		smallest = leftIndex
	}

	// Check if right child exists and is smaller than the smallest found so far.
	if rightIndex < len(h.Priorities) && h.Priorities[rightIndex] < h.Priorities[smallest] {
		smallest = rightIndex
	}

	// If the smallest element is not the current node, swap them.
	if smallest != index {
		h.swap(index, smallest)
		h.heapify(smallest)
	}
}

// Swap elements and update indexMap
func (h *Heap) swap(i, j int) {
	h.Priorities[i], h.Priorities[j] = h.Priorities[j], h.Priorities[i]
	h.Values[i], h.Values[j] = h.Values[j], h.Values[i]
	h.IndexMap[h.Values[i]], h.IndexMap[h.Values[j]] = h.IndexMap[h.Values[j]], h.IndexMap[h.Values[i]]
}

// Transform arbitrary array into min heap.
func (h *Heap) buildMinHeap() {
	n := len(h.Priorities)
	for i := n/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

func NewHeap(priorities []int, values []int) Heap {
	h := Heap{
		Priorities: make([]int, len(priorities)),
		Values:     make([]int, len(values)),
		IndexMap:   make(map[int]int, len(values)),
	}

	copy(h.Priorities, priorities)
	copy(h.Values, values)

	// Initialize indexMap
	for i, val := range values {
		h.IndexMap[val] = i
	}

	h.buildMinHeap()
	return h
}

func (h *Heap) Len() int {
	return len(h.Priorities)
}

// Peek the minimum element without popping it.
func (h *Heap) PeekMin() (prio int, val int, err error) {
	if len(h.Priorities) == 0 {
		return 0, 0, fmt.Errorf("PeekMin error: heap is empty")
	}
	return h.Priorities[0], h.Values[0], nil
}

// Push a value with priority to the heap.
func (h *Heap) Push(priority int, value int) {
	h.Priorities = append(h.Priorities, priority)
	h.Values = append(h.Values, value)
	index := len(h.Priorities) - 1
	h.IndexMap[value] = index

	for index > 0 && h.Priorities[index] < h.Priorities[parent(index)] {
		h.swap(index, parent(index))
		index = parent(index)
	}
}

// Pop the min value from the heap.
func (h *Heap) PopMin() (prio int, val int, err error) {
	if len(h.Priorities) == 0 {
		return 0, 0, fmt.Errorf("PopMin error: heap is empty")
	}

	// Min value to be popped.
	minPriority := h.Priorities[0]
	minValue := h.Values[0]

	// Move the last element to the root.
	h.Priorities[0] = h.Priorities[len(h.Priorities)-1]
	h.Values[0] = h.Values[len(h.Values)-1]
	h.IndexMap[h.Values[0]] = 0

	h.Priorities = h.Priorities[:len(h.Priorities)-1]
	h.Values = h.Values[:len(h.Values)-1]

	if len(h.Priorities) > 0 {
		h.heapify(0)
	}
	delete(h.IndexMap, minValue)
	return minPriority, minValue, nil
}

// Decrease the priority of an element and adjust its place in the queue.
func (h *Heap) DecreasePrio(value int, newPriority int) error {
	index, exists := h.IndexMap[value]
	if !exists {
		return fmt.Errorf("DecreasePrio error: value not found in the heap")
	}

	if newPriority > h.Priorities[index] {
		return fmt.Errorf("DecreasePrio error: new priority is greater than current priority")
	}

	// Decrease the priority.
	h.Priorities[index] = newPriority

	// Bubble up the element to restore the heap property.
	for index > 0 && h.Priorities[index] < h.Priorities[parent(index)] {
		h.swap(index, parent(index))
		index = parent(index)
	}

	return nil
}

// Print the array as a tree.
func printAsTree(data []int) {
	n := len(data)
	if n == 0 {
		fmt.Println("Heap is empty")
		return
	}

	depth := int(math.Ceil(math.Log2(float64(n + 1)))) // Calculate the depth of the tree

	// Print each level
	for level := 0; level < depth; level++ {
		startIndex := int(math.Pow(2, float64(level))) - 1 // Start index for this level
		endIndex := int(math.Min(float64(startIndex+int(math.Pow(2, float64(level)))), float64(n)))

		// Space before first element to center the elements under their parent
		prefixSpaces := int(math.Pow(2, float64(depth-level-1))) - 1
		// Space between the elements at this level
		inbetweenSpaces := int(math.Pow(2, float64(depth-level))) - 1

		// Print spaces before first element at this level
		fmt.Print(strings.Repeat(" ", prefixSpaces))

		// Print elements at this level
		for i := startIndex; i < endIndex; i++ {
			fmt.Printf("%v", data[i])
			if i < endIndex-1 {
				fmt.Print(strings.Repeat(" ", inbetweenSpaces))
			}
		}
		fmt.Println() // Move to the next level
	}
}
