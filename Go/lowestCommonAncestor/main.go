package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := GenerateTree()

	p := &TreeNode{
		Val:   5,
	}

	q := &TreeNode{
		Val:   4,
	}

	node := lowestCommonAncestor(root, p, q)
	println("lowestCommonAncestor value: ", node.Val)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var pAncestors, qAncestors []*TreeNode

	deepSearch(root, p, &pAncestors)

	deepSearch(root, q, &qAncestors)

	i := searchLowerCommonAncestor(pAncestors, qAncestors)

	return pAncestors[len(pAncestors)-i]
}

func deepSearch(root, p *TreeNode, parents *[]*TreeNode) bool {
	if root == nil {
		return false
	}

	// 比较根节点
	if root.Val == p.Val {
		*parents = append(*parents, root)
		return true
	}

	result := deepSearch(root.Left, p, parents)
	if result {
		*parents = append(*parents, root)
		return true
	}

	result = deepSearch(root.Right, p, parents)
	if result {
		*parents = append(*parents, root)
		return true
	}

	return false
}

func GenerateTree() *TreeNode {
	node0 := &TreeNode {
		Val: 0,
	}

	node1 := &TreeNode {
		Val: 1,
	}

	node2 := &TreeNode {
		Val: 2,
	}

	node3 := &TreeNode {
		Val: 3,
	}

	node4 := &TreeNode {
		Val: 4,
	}

	node5 := &TreeNode {
		Val: 5,
	}

	node6 := &TreeNode {
		Val: 6,
	}

	node7 := &TreeNode {
		Val: 7,
	}

	node8 := &TreeNode {
		Val: 8,
	}

	node3.Left = node5
	node3.Right = node1

	node5.Left = node6
	node5.Right = node2

	node1.Left = node0
	node1.Right = node8

	node2.Left = node7
	node2.Right = node4

	return node3
}

func searchLowerCommonAncestor(pAncestors, qAncestors []*TreeNode) int {
	lowerLength :=  len(pAncestors)
	if len(pAncestors) > len(qAncestors) {
		lowerLength = len(qAncestors)
	}
	
	for i := 1; i <= lowerLength; i ++ {
		if pAncestors[len(pAncestors)-i] != qAncestors[len(qAncestors)-i] {
			return i-1
		}

		if i == lowerLength {
			return lowerLength
		}
	}

	return -1
}
