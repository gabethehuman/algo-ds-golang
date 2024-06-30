package main

import (
	"fmt"
	"math"
	"strings"
)

// Min heap, which can be used as a priority queue. All the heap operations work on `Priorities`,
// but we can associate `Values` with them. If `Values` are not needed, we can initialize them
// as anything and ignore them.
type Heap struct {
	Priorities []int
	Values     []int
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
	return (i - 1) / 2 // golang automatically floors integers
}

// Ensure that the subtree rooted at 'index' satisfies the min-heap property.
func heapify(priorities []int, values []int, index int) (prio []int, val []int) {
	smallest := index
	leftIndex := left(index)
	rightIndex := right(index)

	// Check if left child exists and is smaller than the current node.
	if leftIndex < len(priorities) && priorities[leftIndex] < priorities[smallest] {
		smallest = leftIndex
	}

	// Check if right child exists and is smaller than the smallest found so far.
	if rightIndex < len(priorities) && priorities[rightIndex] < priorities[smallest] {
		smallest = rightIndex
	}

	// If the smallest element is not the current node, swap them.
	if smallest != index {
		priorities[index], priorities[smallest] = priorities[smallest], priorities[index]
		values[index], values[smallest] = values[smallest], values[index]
		// Recursively heapify the sub-tree rooted at the new position of the smallest element.
		priorities, values = heapify(priorities, values, smallest)
	}

	return priorities, values
}

// Transform arbitrary array into min heap.
func buildMinHeap(priorities []int, values []int) (prio []int, val []int) {
	n := len(priorities)
	// Start from the last parent node and move to the root.
	for i := n/2 - 1; i >= 0; i-- {
		priorities, values = heapify(priorities, values, i)
	}
	return priorities, values
}

func NewEmpty() Heap {
	return Heap{}
}

func New(priorities []int, values []int) Heap {
	priorities, values = buildMinHeap(priorities, values)
	return Heap{Priorities: priorities, Values: values}
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

	for index > 0 && h.Priorities[index] < h.Priorities[parent(index)] {
		h.Priorities[index], h.Priorities[parent(index)] = h.Priorities[parent(index)], h.Priorities[index]
		h.Values[index], h.Values[parent(index)] = h.Values[parent(index)], h.Values[index]
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
	h.Priorities = h.Priorities[:len(h.Priorities)-1]
	h.Values = h.Values[:len(h.Values)-1]

	// Heapify from the root down. Left and right subtrees are already valid heaps
	// so one call to heapify is sufficient.
	if len(h.Priorities) > 0 {
		h.Priorities, h.Values = heapify(h.Priorities, h.Values, 0)
	}
	return minPriority, minValue, nil
}

// Decrease the priority of an element and adjust its place in the queue.
func (h *Heap) DecreasePrio(index int, newPriority int) error {
	if index < 0 || index >= len(h.Priorities) {
		return fmt.Errorf("DecreasePrio error: index out of range")
	}

	if newPriority > h.Priorities[index] {
		return fmt.Errorf("DecreasePrio error: new priority is greater than current priority")
	}

	// Decrease the priority.
	h.Priorities[index] = newPriority

	// Bubble up the element to restore the heap property.
	for index > 0 && h.Priorities[index] < h.Priorities[parent(index)] {
		h.Priorities[index], h.Priorities[parent(index)] = h.Priorities[parent(index)], h.Priorities[index]
		h.Values[index], h.Values[parent(index)] = h.Values[parent(index)], h.Values[index]
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

func main() {
}
