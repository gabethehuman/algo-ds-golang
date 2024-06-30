package main

import (
	"reflect"
	"sort"
	"testing"
)

// TestNewEmpty tests the creation of an empty BST.
func TestNewEmpty(t *testing.T) {
	bst := NewEmpty()
	if bst.Root != nil {
		t.Errorf("Expected new tree to have nil root, got %v", bst.Root)
	}
	if bst.Size != 0 {
		t.Errorf("Expected new tree to have size 0, got %d", bst.Size)
	}
}

// TestNewFromSlice tests the creation of a BST from a slice.
func TestNewFromSlice(t *testing.T) {
	values := []int{7, 2, 4, 9, 1, 5}
	bst := NewFromSlice(values)
	expectedSize := uint(len(values))
	if bst.Size != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, bst.Size)
	}

	// Check if tree is valid BST
	if !bst.IsValidBST() {
		t.Errorf("Expected valid BST from slice, got invalid BST")
	}
}

// TestBST_Insert tests inserting elements into the BST.
func TestBST_Insert(t *testing.T) {
	bst := NewEmpty()
	errors := bst.Insert(5, 3, 8, 1, 4, 7)
	if errors != nil {
		t.Errorf("Insert returned error: %v", errors)
	}
	if bst.Size != 6 {
		t.Errorf("Expected size 6, got %d", bst.Size)
	}

	// Insert duplicate and check for error
	err := bst.insertSingle(3)
	if err == nil {
		t.Errorf("Expected error when inserting duplicate, got nil")
	}
}

// TestBST_Contains checks the Contains method.
func TestBST_Contains(t *testing.T) {
	bst := NewFromSlice([]int{5, 3, 8, 1, 4, 7})
	if !bst.Contains(4) {
		t.Errorf("Expected to contain 4, but it did not")
	}
	if bst.Contains(10) {
		t.Errorf("Did not expect to contain 10, but it does")
	}
}

// TestBST_MinMax tests the Min and Max methods.
func TestBST_MinMax(t *testing.T) {
	bst := NewFromSlice([]int{5, 3, 8, -1, 4, 7})
	min, minErr := bst.Min()
	if minErr != nil || min != -1 {
		t.Errorf("Expected min 1, got %d, error: %v", min, minErr)
	}

	max, maxErr := bst.Max()
	if maxErr != nil || max != 8 {
		t.Errorf("Expected max 8, got %d, error: %v", max, maxErr)
	}
}

// TestBST_Height checks the height calculation of the BST.
func TestBST_Height(t *testing.T) {
	bst := NewFromSlice([]int{5, 3, 8, 1, 4, 7})
	expectedHeight := uint(3)
	if h := bst.Height(); h != expectedHeight {
		t.Errorf("Expected height %d, got %d", expectedHeight, h)
	}
}

// TestBST_IsBalanced checks if the BST remains balanced.
func TestBST_IsBalanced(t *testing.T) {
	bst := NewFromSlice([]int{5, 3, 8, 1, 4, 7})
	if !bst.IsBalanced() {
		t.Errorf("Expected BST to be balanced")
	}
}

// TestBST_InOrderRecursive tests the in-order traversal.
func TestBST_InOrderRecursive(t *testing.T) {
	bst := NewFromSlice([]int{5, 3, 8, 1, 4, 7})
	expected := []int{1, 3, 4, 5, 7, 8}
	result := bst.InOrderRecursive()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected in-order %v, got %v", expected, result)
	}
}

// TestBST_PreOrderRecursive tests the pre-order traversal.
func TestBST_PreOrderRecursive(t *testing.T) {
	bst := NewFromSlice([]int{5, 3, 8, 1, 4, 7})
	expected := []int{5, 3, 1, 4, 8, 7}
	result := bst.PreOrderRecursive()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected pre-order %v, got %v", expected, result)
	}
}

// TestBST_PostOrderRecursive tests the post-order traversal.
func TestBST_PostOrderRecursive(t *testing.T) {
	bst := NewFromSlice([]int{5, 3, 8, 1, 4, 7})
	expected := []int{1, 4, 3, 7, 8, 5}
	result := bst.PostOrderRecursive()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected post-order %v, got %v", expected, result)
	}
}

// TestBST_Delete tests the deletion of nodes.
func TestBST_Delete(t *testing.T) {
	bst := NewFromSlice([]int{5, 3, 8, 1, 4, 7})
	err := bst.Delete(3)
	if err != nil {
		t.Errorf("Delete returned error: %v", err)
	}
	if bst.Contains(3) {
		t.Errorf("Expected not to contain 3 after deletion")
	}
	if bst.Size != 5 {
		t.Errorf("Expected size 5 after deletion, got %d", bst.Size)
	}
}

