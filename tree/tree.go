// Package tree represent the represantation of a tree data structure
package tree

type elem interface {
	// return -1 if elem is less than b, 1 if more than b, 0 if equal
	Compare(b elem) int
}

type node struct {
	value  elem
	parent *node
	left   *node
	right  *node
}

// BST is the actual type exposed to the user
type BST struct {
	root *elem
}

// New should create and return an empty binary search tree
func New() *BST {
	return &BST{nil}
}

// Insert should insert an element into the tree
func (t *BST) Insert(v elem) {}

// Get should return a tree node given its value
func (t *BST) Find(v elem) elem { return nil }

func (t *BST) Clear() {}

func (t *BST) Height() int { return 0 }

func (t *BST) Len() int {}

func (t *BST) Remove(v elem) elem {}
