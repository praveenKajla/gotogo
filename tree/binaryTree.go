package tree

import (
	"fmt"
	"io"
	"os"
)

type Node struct {
	left  *Node
	right *Node
	data  int64
}

type BinaryTree struct {
	root *Node
}

func (t *BinaryTree) insert(data int64) *BinaryTree {
	if t.root == nil {
		t.root = &Node{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t

}
func (n *Node) insert(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &Node{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}

}

func print(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

//BinaryTreeExample " "
func BinaryTreeExample() {
	tree := &BinaryTree{}
	tree.insert(100).insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)
	print(os.Stdout, tree.root, 0, 'M')

}
