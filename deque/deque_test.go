package main

import (
	"testing"
)

// TestNewDeque checks if a new deque is initialized correctly.
func TestNewDeque(t *testing.T) {
	deque := NewDeque[int]()
	if deque.Len() != 0 || deque.Cap() != initBufferSize {
		t.Errorf("NewDeque failed to initialize correctly with length %d and capacity %d", deque.Len(), deque.Cap())
	}
}

// TestPushBack tests pushing elements to the back of the deque.
func TestPushBack(t *testing.T) {
	deque := NewDeque[int]()
	for i := 0; i < 10; i++ {
		deque.PushBack(i)
		if deque.PeekBack() != i {
			t.Errorf("PushBack failed at index %d", i)
		}
	}
}

// TestPushFront tests pushing elements to the front of the deque.
func TestPushFront(t *testing.T) {
	deque := NewDeque[int]()
	for i := 0; i < 10; i++ {
		deque.PushFront(i)
		if deque.PeekFront() != i {
			t.Errorf("PushFront failed at index %d", i)
		}
	}
}

// TestPopBack tests popping elements from the back of the deque.
func TestPopBack(t *testing.T) {
	deque := NewDeque[int]()
	for i := 0; i < 10; i++ {
		deque.PushBack(i)
	}
	for i := 9; i >= 0; i-- {
		if val := deque.PopBack(); val != i {
			t.Errorf("PopBack failed, expected %d got %d", i, val)
		}
	}
}

// TestPopFront tests popping elements from the front of the deque.
func TestPopFront(t *testing.T) {
	deque := NewDeque[int]()
	for i := 0; i < 10; i++ {
		deque.PushBack(i)
	}
	for i := 0; i < 10; i++ {
		if val := deque.PopFront(); val != i {
			t.Errorf("PopFront failed, expected %d got %d", i, val)
		}
	}
}

// TestPeekFrontAndPeekAtConsistency checks consistency between PeekFront and PeekAt.
func TestPeekFrontAndPeekAtConsistency(t *testing.T) {
	deque := NewDeque[int]()
	for i := 0; i < 10; i++ {
		deque.PushBack(i)
	}
	if deque.PeekFront() != deque.PeekAt(0) {
		t.Errorf("PeekFront and PeekAt are inconsistent, got %d and %d", deque.PeekFront(), deque.PeekAt(0))
	}
}

// TestResize checks if the deque resizes correctly when exceeding initial capacity.
func TestResize(t *testing.T) {
	deque := NewDeque[int]()
	for i := 0; i < initBufferSize+1; i++ {
		deque.PushBack(i)
	}
	if deque.Cap() <= initBufferSize {
		t.Errorf("Resize failed, capacity %d should be greater than %d", deque.Cap(), initBufferSize)
	}
}

// TestShrink checks if the deque shrinks correctly when appropriate.
func TestShrink(t *testing.T) {
	deque := NewDeque[int]()
	for i := 0; i < 100; i++ {
		deque.PushBack(i)
	}
	for i := 0; i < 95; i++ {
		deque.PopFront()
	}
	deque.shrinkIfNeeded()
	if expected := max(initBufferSize, deque.Cap()/2); deque.Cap() != expected {
		t.Errorf("Shrink failed, capacity %d should be about %d", deque.Cap(), expected)
	}
}

// TestIsEmpty checks the IsEmpty function.
func TestIsEmpty(t *testing.T) {
	deque := NewDeque[int]()
	if !deque.IsEmpty() {
		t.Errorf("IsEmpty failed, expected true got false")
	}
	deque.PushBack(1)
	if deque.IsEmpty() {
		t.Errorf("IsEmpty failed, expected false got true")
	}
}

// TestClear checks the Clear method.
func TestClear(t *testing.T) {
	deque := NewDeque[int]()
	for i := 0; i < 10; i++ {
		deque.PushBack(i)
	}
	deque.Clear()
	if !deque.IsEmpty() {
		t.Errorf("Clear failed, deque is not empty after clear")
	}
	if deque.Len() != 0 {
		t.Errorf("Clear failed, length after clear is not 0")
	}
}

// TestIndex checks finding the index of an element.
func TestIndex(t *testing.T) {
	deque := NewDeque[int]()
	for i := 0; i < 10; i++ {
		deque.PushBack(i)
	}
	if idx := deque.Index(func(x int) bool { return x == 5 }); idx != 5 {
		t.Errorf("Index failed, expected 5 got %d", idx)
	}
	if idx := deque.Index(func(x int) bool { return x == 100 }); idx != -1 {
		t.Errorf("Index failed, expected -1 got %d", idx)
	}
}

