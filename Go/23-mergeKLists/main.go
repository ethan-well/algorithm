package main

func main() {
	l := mergeKLists(generateLists())
	for l != nil {
		println(l.Val)
		l = l.Next
	}
}

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	resultList := &ListNode{Val: 0}
	resultListBack := resultList

	for len(lists) != 0 {
		if lists[0] == nil {
			lists = append(lists[1:])
			continue
		}

		minNode := lists[0].Val
		minNodeIndex := 0
		for i := 0; i < len(lists); i ++ {
			if lists[i] == nil {
				if i + 1 < len(lists) {
					lists = append(lists[0:i+1], lists[i+1:]...)
				} else {
					lists = append(lists[:i])
				}

				continue
			}

			if lists[i].Val <= minNode {
				minNode = lists[i].Val
				minNodeIndex = i
			}
		}

		resultList.Next =  lists[minNodeIndex]
		resultList = resultList.Next

		lists[minNodeIndex] = lists[minNodeIndex].Next
		if lists[minNodeIndex] == nil {
			if minNodeIndex + 1 < len(lists) {
				lists = append(lists[:minNodeIndex], lists[(minNodeIndex+1):]...)
			} else {
				lists = append(lists[:minNodeIndex])
			}
		}
	}

	return resultListBack.Next
}


func generateLists() []*ListNode {
	l1 := &ListNode{Val:  1}
	l2 := &ListNode{Val:  4}
	l3 := &ListNode{Val:  5}
	l1.Next = l2
	l2.Next = l3

	l21 := &ListNode{Val:  1}
	l22 := &ListNode{Val:  3}
	l23 := &ListNode{Val:  4}
	l21.Next = l22
	l22.Next = l23


	return []*ListNode{l1, l21}
}
