package reverseList

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	var listNode1 ListNode
	var listNode2 ListNode
	var listNode3 ListNode
	var listNode4 ListNode

	listNode1.Val = 1
	listNode1.Next = nil
	listNode2.Val = 2
	listNode2.Next = &listNode1
	listNode3.Val = 3
	listNode3.Next = &listNode2
	listNode4.Val = 4
	listNode4.Next = &listNode3

	result := reverseList(&listNode4)

	for result.Next != nil {
		result = result.Next
	}

	if result.Val != 4 {
		t.Errorf("everseList(&listNode4) get %d, expected 4", result.Val)
	}
}
