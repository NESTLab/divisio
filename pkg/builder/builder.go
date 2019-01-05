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

	numNodes := randGen.Intn(5) + 5

	for ii := 0; ii < numNodes; ii++ {
		g.AddNodeRand(randGen, fmt.Sprintf("%d", ii))
	}

	for ii := 0; ii < 3*numNodes; ii++ {
		randWeight := randGen.Intn(100)
		g.AddEdgeRand(randGen, randWeight)
	}

	fmt.Println("New graph built")
	fmt.Println("Number of Nodes: " + fmt.Sprintf("%d", numNodes))
	fmt.Println("Number of Edges: " + fmt.Sprintf("%d", 3*numNodes))

	return &g
}
