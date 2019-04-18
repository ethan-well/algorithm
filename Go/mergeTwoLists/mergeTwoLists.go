package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{}
	currentNode := &result
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			currentNode.Next = l2
			l2 = l2.Next
		} else {
			currentNode.Next = l1
			l1 = l1.Next
		}

		currentNode = currentNode.Next
	}

	if l1 == nil {
		currentNode.Next = l2
	}

	if l2 == nil {
		currentNode.Next = l1
	}

	return result.Next
}

func main() {
	l1Node1 := ListNode{4, nil}
	l1Node2 := ListNode{2, &l1Node1}
	l1Node3 := ListNode{1, &l1Node2}

	l2Node1 := ListNode{4, nil}
	l2Node2 := ListNode{3, &l2Node1}
	l2Node3 := ListNode{1, &l2Node2}

	nodeNumber := mergeTwoLists(&l1Node3, &l2Node3)

	for node := nodeNumber; node != nil; node = node.Next {
		fmt.Printf("%d ", (*node).Val)
	}
}
