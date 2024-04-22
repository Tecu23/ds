package bst

import (
	"fmt"
	"testing"
)

type test int

func (v test) Less(than elem) bool {
	thanT, ok := than.(test)

	if !ok {
		fmt.Errorf("%v is not of type test", than)
	}

	if v < thanT {
		return true
	}

	return false
}

func TestInsert(t *testing.T) {
	tree := New()

	arr := []test{10, 5, 3, 2, 5, 11, 23, 2}

	for i := range arr {
		tree.Insert(arr[i])
	}

	inorder(tree.root)
}
