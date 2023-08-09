package main

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	key   int
	left  *Node
	right *Node
}

//Tree
func (t *Tree) insertNode(data int) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insertNode(data)
	}
}

//Node
func (n *Node) insertNode(data int) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insertNode(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insertNode(data)
		}
	}
}

func main() {

	var t Tree
	t.insertNode(10)
	t.insertNode(49)
	t.insertNode(37)
	t.insertNode(16)
	t.insertNode(83)
	t.insertNode(79)
	t.insertNode(15)
	t.insertNode(47)
	t.insertNode(72)

	fmt.Println("Preorder: ")
	PreOrder(t.root)
	fmt.Println()

	fmt.Println("Inorder: ")
	InOrder(t.root)
	fmt.Println()

	fmt.Println("Postorder: ")
	PostOrder(t.root)
	fmt.Println()

	fmt.Println("Search the data! isthere? ")
	fmt.Println(t.root.Search(10))

}

func PreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.key)
		PreOrder(n.left)
		PreOrder(n.right)
	}
}

func InOrder(n *Node) {
	if n == nil {
		return
	} else {
		InOrder(n.left)
		fmt.Printf("%d ", n.key)
		InOrder(n.right)
	}
}
func PostOrder(n *Node) {
	if n == nil {
		return
	} else {
		PostOrder(n.left)
		PostOrder(n.right)
		fmt.Printf("%d ", n.key)
	}
}

func (n *Node) Search(val int) bool {
	if n == nil {
		return false
	}
	if n.key == val {
		return true
	}
	if n.key < val {
		return n.right.Search(val)
	} else {
		return n.left.Search(val)
	}

}
