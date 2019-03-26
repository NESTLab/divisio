package graph

import (
	"fmt"
)

//AddNode adds the provided node n to the graph g
func (g *Graph) AddNode(n *Node) {
	g.Lock.Lock()
	defer g.Lock.Unlock()
	g.Nodes = append(g.Nodes, n)

}

//AddEdge connects the two provided nodes with an Edge of weight weight
func (g *Graph) AddEdge(n1, n2 *Node, weight int) {
	g.Lock.Lock()
	defer g.Lock.Unlock()
	if g.Edges == nil {
		g.Edges = make(map[string][]Edge)
	}

	g.Edges[n1.Name] = append(g.Edges[n1.Name], Edge{
		ToNode: n2.Name,
		Weight: weight,
	})

	g.Edges[n2.Name] = append(g.Edges[n2.Name], Edge{
		ToNode: n1.Name,
		Weight: weight,
	})

}

//PrintConnections prints one node per row, with all their connections listed to the right, with weight in brackets
func (g *Graph) PrintConnections() {
	g.Lock.RLock()
	defer g.Lock.RUnlock()
	s := ""
	for i := 0; i < len(g.Nodes); i++ {

		s += g.Nodes[i].Name
		if g.Nodes[i].Rate > 0 {
			s += "T"
		} else {
			s += "C"
		}
		near := g.Edges[g.Nodes[i].Name]
		s += " -> "

		for j := 0; j < len(near); j++ {
			s += fmt.Sprintf("%s[%d] ", near[j].ToNode, near[j].Weight)
		}
		s += "\n"
	}
	fmt.Println(s)

}

//GetNode returns the pointer to a Node via it's string name
func (g *Graph) GetNode(name string) *Node {
	g.Lock.RLock()
	defer g.Lock.RUnlock()

	for ii := 0; ii < len(g.Nodes); ii++ {
		if g.Nodes[ii].Name == name {
			return g.Nodes[ii]
		}
	}
	return nil
}

//GetEdge returns the pointer to the Edge that corresponds to the two provided names
func (g *Graph) GetEdge(start, end string) Edge {
	g.Lock.RLock()
	defer g.Lock.RUnlock()
	edgeList := g.Edges[start]

	for _, edge := range edgeList {
		if edge.ToNode == end {
			return edge
		}
	}

	return Edge{"", 0}
}

func (g *Graph) GetEdges(start string) []Edge {
	g.Lock.RLock()
	defer g.Lock.RUnlock()

	return g.Edges[start]
}
