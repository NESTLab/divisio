package graph

import (
	"fmt"
	"math/rand"
)

func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	g.Nodes = append(g.Nodes, n)
	g.lock.Unlock()
}

func (g *Graph) AddNodeRand(seed *rand.Rand, name string) {
	g.lock.Lock()
	difficulty := seed.Intn(50) + 50
	rate := seed.Intn(50) + 50
	n := MakeNode(name, difficulty, rate)
	g.Nodes = append(g.Nodes, n)
	g.lock.Unlock()
}

func (g *Graph) AddEdge(n1, n2 *Node, weight int) {
	g.lock.Lock()
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

	g.lock.Unlock()

}

func (g *Graph) AddEdgeRand(seed *rand.Rand, weight int) {
	g.lock.Lock()
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

	g.lock.Unlock()
}

func (g *Graph) PrintConnections() {
	g.lock.RLock()
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
	g.lock.RUnlock()
}