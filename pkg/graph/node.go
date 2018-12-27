package graph

func (n *Node) Name() string {
	return n.name
}

func MakeNode(name string, difficulty int, rate int) Node {
	return Node{
		name:       name,
		difficulty: difficulty,
		rate:       rate,
		crossings:  0,
	}
}
