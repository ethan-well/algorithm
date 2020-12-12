package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := generateLink(2,4,3)
	l2 := generateLink(5,6,4)
	r := addTwoNumbers(l1, l2)

	for r != nil {
		println(r.Val)
		r = r.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l1Raw := l1
	var v1, v2, nv int
	var l1End,l2End bool
	for l1 != nil {

		if !l1End {
			v1 = l1.Val
		} else {
			v1 = 0
		}

		if l1.Next == nil {
			l1.Next = l2.Next
			l1End = true
		}

		if !l2End {
			v2 = l2.Val
		} else {
			v2 = 0
		}

		if l2.Next == nil {
			l2.Next = l1.Next
			l2End = true
		}

		sum :=  v1 + v2 + nv

		l1.Val = sum % 10

		// 下一位的进位
		nv = (sum) / 10

		l2 = l2.Next
		l1 = l1.Next
	}

	if nv != 0 {
		node := &ListNode{
			Val: nv,
		}
		l1.Next = node
	}

	return l1Raw
}

func generateLink(nums ...int) *ListNode {
	var head, tail  *ListNode
	for _, n := range nums {
		node := &ListNode{Val:  n}
		if head == nil {
			head = node
			tail = head
		} else {
			tail.Next = node
			tail = tail.Next
		}
	}

	return head
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//
//	var v1, v2, nv int
//
//	for l1 != nil || l2 != nil {
//
//		if l1 != nil {
//			v1 = l1.Val
//			l1 = l1.Next
//		} else {
//			l1.Next = l2.Next
//			v1 = 0
//		}
//
//		if l2 != nil {
//			v2 = l2.Val
//			l2 = l2.Next
//		} else {
//			l2.Next = l1.Next
//			v2 = 0
//		}
//
//		sum :=  v1 + v2 + nv
//
//		node := &ListNode{
//			Val: sum % 10,
//		}
//
//		l1.Val == sum % 10
//
//		// 下一位的进位
//		nv = (sum) / 10
//	}
//
//	if nv != 0 {
//		node := &ListNode{
//			Val: nv,
//		}
//		tail.Next = node
//	}
//
//	return l1
//}
