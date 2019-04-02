package builder

import (
	"fmt"
	"github.com/NESTLab/divisio.git/pkg/graph"
	"github.com/NESTLab/divisio.git/pkg/stream"
	"log"
	"math/rand"
	"strconv"
	"time"
)

//GraphBuilderRand takes in an instance of rand.Rand and returns a reference to a graph.
func (c Config) GraphBuilderRand(randObj *rand.Rand) *graph.Graph {
	var g graph.Graph

	numNodes := randObj.Intn(c.ExtraNumNodes) + c.BaseNumNodes

	for ii := 0; ii < numNodes; ii++ {
		name := strconv.Itoa(ii)

		if randObj.Float64() < c.ChanceCrossRoads {
			c.AddNodeRand(&g, randObj, name, true)
		} else {
			c.AddNodeRand(&g, randObj, name, false)
		}
		if ii != 0 {
			for {
				whichNode := randObj.Intn(len(g.Nodes))
				whichNodeName := g.Nodes[whichNode].Name
				var EdgeExists bool
				for _, whichNodeEdge := range g.GetEdges(whichNodeName) {
					if whichNodeEdge.ToNode == name {
						EdgeExists = true
					}
				}
				for _, nameEdge := range g.GetEdges(name) {
					if nameEdge.ToNode == whichNodeName {
						EdgeExists = true
					}
				}
				if whichNodeName != name && !EdgeExists {
					edgeWeight := randObj.Intn(c.ExtraEdgeWeight) + c.BaseEdgeWeight
					g.AddEdge(g.GetNode(whichNodeName), g.GetNode(name), edgeWeight)
					break
				}
			}
		}

	}

	numEdges := randObj.Intn(c.ExtraNumEdges) + c.BaseNumEdges - numNodes

	for ii := 0; ii < numEdges; ii++ {
		randWeight := randObj.Intn(c.ExtraEdgeWeight) + c.BaseEdgeWeight
		c.AddEdgeRand(&g, randObj, randWeight)
	}

	fmt.Println("New graph built")
	//fmt.Println("Number of Nodes: " + fmt.Sprintf("%d", numNodes))
	//fmt.Println("Number of Edges: " + fmt.Sprintf("%d", numEdges))

	return &g
}

//AddEdgeRand chooses two random nodes via randObj and with provided weight
func (c Config) AddEdgeRand(g *graph.Graph, randObj *rand.Rand, weight int) {
	g.Lock.Lock()
	defer g.Lock.Unlock()
	if g.Edges == nil {
		g.Edges = make(map[string][]graph.Edge)
	}
	var n1, n2 *graph.Node

	startIndexOuter := randObj.Intn(len(g.Nodes))
	for jj := 0; jj < len(g.Nodes); jj++ {
		indexOuter := startIndexOuter + jj
		if indexOuter >= len(g.Nodes) {
			indexOuter = indexOuter - len(g.Nodes)
		}
		n1 = g.Nodes[indexOuter]
		startIndexInner := randObj.Intn(len(g.Nodes))
		for ii := 0; ii < len(g.Nodes); ii++ {
			indexInner := startIndexInner + ii
			if indexInner >= len(g.Nodes) {
				indexInner = indexInner - len(g.Nodes)
			}
			n2 = g.Nodes[indexInner]
			var EdgeExists bool
			for _, n1Edge := range g.Edges[n1.Name] {
				if n1Edge.ToNode == n2.Name {
					EdgeExists = true
				}
			}
			for _, n2Edge := range g.Edges[n2.Name] {
				if n2Edge.ToNode == n1.Name {
					EdgeExists = true
				}
			}
			if n1.Name != n2.Name && !EdgeExists {
				g.Edges[n1.Name] = append(g.Edges[n1.Name], graph.Edge{
					ToNode: n2.Name,
					Weight: weight,
				})

				g.Edges[n2.Name] = append(g.Edges[n2.Name], graph.Edge{
					ToNode: n1.Name,
					Weight: weight,
				})
				break
			}
		}

	}

}

//AddNodeRand adds a node to the graph g that is randomly generated
//randObj is the seeded rand object that provides the random numbers
//name is the name of the node, usually just its index
//isCrossroads denotes whether the rate and difficulty are set to random variables (false), or 0 (true)
func (c Config) AddNodeRand(g *graph.Graph, randObj *rand.Rand, name string, isCrossroads bool) {
	g.Lock.Lock()
	defer g.Lock.Unlock()
	var difficulty int
	var rate int

	if isCrossroads {
		difficulty = 0
		rate = 0
	} else {
		difficulty = randObj.Intn(c.ExtraNodeDiff) + c.BaseNodeDiff
		rate = randObj.Intn(c.ExtraNodeRate) + c.BaseNodeRate
	}

	x := randObj.Intn(100)
	y := randObj.Intn(100)
	n := graph.MakeNode(name, difficulty, rate, x, y)
	g.Nodes = append(g.Nodes, n)

}

//GenerateGraphs generates numGraphsToMake of random graphs, and stores them at path
func (c Config) GenerateGraphs(numGraphsToMake int, path string) map[string]*graph.Graph {
	graphs := make(map[string]*graph.Graph)

	//Make the random generator
	randSeed := rand.NewSource(time.Now().Unix())
	randObj := rand.New(randSeed)

	//Each loop creates a new graph
	for ii := 0; ii < numGraphsToMake; ii++ {
		//Build the graph
		g := c.GraphBuilderRand(randObj)

		name := strconv.Itoa(ii)
		graphs[name] = g
	}

	err := stream.GraphOut(graphs, path)
	if err != nil {
		log.Fatalln(err)
	}
	return graphs
}
