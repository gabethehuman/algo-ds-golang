package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestMergeSortImplementations(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
	}{
		{"Empty slice", []int{}},
		{"Single element", []int{1}},
		{"Two elements", []int{2, 1}},
		{"Already sorted", []int{1, 2, 3, 4, 5}},
		{"Reverse sorted", []int{5, 4, 3, 2, 1}},
		{"Duplicate elements", []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}},
		{"Negative numbers", []int{-5, -2, -8, 0, -1, 6, 3}},
	}

	// Add random test cases
	for i := 0; i < 5; i++ {
		size := rand.Intn(10000) + 1 // Random size between 1 and 10000
		randomSlice := generateRandomSlice(size)
		testCases = append(testCases, struct {
			name  string
			input []int
		}{
			name:  fmt.Sprintf("Random slice %d", i+1),
			input: randomSlice,
		})
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create copies of the input slice for each sort method
			sliceForMergeSort := make([]int, len(tc.input))
			sliceForMergeSortConcurrent := make([]int, len(tc.input))
			sliceForBuiltinSort := make([]int, len(tc.input))
			copy(sliceForMergeSort, tc.input)
			copy(sliceForMergeSortConcurrent, tc.input)
			copy(sliceForBuiltinSort, tc.input)

			// Perform sorts
			MergeSort(sliceForMergeSort)
			MergeSortConcurrent(sliceForMergeSortConcurrent, 256)
			sort.Ints(sliceForBuiltinSort)

			// Compare results
			if !slicesEqual(sliceForMergeSort, sliceForBuiltinSort) {
				t.Errorf("MergeSort result doesn't match built-in sort for input: %v", tc.input)
			}
			if !slicesEqual(sliceForMergeSortConcurrent, sliceForBuiltinSort) {
				t.Errorf("MergeSortConcurrent result doesn't match built-in sort for input: %v", tc.input)
			}
		})
	}
}

func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(1000) - 500 // Random integers between -500 and 499
	}
	return slice
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
