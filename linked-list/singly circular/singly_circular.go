package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func (node *Node) Insert(data int) {
	if node.Next == nil {
		node.Next = &Node{Data: data}
		node.Next.Next = node.Next
		return
	}

	header := node.Next
	node = node.Next
	for node.Next != header {
		node = node.Next
	}

	node.Next = &Node{Data: data, Next: header}
}

func (node *Node) InsertK(index, data int) {
	count := 0
	for node.Next != nil {
		count++
		if count == index {
			node.Next.Next = &Node{Data: data, Next: node.Next.Next}
		} else {
			node = node.Next
		}
	}
}

func (node *Node) Lookup() {
	if node.Next == nil {
		fmt.Println(" ")
		return
	}

	header := node.Next
	node = node.Next
	for ; node.Next != header; node = node.Next {
		fmt.Printf("%d ", node.Data)
	}

	fmt.Printf("%d\n", node.Data)
}

// Can not delete the header of circular singly list
// Or how to implemente?
func (node *Node) Delete(data int) {
	if node.Next == nil {
		fmt.Println(" ")
		return
	}

	if node.Next.Next == node.Next {
		if node.Next.Data == data {
			node.Next = nil
		}
		return
	}

	header := node.Next
	for node.Next.Next != header {
		if node.Next.Data == data {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}

	if node.Next.Data == data {
		node.Next = header
	}
}

func (node *Node) Remove() {
	if node.Next == nil {
		return
	}

	if node.Next.Next == node.Next {
		node.Next = nil
		return
	}

	header := node.Next
	for node.Next.Next != header {
		node = node.Next
	}

	node.Next = header
}

func (node *Node) Update(index, data int) {
	count := 0
	for node.Next != nil {
		count++
		if count == index {
			node.Next.Data = data
		} else {
			node = node.Next
		}
	}
}

func main() {
	header := &Node{}
	header.Insert(2)
	header.Insert(3)
	header.Insert(1)
	header.Insert(5)
	header.Insert(1)
	header.Insert(4)
	header.Lookup()

	/*
		header.Remove()
		header.Remove()
		header.Remove()
		header.Remove()
		header.Lookup()

		header.Remove()
		header.Lookup()

		header.Remove()
		header.Lookup()
	*/

	header.Delete(1)
	header.Lookup()
	header.Delete(2)
	header.Lookup()
	header.Delete(4)
	header.Lookup()
	header.Delete(3)
	header.Lookup()
	/*
		header.Delete(1)
		header.Lookup()

		header.Delete(2)
		header.Lookup()

		header.Delete(4)
		header.Lookup()

		header.Remove()
		header.Lookup()

		header.Remove()
		header.Lookup()

		header.Insert(1)
		header.InsertK(1, 2)
		header.Lookup()
		header.InsertK(2, 3)
		header.Lookup()
		header.InsertK(2, 4)
		header.Lookup()

		header.Update(1, 2)
		header.Lookup()
		header.Update(4, 2)
		header.Lookup()
		header.Update(3, 2)
		header.Lookup()
	*/
}
