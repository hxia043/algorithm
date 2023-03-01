package main

import (
	"fmt"
)

type Node struct {
	Pre  *Node
	Data int
	Next *Node
}

func (header *Node) Insert(data int) {
	node := header
	for node.Next != nil {
		node = node.Next
	}

	node.Next = &Node{Pre: node, Data: data}
}

func (header *Node) Lookup() {
	if header.Next == nil {
		return
	}

	node := header.Next
	for ; node.Next != nil; node = node.Next {
		fmt.Printf("%d ", node.Data)
	}

	fmt.Println(node.Data)
}

func (header *Node) Len() int {
	count := 0
	for node := header; node.Next != nil; node = node.Next {
		count++
	}

	return count
}

func (header *Node) InsertK(k, data int) {
	if header.Next == nil && k == 1 {
		header.Next = &Node{Pre: header, Data: data}
	}

	length := header.Len()
	if k > length {
		fmt.Printf("Error: insert data %d failed cause %d is out of the list length %d\n", data, k, length)
		return
	} else if k < 1 {
		fmt.Printf("Error: %d is not in the range of [1, %d]\n", k, length)
		return
	}

	for node, count := header, 0; node.Next != nil; {
		node = node.Next
		count++
		if count == k {
			if node.Next != nil {
				node.Next = &Node{Pre: node, Data: data, Next: node.Next}
				node.Next.Next.Pre = node.Next
			} else {
				node.Next = &Node{Pre: node, Data: data, Next: nil}
			}
		}
	}
}

func (header *Node) Remove() {
	if header.Next == nil {
		return
	}

	node := header
	for node.Next.Next != nil {
		node = node.Next
	}

	node.Next, node.Next.Pre = nil, nil
}

//func (header *Node) LookupPre() {}

// node.Next.Next
func (header *Node) DeleteK(k int) {
	if header.Next == nil {
		fmt.Println("empty list")
		return
	}

	length := header.Len()
	if k > length {
		fmt.Printf("Error: delete data failed cause %d is out of the list length %d\n", k, length)
		return
	} else if k < 1 {
		fmt.Printf("Error: %d is not in the range of [1, %d]\n", k, length)
		return
	}

	count, node := 0, header
	for ; node.Next.Next != nil; node = node.Next {
		count++
		if count == k {
			node.Next = node.Next.Next
			node.Next.Pre = node
			return
		}
	}

	count++
	if count == k {
		node.Next.Pre = nil
		node.Next = nil
	}
}

func main() {
	header := &Node{}

	header.Insert(0)
	header.Insert(1)
	header.Insert(2)
	header.Insert(3)
	header.Insert(4)
	header.Lookup()

	header.InsertK(3, 8)
	header.Lookup()
	header.InsertK(6, 10)
	header.Lookup()

	header.Remove()
	header.Lookup()
	header.Remove()
	header.Lookup()

	header.DeleteK(1)
	header.Lookup()
	header.DeleteK(4)
	header.Lookup()
}
