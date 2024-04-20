// Package queue represent a queue
package queue

const minLen = 16

// ArrQueue is a representaion of the queue as an array
type ArrQueue struct {
	arr                []*interface{}
	head, tail, length int
}

// ArrNew should create a newQ
func ArrNew() *ArrQueue {
	return &ArrQueue{
		arr: make([]*interface{}, minLen),
	}
}

// ArrLen should return the length of the q
func (q *ArrQueue) ArrLen() int {
	return q.length
}

func (q *ArrQueue) resize() {
	newArr := make([]*interface{}, q.length<<1)

	if q.tail > q.head {
		copy(newArr, q.arr[q.head:q.tail])
	} else {
		n := copy(newArr, q.arr[q.head:q.tail])
		copy(newArr[n:], q.arr[q.head:q.tail])
	}

	q.head = 0
	q.tail = q.length
	q.arr = newArr
}

// ArrEnqueue should add an element to the start of the queue
func (q *ArrQueue) ArrEnqueue(v interface{}) {
	if q.length == len(q.arr) {
		q.resize()
	}

	q.arr[q.tail] = &v

	q.tail = (q.tail + 1) & (len(q.arr) - 1)
	q.length++
}

// ArrPeek should return the first element in the queue but not remove it
func (q *ArrQueue) ArrPeek() interface{} {
	if q.length <= 0 {
		return nil
	}

	return (q.arr[q.head])
}

// ArrDequeue should remove & return the first element in the queue
func (q *ArrQueue) ArrDequeue() interface{} {
	if q.length <= 0 {
		return nil
	}

	n := q.arr[q.head]
	q.arr[q.head] = nil

	q.head = (q.head + 1) & (len(q.arr) - 1)
	q.length--

	if len(q.arr) > minLen && (q.length<<2) == len(q.arr) {
		q.resize()
	}

	return n
}
