package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func preOrderNoRecursive(root *Node) []int {
	var result []int
	nodes := []*Node{}
	node := root
	for node != nil || len(nodes) != 0 {
		if node != nil {
			result = append(result, node.Data)
			nodes = append(nodes, node)
			node = node.Left
		} else {
			node = nodes[len(nodes)-1].Right
			nodes = nodes[:len(nodes)-1]
		}

	}
	return result
}

func main() {
	node7 := Node{7, nil, nil}
	node8 := Node{8, nil, nil}
	node5 := Node{5, &node7, &node8}
	node4 := Node{4, nil, nil}
	node6 := Node{6, nil, nil}
	node2 := Node{2, &node4, &node5}
	node3 := Node{3, nil, &node6}
	node1 := Node{1, &node2, &node3}

	fmt.Println(preOrderNoRecursive(&node1))
}
