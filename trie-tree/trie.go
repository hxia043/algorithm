package main

import "fmt"

type TrieNode struct {
	Char   rune
	Childs [26]*TrieNode
}

func CreateTrieNode(char rune) *TrieNode {
	return &TrieNode{Char: char}
}

func (node *TrieNode) Insert(text string) {
	for _, char := range text {
		sub := int(char - 'a')
		if node.Childs[sub] == nil {
			node.Childs[sub] = &TrieNode{Char: char}
		}

		node = node.Childs[sub]
	}
}

func (node *TrieNode) Lookup(text string) bool {
	lookupBytes := []byte{}
	for _, char := range text {
		sub := int(char - 'a')
		if node.Childs[sub] == nil {
			return false
		}

		lookupBytes = append(lookupBytes, byte(node.Childs[sub].Char))
		node = node.Childs[sub]
	}

	fmt.Println(string(lookupBytes))
	return true
}

func (root *TrieNode) lookupWrapper(text string) {
	if root.Lookup(text) {
		fmt.Printf("%s is not exist.\n", text)
	}
}

func main() {
	root := CreateTrieNode('/')
	root.Insert("hello")
	root.Insert("her")
	root.Insert("hi")
	root.Insert("how")
	root.Insert("seo")
	root.Insert("so")

	root.lookupWrapper("hello")
	root.lookupWrapper("hallo")
}
