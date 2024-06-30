## Union Find (Disjoint Set)

### Features

* Array-based storage: Uses arrays to keep track of parent elements and sizes of the sets.
* Union by size: `Union` connects two sets by attaching the smaller tree to the root of the larger tree.
* Path Compression: Flattens the structure of the tree whenever `Find` is called, ensuring very efficient subsequent operations.
* Union Find functions:
    - `NewUnionFind(n int)` creates a disjoint set with `n` elements.
    - `Find(k int)` finds a representative of an element. Iterative and recursive implementations are provided.
    - `NewSet()` creates a new set with 1 element and adds it to `UnionFind`.
    - `Union(k int, l int)` merges sets which `k` and `l` belong to.

### Limitations

* The implementation supports only `int` values for elements.

### Usage
```golang
u := NewUnionFind(0) // Initialize an empty Union-Find structure

// Add new sets (you could use NewUnionFind(3) for the same effect)
u.NewSet()
u.NewSet()
u.NewSet()

// Union operations
u.Union(0, 1)
u.Union(1, 2)

// Find representative of an element
rep := u.Find(2)
```