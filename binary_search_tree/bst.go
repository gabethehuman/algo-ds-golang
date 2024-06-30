package main

import (
	"fmt"
	"math"
	"sort"
)

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
}

type BST struct {
	Root *Node
	Size uint
}

func NewEmpty() *BST {
	return &BST{}
}

func NewFromSlice(values []int) *BST {
	sort.Ints(values)
	return &BST{Root: newFromSlice(values, nil), Size: uint(len(values))}
}

// Helper function to build a tree from a sorted slice.
// Use the middle element as the root and assign the left and right subtrees recursively.
// Current node is passed as a parameter to correctly assign parents.
func newFromSlice(values []int, current *Node) *Node {
	if len(values) == 0 {
		return nil
	}
	mid := len(values) / 2
	root := &Node{Value: values[mid], Parent: current}
	root.Left = newFromSlice(values[:mid], root)
	root.Right = newFromSlice(values[mid+1:], root)
	return root
}

// Insert a value into the BST by going left or right depending on the value.
// If the value already exists, return an error.
func (b *BST) insertSingle(value int) error {
	if b.Root == nil {
		b.Root = &Node{Value: value}
		b.Size = 1
		return nil
	}

	// Go left or right depending on the value. When you get to nil, insert there.
	// Return an error if the value already exists.
	current := b.Root
	for {
		if value < current.Value {
			if current.Left == nil {
				current.Left = &Node{Value: value, Parent: current}
				b.Size++
				return nil
			}
			current = current.Left
		} else if value > current.Value {
			if current.Right == nil {
				current.Right = &Node{Value: value, Parent: current}
				b.Size++
				return nil
			}
			current = current.Right
		} else {
			return fmt.Errorf("Insert error: value %d already exists", value)
		}
	}
}

// Insert multiple values into the BST.
func (b *BST) Insert(values ...int) error {
	for _, value := range values {
		err := b.insertSingle(value)
		if err != nil {
			return err
		}
	}
	return nil
}

// Check if the value exists in the BST.
func (b *BST) Contains(value int) bool {
	current := b.Root
	for current != nil {
		if value < current.Value {
			current = current.Left
		} else if value > current.Value {
			current = current.Right
		} else {
			return true
		}
	}
	return false
}

