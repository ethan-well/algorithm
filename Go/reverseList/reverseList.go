package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head_next := head.Next

	new_head := reverseList(head_next)

	head_tem := new_head

	for head_tem.Next != nil {
		head_tem = head_tem.Next
	}
	head_tem.Next = head
	head.Next = nil

	return new_head
}
