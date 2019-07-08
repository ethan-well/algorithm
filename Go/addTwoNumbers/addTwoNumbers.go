package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nodeL1 := ListNode{3, nil}
	nodeL2 := ListNode{4, &nodeL1}
	nodeL3 := ListNode{2, &nodeL2}

	nodeR1 := ListNode{4, nil}
	nodeR2 := ListNode{6, &nodeR1}
	nodeR3 := ListNode{5, &nodeR2}

	// fmt.Printf("%v, %v", nodeL3, nodeR3)

	newNode := addTwoNumbers(&nodeL3, &nodeR3)
	fmt.Printf("%v", newNode)

	for newNode != nil {
		fmt.Printf("%d \n", newNode.Val)
		newNode = newNode.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var new_node, pre_node, first_node ListNode
	var first_node_ptr *ListNode
	fmt.Printf("firs noed %v", first_node)
	var value, to_next, current, vl, vr int

	for l1 != nil || l2 != nil {
		fmt.Printf("l1 = %v\n", l1)
		fmt.Printf("l2 = %v\n", l2)

		if l1 == nil {
			vl = 0
		} else {
			vl = l1.Val
		}

		if l2 == nil {
			vr = 0
		} else {
			vr = l2.Val
		}

		value = vl + vr + to_next
		to_next = value / 10
		current = value % 10
		fmt.Printf("value = %d, to_next = %d, current = %d\n", value, to_next, current)
		new_node = ListNode{current, nil}
		fmt.Printf("new_node: %v\n", new_node)

		if first_node_ptr == nil {
			fmt.Printf("xxxxxxxxxxxx: %v", first_node_ptr)
			first_node_ptr = &new_node
		}
		pre_node.Next = &new_node
		pre_node = new_node
		l1 = l1.Next
		l2 = l2.Next
	}

	return first_node_ptr
}
