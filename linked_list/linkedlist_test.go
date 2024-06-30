package main

import (
	"testing"
)

// TestNewList verifies NewList creates an empty list.
func TestNewList(t *testing.T) {
	list := NewList[int]()
	if list.First != nil || list.Last != nil || list.Length != 0 {
		t.Errorf("Expected empty list, got first: %v, last: %v, length: %d", list.First, list.Last, list.Length)
	}
}

// TestInsertFirst verifies inserting at the beginning of the list.
func TestInsertFirst(t *testing.T) {
	list := NewList[int]()
	list.InsertFirst(1)
	if list.First.Value != 1 || list.Last.Value != 1 || list.Length != 1 {
		t.Errorf("InsertFirst failed, expected first and last values of 1, got first: %v, last: %v", list.First.Value, list.Last.Value)
	}
}

// TestInsertLast verifies inserting at the end of the list.
func TestInsertLast(t *testing.T) {
	list := NewList[int]()
	list.InsertLast(2)
	if list.First.Value != 2 || list.Last.Value != 2 || list.Length != 1 {
		t.Errorf("InsertLast failed, expected first and last values of 2, got first: %v, last: %v", list.First.Value, list.Last.Value)
	}
}

// TestDeleteFirst verifies deleting the first element of the list.
func TestDeleteFirst(t *testing.T) {
	list := NewList[int]()
	list.InsertFirst(1)
	list.DeleteFirst()
	if list.First != nil || list.Last != nil || list.Length != 0 {
		t.Errorf("DeleteFirst failed, expected empty list, got first: %v, last: %v, length: %d", list.First, list.Last, list.Length)
	}
}

// TestDeleteLast verifies deleting the last element of the list.
func TestDeleteLast(t *testing.T) {
	list := NewList[int]()
	list.InsertLast(2)
	list.DeleteLast()
	if list.First != nil || list.Last != nil || list.Length != 0 {
		t.Errorf("DeleteLast failed, expected empty list, got first: %v, last: %v, length: %d", list.First, list.Last, list.Length)
	}
}

// TestInsertAt verifies inserting at a specific index.
func TestInsertAt(t *testing.T) {
	list := NewList[int]()
	err := list.InsertAt(0, 1) // Should succeed.
	if err != nil {
		t.Errorf("InsertAt(0, 1) returned an unexpected error: %v", err)
	}

	// Attempt to insert into an out-of-bounds index, which should fail.
	err = list.InsertAt(2, 2)
	if err == nil {
		t.Errorf("Expected InsertAt to fail for index out of bounds, but it succeeded")
	}

	if list.First.Value != 1 || list.Last.Value != 1 || list.Length != 1 {
		t.Errorf("After InsertAt, expected value of 1, got first: %v, last: %v", list.First.Value, list.Last.Value)
	}
}
