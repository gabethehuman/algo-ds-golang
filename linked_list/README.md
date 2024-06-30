## Doubly Linked List

### Features
* Generic implementation: Works with any data type.
* Bidirectional traversal: Supports forward and backward traversal through the list.
* Multiple insertion methods: `InsertFirst`, `InsertLast`, `InsertAt`
* Deletion operations: `DeleteFirst`, `DeleteLast`

### Limitations
* Many more methods could be implemented, such as `InsertBefore`, `InsertAfter`, `Move`, etc. Those could accept indices or references to elements.
* Methods could return a reference to a given element, for example `InsertAt(index, value)` could return a reference to the entry, so that the user can easily remove it later, or start iteration from there. This implementation does not have that feature.

### Usage
```golang
// Create a new list of integers
list := NewList[int]()

// Insert elements
list.InsertFirst(1)
list.InsertLast(2)
list.InsertAt(1, 3) // Inserts 3 at index 1

// Print the list
list.PrintList()

// Delete elements
list.DeleteFirst()
list.DeleteLast()
```
