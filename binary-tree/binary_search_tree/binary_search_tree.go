package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTreeNode(data int) *TreeNode {
	return &TreeNode{Val: data}
}

func (node *TreeNode) Insert(data int) {
	if data < node.Val {
		if node.Left == nil {
			node.Left = &TreeNode{Val: data}
			return
		} else {
			node.Left.Insert(data)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Val: data}
			return
		} else {
			node.Right.Insert(data)
		}
	}
}

func PreOrderTraverse(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.Val)
	PreOrderTraverse(node.Left)
	PreOrderTraverse(node.Right)
}

func MidOrderTraverse(node *TreeNode) {
	if node == nil {
		return
	}

	MidOrderTraverse(node.Left)
	fmt.Printf("%d ", node.Val)
	MidOrderTraverse(node.Right)
}

func PostOrderTraverse(node *TreeNode) {
	if node == nil {
		return
	}

	PostOrderTraverse(node.Left)
	PostOrderTraverse(node.Right)

	fmt.Printf("%d ", node.Val)
}

func (node *TreeNode) Find(data int) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Val == data {
		return node
	}

	if data < node.Val {
		return node.Left.Find(data)
	} else {
		return node.Right.Find(data)
	}
}

func getAndDeleteMinFrom(node *TreeNode) int {
	pre := node
	cur := pre.Right
	if cur.Left == nil {
		pre = cur.Right
	} else {
		pre = cur
		cur = cur.Left
	}

	for ; cur.Left != nil; cur = cur.Left {
		pre = cur
	}

	pre.Left = cur.Right
	return cur.Val
}

func (head *TreeNode) Delete(data int) {
	pre := head
	cur := pre
	for cur != nil {
		if cur.Val == data {
			if cur.Left != nil && cur.Right != nil {
				min := getAndDeleteMinFrom(cur)
				cur.Val = min
				continue
			}

			var child *TreeNode
			if cur.Left != nil {
				child = cur.Left
			} else if cur.Right != nil {
				child = cur.Right
			}

			if pre.Left == cur {
				pre.Left = child
				break
			} else if pre.Right == cur {
				pre.Right = child
				cur = child
				continue
			} else {
				pre = child
				break
			}
		}

		if data < cur.Val {
			pre = cur
			cur = cur.Left
		} else {
			pre = cur
			cur = cur.Right
		}
	}
}

func main() {
	root := CreateTreeNode(33)
	root.Insert(16)
	root.Insert(50)
	root.Insert(13)
	root.Insert(18)
	root.Insert(34)
	root.Insert(58)
	root.Insert(15)
	root.Insert(17)
	root.Insert(25)
	root.Insert(51)
	root.Insert(66)
	root.Insert(19)
	root.Insert(27)
	root.Insert(55)
	root.Insert(18)

	PreOrderTraverse(root)
	fmt.Println()
	MidOrderTraverse(root)
	fmt.Println()
	PostOrderTraverse(root)
	fmt.Println()

	if node := root.Find(15); node != nil {
		fmt.Println(node.Val)
	}

	root.Delete(13)
	MidOrderTraverse(root)
	fmt.Println()

	root.Delete(27)
	MidOrderTraverse(root)
	fmt.Println()

	root.Delete(18)
	MidOrderTraverse(root)
	fmt.Println()
}
