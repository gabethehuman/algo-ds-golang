package main

import (
	"math"
)

// Merge two already sorted subsequences of A into one sorted sequence.
// The subsequences are defined by indices:
// First subsequence: A[left] to A[middle]
// Second subsequence: A[middle+1] to A[right]
func merge(A []int, left int, middle int, right int) []int {
	leftElements := middle - left + 1
	rightElements := right - middle

	// Make left and right buffers, their length is +1 for the inf value at the end.
	L := make([]int, leftElements+1)
	R := make([]int, rightElements+1)

	for i := 0; i < leftElements; i++ {
		L[i] = A[left+i]
	}
	for i := 0; i < rightElements; i++ {
		R[i] = A[middle+i+1]
	}

	// Add large values at the end, this way if L and R have different lengths, we don't have to
	// do ifs, just a simple comparison is enough.
	L[leftElements] = math.MaxInt32
	R[rightElements] = math.MaxInt32

	i := 0
	j := 0

	// Combine the left and right slices into the original slice, overriding the values of the original
	for k := left; k <= right; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
	return A
}

// Recursive merge sort implementation.
func mergeSort(A []int, left int, right int) {
	if left >= right {
		return
	}
	middle := (left + right) / 2
	mergeSort(A, left, middle)    // Sort first half.
	mergeSort(A, middle+1, right) // Sort second half.
	merge(A, left, middle, right) // Merge the sorted halves.
}

// Wrapper function to start the merge sort.
func MergeSort(A []int) {
	mergeSort(A, 0, len(A)-1)
}

func main() {
}
