package graph

//MakeNode is just a wrapper for the normal Node{} method of building
func MakeNode(name string, difficulty int, rate int, x, y int) *Node {
	return &Node{
		Name:       name,
		Difficulty: difficulty,
		Rate:       rate,
		Pos:        Location{x, y},
	}
}

//
func (n Node) IsTask() bool {
	return n.Difficulty > 0
}
