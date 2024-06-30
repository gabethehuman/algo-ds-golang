## Prefix tree (trie)

### Features
* Trie operations: `Insert`, `StartsWith`, `Search`, `Delete`.
* Helper functions: `Print` to print the trie row by row, with information about the children, the parent and the value.
* Node representation: `Node` contains `isCompleteWord` flag, `value`, which is a letter, and `children` which is a hashmap, where each key is the letter of a child, and value is a pointer to that child node.

### Limitations
* The implementation only supports `string` values.

### Usage
```golang
trie := NewEmpty()

// Insert
words := []string{"apple", "banana", "grape", "apricot", "orange"}
for _, word := range words {
    trie.Insert(word)
}

// Search
found := trie.Search("apple") // true
notFound := trie.Search("strawberry") // false

// StartsWith
startTrue := trie.StartsWith("app") // true
startFalse := trie.StartsWith("add") // false

// Delete and print
err := trie.Delete("grape")
if err == nil {
    trie.Print()
}
```

