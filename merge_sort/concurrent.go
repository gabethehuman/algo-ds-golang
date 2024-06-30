package main

import (
	"sync"
)

// Performs a concurrent merge of two sorted subarrays of T into A.
// It recursively divides the problem and uses goroutines for parallel execution.
// Parameters:
//
//	T: The source array containing two sorted subarrays to be merged
//	p1, r1: Start and end indices of the first subarray in T
//	p2, r2: Start and end indices of the second subarray in T
//	A: The destination array where the merged result will be stored
//	p3: Start index in A where the merged result should be placed
//	threshold: Minimum size of subproblem to switch to sequential merge
//	wg: Pointer to sync.WaitGroup for synchronization
func mergeConcurrent(T []int, p1 int, r1 int, p2 int, r2 int, A []int, p3 int, threshold int, wg *sync.WaitGroup) {
	defer wg.Done()
	if p1 > r1 {
		copy(A[p3:p3+r2-p2+1], T[p2:r2+1])
		return
	}
	if p2 > r2 {
		copy(A[p3:p3+r1-p1+1], T[p1:r1+1])
		return
	}
	if r1-p1+r2-p2+2 <= threshold {
		sequentialMerge(T, p1, r1, p2, r2, A, p3)
		return
	}
	mid1 := (p1 + r1) / 2
	midVal := T[mid1]
	low, high := p2, r2
	for low <= high {
		mid2 := (low + high) / 2
		if T[mid2] <= midVal {
			low = mid2 + 1
		} else {
			high = mid2 - 1
		}
	}
	mid2 := low - 1
	mid3 := p3 + (mid1 - p1) + (mid2 - p2 + 1)
	A[mid3] = midVal
	wg.Add(2)
	go mergeConcurrent(T, p1, mid1-1, p2, mid2, A, p3, threshold, wg)
	mergeConcurrent(T, mid1+1, r1, mid2+1, r2, A, mid3+1, threshold, wg)
}

// sequentialMerge performs a sequential merge of two sorted subarrays of T into A.
// This function is used when the problem size falls below the threshold for parallel execution.
// Parameters:
//
//	T: The source array containing two sorted subarrays to be merged
//	p1, r1: Start and end indices of the first subarray in T
//	p2, r2: Start and end indices of the second subarray in T
//	A: The destination array where the merged result will be stored
//	p3: Start index in A where the merged result should be placed
func sequentialMerge(T []int, p1 int, r1 int, p2 int, r2 int, A []int, p3 int) {
	i, j, k := p1, p2, p3
	for i <= r1 && j <= r2 {
		if T[i] <= T[j] {
			A[k] = T[i]
			i++
		} else {
			A[k] = T[j]
			j++
		}
		k++
	}
	for i <= r1 {
		A[k] = T[i]
		i++
		k++
	}
	for j <= r2 {
		A[k] = T[j]
		j++
		k++
	}
}

// Performs a concurrent merge sort on the given array.
// It uses goroutines to sort subarrays in parallel when the problem size is above the threshold.
// Parameters:
//
//	A: The array to be sorted
//	left, right: The range of indices to be sorted (inclusive)
//	threshold: Minimum size of subproblem to switch to sequential merge sort
func mergeSortConcurrent(A []int, left int, right int, threshold int) {
	if right-left+1 <= threshold {
		mergeSort(A, left, right) // Use sequential merge sort when below threshold.
		return
	}
	middle := (left + right) / 2
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		mergeSortConcurrent(A, left, middle, threshold)
		wg.Done()
	}()
	go func() {
		mergeSortConcurrent(A, middle+1, right, threshold)
		wg.Done()
	}()
	wg.Wait()
	temp := make([]int, right-left+1)
	var mergeWg sync.WaitGroup
	mergeWg.Add(1)
	mergeConcurrent(A, left, middle, middle+1, right, temp, 0, threshold, &mergeWg)
	mergeWg.Wait()
	copy(A[left:right+1], temp)
}

// Wrapper function to start the concurrent merge sort.
func MergeSortConcurrent(A []int, threshold int) {
	mergeSortConcurrent(A, 0, len(A)-1, threshold)
}
