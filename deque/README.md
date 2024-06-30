## Deque (Double-Ended Queue)

### Features
* Efficient: Uses ring buffer to store data. Compared to the simples implementation using slice, ring buffer is more efficient as it avoids copying when pushing and popping from the front. It also has better cache locality and smaller overhead compared to linked list implementation.
* Generic implementation: Supports any data type.
* Dynamic resizing: Automatically doubles the buffer size when needed.
* Shrink capability: Shrinks the buffer size when less than 25% is occupied, but not below the initial buffer size.
* Double-Ended operations:
    - `PushBack`: Adds an element to the end of the deque.
    - `PushFront`: Adds an element to the front of the deque.
    - `PopBack`: Removes and returns the element from the end.
    - `PopFront`: Removes and returns the element from the front.
* Peek operations:
    - `PeekBack`: Returns the last element without removing it.
    - `PeekFront`: Returns the first element without removing it.
    - `PeekAt`: Returns the element at a specific index.
* Utility functions:
    - `Cap`: Returns the total capacity of the deque.
    - `Len`: Returns the number of elements currently in the deque.
    - `IsEmpty`: Checks if the deque is empty.
    - `Print`: Displays all elements in the deque.
    - `Clear`: Clears all the elements without changing the capacity.
    - `Index`: Finds the first index matching a specific predicate.
    - `SetAt`: Replaces the element at a specified index.

### Limitations
- No `Insert` operation. While technically possible, it would require shifting elements to maintain order, which contradicts the primary design principle of a deque. Also no `Delete` for removing arbitrary element.
- No `Rotate` or `Reverse` operations, which might be present in various implementations of deque.
- Pushing and popping works only for one value at a time, doesn't handle multiple values at once.
- Provided the size of the buffer is a power of 2, it would be possible to calculate indices more efficiently, using bit operations, at the cost of clarity.

### Usage
```golang
deque := NewDeque[int]() // Create a new deque for integers

// Add elements
deque.PushBack(10)
deque.PushFront(20)
fmt.Println("After pushes:")
deque.Print()

// Peek elements
back, front := deque.PeekBack(), deque.PeekFront()
fmt.Println("Back element:", back, "Front element:", front)

// Remove elements
popBack := deque.PopBack()
popFront := deque.PopFront()
fmt.Println("Popped Back:", popBack, "Popped Front:", popFront)

// Check length and capacity
fmt.Println("Current Length:", deque.Len())
fmt.Println("Current Capacity:", deque.Cap())

// Find indices.
deque.PushBack(15)
deque.PushBack(30)
index := deque.Index(func(x int) bool { return x > 20 })
fmt.Println("Index of first element greater than 20:", index)

// Set value at index
deque.SetAt(1, 25)
fmt.Println("After setting value at index 1 to 25:")
deque.Print()

// Clear the deque
deque.Clear()
fmt.Println("After clearing the deque:")
deque.Print()
```