// Get the node with the min value, starting from the node n.
func minNode(n *Node) *Node {
	current := n
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Get the node with the max value, starting from the node n.
func maxNode(n *Node) *Node {
	current := n
	for current.Right != nil {
		current = current.Right
	}
	return current
}

// Get the minimum value in the BST by going left all the way down.
func (b *BST) Min() (int, error) {
	if b.Root == nil {
		return 0, fmt.Errorf("Min error: tree is empty")
	}
	return minNode(b.Root).Value, nil
}

// Get the maximum value in the BST by going right all the way down.
func (b *BST) Max() (int, error) {
	if b.Root == nil {
		return 0, fmt.Errorf("Max error: tree is empty")
	}
	return maxNode(b.Root).Value, nil
}

// Get the height of the tree.
func (b *BST) Height() uint {
	return height(b.Root)
}

// Helper function to get the height of the tree starting from the root and following recursively.
func height(n *Node) uint {
	if n == nil {
		return 0
	}
	return max(height(n.Left), height(n.Right)) + 1
}

// Check if the BST is balanced by comparing the height difference between the left and right subtrees.
func (b *BST) IsBalanced() bool {
	if b.Root == nil {
		return true
	}

	// Get the height difference between the left and right subtrees.
	// If the difference is greater than 1, the BST is not balanced.
	leftHeight := height(b.Root.Left)
	rightHeight := height(b.Root.Right)

	heightDiff := func(x int) int { // abs closure
		if x < 0 {
			return -x
		}
		return x
	}(int(leftHeight) - int(rightHeight))

	return heightDiff <= 1
}

// Recursive in-order traversal.
func inOrderRecursive(n *Node, result []int) []int {
	if n == nil {
		return result
	}
	result = inOrderRecursive(n.Left, result)
	result = append(result, n.Value)
	result = inOrderRecursive(n.Right, result)
	return result
}

// Recursive pre-order traversal.
func preOrderRecursive(n *Node, result []int) []int {
	if n == nil {
		return result
	}
	result = append(result, n.Value)
	result = preOrderRecursive(n.Left, result)
	result = preOrderRecursive(n.Right, result)
	return result
}

// Recursive post-order traversal.
func postOrderRecursive(n *Node, result []int) []int {
	if n == nil {
		return result
	}
	result = postOrderRecursive(n.Left, result)
	result = postOrderRecursive(n.Right, result)
	result = append(result, n.Value)
	return result
}

// Pre-order recursive traversal (node, left, right).
func (b *BST) InOrderRecursive() []int {
	return inOrderRecursive(b.Root, []int{})
}

// Pre-order recursive traversal (node, left, right).
func (b *BST) PreOrderRecursive() []int {
	return preOrderRecursive(b.Root, []int{})
}

// Post-order recursive traversal (left, right, node).
func (b *BST) PostOrderRecursive() []int {
	return postOrderRecursive(b.Root, []int{})
}

// In-order iterative traversal (left, node, right).
func (b *BST) InOrderIterative() []int {
	var result []int
	if b.Root == nil {
		return result
	}

	stack := NewStack[*Node]() // Stack to hold visited nodes.

	// Visit the left-most node and append every visited node to the stack.
	// Upon reaching the end, pop from the visited stack and go right.
	current := b.Root
	for stack.Length() > 0 || current != nil {
		if current != nil {
			// Current is not nil, meaning we can still go to the left child.
			stack.Push(current)
			current = current.Left
		} else {
			// Upon reaching a leaf, visit it and go right.
			current = stack.Pop()
			result = append(result, current.Value)
			current = current.Right
		}
	}
	return result
}

// Pre-order iterative traversal (node, left, right).
func (b *BST) PreOrderIterative() []int {
	var result []int
	if b.Root == nil {
		return result
	}

	// Stack to keep track of nodes, initialize with the root.
	stack := NewStack[*Node]()
	stack.Push(b.Root)

	for stack.Length() > 0 {
		// Process the current node.
		node := stack.Pop()
		result = append(result, node.Value)

		// Push right child first so that the left child is processed first.
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
	return result
}

// Post-order iterative traversal(left, right, node).
func (b *BST) PostOrderIterative() []int {
	var result []int
	if b.Root == nil {
		return result
	}

	stack := NewStack[*Node]() // Stack to keep track of nodes.
	var lastVisited *Node

	current := b.Root
	for stack.Length() > 0 || current != nil {
		if current != nil {
			stack.Push(current)
			current = current.Left
		} else {
			peek := stack.Peek()
			// If right child exists and we came from the left child, then move right.
			if peek.Right != nil && lastVisited != peek.Right {
				current = peek.Right
			} else {
				result = append(result, peek.Value)
				lastVisited = stack.Pop()
			}
		}
	}
	return result
}

// Level-order traversal (breadth frist).
func (b *BST) LevelOrder() []int {
	var result []int
	if b.Root == nil {
		return result
	}

	queue := NewQueue[*Node]()
	queue.Enqueue(b.Root)

	for queue.Length() > 0 {
		current := queue.Dequeue()
		result = append(result, current.Value)
		if current.Left != nil {
			queue.Enqueue(current.Left)
		}
		if current.Right != nil {
			queue.Enqueue(current.Right)
		}
	}
	return result
}

// Find the successor of a node, that is a node with the smallest value greater than the given node's value.
func (b *BST) Successor(value int) (int, error) {
	// Find the node with the given value, if doesn't exit, return an error.
	current := b.Root
	for current != nil {
		if value < current.Value {
			current = current.Left
		} else if value > current.Value {
			current = current.Right
		} else {
			break
		}
	}
	if current == nil {
		return 0, fmt.Errorf("Successor error: did not find value %d", value)
	}

	// If current has a right child, then is the leftmost node in the right child's subtree.
	// If current has no right child, move up the tree until we find a node that is a left child of its parent.
	// The successor is the parent of this newly found node.
	if current.Right != nil {
		current = current.Right
		for current.Left != nil {
			current = current.Left
		}
		return current.Value, nil
	} else {
		for current.Parent != nil {
			if current == current.Parent.Left {
				return current.Parent.Value, nil
			}
			current = current.Parent
		}
	}
	return 0, fmt.Errorf("Successor error: no successor of %d in the tree, it is the rightmost node", value)
}

// Find the successor of a node, that is a node with the largest value smaller than the given node's value.
func (b *BST) Predecessor(value int) (int, error) {
	// Find the node with the given value, if doesn't exit, return an error.
	current := b.Root
	for current != nil {
		if value < current.Value {
			current = current.Left
		} else if value > current.Value {
			current = current.Right
		} else {
			break
		}
	}
	if current == nil {
		return 0, fmt.Errorf("Predecessor error: did not find value %d", value)
	}

	// If current has a left child, then is the rightmost node in the left child's subtree.
	// If current has no left child, move up the tree until we find a node that is a right child of its parent.
	// The predecessor is the parent of this newly found node.
	if current.Left != nil {
		current = current.Left
		for current.Right != nil {
			current = current.Right
		}
		return current.Value, nil
	} else {
		for current.Parent != nil {
			if current == current.Parent.Right {
				return current.Parent.Value, nil
			}
			current = current.Parent
		}
	}
	return 0, fmt.Errorf("Predecessor error: no predecessor of %d in the tree, it is the leftmost node", value)
}

// Check if a BST is valid using in-order traversal and checking if the result is sorted.
func (b *BST) IsValidBST() bool {
	return sort.IntsAreSorted(b.InOrderIterative())
}

// Helper function to delete a node from the BST.
// If no with the given value is found, return an error.
// If node to be deleted is a leaf node, just delete it by setting its parent to nil and its parent left/right to nil.
// If node to be deleted has one child, the child takes the nodes place by moving pointers around.
// If node to be deleted has two children, find the successor of the node, assign the successor's value to the
// current node, then delete the successor (it will have 0 or 1 children).
func (b *BST) deleteNode(node *Node) error {
	if node == nil {
		return fmt.Errorf("Delete error: node is nil")
	}

	// Case 1: Node is a leaf node.
	if node.Left == nil && node.Right == nil {
		if node.Parent == nil {
			// If left and right children are nil, and also parent is nil, then it's the only node in the tree.
			b.Root = nil
		} else {
			// Check if current node is a left child or a right child, then delete the connection between the nodes.
			if node == node.Parent.Left {
				node.Parent.Left = nil
			} else {
				node.Parent.Right = nil
			}
		}
		b.Size--
		return nil
	}

	// Case 2: Node has exactly one child.
	if (node.Left == nil) != (node.Right == nil) {

		// Determine which child (left or right) is non-nil.
		child := node.Right
		if node.Left != nil {
			child = node.Left
		}

		if node.Parent == nil {
			// Handle the case where current node is the root with only one child.
			b.Root = child
			b.Root.Parent = nil
		} else {
			// Connect the parent to the child, leaving current node out.
			if node == node.Parent.Left {
				node.Parent.Left = child
			} else {
				node.Parent.Right = child
			}
			child.Parent = node.Parent
		}
		b.Size--
		return nil
	}

	// Case 3: Node has two children.
	successor := minNode(node.Right)
	node.Value = successor.Value   // Transfer the value from successor to node.
	return b.deleteNode(successor) // Recursively delete the successor, which will have 0 or 1 children.
}

// Delete a node with the given value from the BST.
func (b *BST) Delete(value int) error {
	if b.Root == nil {
		return fmt.Errorf("Delete error: tree is empty")
	}

	// Find the node to be deleted.
	current := b.Root
	for current != nil {
		if value < current.Value {
			current = current.Left
		} else if value > current.Value {
			current = current.Right
		} else {
			break // Node to delete is found.
		}
	}

	if current == nil {
		return fmt.Errorf("Delete error: did not find value %d", value)
	}

	// Use the helper function to delete the node.
	return b.deleteNode(current)
}

// Rebalance the tree using in-order traversal, and then creating a new, balanced tree.
func (b *BST) Rebalance() {
	values := b.InOrderIterative() // Already sorted.
	b.Root = newFromSlice(values, nil)
	// Size remains the same, no need to modify b.Size.
}

// Rotate the tree clockwise around the node, call it node A.
// The left child of A (node B) becomes the new root, and A becomes the right child of B.
// The right child of B (node E) becomes the left child of A.
//
//	    A             B
//	   / \           / \
//	  B   C  ---->  D   A
//	 / \               / \
//	D   E             E   C
func (n *Node) rightRotation() *Node {

	// Select B as the new root and connect it to the upper side of the tree.
	newRoot := n.Left
	newRoot.Parent = n.Parent
	if n.Parent != nil { // Check if A was left or right child and plug B into that slot instead.
		if n.Parent.Left == n {
			n.Parent.Left = newRoot
		} else {
			n.Parent.Right = newRoot
		}
	}

	// Connect E as a left child of A, if E is not nil, then update E's parent as well.
	n.Left = newRoot.Right
	if newRoot.Right != nil {
		newRoot.Right.Parent = n
	}

	// Move A to the right child of B.
	newRoot.Right = n
	n.Parent = newRoot

	return newRoot
}

// Rotate the tree counter-clockwise around the node, call it node A.
// The right child of A (node C) becomes the new root, and A becomes the left child of C.
// The left child of C (node D) becomes the right child of A.
//
//	  A                 C
//	 / \               / \
//	B   C    ---->    A   E
//	   / \           / \
//	  D   E         B   D
func (n *Node) leftRotation() *Node {
	// Select C as the new root and connect it to the upper side of the tree.
	newRoot := n.Right
	newRoot.Parent = n.Parent
	if n.Parent != nil { // Check if A was left or right child and plug C into that slot instead.
		if n.Parent.Left == n {
			n.Parent.Left = newRoot
		} else {
			n.Parent.Right = newRoot
		}
	}

	// Connect D as a right child of A, if E is not nil, then update E's parent as well.
	n.Right = newRoot.Left
	if newRoot.Left != nil {
		newRoot.Left.Parent = n
	}

	// Move A to the right child of C.
	newRoot.Left = n
	n.Parent = newRoot

	return newRoot
}

// Rebalance the tree using Day–Stout–Warren algorithm.
// This is done in-place by first converting the tree into a list, called a vine or a backbone,
// which is then transformed into a tree using tree rotations.
// DSW algorithm results in a complete binary tree, that is a tree where the last level is filled left to right.
func (b *BST) RebalanceDSW() {
	// First step is converting the tree into a list, where every node has only the right child.
	// We keep moving to the left child if it exists, perform a right rotation, and then move to the right child.
	current := b.Root
	for current != nil {
		for current.Left != nil {
			current = current.rightRotation()
			if current.Parent == nil {
				// This is to ensure the root of the vine is correct. rightRotation already takes care of reassigning
				// parents, so if after a rotation a node has no parent, it has to be the root of the tree.
				b.Root = current
			}
		}
		current = current.Right
	}
	// At this point, b is a tree with only right children, effectively it's a tree degenerated to a list.

	// Calculate the number of initial rotations that have to be performed.
	// This number has the following interpretation: if a perfect tree can be built from b.Size nodes this number is 0.
	// Otherwise this number is equal to the number of leaves in the last level of the tree.
	// For instance, if b.Size == 5, then numInitialRotations == 2, if b.Size == 7, then numInitialRotations == 0.
	numInitialRotations := int(b.Size) - int(math.Pow(2, math.Floor(math.Log2(float64(b.Size+1))))) + 1

	// Perform initial rotations starting from the root.
	// This step ensures the bottom level is filled left to right.
	current = b.Root
	for i := 0; i < numInitialRotations; i++ {
		current = current.leftRotation()
		if current.Parent == nil {
			b.Root = current
		}
		if current.Right != nil {
			current = current.Right
		}
	}

	// Balance the tree by performing left rotations, jumping every two nodes on the vine.
	for m := int(b.Size) - numInitialRotations; m > 1; m /= 2 {
		current = b.Root
		for i := 0; i < m/2; i++ {
			current = current.leftRotation()
			if current.Parent == nil {
				b.Root = current
			}
			if current.Right != nil {
				current = current.Right
			}
		}
	}
}

// Helper recursive function to print the tree.
func printTree(node *Node, prefix string, isTail bool) {
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		printTree(node.Right, newPrefix, false)
	}
	fmt.Print(prefix)
	if isTail {
		fmt.Printf("└── %d\n", node.Value)
	} else {
		fmt.Printf("┌── %d\n", node.Value)
	}
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		printTree(node.Left, newPrefix, true)
	}
}

// Print the whole BST.
func (bst *BST) Print() {
	if bst.Root != nil {
		printTree(bst.Root, "", true)
	} else {
		fmt.Println("Tree is empty")
	}
}

func main() {
}
