package breadthfirst

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBreadthFirst(t *testing.T) {
	node7 := Node{7, nil, nil}
	node8 := Node{8, nil, nil}
	node5 := Node{5, &node7, &node8}
	node4 := Node{4, nil, nil}
	node6 := Node{6, nil, nil}
	node2 := Node{2, &node4, &node5}
	node3 := Node{3, nil, &node6}
	node1 := Node{1, &node2, &node3}

	expect := []int{1, 2, 3, 4, 5, 6, 7, 8}

	if result := breadthFirst(node1); reflect.DeepEqual(result, expect) {
		fmt.Println(result)
		t.Log("Succeed!")
	} else {
		t.Error("Failed!")
		t.Errorf("Expect: %v, get: %v", expect, result)
	}
}
