package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) Insert(flag string, val int) {
	if node == nil {
		return
	}

	if flag == "left" {
		if node.Left != nil {
			fmt.Println("insert left node failed")
		}

		node.Left = &TreeNode{Val: val}
	} else if flag == "right" {
		if node.Right != nil {
			fmt.Println("insert right node failed")
		}

		node.Right = &TreeNode{Val: val}
	} else {
		fmt.Println("unknown flag")
	}
}

func NewTree(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func (node *TreeNode) MidOrderTraverse() {
	if node == nil {
		return
	}

	node.Left.MidOrderTraverse()
	fmt.Println(node.Val)
	node.Right.MidOrderTraverse()
}

func (node *TreeNode) Depth() int {
	return max(node)
}

func max(node *TreeNode) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return 1
	}

	left := max(node.Left) + 1
	right := max(node.Right) + 1
	if left > right {
		return left
	} else {
		return right
	}
}

func main() {
	root := NewTree(3)

	root.Insert("left", 9)
	root.Insert("right", 20)
	root.Right.Insert("left", 15)
	root.Right.Insert("right", 7)
	root.Right.Left.Insert("left", 8)
	root.Right.Left.Left.Insert("left", 10)

	root.MidOrderTraverse()
	fmt.Println(root.Depth())
}
