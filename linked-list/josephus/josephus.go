package main

import "fmt"

type People struct {
	Id   int
	Next *People
}

func (header *People) Insert(id int) {
	if header.Next == nil {
		header.Next = header
	}

	people := header
	for people.Next != header {
		people = people.Next
	}

	people.Next = &People{Id: id, Next: header}
}

func (header *People) Lookup() {
	people := header
	for ; people.Next != header; people = people.Next {
		fmt.Printf("%d ", people.Id)
	}

	fmt.Printf("%d %d\n", people.Id, people.Next.Id)
}

func (header *People) Josephus(k int) (int, error) {
	count := 1
	for people := header; people.Next != header; people = people.Next {
		count++
	}

	if k <= 1 || k > count {
		return 0, fmt.Errorf("unexpect k: %d defined, k is defined in the range of [2, %d]", k, count)
	}

	index, people := 1, header
	for people.Next != header {
		index++
		if index == k {
			people.Next = people.Next.Next
			people = people.Next
			header = people
			index = 1
		} else {
			people = people.Next
		}
	}

	return people.Id, nil
}

func (header *People) Length() int {
	length := 1
	for people := header; people.Next != header; people = people.Next {
		length++
	}

	return length
}

func (header *People) Index(k int) int {
	index, length := 0, header.Length()
	if index = k % length; index == 0 {
		index = length
	}

	return index
}

func (header *People) Josephus2(k int) (int, error) {
	if k <= 1 {
		return 0, fmt.Errorf("unexpect k: %d defined, k should more than 1", k)
	}

	count, people := 1, header
	index := people.Index(k)
	for people.Next != header {
		count++
		if count == index {
			people.Next = people.Next.Next
			people = people.Next
			header = people
			count, index = 1, people.Index(k)
		} else {
			people = people.Next
		}
	}

	return people.Id, nil
}

func main() {
	header := &People{Id: 0}
	header.Insert(1)
	header.Insert(2)
	header.Insert(3)
	header.Insert(4)
	header.Insert(5)

	//header.Lookup()

	/*
		id, err := header.Josephus(3)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(id)
	*/

	id, err := header.Josephus2(3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id)
}
