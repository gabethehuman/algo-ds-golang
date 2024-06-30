package main

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	Next  *Node[T]
	Prev  *Node[T]
	Value T
}

type DoublyLinkedList[T any] struct {
	First  *Node[T]
	Last   *Node[T]
	Length uint
}

func NewList[T any]() *DoublyLinkedList[T] {
	l := DoublyLinkedList[T]{}
	return &l
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{Value: value}
}

// Pass list first to last and print the content.
func (list *DoublyLinkedList[T]) PrintList() {

	if list.First == nil {
		fmt.Println("empty")
		return
	}

	var builder strings.Builder

	for n := list.First; n != nil; n = n.Next {
		builder.WriteString(fmt.Sprintf("%v ", n.Value))
	}

	// Print the concatenated string of elements
	fmt.Println(strings.TrimSpace(builder.String()))
}

func (list *DoublyLinkedList[T]) InsertFirst(value T) {
	node := NewNode(value)
	if list.First == nil {
		// Empty list case.
		list.First = node
		list.Last = node
	} else {
		node.Next = list.First
		list.First.Prev = node
		list.First = node
	}
	list.Length++
}

func (list *DoublyLinkedList[T]) InsertLast(value T) {
	node := NewNode(value)
	if list.Last == nil {
		// Empty list case.
		list.First = node
		list.Last = node
	} else {
		node.Prev = list.Last
		list.Last.Next = node
		list.Last = node
	}
	list.Length++
}

func (list *DoublyLinkedList[T]) DeleteFirst() {
	if list.First == nil {
		// Empty list
		return
	} else if list.First == list.Last {
		// List with 1 element
		list.First = nil
		list.Last = nil
	} else {
		// List with 2 or more elements
		list.First = list.First.Next
		list.First.Prev = nil
	}
	list.Length--
}

func (list *DoublyLinkedList[T]) DeleteLast() {
	if list.Last == nil {
		// Empty list
		return
	} else if list.First == list.Last {
		list.First = nil
		list.Last = nil
	} else {
		list.Last = list.Last.Prev
		list.Last.Next = nil
	}
	list.Length--
}

// Insert value at a specified index thats within the length of the array.
// Return an error if index out of bounds.
func (list *DoublyLinkedList[T]) InsertAt(index uint, value T) error {
	if index > list.Length {
		return fmt.Errorf("InsertAt error: index out of bounds")
	}
	if index == 0 {
		list.InsertFirst(value)
		return nil
	}
	if index == list.Length {
		list.InsertLast(value)
		return nil
	}

	newNode := NewNode(value)

	// Navigate to the node just before the insertion point.
	current := list.First
	for i := uint(0); i < index-1; i++ {
		current = current.Next
	}

	// Adjust pointers to insert newNode.
	newNode.Next = current.Next
	if current.Next != nil { // Safeguard for the scenario when inserting before the last element.
		current.Next.Prev = newNode
	}
	newNode.Prev = current
	current.Next = newNode

	list.Length++
	return nil
}

func main() {
}
