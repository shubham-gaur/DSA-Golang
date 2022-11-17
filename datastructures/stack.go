package datastructures

import (
	"fmt"
)

// Stack represents a stack of elements.
type Stack[T any] struct {
	stack []T
}

// NewStack returns a new stack.
func NewStack[T any](initialCapacity ...int) Stack[T] {
	capacity := 512

	if len(initialCapacity) > 0 {
		capacity = initialCapacity[0]
	}

	return Stack[T]{
		stack: make([]T, 0, capacity),
	}
}

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int {
	return len(s.stack)
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

// Push pushes the element onto the stack.
func (s *Stack[T]) Push(value T) {
	s.stack = append(s.stack, value)
}

// Pop removes and returns the top element of the stack.
func (s *Stack[T]) Pop() (zero T, err error) {
	if s.IsEmpty() {
		err = fmt.Errorf("can not perform `Stack.Pop()`, stack is empty")
		return zero, err
	}

	// Get the index of the top most element.
	index := len(s.stack) - 1
	// Index into the slice and obtain the last element.
	elm := s.stack[index]
	// Remove it from the stack by slicing it off.
	s.stack = s.stack[:index]

	return elm, nil
}
