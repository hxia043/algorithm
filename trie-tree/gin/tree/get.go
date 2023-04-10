package tree

import "gin-trie/node"

type GetTree struct {
	patternCount int
	Root         *node.Node
}

func (t *GetTree) isPatternExist(pattern string) bool {
	node := t.Root
	parts := parsePattern(pattern)
	isMatched := false

	for _, part := range parts {
		node, isMatched = node.Match(part)
		if !isMatched {
			return false
		}
	}

	return true
}

func (t *GetTree) Insert(pattern string) {
	if !t.isPatternExist(pattern) {
		t.patternCount++
	}

	node := t.Root
	parts := parsePattern(pattern)
	isMatched := false

	for i, part := range parts {
		node, isMatched = node.Match(part)
		if !isMatched {
			n := node.New("", part, nil, part[0] == ':' || part[0] == '*', false)

			if len(parts) == i+1 {
				n.Pattern = pattern
				n.IsLeaf = true
			}

			node.Children = append(node.Children, n)
			node = n
		}
	}
}

func (t *GetTree) Select(pattern string) bool {
	node := t.Root
	parts := parsePattern(pattern)
	isSelected := false
	for i, part := range parts {
		node, isSelected = node.Select(part)
		if !isSelected {
			return false
		}

		if len(parts) == i+1 {
			if !node.IsLeaf {
				return false
			}
		}
	}

	return true
}

func (t *GetTree) PatternCount() int {
	return t.patternCount
}
