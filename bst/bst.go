// Package bst represent the represantation of a binary search tree data structure
package bst

import "fmt"

type elem interface {
	Less(than elem) bool
}

type node struct {
	value  elem
	parent *node
	left   *node
	right  *node
}

// BST is the actual type exposed to the user
type BST struct {
	root *node
}

// New should create and return an empty binary search tree
func New() *BST {
	return &BST{nil}
}

// Insert should insert an element into the tree
func (t *BST) Insert(v elem) {
	n := node{
		value:  v,
		parent: nil,
		left:   nil,
		right:  nil,
	}

	if t.root == nil {
		t.root = &n
	} else {
		curr := t.root
		var parent *node
		for curr != nil {
			parent = curr
			if v.Less(curr.value) {
				curr = curr.left
			} else {
				curr = curr.right
			}
		}
		if v.Less(parent.value) {
			n.parent = parent
			parent.left = &n

		} else {
			n.parent = parent
			parent.right = &n
		}
	}
}

// Find should return a tree node given its value
func (t *BST) Find(v elem) elem {
	n := t.root.search(v)

	if n == nil {
		return nil
	}

	return n.value
}

// Min should return the minimum value in the tree
func (t *BST) Min() elem {
	curr := t.root

	if curr == nil {
		return nil
	}

	for curr.left != nil {
		curr = curr.left
	}

	return curr.value
}

// Max should return the minimum value in the tree
func (t *BST) Max() elem {
	curr := t.root

	if curr == nil {
		return nil
	}

	for curr.right != nil {
		curr = curr.right
	}

	return curr.value
}

// Clear should clear the tree
func (t *BST) Clear() {}

// Height should return the height of the tree
func (t *BST) Height() int { return 0 }

// Len should return the length of the tree (length of nodes)
func (t *BST) Len() int { return 0 }

// Remove should remove an element from the tree
func (t *BST) Remove(v elem) elem { return nil }

/*


# PRIVATE METHODS


*/

func (n *node) search(v elem) *node {
	if n == nil {
		return nil
	}
	if n.value == v {
		return n
	}
	if v.Less(n.value) {
		return n.left.search(v)
	}

	return n.right.search(v)
}

func preorder(node *node) {
	if node == nil {
		return
	}
	fmt.Println(node.value)
	inorder(node.left)
	inorder(node.right)
}

func inorder(node *node) {
	if node == nil {
		return
	}
	inorder(node.left)
	fmt.Println(node.value)
	inorder(node.right)
}

func postorder(node *node) {
	if node == nil {
		return
	}
	inorder(node.left)
	inorder(node.right)
	fmt.Println(node.value)
}
