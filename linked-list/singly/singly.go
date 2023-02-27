package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func (node *Node) Insert(data int) {
	for node.Next != nil {
		node = node.Next
	}

	node.Next = &Node{Data: data}
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
	for ; node.Next != nil; node = node.Next {
		fmt.Printf("%d", node.Next.Data)
	}

	fmt.Printf("\n")
}

func (node *Node) Delete(data int) {
	for node.Next != nil {
		if node.Next.Data == data {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
}

func (node *Node) Remove() {
	if node.Next == nil {
		return
	}

	for node.Next.Next != nil {
		node = node.Next
	}

	node.Next = nil
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
}