// TestBST_Rebalance tests the rebalancing of the BST.
func TestBST_Rebalance(t *testing.T) {
	bst := NewFromSlice([]int{5, 3, 8, 1, 4, 7})
	bst.Insert(0, 2, 6, 9) // Make the tree a bit unbalanced
	bst.Rebalance()
	if !bst.IsBalanced() {
		t.Errorf("Expected BST to be balanced after rebalancing")
	}
}

// TestBST_SuccessorPredecessor tests finding successor and predecessor.
func TestBST_SuccessorPredecessor(t *testing.T) {
	bst := NewFromSlice([]int{5, 3, 8, 1, 4, 7})
	successor, sucErr := bst.Successor(4)
	if sucErr != nil || successor != 5 {
		t.Errorf("Expected successor 5, got %d, error: %v", successor, sucErr)
	}

	predecessor, predErr := bst.Predecessor(4)
	if predErr != nil || predecessor != 3 {
		t.Errorf("Expected predecessor 3, got %d, error: %v", predecessor, predErr)
	}
}

// TestBST_Validity tests the BST property across various operations.
func TestBST_Validity(t *testing.T) {
	bst := NewEmpty()
	bst.Insert(5, 3, 8, 1, 4, 7)
	if !bst.IsValidBST() {
		t.Errorf("Expected valid BST")
	}

	bst.Insert(6) // insert a new element
	if !bst.IsValidBST() {
		t.Errorf("BST validity failed after insertion")
	}

	bst.Delete(1) // delete a leaf node
	if !bst.IsValidBST() {
		t.Errorf("BST validity failed after deletion")
	}
}

// TestBST_InOrderIterative tests the in-order iterative traversal.
func TestBST_InOrderIterative(t *testing.T) {
	bst := NewFromSlice([]int{5, 3, 8, 1, 4, 7})
	expected := []int{1, 3, 4, 5, 7, 8}
	result := bst.InOrderIterative()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected in-order iterative %v, got %v", expected, result)
	}
}

// TestBST_PreOrderIterative tests the pre-order iterative traversal.
func TestBST_PreOrderIterative(t *testing.T) {
	bst := NewFromSlice([]int{5, 3, 8, 1, 4, 7})
	expected := []int{5, 3, 1, 4, 8, 7}
	result := bst.PreOrderIterative()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected pre-order iterative %v, got %v", expected, result)
	}
}

// TestBST_PostOrderIterative tests the post-order iterative traversal.
func TestBST_PostOrderIterative(t *testing.T) {
	bst := NewFromSlice([]int{5, 3, 8, 1, 4, 7})
	expected := []int{1, 4, 3, 7, 8, 5}
	result := bst.PostOrderIterative()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected post-order iterative %v, got %v", expected, result)
	}
}

// TestBST_LevelOrder tests the level-order (BFS) traversal.
func TestBST_LevelOrder(t *testing.T) {
	bst := NewFromSlice([]int{5, 3, 8, 1, 4, 7})
	expected := []int{5, 3, 8, 1, 4, 7}
	result := bst.LevelOrder()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected level-order %v, got %v", expected, result)
	}
}

// TestBST_TraversalConsistency checks consistency between recursive and iterative traversal methods.
func TestBST_TraversalConsistency(t *testing.T) {
	bst := NewFromSlice([]int{5, 3, 8, 1, 4, 7})
	if !reflect.DeepEqual(bst.InOrderRecursive(), bst.InOrderIterative()) {
		t.Errorf("InOrder traversal mismatch between recursive (%v) and iterative (%v)", bst.InOrderRecursive(), bst.InOrderIterative())
	}
	if !reflect.DeepEqual(bst.PreOrderRecursive(), bst.PreOrderIterative()) {
		t.Errorf("PreOrder traversal mismatch between recursive (%v) and iterative (%v)", bst.PreOrderRecursive(), bst.PreOrderIterative())
	}
	if !reflect.DeepEqual(bst.PostOrderRecursive(), bst.PostOrderIterative()) {
		t.Errorf("PostOrder traversal mismatch between recursive (%v) and iterative (%v)", bst.PostOrderRecursive(), bst.PostOrderIterative())
	}
}

