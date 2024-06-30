## Merge sortwith additional parallel implementation

### Features
* `MergeSort`: A sequential version of merge sort.
* `MergeSortConcurrent`: A parallel version of merge sort using goroutines. There's a subproblem threshold, below which a sequential version of `merge` and `mergeSort` will be used.

### Limitations
* The parallel version requires a manual setting of the subproblem threshold. For very large slices, this could lead to the creation of thousands of goroutines, resulting in significant overhead and decreased performance. Users must manually specify an appropriate threshold. With a threshold of 256, this implementation outperforms Go's built-in sort for slices exceeding 10,000 elements.
* There are alternative methods to reduce the number of goroutines. These include using external packages or implementing a semaphore-like mechanism with channels. However, these approaches proved challenging to implement, likely due to the recursive nature of the algorithm. Another potential strategy to reduce overhead would be to limit the recursion depth.
* Sorts only integer slices.

### Usage
```golang
s := []int{9, 5, 1, 3, 5, 8, 0, 1, 2}

MergeSort(s) // Regular merge sort in-place.
fmt.Println(s)

// Concurrent version, 256 is the threshold, below which a sequential version will be used
// instead of concurrent, to reduce the number of goroutines running.
MergeSortConcurrent(s, 256) 
fmt.Println(s)
```

