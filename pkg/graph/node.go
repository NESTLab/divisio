package graph

func MakeNode(name string, difficulty int, rate int, x, y int) *Node {
	return &Node{
		Name:       name,
		Difficulty: difficulty,
		Rate:       rate,
		Pos:        Location{x, y},
	}
}
