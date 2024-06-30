## Binary Search Tree

### Features

* Dynamic insertion: `Insert` for adding single or multiple integers, ensuring all elements are unique.
* Deletion: `Delete` for removing nodes by value, addressing leaf nodes, single-child nodes, and two-children nodes.
* Search operations: `Contains` for existence checks, `Min` for the smallest value, and `Max` for the largest value.
* Traversal methods: Various traversal approaches are supported:
    - In-Order: `InOrderRecursive` and `InOrderIterative`
    - Pre-Order: `PreOrderRecursive` and `PreOrderIterative`
    - Post-Order: `PostOrderRecursive` and `PostOrderIterative`
    - Level-Order: `LevelOrder` for breadth-first traversal
* Utility functions:
    - `Height` for obtaining the tree height.
    - `IsBalanced` for checking tree balance.
    - `Successor` and `Predecessor` for finding the next larger or smaller elements.
    - `IsValidBST` for confirming the proper BST structure.
* Rebalancing:
    - `Rebalance` uses in-order traversal to collect all values and create new, balanced tree in place of the old tree.
    - `RebalanceDSW` uses Day-Stout-Warren algorithm to rebalance the tree in place, using O(1) space, and results in a complete binary tree, that is a tree which has the bottom level filled left to right.
* Visualization: `Print` for a visual display of the BST structure.

### Limitations

* Designed specifically for integer values; does not support generic types out of the box.
* No automatic balancing like AVL trees or Red-Black trees; manual rebalancing is required.
* Deletion and insertion operations can potentially unbalance the tree, leading to degraded performance (e.g., O(n) time complexity in skewed trees).

### Usage


```golang
// Create a new BST from a slice of integers
values := []int{7, 2, 4, 9, 1, 5, 8}
bst := NewFromSlice(values)
fmt.Println("Original BST:")
bst.Print()

// Insert new values
err := bst.Insert(6, 10)
if err != nil {
    fmt.Println("Insert error:", err)
}
fmt.Println("BST after inserting 6 and 10:")
bst.Print()

// Check if a value exists
fmt.Println("Contains 5:", bst.Contains(5))
fmt.Println("Contains 100:", bst.Contains(100))

// Get minimum and maximum value
minVal, err := bst.Min()
if err != nil {
    fmt.Println("Min error:", err)
} else {
    fmt.Println("Minimum value:", minVal)
}

maxVal, err := bst.Max()
if err != nil {
    fmt.Println("Max error:", err)
} else {
    fmt.Println("Maximum value:", maxVal)
}

// Perform different types of tree traversals, both Iterative and Recursive are available
fmt.Println("In-order traversal:", bst.InOrderIterative())
fmt.Println("Pre-order traversal:", bst.PreOrderRecursive())
fmt.Println("Post-order traversal:", bst.PostOrderRecursive())
fmt.Println("Level-order traversal:", bst.LevelOrder())

// Find successor and predecessor
successor, err := bst.Successor(4)
if err != nil {
    fmt.Println("Successor error:", err)
} else {
    fmt.Println("Successor of 4:", successor)
}

predecessor, err := bst.Predecessor(4)
if err != nil {
    fmt.Println("Predecessor error:", err)
} else {
    fmt.Println("Predecessor of 4:", predecessor)
}

// Check if BST is balanced and valid
fmt.Println("Is the tree balanced?", bst.IsBalanced())
fmt.Println("Is the tree a valid BST?", bst.IsValidBST())

// Delete a value
err = bst.Delete(4)
if err != nil {
    fmt.Println("Delete error:", err)
}
fmt.Println("BST after deleting 4:")
bst.Print()

// Rebalance the tree
bst.RebalanceDSW() // or bst.Rebalance()
fmt.Println("BST after rebalancing:")
bst.Print()

// Check height of the tree
fmt.Println("Height of tree:", bst.Height())
```