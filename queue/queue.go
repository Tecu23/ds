// Package queue is the simple implementation of a queue. A queue is a first in first out
// data structure.
package queue

type node struct {
	value interface{}
	next  *node
}

// Queue represents the queue type
type Queue struct {
	head, tail *node
	length     int
}

// New should create and return an empty queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

// Len should return the number of items in the queue
func (q *Queue) Len() int {
	return q.length
}

// Enqueue should add an element to the back of the queue
func (q *Queue) Enqueue(value interface{}) {

	n := &node{value, nil}

	if q.length == 0 {
		q.head = n
		q.tail = n
	} else {
		q.tail.next = n
		q.tail = n
	}

	q.length++
}

// Dequeue should remove and return the first element in the queue
func (q *Queue) Dequeue() interface{} {
	if q.length == 0 {
		return nil
	}
	n := q.head

	if q.length == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}

	q.length--

	return n.value
}

// Peek should return the first item in the queue but not remove it
func (q *Queue) Peek() interface{} {
	if q.length == 0 {
		return nil
	}

	return q.head.value
}
