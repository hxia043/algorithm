package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) Insert(data int) {
	node := head
	for ; node.Next != nil; node = node.Next {
	}

	node.Next = &ListNode{Val: data}
}

func (head *ListNode) Lookup() {
	node := head
	for ; node.Next != nil; node = node.Next {
		fmt.Printf("%d ", node.Val)
	}

	fmt.Println(node.Val)
}

func getLen(head *ListNode) int {
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	return length
}

func findMid(head *ListNode, length int) *ListNode {
	for count, node := 0, head; node != nil; node = node.Next {
		count++
		if count == length/2 {
			return node
		}
	}

	return nil
}

func reverseList(mid *ListNode) *ListNode {
	var pre *ListNode = nil
	for temp, cur := mid, mid.Next; cur != nil; cur = temp {
		temp = cur.Next
		cur.Next = pre
		pre = cur
	}

	return pre
}

func isPalindromeInstance(head, mid, rear *ListNode) bool {
	isPalindrome := true
	for hnode, rnode := head, rear; hnode != mid.Next; hnode, rnode = hnode.Next, rnode.Next {
		if hnode.Val != rnode.Val {
			isPalindrome = false
			break
		}
	}

	return isPalindrome
}

func restoreList(mid, rear *ListNode) {
	var rpre *ListNode = nil
	var temp *ListNode = nil
	fmt.Println(mid)
	for rcur := rear; rcur != mid.Next; rcur = temp {
		fmt.Println(rcur.Next)
		temp = rcur.Next
		rcur.Next = rpre
		rpre = rcur
	}
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	listLen := getLen(head)
	if listLen == 1 {
		return true
	}

	mid := findMid(head, listLen)
	rear := reverseList(mid)
	isPalindromeResult := isPalindromeInstance(head, mid, rear)
	restoreList(mid, rear)

	for node := head; node != nil; node = node.Next {
		fmt.Printf("%d ", node.Val)
	}

	return isPalindromeResult
}

func main() {
	header := &ListNode{}
	header.Insert(1)
	header.Insert(2)
	header.Insert(3)
	header.Insert(4)
	header.Insert(5)
	header.Insert(6)
	header.Lookup()

	fmt.Println(isPalindrome(header))
}
