package traversal

import (
	"reflect"
	"testing"
)

func TestTraversal(t *testing.T) {
	node7 := Node{7, nil, nil}
	node8 := Node{8, nil, nil}
	node5 := Node{5, &node7, &node8}
	node4 := Node{4, nil, nil}
	node6 := Node{6, nil, nil}
	node2 := Node{2, &node4, &node5}
	node3 := Node{3, nil, &node6}
	node1 := Node{1, &node2, &node3}

	expect := []int{4, 2, 7, 5, 8, 1, 3, 6}
	var result []int
	traversal(&node1, &result)
	if reflect.DeepEqual(result, expect) {
		t.Log("Succeed!")
	} else {
		t.Errorf("Failed, expected: %v, get: %v", expect, result)
	}
}
