package redblacktree

var Red = true
var Black = false

type RBTNode struct {
	Color bool
	Value int64
	Key int64

	Parent *RBTNode
	Left *RBTNode
	Right *RBTNode
}

func (n *RBTNode) ISRed() bool {
	return n.Color == Red
}


// 右旋
/*
				  p							    p
		         / \						   / \
		        n   ...						   l  ...
		    /      \						 /  \
		   l         r						ll    n
		 /  \	   /  \							/  \
        ll   lr   rl   rr                      lr    r
												   /  \
												  rl  rr
 */
func rightRotate(n *RBTNode) {
	if n == nil {
		return
	}

	if n.Left == nil {
		return
	}

	tmpN := n

	// 移动 n 的 left 节点（上图的 l 结点）
	n.Left.Parent = n.Parent
	if n.Parent != nil {

		if n.Parent.Left == n {
			n.Parent.Left = n.Left
		}

		if n.Parent.Right ==  nil {
			n.Parent.Right = n.Left
		}
	}

	// 移动 n 节点
	n.Left = n.Left.Right // lr 移动到 n 的右节点
	n.Parent = tmpN.Left //  n 的父节点指向 l
	n.Parent.Right = n // l 的 right 节点指向 n
}

// 左旋
/*
				    p							  p
		           / \						     / \
				  n   ...						r  ...
			   /    \						   / \
		      l      r						  n   rr
		     / \	/ \						 / \
            ll  lr rl  rr                    l  rl
											/ \
											ll lr
*/
func leftRotate(n *RBTNode) {
	if n == nil {
		return
	}

	if n.Right == nil {
		return
	}

	tmpN := n

	n.Right.Parent = n.Parent
	if n.Parent != nil {
		if n.Parent.Left != nil {
			n.Parent.Left = n.Right
		}

		if n.Parent.Right != nil {
			n.Parent.Right = n.Right
		}
	}

	n.Right = n.Right.Left
	n.Parent = tmpN.Right
	n.Parent.Left = n
}

// 插入节点

// 空树
//插入第一个节点，插入节点红色，插入后变色为黑丝

// 插入到 1 节点
// 插入一个节点，节点变成了二节点，插入的节点变成红色，或者左旋右旋变成


// 插入到 2 节点
// 插入一个节点，节点变成了三节点, 结点需要变颜色

// 插入到 3 节点
// 插入节点后，节点变成了四节点，此时节点需要裂变

func put(root *RBTNode, value int64) {
	n := newNode(value)
	if root == nil {
		root = n
		return
	}

	parent := root.Parent
	for root != nil {
		parent = root

		// 已经存在相同的值，不做任何操作
		// 这里没有考虑 key 相同，value 不相同的情况，默认只有 value
		if value == root.Value {
			return
		}

		if value < root.Value {
			root = root.Left
			continue
		}

		if value > root.Value {
			root = root.Right
			continue
		}
	}

	node := newNode(value)
	node.Parent = parent
	if value < parent.Value {
		parent.Left = n
	} else {
		parent.Right = n
	}

	fixAfterPut(node);
}

func fixAfterPut(node *RBTNode) {
	if node == nil {
		return
	}

	node.Color = Red

	// 树只有一个元素, 这个结点就是跟节点，节点改变成黑色
	if node.Parent == nil {
		node.Color = Black
	}

	// 1:
	// 2-3-4：原本节点是一个 2 节点(节点本身一个元素)，新元素加入后，节点变成了一个 3 节点（节点变成了 2 个元素）
	// 红黑树：原本父节点是黑节点，新增一个红节点 => 上黑下红节点 --- 仍然是平衡的红黑树，不用调整

	// 2:
	// 2-3-4：原本节点是一个 3 节点（节点本身两个元素），新元素加入后，节点变成了一个 4 节点（节点变成了 3 个元素）
	// 插入位置，左1，右1，中间： 2,3 + 4 => 2,3，4 => 黑红红
	// 2,3 + 1 => 1,2,3 红黑红 => 不用调整
	// 2，3 + 2.5 => 2,2.5,3 => 黑红红
	// 红黑树：原本节点为 黑红， 加入一个新节点后变成了 黑红红，红黑树不平衡，此时需要调整为 红黑红，也就是排序后，出现两边红，的情况需要调整，
	// 中间黑，两边红不用调整
	if node.Parent.ISRed() {

	}



	// 3:
	// 2-3-4: 新增一个节点到 4 节点，节点需要分裂，中间节点升为父亲节点，新节点和左节点或者右节点合并
	// 红黑树：新增黑色节点 + 爷爷节点黑，父节点和叔叔节点红色 => 调整爷爷节点变红，父亲和叔叔节点变黑，爷爷节点是根节点改颜色为黑
}


func newNode(value int64) *RBTNode {
	return &RBTNode{
		Color:  Red,
		Value:  value,
		Key:    value,
		Parent: nil,
		Left:   nil,
		Right:  nil,
	}
}