// TestBST_ErrorHandling tests error scenarios across various functions.
func TestBST_ErrorHandling(t *testing.T) {
	// Test error when deleting from an empty tree.
	emptyBST := NewEmpty()
	if err := emptyBST.Delete(10); err == nil {
		t.Errorf("Expected error when deleting from empty tree, got nil")
	}

	// Test error when finding min and max in an empty tree.
	if _, err := emptyBST.Min(); err == nil {
		t.Errorf("Expected error when finding min in empty tree, got nil")
	}
	if _, err := emptyBST.Max(); err == nil {
		t.Errorf("Expected error when finding max in empty tree, got nil")
	}

	// Test error when attempting to delete a non-existent value.
	bst := NewFromSlice([]int{10, 5, 20, 15, 25})
	if err := bst.Delete(999); err == nil {
		t.Errorf("Expected error when deleting non-existent value, got nil")
	}

	// Test error when inserting duplicate values.
	if err := bst.Insert(10); err == nil {
		t.Errorf("Expected error when inserting duplicate value 10, got nil")
	}

	// Test error on successor and predecessor queries for non-existent values.
	if _, err := bst.Successor(999); err == nil {
		t.Errorf("Expected error when finding successor of non-existent value, got nil")
	}
	if _, err := bst.Predecessor(999); err == nil {
		t.Errorf("Expected error when finding predecessor of non-existent value, got nil")
	}

	// Test error for successor and predecessor when no valid successor/predecessor exists.
	bst = NewFromSlice([]int{10})
	if _, err := bst.Successor(10); err == nil {
		t.Errorf("Expected error when finding successor of the highest value, got nil")
	}
	if _, err := bst.Predecessor(10); err == nil {
		t.Errorf("Expected error when finding predecessor of the lowest value, got nil")
	}
}

// Test the RebalanceDSW function.
func TestRebalanceDSW(t *testing.T) {

	checkInOrderTraversal := func(tree *BST, values []int) bool {
		inOrderValues := tree.InOrderIterative()
		sort.Ints(values)
		if len(inOrderValues) != len(values) {
			return false
		}
		for i, v := range inOrderValues {
			if v != values[i] {
				return false
			}
		}
		return true
	}

	// Test case 1: Specific values as mentioned in the prompt.
	tree1 := NewEmpty()
	values1 := []int{10, 5, 20, 7, 15, 30, 25, 40, 23}
	tree1.Insert(values1...)
	tree1.RebalanceDSW()

	if !tree1.IsBalanced() {
		t.Errorf("Tree is not balanced after RebalanceDSW for values: %v", values1)
	}
	if !tree1.IsValidBST() {
		t.Errorf("Tree is not balanced a valid BST")
	}
	if !checkInOrderTraversal(tree1, values1) {
		t.Errorf("Tree in-order traversal does not match the sorted values: %v", values1)
	}

	// Test case 2: Already balanced tree.
	tree2 := NewEmpty()
	values2 := []int{10, 5, 15}
	tree2.Insert(values2...)
	tree2.RebalanceDSW()

	if !tree2.IsBalanced() {
		t.Errorf("Tree is not balanced after RebalanceDSW for values: %v", values2)
	}
	if !tree2.IsValidBST() {
		t.Errorf("Tree is not balanced a valid BST")
	}
	if !checkInOrderTraversal(tree2, values2) {
		t.Errorf("Tree in-order traversal does not match the sorted values: %v", values2)
	}

	// Test case 3: Single node tree.
	tree3 := NewEmpty()
	values3 := []int{10}
	tree3.Insert(values3...)
	tree3.RebalanceDSW()

	if !tree3.IsBalanced() {
		t.Errorf("Tree is not balanced after RebalanceDSW for values: %v", values3)
	}
	if !tree3.IsValidBST() {
		t.Errorf("Tree is not balanced a valid BST")
	}
	if !checkInOrderTraversal(tree3, values3) {
		t.Errorf("Tree in-order traversal does not match the sorted values: %v", values3)
	}

	// Test case 4: Large tree.
	tree4 := NewEmpty()
	values4 := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45, 55, 65, 75, 85}
	tree4.Insert(values4...)
	tree4.RebalanceDSW()

	if !tree4.IsBalanced() {
		t.Errorf("Tree is not balanced after RebalanceDSW for values: %v", values4)
	}
	if !tree4.IsValidBST() {
		t.Errorf("Tree is not balanced a valid BST")
	}
	if !checkInOrderTraversal(tree4, values4) {
		t.Errorf("Tree in-order traversal does not match the sorted values: %v", values4)
	}
}
