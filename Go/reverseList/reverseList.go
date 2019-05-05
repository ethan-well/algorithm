package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// fmt.Printf("head, %d", head.Val)

	if head.Next == nil {
		return head
	}

	head_next := head.Next
	// fmt.Printf("value: %d, opt: %d\n", (*head_next).Val, (*head_next).Next)
	// fmt.Println("reverseList(head_next)", reverseList(head_next))

	new_head := reverseList(head_next)

	for new_head.Next != nil {
		new_head = new_head.Next
	}
	new_head.Next = head

	head.Next = nil
	return head_next
}

func main() {
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

	for node := &listNode4; node != nil; node = (*node).Next {
		fmt.Println((*node).Val)
	}

	result := reverseList(&listNode4)

	for node := result; node != nil; node = (*node).Next {
		fmt.Println((*node).Val)
	}
}
