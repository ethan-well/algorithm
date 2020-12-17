package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l := generateList()

	l = reverseKGroup(l, 8)

	for l != nil {
		println(l.Val)
		l = l.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	if head == nil || head.Next == nil || k <= 1 {
		return head
	}

	var headBackup, preTails *ListNode
	nodeArr := make([]*ListNode, k)
	var count int

	for head != nil {

		subLen := 0
		for ; subLen < k; subLen ++ {
			if head == nil {
				break
			}

			nodeArr[subLen] = head

			head = head.Next
		}

		count ++

		println("count", count)
		for i := 0; i < subLen; i ++ {
			println("value: ", nodeArr[i].Val)
		}
		println("---------")


		if subLen < k && count == 1 {
			return nodeArr[0]
		}

		if subLen == k {
			for i := k-1; i > 0; i -- {
				nodeArr[i].Next = nodeArr[i-1]
			}
			nodeArr[0].Next = nil

			if count == 1 {
				headBackup = nodeArr[k-1]
				preTails = nodeArr[0]
			} else {
				preTails.Next = nodeArr[k-1]
			}

			preTails = nodeArr[0]

		} else {
			preTails.Next = nodeArr[0]
		}
	}

	return headBackup
}

func generateList() *ListNode {
	l1 := &ListNode{Val:  1}
	l2 := &ListNode{Val:  2}
	l3 := &ListNode{Val:  3}
	l4 := &ListNode{Val:  4}
	//l5 := &ListNode{Val:  5}
	//
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	//l4.Next = l5

	return l1
}
