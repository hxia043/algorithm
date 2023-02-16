package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func (node *Node) Insert(data int) {
	for ; node.Next != nil; node = node.Next {
	}

	node.Next = &Node{Data: data}
}

func (node *Node) Lookup() {
	for ; node.Next != nil; node = node.Next {
		fmt.Printf("%d ", node.Data)
	}

	// lookup tail node
	fmt.Printf("%d\n", node.Data)
}

func (node *Node) Delete(data int) {
	// only one node
	if node.Next == nil {
		if node.Data == data {
			fmt.Println("empty linked list")
		}

		return
	}

	// delete header of node
	if node.Data == data {
		node.Next.Lookup()
	}

	for node.Next != nil {
		if node.Next.Data == data {
			node.Next = node.Next.Next
			continue
		}

		node = node.Next
	}
}

func main() {
	// insert: {4, 5, 6, 3, 1, 2}
	header := &Node{Data: 4}
	header.Insert(5)
	header.Insert(6)
	header.Insert(3)
	header.Insert(1)
	header.Insert(2)

	// lookup header
	header.Lookup()

	// delete the specified data
	header.Delete(3)
	header.Lookup()

	// delete the tail node
	header.Delete(2)
	header.Lookup()

	// delete the header node
	header.Delete(4)
	header.Lookup()
}
