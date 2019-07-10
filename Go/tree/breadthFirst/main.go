package breadthfirst

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func breadthFirst(node Node) []int {
	var result []int
	var nodes = []Node{node}

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		result = append(result, node.Data)
		if node.Left != nil {
			nodes = append(nodes, *node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, *node.Right)
		}
	}
	return result
}
