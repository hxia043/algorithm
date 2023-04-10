package tree

import (
	"gin-trie/node"
	"gin-trie/types"
	"strings"
)

type Tree interface {
	Insert(pattern string)
	Select(pattern string) bool
	PatternCount() int
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := []string{}

	for _, s := range vs {
		if s != "" {
			parts = append(parts, s)
		}
	}

	return parts
}

func NewTree(method string) Tree {
	var t Tree
	switch method {
	case types.GETROUTEMETHOD:
		t = &GetTree{Root: node.Init("", "", nil, false, false)}
	case types.POSTROUTERMETHOD:
		t = &PostTree{Root: node.Init("", "", nil, false, false)}
	default:
		return nil
	}

	return t
}
