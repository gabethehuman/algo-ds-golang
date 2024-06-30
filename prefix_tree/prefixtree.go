package main

import (
	"fmt"
)

type Node struct {
	children       map[string]*Node
	isCompleteWord bool
	value          string
}

type Trie struct {
	Root *Node
}

// Get a new empty node.
func newNode(value string) *Node {
	return &Node{children: make(map[string]*Node), isCompleteWord: false, value: value}
}

// Get a new empty trie.
func NewEmpty() Trie {
	return Trie{Root: newNode("")}
}

// Insert any number of words into the trie.
func (trie *Trie) Insert(words ...string) error {
	for _, word := range words {
		current := trie.Root
		for _, letter := range word {
			letterStr := string(letter)
			next, exists := current.children[letterStr]
			if exists {
				current = next
			} else {
				newChild := newNode(letterStr)
				current.children[letterStr] = newChild
				current = newChild
			}
		}
		current.isCompleteWord = true
	}
	return nil
}

// Check if a word is in the trie.
func (trie *Trie) Search(word string) bool {
	current := trie.Root
	for _, letter := range word {
		next, exists := current.children[string(letter)]
		if exists {
			current = next
		} else {
			return false
		}
	}
	return current.isCompleteWord
}

// Check if there is any word in the trie that starts with a given prefix.
func (trie *Trie) StartsWith(word string) bool {
	current := trie.Root
	for _, letter := range word {
		next, exists := current.children[string(letter)]
		if exists {
			current = next
		} else {
			return false
		}
	}
	return true
}

// Delete a word from the trie.
func (trie *Trie) Delete(word string) error {
	nodeStack := []*Node{}

	// First, find all the nodes belonging to the word, if such word doesn't exist in the trie, return an error.
	current := trie.Root
	nodeStack = append(nodeStack, current) // Add the Root node manually.

	for _, letter := range word {
		next, exists := current.children[string(letter)]
		if exists {
			current = next
			nodeStack = append(nodeStack, current) // Append to the stack for backtracking later.
		} else {
			return fmt.Errorf("Delete error: word '%v' not found", word)
		}
	}

	// If the word does not mark the end of any word, it's not a complete word in the trie.
	if !current.isCompleteWord {
		return fmt.Errorf("Delete error: word '%v' not found", word)
	}

	// Unmark the end of word
	current.isCompleteWord = false

	// Backtrack through the nodes, and delete all the children references corresponding to a letter.
	for len(nodeStack) > 1 { // Stop when there's 1 node on the stack, this is the Root (value "") node, leave it be.
		// Pop from the stack
		currentNode := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]

		// Check if the current node has any children or if it is marked as a complete word
		// Only delete if the current node has no children and is not a complete word
		if len(currentNode.children) == 0 && !currentNode.isCompleteWord {
			parent := nodeStack[len(nodeStack)-1]
			delete(parent.children, currentNode.value)
		} else {
			break
		}
	}

	return nil
}

// Print the trie row by row.
func (trie *Trie) Print() {
	type NodeInfo struct {
		node   *Node
		parent *Node
		level  int
	}

	queue := []NodeInfo{{node: trie.Root, parent: nil, level: 0}}

	for len(queue) > 0 {
		// Dequeue the first element
		currentInfo := queue[0]
		queue = queue[1:]

		parentValue := "nil"
		if currentInfo.parent != nil {
			parentValue = currentInfo.parent.value
		}

		// Collect children information
		childrenValues := []string{}
		for key := range currentInfo.node.children {
			childrenValues = append(childrenValues, key)
		}

		fmt.Printf("Level: %d, Parent: %s, Value: %s, IsCompleteWord: %t, Children: %v\n",
			currentInfo.level, parentValue, currentInfo.node.value, currentInfo.node.isCompleteWord, childrenValues)

		// Enqueue children nodes
		for _, child := range currentInfo.node.children {
			queue = append(queue, NodeInfo{node: child, parent: currentInfo.node, level: currentInfo.level + 1})
		}
	}
}

func main() {
}
