package graph

import (
	"fmt"
	"math/rand"
)

func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.Nodes = append(g.Nodes, n)

}

func (g *Graph) AddNodeRand(seed *rand.Rand, name string, isCrossroads bool) {
	g.lock.Lock()
	defer g.lock.Unlock()
	var difficulty int
	var rate int

	if isCrossroads {
		difficulty = 0
		rate = 0
	} else {
		difficulty = seed.Intn(50) + 50
		rate = seed.Intn(50) + 50
	}

	x := seed.Intn(100)
	y := seed.Intn(100)
	n := MakeNode(name, difficulty, rate, x, y)
	g.Nodes = append(g.Nodes, n)

}

func (g *Graph) AddEdge(n1, n2 *Node, weight int) {
	g.lock.Lock()
	defer g.lock.Unlock()
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

func (g *Graph) AddEdgeRand(seed *rand.Rand, weight int) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.Edges == nil {
		g.Edges = make(map[string][]Edge)
	}
	n1 := g.Nodes[seed.Intn(len(g.Nodes))]
	n2 := g.Nodes[seed.Intn(len(g.Nodes))]

	for n1 == n2 {
		n2 = g.Nodes[seed.Intn(len(g.Nodes))]
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

func (g *Graph) PrintConnections() {
	g.lock.RLock()
	defer g.lock.RUnlock()
	s := ""
	for i := 0; i < len(g.Nodes); i++ {
		s += g.Nodes[i].Name + " -> "
		near := g.Edges[g.Nodes[i].Name]

		for j := 0; j < len(near); j++ {
			s += fmt.Sprintf("%s[%d] ", near[j].ToNode, near[j].Weight)
		}
		s += "\n"
	}
	fmt.Println(s)

}

func (g *Graph) GetNode(name string) *Node {
	g.lock.RLock()
	defer g.lock.RUnlock()

	for ii := 0; ii < len(g.Nodes); ii++ {
		if g.Nodes[ii].Name == name {
			return g.Nodes[ii]
		}
	}
	return nil
}

func (g *Graph) GetEdge(start, end string) Edge {
	g.lock.RLock()
	defer g.lock.RUnlock()
	edgeList := g.Edges[start]

	for _, edge := range edgeList {
		if edge.ToNode == end {
			return edge
		}
	}

	return Edge{"", 0}
}
