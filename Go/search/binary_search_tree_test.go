package search

import (
	"testing"
)

type BST struct {
	Value int64
	Left *BST
	Right *BST
}

func gTree() *BST {
	root := &BST{Value: 41}

	arr := []int64{22, 15, 88, 99, 11, 33, 66, 77}

	for _, a := range arr {
		insert(root, a)
	}

	return root
}

func insert(root *BST, value int64) {
	if root == nil {
		return
	}

	if value == root.Value {
		return
	}

	if value > root.Value  {
		if root.Right == nil {
			root.Right = &BST{Value: value}
		} else {
			insert(root.Right, value)
		}
	} else {
		if root.Left == nil {
			root.Left = &BST{Value: value}
		} else {
			insert(root.Left, value)
		}
	}
}

func TestBinarySearchTree(t *testing.T) {
	root := gTree()

	t.Logf("result: %t", contain(root, 15))

	//inOrderV(t, root)

	levelOrder(t, root)
}

func contain(root *BST, key int64) bool {
	if root == nil {
		return false
	}

	if root.Value == key {
		return true
	} else if root.Value < key {
		return contain(root.Right, key)
	} else {
		return contain(root.Left, key)
	}
}

func PreValue(t *testing.T, root *BST)  {
	if root == nil {
		return
	}

	t.Logf("value: %d\n", root.Value)

	PreValue(t, root.Left)

	PreValue(t, root.Right)
}


func inOrderV(t *testing.T, root *BST)  {
	if root == nil {
		return
	}

	inOrderV(t, root.Left)

	t.Logf("value: %d\n", root.Value)

	inOrderV(t, root.Right)
}

func levelOrder(t  *testing.T, root *BST) {
	queue := make([]*BST,0)

	queue = append(queue, root)

	for len(queue) > 0 {
		q := queue[0]
		t.Logf("value: %d", q.Value)

		queue = queue[1:]

		if q.Left != nil {
			queue = append(queue, q.Left)
		}

		if q.Right != nil {
			queue = append(queue, q.Right)
		}
	}
}