package graph

import (
	"fmt"
	"math/rand"
)

//AddNode adds the provided node n to the graph g
func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.Nodes = append(g.Nodes, n)

}

//AddNodeRand adds a node to the graph g that is randomly generated
//randObj is the seeded rand object that provides the random numbers
//name is the name of the node, usually just its index
//isCrossroads denotes whether the rate and difficulty are set to random variables (false), or 0 (true)
func (g *Graph) AddNodeRand(randObj *rand.Rand, name string, isCrossroads bool) {
	g.lock.Lock()
	defer g.lock.Unlock()
	var difficulty int
	var rate int

	if isCrossroads {
		difficulty = 0
		rate = 0
	} else {
		difficulty = randObj.Intn(50) + 50
		rate = randObj.Intn(50) + 50
	}

	x := randObj.Intn(100)
	y := randObj.Intn(100)
	n := MakeNode(name, difficulty, rate, x, y)
	g.Nodes = append(g.Nodes, n)

}

//AddEdge connects the two provided nodes with an Edge of weight weight
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

//AddEdgeRand chooses two random nodes via randObj and with provided weight
func (g *Graph) AddEdgeRand(randObj *rand.Rand, weight int) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.Edges == nil {
		g.Edges = make(map[string][]Edge)
	}
	n1 := g.Nodes[randObj.Intn(len(g.Nodes))]
	n2 := g.Nodes[randObj.Intn(len(g.Nodes))]

	for n1 == n2 {
		n2 = g.Nodes[randObj.Intn(len(g.Nodes))]
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
	g.lock.RLock()
	defer g.lock.RUnlock()
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
	g.lock.RLock()
	defer g.lock.RUnlock()

	for ii := 0; ii < len(g.Nodes); ii++ {
		if g.Nodes[ii].Name == name {
			return g.Nodes[ii]
		}
	}
	return nil
}

//GetEdge returns the pointer to the Edge that corresponds to the two provided names
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

func (g *Graph) GetEdges(start string) []Edge {
	g.lock.RLock()
	defer g.lock.RUnlock()

	return g.Edges[start]
}
