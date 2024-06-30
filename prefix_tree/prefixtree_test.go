package main

import (
	"testing"
)

// TestNewEmpty ensures that a new trie is correctly initialized.
func TestNewEmpty(t *testing.T) {
	trie := NewEmpty()
	if trie.Root == nil {
		t.Fatalf("Expected root node to be initialized")
	}
}

// TestInsertAndSearch tests basic insertion and search functionality.
func TestInsertAndSearch(t *testing.T) {
	trie := NewEmpty()

	words := []string{"apple", "banana", "grape", "apricot", "orange"}
	for _, word := range words {
		trie.Insert(word)
	}

	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Expected to find word '%s'", word)
		}
	}

	nonexistentWords := []string{"berry", "grapefruit", "melon"}
	for _, word := range nonexistentWords {
		if trie.Search(word) {
			t.Errorf("Expected not to find word '%s'", word)
		}
	}
}

// TestStartsWith tests the functionality of the StartsWith method.
func TestStartsWith(t *testing.T) {
	trie := NewEmpty()
	trie.Insert("apple", "applet", "banana", "bandana", "band")

	prefixes := []string{"app", "ban", "band", "appl"}
	for _, prefix := range prefixes {
		if !trie.StartsWith(prefix) {
			t.Errorf("Expected prefix '%s' to be found", prefix)
		}
	}

	nonexistentPrefixes := []string{" bana", "cat", "dog"}
	for _, prefix := range nonexistentPrefixes {
		if trie.StartsWith(prefix) {
			t.Errorf("Expected prefix '%s' not to be found", prefix)
		}
	}
}

// TestDelete tests the functionality of the Delete method.
func TestDelete(t *testing.T) {
	trie := NewEmpty()
	trie.Insert("apple", "app", "apricot", "banana", "bandana")

	deleteTests := []struct {
		word           string
		expectedError  bool
		expectedSearch bool
	}{
		{"apple", false, false},
		{"app", false, false},
		{"bandana", false, false},
		{"nonexistent", true, false},
	}

	for _, test := range deleteTests {
		err := trie.Delete(test.word)
		if (err != nil) != test.expectedError {
			t.Errorf("Delete(%s) error = %v, expectedError %v", test.word, err, test.expectedError)
		}
		if trie.Search(test.word) != test.expectedSearch {
			t.Errorf("Expected search result for '%s' to be %v", test.word, test.expectedSearch)
		}
	}
}

// TestDeleteWithPrefixes tests deleting words that share common prefixes.
func TestDeleteWithPrefixes(t *testing.T) {
	trie := NewEmpty()
	trie.Insert("app", "apple", "applet", "application")

	err := trie.Delete("apple")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if trie.Search("apple") {
		t.Errorf("Expected 'apple' to be deleted")
	}

	if !trie.Search("app") {
		t.Errorf("Expected 'app' to still exist")
	}

	if !trie.Search("applet") {
		t.Errorf("Expected 'applet' to still exist")
	}

	if !trie.Search("application") {
		t.Errorf("Expected 'application' to still exist")
	}
}

// TestDeleteRootOnlyWord tests deleting the root word when it's the only word in the trie.
func TestDeleteRootOnlyWord(t *testing.T) {
	trie := NewEmpty()
	trie.Insert("root")

	err := trie.Delete("root")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if trie.Search("root") {
		t.Errorf("Expected 'root' to be deleted")
	}

	if len(trie.Root.children) != 0 {
		t.Errorf("Expected root node to have no children after deletion")
	}
}

// TestEmptyTrieSearch tests search operations on an empty trie.
func TestEmptyTrieSearch(t *testing.T) {
	trie := NewEmpty()

	if trie.Search("anything") {
		t.Errorf("Expected 'anything' to not be found in an empty trie")
	}

	if trie.StartsWith("a") {
		t.Errorf("Expected prefix 'a' to not be found in an empty trie")
	}
}

// TestInsertEmptyString tests inserting and searching for an empty string.
func TestInsertEmptyString(t *testing.T) {
	trie := NewEmpty()
	trie.Insert("")

	if !trie.Search("") {
		t.Errorf("Expected empty string to be found")
	}
}

// TestComplexCases tests a combination of insert, delete, and search operations.
func TestComplexCases(t *testing.T) {
	trie := NewEmpty()
	trie.Insert("apple", "app", "applet", "banana", "bandana", "band")

	err := trie.Delete("banana")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if trie.Search("banana") {
		t.Errorf("Expected 'banana' to be deleted")
	}

	if !trie.Search("bandana") {
		t.Errorf("Expected 'bandana' to still exist")
	}

	if !trie.Search("band") {
		t.Errorf("Expected 'band' to still exist")
	}

	err = trie.Delete("band")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if trie.Search("band") {
		t.Errorf("Expected 'band' to be deleted")
	}

	if !trie.Search("bandana") {
		t.Errorf("Expected 'bandana' to still exist")
	}
}
