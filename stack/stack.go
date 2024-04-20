// Package stack is the simple implementation of a stack
package stack

type node struct {
	value interface{}
	prev  *node
}

// Stack type is a linked list representation of a stack
type Stack struct {
	top    *node
	length int
}

// New should create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Len should return the length of a stack
func (s *Stack) Len() int {
	return s.length
}

// Push should add an element to the stack
func (s *Stack) Push(v interface{}) {
	n := &node{v, s.top}

	s.top = n
	s.length++
}

// Pop should remove the top element of the stack and return its value
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

// Peek should return the value of the top element in the stack
func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}

	return s.top.value
}
