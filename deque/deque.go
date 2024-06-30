package main

import (
	"fmt"
	"slices"
)

const initBufferSize int = 32

type Deque[T any] struct {
	buffer        []T
	head          int // Points to the first element of the queue.
	tail          int // Points to the next empty slot after the last element.
	count         int
	minBufferSize int
}

func NewDeque[T any]() *Deque[T] {
	buffer := make([]T, initBufferSize)
	return &Deque[T]{buffer: buffer, head: 0, tail: 0, count: 0, minBufferSize: initBufferSize}
}

// Capacity of the queue.
func (q *Deque[T]) Cap() int {
	return len(q.buffer)
}

// Number of currently occupied elements in the queue.
func (q *Deque[T]) Len() int {
	return q.count
}

// Check if the queue is empty.
func (q *Deque[T]) IsEmpty() bool {
	return q.count == 0
}

// Print all elements of the queue, starting from the front.
func (q *Deque[T]) Print() {
	if q.tail > q.head {
		fmt.Println(q.buffer[q.head:q.tail])
	} else {
		fmt.Println(slices.Concat(q.buffer[q.head:], q.buffer[:q.tail]))
	}
}

// Double the buffer in size if necessary.
func (q *Deque[T]) resizeIfNeeded() {
	if q.count == len(q.buffer) {
		newBuffer := make([]T, 2*len(q.buffer))

		if q.tail > q.head { // Simpler case when we don't have any wrap-around, copy from head to tail.
			copy(newBuffer, q.buffer[q.head:q.tail])
		} else { // Handle wrap-around by copying from head to end, then start to tail.
			numCopied := copy(newBuffer, q.buffer[q.head:])
			copy(newBuffer[numCopied:], q.buffer[:q.tail])
		}

		q.head = 0
		q.tail = q.count
		q.buffer = newBuffer
	}
}

// Shrink buffer by half if less that 25 percent is occupied.
func (q *Deque[T]) shrinkIfNeeded() {
	if len(q.buffer) > q.minBufferSize && q.count <= len(q.buffer)/4 {

		newSize := max(len(q.buffer)/2, q.minBufferSize)
		newBuffer := make([]T, newSize)

		if q.tail > q.head { // Simpler case when we don't have any wrap-around, copy from head to tail.
			copy(newBuffer, q.buffer[q.head:q.tail])
		} else { // Handle wrap-around by copying from head to end, then start to tail.
			numCopied := copy(newBuffer, q.buffer[q.head:])
			copy(newBuffer[numCopied:], q.buffer[:q.tail])
		}

		q.head = 0
		q.tail = q.count
		q.buffer = newBuffer
	}
}

// Push value to the end of the queue.
func (q *Deque[T]) PushBack(value T) {
	q.resizeIfNeeded()
	q.buffer[q.tail] = value              // Tail points to empty slot after the last element, put the value there.
	q.tail = (q.tail + 1) % cap(q.buffer) // Move tail to next position, with circular wrapping
	q.count++
}

// Push value to the front of the queue.
func (q *Deque[T]) PushFront(value T) {
	q.resizeIfNeeded()
	q.head = (q.head + cap(q.buffer) - 1) % cap(q.buffer) // Head is the first occupied element, so first move it.
	q.buffer[q.head] = value                              // Only then put the value there.
	q.count++
}

// Pop value from the back of the queue.
// Panic if the deque is empty.
func (q *Deque[T]) PopBack() T {
	if q.count <= 0 {
		panic("PopBack panic: deque length is 0.")
	}

	defer q.shrinkIfNeeded()

	q.tail = (q.tail + cap(q.buffer) - 1) % cap(q.buffer) // Tail is next after the last element, first move it.
	value := q.buffer[q.tail]                             // Only then get the value.
	q.count--
	return value
}

// Pop value from the front of the queue.
// Panic if the deque is empty.
func (q *Deque[T]) PopFront() T {
	if q.count <= 0 {
		panic("PopFront panic: deque length is 0.")
	}

	defer q.shrinkIfNeeded()

	value := q.buffer[q.head]             // Head is the first element, so first get the value.
	q.head = (q.head + 1) % cap(q.buffer) // Only then move it.
	q.count--
	return value
}

// Get value from the back of the queue without removing it.
// Panic if the deque is empty.
func (q *Deque[T]) PeekBack() T {
	if q.count <= 0 {
		panic("PeekBack panic: deque length is 0.")
	}
	return q.buffer[(q.tail+cap(q.buffer)-1)%cap(q.buffer)]
}

// Get value from the front of the queue without removing it.
// Panic if the deque is empty.
func (q *Deque[T]) PeekFront() T {
	if q.count <= 0 {
		panic("PeekFront panic: deque length is 0.")
	}
	return q.buffer[q.head]
}

// Get value at the specified index, counting from the front of the deque.
// Panic if the deque is empty, the index is negative, or the index is out of bounds.
func (q *Deque[T]) PeekAt(index int) T {
	if q.count <= 0 {
		panic("PeekAt panic: deque length is 0.")
	}
	if index < 0 {
		panic("PeekAt panic: negative indices not supported.")
	}
	if index >= q.count {
		panic(fmt.Sprintf("PeekAt panic: index %d outside range %d", index, q.count-1))
	}
	bufferIndex := (q.head + index) % len(q.buffer)
	return q.buffer[bufferIndex]
}

// Clear the deque without changing the capacity (buffer length).
// Shrinking may happen later, when we pop elements from the queue.
func (q *Deque[T]) Clear() {
	q.head = 0
	q.tail = 0
	q.count = 0
}

// Get the first index, for which the specified predicate is true.
// Return -1 if there is no such element found.
func (q *Deque[T]) Index(f func(T) bool) int {
	for i := 0; i < q.count; i++ {
		bufferIndex := (q.head + i) % len(q.buffer)
		if f(q.buffer[bufferIndex]) {
			return i
		}
	}
	return -1
}

// Replace value at the specified index with a new value.
// Panic if the deque is empty, the index is negative, or the index is out of bounds.
func (q *Deque[T]) SetAt(index int, value T) {
	if q.count <= 0 {
		panic("SetAt panic: deque length is 0.")
	}
	if index < 0 {
		panic("SetAt panic: negative indices not supported.")
	}
	if index >= q.count {
		panic(fmt.Sprintf("SetAt panic: index %d outside range %d", index, q.count-1))
	}
	bufferIndex := (q.head + index) % len(q.buffer)
	q.buffer[bufferIndex] = value
}

func main() {
}
