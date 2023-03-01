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

/*
	func (head *ListNode) Reverse() *ListNode {
		var reverseHead *ListNode
		node := head
		for ; node.Next != nil; node = node.Next {
			&ListNode{Val: node.Val}
		}
	}
*/

func (head *ListNode) Len() int {
	node := head
	count := 1
	for ; node.Next != nil; node = node.Next {
		count++
	}

	return count
}

func (head *ListNode) LookupK(k int) (*ListNode, error) {
	if !(1 <= k && k <= head.Len()) {
		return nil, fmt.Errorf("unexpect lookup value %d, value is in the range of [1, %d]", k, head.Len())
	}

	if head.Next == nil && k == 1 {
		return head, nil
	}

	count, node := 0, head
	for ; node.Next != nil; node = node.Next {
		count++
		if count == k {
			return node, nil
		}
	}

	count++
	if count == k {
		return node, nil
	}

	return nil, nil
}

func (head *ListNode) ReverseByExchange() {
	length := head.Len()
	step, node := 0, head
	for ; node.Next != nil; node = node.Next {
		step++
		if step > length/2 {
			return
		}

		eNode, err := head.LookupK(length - step + 1)
		if err != nil {
			return
		}

		if eNode == nil {
			return
		}

		eNode.Val, node.Val = node.Val, eNode.Val
	}
}

func (head *ListNode) ReverseByStack() *ListNode {
	lists := []int{}
	node := head
	for ; node.Next != nil; node = node.Next {
		lists = append(lists, node.Val)
	}
	lists = append(lists, node.Val)

	length := len(lists)
	rhead := &ListNode{Val: lists[length-1]}
	for index := length - 1; index != 0; index-- {
		rhead.Insert(lists[index-1])
	}

	return rhead
}

func (head *ListNode) ReverseByDoublePointer() *ListNode {
	var pre *ListNode = nil
	var temp *ListNode = nil
	for cur := head; cur != nil; cur = temp {
		temp = cur.Next
		cur.Next = pre
		pre = cur
	}

	return pre
}

func main() {
	header := ListNode{}
	header.Insert(1)
	header.Insert(2)
	header.Insert(3)
	header.Insert(4)
	header.Insert(5)
	header.Insert(6)
	header.Lookup()

	/*
		header.ReverseByExchange()
		header.Lookup()

		rheader := header.ReverseByStack()
		rheader.Lookup()
	*/

	rear := header.ReverseByDoublePointer()
	rear.Lookup()
}
