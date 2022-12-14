package tree

import (
	"fmt"
	"strconv"
)

type Tree struct {
	Root *Node
}

type Node struct {
	key   int64
	left  *Node
	right *Node
}

func (t *Tree) InsertRoot(data int64) {

	s := &Node{key: data}
	if t.Root == nil {
		t.Root = s
	} else {
		t.Root.InsertNode(data)
	}
}
func (n *Node) InsertNode(data int64) {

	s := &Node{key: data}
	if data <= n.key {
		if n.left == nil {
			n.left = s
		} else {
			n.left.InsertNode(data)
		}
	} else {
		if n.right == nil {
			n.right = s
		} else {
			n.right.InsertNode(data)
		}
	}
}

func Search(n *Node, data int64) string {
	if n == nil {
		fmt.Println("false")
		return "false"
	}
	if n.key == data {
		fmt.Println("true")
		return "true"
	}
	if n.key < data {
		return Search(n.right, data)
	}
	if n.key > data {
		return Search(n.left, data)
	}
	return "false"
}
func minValued(n *Node) *Node {
	temp := n
	for nil != temp && temp.left != nil {
		temp = temp.left
	}
	return temp
}
func Delete(n *Node, data int64) *Node {

	if nil == n {
		return n
	}
	if n.key > data {
		n.left = Delete(n.left, data)
	}
	if n.key < data {
		n.right = Delete(n.right, data)
	}
	if n.key == data {
		if n.left == nil && n.right == nil {
			n = nil
			return n
		}
		if n.left == nil && n.right != nil {
			temp := n.right
			n = nil
			n = temp
			return n
		}
		if n.left != nil && n.right == nil {
			temp := n.left
			n = nil
			n = temp
			return n
		}

		tempNode := minValued(n.right)
		n.key = tempNode.key
		n.right = Delete(n.right, tempNode.key)
	}
	return n
}
func PrintPreOrder(n *Node, s *string) string {
	if n == nil {
		return ""
	} else {
		x := strconv.FormatInt(n.key, 10)
		*s += x
		*s += " "
		fmt.Printf("%d ", n.key)
		PrintPreOrder(n.left, s)
		PrintPreOrder(n.right, s)

	}
	return *s
}
func NewTree() Tree {
	tree := Tree{}
	elements := [9]int64{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, element := range elements {
		tree.InsertRoot(element)
	}
	return tree
}

func printInOrder(n *Node) {
	if n == nil {
		return
	} else {
		printInOrder(n.left)
		fmt.Printf("%d ", n.key)
		printInOrder(n.right)
	}
}

func printPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Printf("%d ", n.key)
	}
}
