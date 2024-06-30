package main

// Union find (disjoint set) data structure.
type UnionFind struct {
	parents []int // An array, where parents[i] is the parent of the i-th element.
	sizes   []int // To keep track of the sizes.
	numSets int   // Number of disjoint sets.
}

func NewUnionFind(num int) UnionFind {
	parents := make([]int, num)
	sizes := make([]int, num)

	for i := 0; i < num; i++ {
		parents[i] = i // At the beginning, each element is its own parent.
		sizes[i] = 1
	}
	return UnionFind{parents: parents, sizes: sizes, numSets: num}
}

// Add a new set with a single element.
func (u *UnionFind) NewSet() {
	u.parents = append(u.parents, len(u.parents))
	u.sizes = append(u.sizes, 1)
	u.numSets += 1
}

// Find the representative of the set `k` belongs to. Apply path compression along the way.
// Uses additional space for recursion stack, but it's simpler than iterative version.
func (u *UnionFind) Find(k int) int {
	if u.parents[k] != k {
		u.parents[k] = u.Find(u.parents[k]) // Path compression.
	}
	return u.parents[k]
}

// Iterative version of `Find`, doesn't require additional space for the stack, but it requires
// two passes from `k` to the root (or representative) of the set.
func (u *UnionFind) FindIterative(k int) int {
	root := k
	// Find the root of the tree.
	for root != u.parents[root] {
		root = u.parents[root]
	}
	// Path compression: go once again node by node, make every node point directly to the root.
	for k != root {
		parent := u.parents[k]
		u.parents[k] = root
		k = parent
	}
	return root
}

// Connect two sets which `k` and `l` belong to. The bigger set always becomes the parent.
// This is so called union-by-size, as opposed to union-by-rank.
func (u *UnionFind) Union(k int, l int) {
	root1 := u.Find(k)
	root2 := u.Find(l)

	if root1 == root2 {
		return
	}

	if u.sizes[root1] >= u.sizes[root2] {
		u.parents[root2] = root1
		u.sizes[root1] += u.sizes[root2]
	} else {
		u.parents[root1] = root2
		u.sizes[root2] += u.sizes[root1]
	}
	u.numSets -= 1
}

func main() {
}
