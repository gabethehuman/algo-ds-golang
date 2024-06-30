package main

import (
	"testing"
)

// TestHeap_Push checks the integrity of the Push operation.
// It validates whether elements are inserted correctly and maintains the min-heap property.
func TestHeap_Push(t *testing.T) {
	h := NewEmpty()
	h.Push(5, 5)
	if len(h.Priorities) != 1 || h.Priorities[0] != 5 || h.Values[0] != 5 {
		t.Errorf("Expected [(5, 5)], got %v", h)
	}

	h.Push(10, 10)
	if len(h.Priorities) != 2 || h.Priorities[0] != 5 || h.Values[0] != 5 || h.Priorities[1] != 10 || h.Values[1] != 10 {
		t.Errorf("Expected [(5, 5), (10, 10)], got %v", h)
	}

	h.Push(3, 3)
	if len(h.Priorities) != 3 || h.Priorities[0] != 3 || h.Values[0] != 3 || h.Priorities[1] != 10 || h.Values[1] != 10 || h.Priorities[2] != 5 || h.Values[2] != 5 {
		t.Errorf("Expected [(3, 3), (10, 10), (5, 5)], got %v", h)
	}

	h.Push(8, 8)
	if len(h.Priorities) != 4 || h.Priorities[0] != 3 || h.Values[0] != 3 || h.Priorities[1] != 8 || h.Values[1] != 8 || h.Priorities[2] != 5 || h.Values[2] != 5 || h.Priorities[3] != 10 || h.Values[3] != 10 {
		t.Errorf("Expected [(3, 3), (8, 8), (5, 5), (10, 10)], got %v", h)
	}
}

// TestHeap_PeekMin verifies the PeekMin function of the heap.
func TestHeap_PeekMin(t *testing.T) {
	h := NewEmpty()
	_, _, err := h.PeekMin()
	if err == nil {
		t.Error("Expected error when peeking from empty heap, got nil")
	}

	h.Push(2, 2)
	h.Push(8, 8)
	minPriority, minValue, err := h.PeekMin()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if minPriority != 2 || minValue != 2 {
		t.Errorf("Expected (2, 2), got (%d, %d)", minPriority, minValue)
	}
}

// TestHeap_PopMin ensures that the PopMin function correctly removes and returns the minimum element.
func TestHeap_PopMin(t *testing.T) {
	h := NewEmpty()
	_, _, err := h.PopMin()
	if err == nil {
		t.Error("Expected error when popping from empty heap, got nil")
	}

	h.Push(5, 5)
	h.Push(3, 3)
	h.Push(10, 10)
	minPriority, minValue, err := h.PopMin()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if minPriority != 3 || minValue != 3 {
		t.Errorf("Expected (3, 3), got (%d, %d)", minPriority, minValue)
	}
	if len(h.Priorities) != 2 || h.Priorities[0] != 5 || h.Priorities[1] != 10 || h.Values[0] != 5 || h.Values[1] != 10 {
		t.Errorf("Expected [(5, 5), (10, 10)], got %v", h)
	}
}

// TestHeap_New tests the New function, which initializes a heap with a given slice of integers.
func TestHeap_New(t *testing.T) {
	priorities := []int{5, 3, 10, 1, 4}
	values := []int{50, 30, 100, 10, 40}
	h := New(priorities, values)
	if len(h.Priorities) != 5 || h.Priorities[0] != 1 || h.Values[0] != 10 {
		t.Errorf("Expected [(1, 10) ...], got %v", h)
	}
}

// TestHeap_Empty checks that a new heap created with NewEmpty is indeed empty.
func TestHeap_Empty(t *testing.T) {
	h := NewEmpty()
	if len(h.Priorities) != 0 {
		t.Errorf("Expected empty heap, got %v", h)
	}
}

// TestHeap_BuildMinHeap tests the buildMinHeap function separately to ensure it can
// convert an arbitrary array into a min-heap.
func TestHeap_BuildMinHeap(t *testing.T) {
	priorities := []int{1, 3, 2, 7, 6, 5, 4}
	values := []int{10, 30, 20, 70, 60, 50, 40}
	priorities, values = buildMinHeap(priorities, values)
	expectedPriorities := []int{1, 3, 2, 7, 6, 5, 4}
	expectedValues := []int{10, 30, 20, 70, 60, 50, 40}
	for i, v := range priorities {
		if v != expectedPriorities[i] || values[i] != expectedValues[i] {
			t.Errorf("Expected %v and %v, got %v and %v", expectedPriorities, expectedValues, priorities, values)
			break
		}
	}
}

// TestHeap_Stress conducts a stress test on the heap's push and pop operations.
// It pushes a large number of elements (0 to 999) into the heap and then pops them all,
// verifying that pop operations return elements in increasing order.
func TestHeap_Stress(t *testing.T) {
	h := NewEmpty()
	for i := 0; i < 1000; i++ {
		h.Push(i, i)
	}

	for i := 0; i < 1000; i++ {
		minPriority, minValue, _ := h.PopMin()
		if minPriority != i || minValue != i {
			t.Errorf("Expected (%d, %d), got (%d, %d)", i, i, minPriority, minValue)
			break
		}
	}
}

// TestHeap_DecreasePrio checks the integrity of the DecreasePrio operation.
// It validates whether the priority of an element is decreased correctly and the min-heap property is maintained.
func TestHeap_DecreasePrio(t *testing.T) {
	h := New([]int{5, 3, 8, 1, 2}, []int{50, 30, 80, 10, 20})

	// Test decreasing priority of an element at a valid index
	err := h.DecreasePrio(2, 0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if h.Priorities[0] != 0 || h.Priorities[2] != 1 {
		t.Errorf("Expected priority 0 at root and 1 at index 2, got %v", h.Priorities)
	}

	// Test decreasing priority with index out of range
	err = h.DecreasePrio(10, 0)
	if err == nil {
		t.Error("Expected error for index out of range, got nil")
	}

	// Test decreasing priority to a value greater than current priority
	err = h.DecreasePrio(1, 4)
	if err == nil {
		t.Error("Expected error for new priority greater than current priority, got nil")
	}

	// Additional test: Decreasing priority of an element and ensuring heap property is maintained
	h = New([]int{10, 15, 20, 17, 25}, []int{100, 150, 200, 170, 250})
	err = h.DecreasePrio(3, 5) // Decrease priority of the element at index 3 to 5
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if h.Priorities[0] != 5 || h.Priorities[1] != 10 || h.Priorities[2] != 20 || h.Priorities[3] != 15 || h.Priorities[4] != 25 {
		t.Errorf("Expected priorities [5, 10, 20, 15, 25], got %v", h.Priorities)
	}
	if h.Values[0] != 170 || h.Values[1] != 100 || h.Values[2] != 200 || h.Values[3] != 150 || h.Values[4] != 250 {
		t.Errorf("Expected values [170, 100, 200, 150, 250], got %v", h.Values)
	}
}
