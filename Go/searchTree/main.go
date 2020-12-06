package main

// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func main() {
	root := GenerateTree()

	println(kthSmallest(root, 4))
}


func kthSmallest(root *TreeNode, k int) int {
	var n = 0
	_, result := deepSearchTree(root, k, &n)
	return result
}

//func deepSearchTree(root *TreeNode, resultArr *[]int, k int) {
//	if root.Left != nil {
//		deepSearchTree(root.Left, resultArr, k)
//	}
//
//	if len(*resultArr) == k {
//		return
//	}
//
//	*resultArr = append(*resultArr, root.Val)
//
//	if root.Right != nil {
//		deepSearchTree(root.Right, resultArr, k)
//	}
//}

func deepSearchTree(root *TreeNode, k int, n *int) (bool, int)  {
	if root.Left != nil {
		bool, v := deepSearchTree(root.Left, k, n)
		if bool {
			return true, v
		}
	}

	*n ++
	if *n == k {
		return true, root.Val
	}

	if root.Right != nil {
		bool, v := deepSearchTree(root.Right, k, n)
		if bool {
			return true, v
		}
	}

	return false, root.Val
}

func GenerateTree() *TreeNode {
	node1 := &TreeNode {
		Val: 5,
	}

	node2 := &TreeNode {
		Val: 3,
	}

	node3 := &TreeNode {
		Val: 6,
	}

	node4 := &TreeNode {
		Val: 2,
	}

	node5 := &TreeNode {
		Val: 4,
	}

	node6 := &TreeNode {
		Val: 1,
	}


	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node4.Left = node6

	return node1
}