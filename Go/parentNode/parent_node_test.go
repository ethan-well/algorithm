package parentNode


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parentMap := make(map[int]*TreeNode)
	var dfsFunc func(*TreeNode)

	dfsFunc = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			parentMap[root.Left.Val] = root
			dfsFunc(root.Left)
		}

		if root.Right != nil {
			parentMap[root.Right.Val] = root
			dfsFunc(root.Right)
		}
	}

	dfsFunc(root)

	searched := make(map[int]struct{})
	for p != nil {
		searched[p.Val] = struct{}{}
		p = parentMap[p.Val]
	}

	for q != nil {
		if _, ok := searched[q.Val]; ok {
			return parentMap[q.Val]
		}

		q = parentMap[q.Val]
	}

	return nil
}
