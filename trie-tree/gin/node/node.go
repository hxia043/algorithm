package node

type Node struct {
	Pattern  string
	Part     string
	Children []*Node
	isWild   bool
	IsLeaf   bool
}

func (n *Node) Match(part string) (*Node, bool) {
	for _, child := range n.Children {
		if child.Part == part {
			n = child
			return n, true
		}
	}

	return n, false
}

func (n *Node) Select(part string) (*Node, bool) {
	for _, child := range n.Children {
		if child.Part == part {
			n = child
			return n, true
		}

		if child.isWild && part != " " {
			n = child
			return n, true
		}
	}

	return n, false
}

func (n *Node) New(pattern, part string, children []*Node, wild, leaf bool) *Node {
	return &Node{
		Pattern:  pattern,
		Part:     part,
		Children: children,
		isWild:   wild,
		IsLeaf:   leaf,
	}
}

func Init(pattern, part string, children []*Node, wild, leaf bool) *Node {
	return &Node{
		Pattern:  pattern,
		Part:     part,
		Children: children,
		isWild:   wild,
		IsLeaf:   leaf,
	}
}
