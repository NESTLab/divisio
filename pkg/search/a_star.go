package search

import (
	"github.com/NESTLab/divisio.git/pkg/graph"
	"math"
)

func HCost(g graph.Graph, start string, end string) int {

	startNode := g.GetNode(start)
	endNode := g.GetNode(end)

	return int(math.Abs(float64(endNode.Pos.Y-startNode.Pos.Y)) + math.Abs(float64(endNode.Pos.X-startNode.Pos.X)))
}

func GCost(g graph.Graph, start string, end string) int {

	weight := g.GetEdge(start, end).Weight
	if weight != 0 {
		return weight
	}

	return -1
}
