package main

import (
	"fmt"
	"testing"
)

// Test insertion into the hashmap.
func TestHashMapInsertion(t *testing.T) {
	hm := NewHashMap(10)
	ok := hm.Insert("key1", "value1")
	if !ok {
		t.Errorf("Insert failed for key1")
	}
}

// Test retrieval from the hashmap.
func TestHashMapRetrieval(t *testing.T) {
	hm := NewHashMap(10)
	hm.Insert("key1", "value1")
	val, ok := hm.Get("key1")
	if !ok || val != "value1" {
		t.Errorf("Get failed for key1, expected 'value1', got '%s'", val)
	}
}

// Test updating an existing key in the hashmap.
func TestHashMapUpdate(t *testing.T) {
	hm := NewHashMap(10)
	hm.Insert("key1", "value1")
	ok := hm.Insert("key1", "newValue1")
	if !ok {
		t.Errorf("Insert failed for key1 on update attempt")
	}
	val, ok := hm.Get("key1")
	if !ok || val != "newValue1" {
		t.Errorf("Update failed for key1, expected 'newValue1', got '%s'", val)
	}
}

// Test hashmap resizing functionality.
func TestHashMapResizing(t *testing.T) {
	hm := NewHashMap(10)
	for i := 0; i < 10; i++ {
		hm.Insert(fmt.Sprint('a'+i), fmt.Sprint('A'+i))
	}
	if len(hm.entries) <= 10 { // assuming resizing doubles the array size
		t.Errorf("Resizing failed, expected size greater than 10, got %d", len(hm.entries))
	}
}

// Test deletion from the hashmap.
func TestHashMapDeletion(t *testing.T) {
	hm := NewHashMap(10)
	hm.Insert("key1", "value1")
	ok := hm.Delete("key1")
	if !ok {
		t.Errorf("Delete failed for key1")
	}
	_, ok = hm.Get("key1")
	if ok {
		t.Errorf("Get succeeded for deleted key1")
	}
}

// Test if deleting a non-existent key gives false.
func TestDeleteNonExistentKey(t *testing.T) {
	hm := NewHashMap(2)
	ok := hm.Delete("nonexistent")
	if ok {
		t.Errorf("Delete reported success for nonexistent key")
	}
}

// TestInsertLargeNumberOfItems ensures that the hashmap can handle a large number of insertions and still retrieve them correctly.
func TestInsertLargeNumberOfItems(t *testing.T) {
	hm := NewHashMap(10) // Start with a small size to force several resizes

	const itemCount = 1000
	for i := 0; i < itemCount; i++ {
		ok := hm.Insert(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
		if !ok {
			t.Fatalf("Insert failed at iteration %d", i)
		}
	}

	for i := 0; i < itemCount; i++ {
		val, ok := hm.Get(fmt.Sprintf("key%d", i))
		if !ok || val != fmt.Sprintf("value%d", i) {
			t.Errorf("Get failed for key%d, expected 'value%d', got '%s'", i, i, val)
		}
	}
}

// TestResizing ensures that the hashmap resizes correctly and maintains data integrity after resizing.
func TestResizing(t *testing.T) {
	hm := NewHashMap(2)

	// These insertions should trigger at least one resize
	hm.Insert("key1", "value1")
	hm.Insert("key2", "value2")
	hm.Insert("key3", "value3")

	// Check size to confirm resize
	expectedMinSize := 4 // Initial size was 2, expect at least one resize
	if len(hm.entries) < expectedMinSize {
		t.Errorf("Expected a minimum size of %d after resizing, got %d", expectedMinSize, len(hm.entries))
	}

	// Verify data integrity after resizing
	for i := 1; i <= 3; i++ {
		val, ok := hm.Get(fmt.Sprintf("key%d", i))
		if !ok || val != fmt.Sprintf("value%d", i) {
			t.Errorf("Data integrity check failed for key%d after resize, expected 'value%d', got '%s'", i, i, val)
		}
	}
}

// TestDeleteUpdatesLoadFactor verifies that deleting entries affects the load factor as expected.
func TestDeleteUpdatesLoadFactor(t *testing.T) {
	hm := NewHashMap(10)
	hm.Insert("key1", "value1")
	hm.Insert("key2", "value2")

	// Initial load factor check
	initialLoadFactor := hm.GetLoadFactor()

	hm.Delete("key1")

	// Load factor after deletion
	updatedLoadFactor := hm.GetLoadFactor()
	if updatedLoadFactor >= initialLoadFactor {
		t.Errorf("Load factor did not decrease after deletion, was %.2f, now %.2f", initialLoadFactor, updatedLoadFactor)
	}
}
