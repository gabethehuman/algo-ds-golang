package main

// ============================
// Queue implementation
// ============================
type Queue[T any] struct {
	elements []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Add an element to the queue.
func (q *Queue[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
}

// Remove and return the first element of the queue.
func (q *Queue[T]) Dequeue() T {
	if len(q.elements) == 0 {
		var zeroValue T
		return zeroValue
	}

	// Memory usage might not be optimal here. The slice queue[1:] does not free
	// the memory used by the first element q[0], but it does not matter for this use case.
	element := q.elements[0]
	q.elements = q.elements[1:]

	return element
}

func (q *Queue[T]) Length() int {
	return len(q.elements)
}

// ============================
// Stack implementation
// ============================
type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push an element to the stack.
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop and return the last element of the stack.
func (s *Stack[T]) Pop() T {
	if len(s.elements) == 0 {
		var zeroValue T
		return zeroValue
	}
	index := len(s.elements) - 1
	popped := s.elements[index]
	s.elements = s.elements[:index]
	return popped
}

func (s *Stack[T]) Length() int {
	return len(s.elements)
}

// Return the last element of the stack without popping it.
func (s *Stack[T]) Peek() T {
	if len(s.elements) == 0 {
		var zeroValue T
		return zeroValue
	}
	return s.elements[len(s.elements)-1]
}
