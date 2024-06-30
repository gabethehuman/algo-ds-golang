package main

import (
	"testing"
)

// Test initialization of UnionFind data structure.
func TestNewUnionFind(t *testing.T) {
	u := NewUnionFind(6)
	if len(u.parents) != 6 {
		t.Errorf("Expected parents array length of 6, got %d", len(u.parents))
	}
	if len(u.sizes) != 6 {
		t.Errorf("Expected sizes array length of 6, got %d", len(u.sizes))
	}
	if u.numSets != 6 {
		t.Errorf("Expected numSets to be 6, got %d", u.numSets)
	}
	for i := 0; i < 6; i++ {
		if u.parents[i] != i {
			t.Errorf("Expected parents[%d] to be %d, got %d", i, i, u.parents[i])
		}
		if u.sizes[i] != 1 {
			t.Errorf("Expected sizes[%d] to be 1, got %d", i, u.sizes[i])
		}
	}
}

// Test Find function with path compression.
func TestFind(t *testing.T) {
	u := NewUnionFind(6)
	u.Union(0, 1)
	u.Union(1, 2)
	if u.Find(2) != 0 {
		t.Errorf("Expected root of 2 to be 0, got %d", u.Find(2))
	}
}

// Test FindIterative function with path compression.
func TestFindIterative(t *testing.T) {
	u := NewUnionFind(6)
	u.Union(0, 1)
	u.Union(1, 2)
	if u.FindIterative(2) != 0 {
		t.Errorf("Expected root of 2 to be 0, got %d", u.FindIterative(2))
	}
}

// Test Union function with union-by-size.
func TestUnion(t *testing.T) {
	u := NewUnionFind(6)
	u.Union(0, 1)
	if u.Find(1) != 0 {
		t.Errorf("Expected root of 1 to be 0, got %d", u.Find(1))
	}
	if u.sizes[0] != 2 {
		t.Errorf("Expected size of root 0 to be 2, got %d", u.sizes[0])
	}
}

// Test multiple Union operations.
func TestMultipleUnions(t *testing.T) {
	u := NewUnionFind(6)
	u.Union(0, 1)
	u.Union(1, 2)
	u.Union(3, 1)
	u.Union(5, 1)
	u.Union(4, 2)
	expectedParents := []int{0, 0, 0, 0, 0, 0}
	for i, parent := range u.parents {
		if parent != expectedParents[i] {
			t.Errorf("Expected parent of %d to be %d, got %d", i, expectedParents[i], parent)
		}
	}
	if u.numSets != 1 {
		t.Errorf("Expected numSets to be 1, got %d", u.numSets)
	}
}

// Test edge cases for UnionFind.
func TestUnionFindEdgeCases(t *testing.T) {
	u := NewUnionFind(1)
	if u.Find(0) != 0 {
		t.Errorf("Expected root of 0 to be 0, got %d", u.Find(0))
	}
	u = NewUnionFind(0)
	if u.numSets != 0 {
		t.Errorf("Expected numSets to be 0, got %d", u.numSets)
	}
}

func TestUnionSameElements(t *testing.T) {
	u := NewUnionFind(5)
	u.Union(1, 1)
	if u.numSets != 5 {
		t.Errorf("Unioning the same element should not change numSets, expected 5, got %d", u.numSets)
	}
	if u.Find(1) != 1 {
		t.Errorf("Expected root of 1 to be 1, got %d", u.Find(1))
	}
}

func TestUnionAlreadyConnectedElements(t *testing.T) {
	u := NewUnionFind(5)
	u.Union(1, 2)
	initialNumSets := u.numSets
	u.Union(1, 2)
	if u.numSets != initialNumSets {
		t.Errorf("Unioning already connected elements should not change numSets, expected %d, got %d", initialNumSets, u.numSets)
	}
}
