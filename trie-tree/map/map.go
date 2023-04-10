package main

import "fmt"

type TrieTree struct {
	TermCount int
	Root      *TrieNode
}

type TrieNode struct {
	Char   rune
	Childs map[rune]*TrieNode
	term   bool
	path   string
}

func CreateTrieTree() *TrieTree {
	return &TrieTree{
		TermCount: 0,
		Root: &TrieNode{
			Char:   '/',
			Childs: make(map[rune]*TrieNode),
		},
	}
}

func (t *TrieTree) Insert(text string) {
	isExist, isPath := t.Lookup(text)
	if !(isExist && isPath) {
		t.TermCount++
	}

	node := t.Root
	for _, c := range text {
		if node.Childs[c] == nil {
			node.Childs[c] = &TrieNode{
				Char:   c,
				Childs: make(map[rune]*TrieNode),
			}
		}

		node = node.Childs[c]
	}
	node.path, node.term = text, true
}

func (t *TrieTree) Lookup(text string) (isExist bool, isPath bool) {
	node := t.Root
	for _, c := range text {
		if node.Childs[c] == nil {
			isExist, isPath = false, false
			return
		}

		node = node.Childs[c]
	}

	if node.term {
		isExist, isPath = true, true
	} else {
		isExist, isPath = true, false
	}

	return
}

func (t *TrieTree) lookupWrapper(text string) {
	isExist, isPath := t.Lookup(text)
	if isExist {
		if isPath {
			fmt.Printf("%s is the path of tree\n", text)
		} else {
			fmt.Printf("%s is the prefix of tree\n", text)
		}
	} else {
		fmt.Printf("%s is not exist in tree\n", text)
	}
}

func (t *TrieTree) FindNode(text string) *TrieNode {
	node := t.Root
	for _, c := range text {
		if node.Childs[c] == nil {
			return nil
		}

		node = node.Childs[c]
	}

	return node
}

func (t *TrieTree) GetTermCount() int {
	return t.TermCount
}

func (t *TrieTree) FuzzySearch(text string) []string {
	isExist, _ := t.Lookup(text)
	if !isExist {
		return nil
	}

	node := t.FindNode(text)
	nodes := []*TrieNode{}
	texts := []string{}

	for _, n := range node.Childs {
		if n != nil {
			nodes = append(nodes, n)
		}
	}

	for l := len(nodes); l != 0; l = len(nodes) {
		i := l - 1
		node = nodes[i]
		nodes = nodes[:i]

		for _, n := range node.Childs {
			if n != nil {
				nodes = append(nodes, n)
			}
		}

		if node.term {
			texts = append(texts, node.path)
		}
	}

	return texts
}

func main() {
	t := CreateTrieTree()
	t.Insert("hello")
	t.Insert("her")
	t.Insert("hi")
	t.Insert("how")
	t.Insert("seo")
	t.Insert("so")
	t.Insert("hello")
	t.Insert("soo")

	t.lookupWrapper("hello")
	t.lookupWrapper("he")
	t.lookupWrapper("hallo")

	for _, text := range t.FuzzySearch("he") {
		fmt.Println(text)
	}

	fmt.Println(t.GetTermCount())
}
