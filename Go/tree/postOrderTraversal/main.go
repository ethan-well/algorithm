package main

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func postOrderTraversaml(node *Node) {
	nodes := []*Node{}
	for node != nil || len(nodes) != 0 {

	}
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

	postOrderTraversaml(&node1)
}
