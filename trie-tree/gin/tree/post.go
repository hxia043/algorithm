package tree

import "gin-trie/node"

type PostTree struct {
	patternCount int
	Root         *node.Node
}

func (t *PostTree) Insert(pattern string) {}

func (t *PostTree) Select(pattern string) bool { return false }

func (t *PostTree) PatternCount() int { return t.patternCount }