// TestSetAt checks setting a value at a specific index.
func TestSetAt(t *testing.T) {
	deque := NewDeque[int]()
	for i := 0; i < 10; i++ {
		deque.PushBack(i)
	}
	deque.SetAt(5, 100)
	if val := deque.PeekAt(5); val != 100 {
		t.Errorf("SetAt failed, expected 100 got %d", val)
	}
}

// TestPeekBackEmpty tests peeking at the back of an empty deque to ensure it panics.
func TestPeekBackEmpty(t *testing.T) {
	d := NewDeque[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when peeking back on empty deque")
		}
	}()
	_ = d.PeekBack()
}

// TestPeekFrontEmpty tests peeking at the front of an empty deque to ensure it panics.
func TestPeekFrontEmpty(t *testing.T) {
	d := NewDeque[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when peeking front on empty deque")
		}
	}()
	_ = d.PeekFront()
}

// TestPopBackEmpty tests popping from the back of an empty deque to ensure it panics.
func TestPopBackEmpty(t *testing.T) {
	d := NewDeque[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when popping back on empty deque")
		}
	}()
	_ = d.PopBack()
}

// TestPopFrontEmpty tests popping from the front of an empty deque to ensure it panics.
func TestPopFrontEmpty(t *testing.T) {
	d := NewDeque[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when popping front on empty deque")
		}
	}()
	_ = d.PopFront()
}

// TestPeekAtOutOfBounds tests accessing an out-of-bounds index to ensure it panics.
func TestPeekAtOutOfBounds(t *testing.T) {
	d := NewDeque[int]()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when peeking at out of bounds index")
		}
	}()
	_ = d.PeekAt(3)
}

// TestSetAtOutOfBounds tests setting a value at an out-of-bounds index to ensure it panics.
func TestSetAtOutOfBounds(t *testing.T) {
	d := NewDeque[int]()
	d.PushBack(1)
	d.PushBack(2)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when setting at out of bounds index")
		}
	}()
	d.SetAt(11, 5)
}

// TestBehavioralConsistency checks the consistency of deque behavior through a series of operations.
func TestBehavioralConsistency(t *testing.T) {
	d := NewDeque[int]()
	// Pushing 0 to 19 at the back
	for i := 0; i < 20; i++ {
		d.PushBack(i)
	}
	// Popping 0 to 9 from the front, leaving 10 to 19
	for i := 0; i < 10; i++ {
		d.PopFront()
	}
	// Pushing 20 to 29 at the front, with 29 being the very front
	for i := 20; i < 30; i++ {
		d.PushFront(i)
	}

	// Check length
	if d.Len() != 20 {
		t.Errorf("Expected deque length of 20, got %d", d.Len())
	}

	// Check consistency of values in deque
	expectedValues := make([]int, 20)
	for i := 0; i < 10; i++ {
		expectedValues[i] = 29 - i // 29 down to 20 (last pushed to the front is at the front)
	}
	for i := 10; i < 20; i++ {
		expectedValues[i] = i // 10 to 19 (remaining after pops)
	}

	// Validate each element
	for i := 0; i < d.Len(); i++ {
		if d.PeekAt(i) != expectedValues[i] {
			t.Errorf("Expected %d at index %d, got %d", expectedValues[i], i, d.PeekAt(i))
		}
	}
}

// TestPushAfterEmpty checks the behavior of the deque after being emptied and then used again.
func TestPushAfterEmpty(t *testing.T) {
	d := NewDeque[int]()
	for i := 0; i < 10; i++ {
		d.PushBack(i)
	}
	for i := 0; i < 10; i++ {
		d.PopFront()
	}
	if !d.IsEmpty() {
		t.Errorf("Deque should be empty after popping all elements")
	}

	// Push again after empty
	d.PushBack(10)
	if d.PeekBack() != 10 {
		t.Errorf("Expected 10 at the back after re-pushing, got %d", d.PeekBack())
	}
}

// TestClearAndReuse checks that the deque can be repeatedly cleared and reused without issues.
func TestClearAndReuse(t *testing.T) {
	d := NewDeque[int]()
	for times := 0; times < 10; times++ {
		for i := 0; i < 100; i++ {
			d.PushBack(i)
		}
		d.Clear()
		if !d.IsEmpty() {
			t.Errorf("Deque not empty after Clear on iteration %d", times)
		}
		if d.Len() != 0 {
			t.Errorf("Deque length not 0 after Clear on iteration %d", times)
		}
	}
}
