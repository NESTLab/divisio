package builder

import (
	"fmt"
	"github.com/NESTLab/divisio.git/pkg/graph"
	"math/rand"
)

func GraphBuilderRand(seed int64) *graph.Graph {
	randSeed := rand.NewSource(seed)
	randGen := rand.New(randSeed)

	var g graph.Graph

	numNodes := randGen.Intn(5) + 10

	for ii := 0; ii < numNodes; ii++ {
		name := fmt.Sprintf("%d", ii)

		//30% chance of being a task. ~2-5 task nodes per graph
		if randGen.Intn(100) > 30 {
			g.AddNodeRand(randGen, name, true)
		} else {
			g.AddNodeRand(randGen, name, false)
		}
	}

	edgeNumFactor := randGen.Intn(2) + 2 //2-4x number of nodes

	for ii := 0; ii < edgeNumFactor*numNodes; ii++ {
		randWeight := randGen.Intn(100)
		g.AddEdgeRand(randGen, randWeight)
	}

	fmt.Println("New graph built")
	fmt.Println("Number of Nodes: " + fmt.Sprintf("%d", numNodes))
	fmt.Println("Number of Edges: " + fmt.Sprintf("%d", edgeNumFactor*numNodes))

	return &g
}
