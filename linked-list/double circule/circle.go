package main

import (
	"fmt"
)

type Node struct {
	Pre  *Node
	Data int
	Next *Node
}

func InitDoubleCircule() *Node {
	header := &Node{}
	header.Next = header
	header.Pre = header

	return header
}

func (header *Node) Insert(data int) {
	node := header
	for node.Next != header {
		node = node.Next
	}

	node.Next = &Node{Pre: node, Data: data, Next: header}
	header.Pre = node.Next
}

func (header *Node) LookupNext() {
	node := header
	for ; node.Next != header; node = node.Next {
		fmt.Printf("%d ", node.Data)
	}

	fmt.Println(node.Data)
}

func (header *Node) LookupPre() {
	node := header
	for ; node.Pre != header; node = node.Pre {
		fmt.Printf("%d ", node.Data)
	}

	fmt.Println(node.Data)
}

func (header *Node) Len() int {
	count := 1
	for node := header; node.Next != header; node = node.Next {
		count++
	}

	return count
}

func (header *Node) InsertK(k, data int) {
	if header.Next == header && k == 1 {
		header.Next = &Node{Pre: header, Data: data, Next: header}
		header.Pre = header.Next
	}

	length := header.Len()
	if k > length {
		fmt.Printf("Error: insert data %d failed cause %d is out of the list length %d\n", data, k, length)
		return
	} else if k < 1 {
		fmt.Printf("Error: %d is not in the range of [1, %d]\n", k, length)
		return
	}

	for node, count := header, 1; node.Next != header; {
		node = node.Next
		count++
		if count == k {
			if node.Next != header {
				node.Next = &Node{Pre: node, Data: data, Next: node.Next}
				node.Next.Next.Pre = node.Next
			} else {
				node.Next = &Node{Pre: node, Data: data, Next: header}
				header.Pre = node.Next
			}
		}
	}
}

func (header *Node) Remove() {
	if header.Next == header {
		header = nil
		return
	}

	node := header
	for node.Next.Next != header {
		node = node.Next
	}

	node.Next, node.Next.Pre = header, nil
}

// node.Next.Next
// header is not allowed to be delete
func (header *Node) DeleteK(k int) {
	if header.Next == header && k == 1 {
		header = nil
		return
	}

	length := header.Len()
	if k > length {
		fmt.Printf("Error: delete data failed cause %d is out of the list length %d\n", k, length)
		return
	} else if k < 2 {
		fmt.Printf("Error: %d is not in the range of [1, %d]\n", k, length)
		return
	}

	count, node := 1, header
	for ; node.Next.Next != header; node = node.Next {
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
		node.Next = header
	}
}

func main() {
	header := InitDoubleCircule()
	header.Insert(1)
	header.Insert(2)
	header.Insert(3)
	header.Insert(4)
	header.LookupNext()
	header.LookupPre()

	header.InsertK(3, 8)
	header.LookupNext()
	header.InsertK(6, 10)
	header.LookupNext()

	header.Remove()
	header.LookupNext()
	header.Remove()
	header.LookupNext()

	header.DeleteK(2)
	header.LookupNext()
	header.DeleteK(4)
	header.LookupNext()
}
