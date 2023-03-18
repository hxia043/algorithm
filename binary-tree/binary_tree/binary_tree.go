package main

import "fmt"

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

func PreOrderTraverse(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Printf("%s ", node.Val)
	PreOrderTraverse(node.Left)
	PreOrderTraverse(node.Right)
}

func MidOrderTraverse(node *TreeNode) {
	if node == nil {
		return
	}

	MidOrderTraverse(node.Left)
	fmt.Printf("%s ", node.Val)
	MidOrderTraverse(node.Right)
}

func PostOrderTraverse(node *TreeNode) {
	if node == nil {
		return
	}

	PostOrderTraverse(node.Left)
	PostOrderTraverse(node.Right)

	fmt.Printf("%s ", node.Val)
}

func main() {
	root := &TreeNode{Val: "a"}
	root.Left = &TreeNode{Val: "b"}
	root.Right = &TreeNode{Val: "c"}
	root.Left.Left = &TreeNode{Val: "d"}
	root.Left.Right = &TreeNode{Val: "e"}
	root.Right.Left = &TreeNode{Val: "f"}
	root.Right.Right = &TreeNode{Val: "g"}

	PreOrderTraverse(root)
	fmt.Println()
	MidOrderTraverse(root)
	fmt.Println()
	PostOrderTraverse(root)
	fmt.Println()
}
