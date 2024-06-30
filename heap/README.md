## Min Heap (with Priority Queue functionality)

### Features
* Min-heap property: Ensures the element at the root is the minimum of the heap (in terms of priority).
* Can be used as a priority queue. All heap operations are done using the `priorities`, but we can also associate `values` with them, which are kept track of. This is useful for example in graph algorithms, where `values` would represent nodes, and `priorities` would represent their associated costs or distances.
* Array-based storage: Utilizes a dynamic array to store the heap elements, which allows for efficient use of space. Parents and children are accessed via indices in this array.
* Heap functions:
    - `PeekMin`: Returns the minimum element without removing it.
    - `Push`: Adds a new element to the heap, preserving the min-heap property.
    - `PopMin`: Removes and returns the minimum element, restructures the heap to maintain the heap property.
    - `DecreasePrio`: Decreases the priority of an element specified by a given index and adjusts its place in the heap.
* Constructors:
    - `NewEmpty`: Initializes an empty heap.
    - `New`: Initializes a heap from slices of integers for priorities and values, and transforms it into a min heap using the `heapify` method.

### Limitations
* The implementation only supports `int` priorities and values.
* Lacks advanced heap operations like `Delete` and `Meld`, which might be present in other implementations. There's also no optimized implementation for combining `PopMin` and `Push`, which could be done with one pass, instead of two.

```golang
// Create a new min heap prio queue from slices, priorities and values respectively.
heap := New([]int{10, 20, 5, 3, 2, 8}, []int{100, 200, 50, 30, 20, 80}) 
emptyHeap := NewEmpty() // Empty heap

// Peek the minimum element
minPriority, minValue, err := heap.PeekMin()
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Printf("Min element: (priority: %d, value: %d)\n", minPriority, minValue)
}

// Add a new element
heap.Push(1, 10)

// Remove the min element
poppedPriority, poppedValue, err := heap.PopMin()
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Printf("Popped element: (priority: %d, value: %d)\n", poppedPriority, poppedValue)
}
```
