package graph

import "fmt"

func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

func (g *Graph) AddEdge(n1, n2 *Node, weight int) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]Edge)
	}

	g.edges[*n1] = append(g.edges[*n1], Edge{
		toNode: n2,
		weight: weight,
	})

	g.edges[*n2] = append(g.edges[*n2], Edge{
		toNode: n1,
		weight: weight,
	})

	g.lock.Unlock()

}

func (g *Graph) PrintConnections() {
	g.lock.RLock()
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].Name() + " -> "
		near := g.edges[*g.nodes[i]]

		for j := 0; j < len(near); j++ {
			s += fmt.Sprintf("%s[%d] ", near[j].toNode.Name(), near[j].weight)
		}
		s += "\n"
	}
	fmt.Println(s)
	g.lock.RUnlock()
}
