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

func (node *TreeNode) MaxDepth() int {
	if node == nil {
		return 0
	}

	queue := []*TreeNode{}
	queue = append(queue, node)
	qs, depth := len(queue), 0

	for qs != 0 {
		for qs != 0 {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			qs--
		}
		depth++
		qs = len(queue)
	}

	return depth
}

func main() {
	root := NewTree(3)

	root.Insert("left", 9)
	root.Insert("right", 20)
	root.Right.Insert("left", 15)
	root.Right.Insert("right", 7)
	root.Right.Left.Insert("left", 8)
	root.Right.Left.Left.Insert("left", 10)

	fmt.Println(root.MaxDepth())
}
