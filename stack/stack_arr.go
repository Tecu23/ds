// Package stack is the simple implementation of a stack
package stack

const minLen = 16

// ArrStack type is a linked list representation of a stack
type ArrStack struct {
	arr         []interface{}
	top, length int
}

// ArrNew should create a new stack
func ArrNew() *ArrStack {
	return &ArrStack{
		arr: make([]interface{}, minLen),
	}
}

// ArrLen should return the length of a stack
func (s *ArrStack) ArrLen() int {
	return s.length
}

func (s *ArrStack) resize() {
	newArr := make([]interface{}, s.length<<1)

	copy(newArr, s.arr[:s.top])

	s.top = s.length
	s.arr = newArr
}

// ArrPush should add an element to the stack
func (s *ArrStack) ArrPush(v interface{}) {
	if s.length == len(s.arr) {
		s.resize()
	}

	s.arr[s.top] = v

	s.top = (s.top + 1) & (len(s.arr) - 1)
	s.length++
}

// ArrPop should remove the top element of the stack and return its value
func (s *ArrStack) ArrPop() interface{} {
	if s.length <= 0 {
		return nil
	}

	n := s.arr[s.top]
	s.arr[s.top] = nil
	s.top = (s.top + 1) & (len(s.arr) - 1)
	s.length--
	return n
}

// ArrPeek should return the value of the top element in the stack
func (s *ArrStack) ArrPeek() interface{} {
	if s.length <= 0 {
		return nil
	}

	return s.arr[s.top]
}
