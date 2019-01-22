package graph

func MakeNode(name string, difficulty int, rate int) *Node {
	return &Node{
		Name:       name,
		Difficulty: difficulty,
		Rate:       rate,
	}
}
