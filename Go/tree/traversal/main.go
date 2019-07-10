package traversal

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func traversal(root *Node, result *[]int) {
	if root != nil {
		traversal(root.Left, result)
		*result = append(*result, root.Data)
		traversal(root.Right, result)
	}
}
