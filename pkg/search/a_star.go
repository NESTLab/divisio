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

func AStarSearch(g graph.Graph, start, end string) (map[string]string, map[string]int) {
	frontier := make(PriorityQueue, 0)
	startNode := &Item{
		Node:     g.GetNode(start),
		priority: 0,
	}
	frontier.Push(startNode)

	cameFrom := make(map[string]string)
	costSoFar := make(map[string]int)

	cameFrom[start] = ""
	costSoFar[start] = 0

	for frontier.Len() > 0 {
		current := frontier.Pop().(*Item)
		currentName := current.Node.Name

		if current.Node.Name == end {
			break
		}

		for ii := 0; ii < len(g.Edges[currentName]); ii++ {
			toName := g.Edges[currentName][ii].ToNode
			thisCost := GCost(g, currentName, toName)
			newCost := costSoFar[currentName] + thisCost

			_, inCostSoFar := costSoFar[toName]
			if !inCostSoFar || newCost < costSoFar[toName] {
				costSoFar[toName] = newCost
				priority := newCost + HCost(g, start, end)
				frontier.Push(&Item{
					Node:     g.GetNode(toName),
					priority: priority,
				})
				cameFrom[toName] = currentName
			}
		}
	}

	return cameFrom, costSoFar
}

func AStarReconstructPath(cameFrom map[string]string, start string, end string, onlyMid bool) []string {
	var current string

	if onlyMid {
		current = cameFrom[end]
	} else {
		current = end
	}

	path := make([]string, 0)

	for current != start {
		path = append(path, current)
		current = cameFrom[current]
	}
	if !onlyMid {
		path = append(path, start)
	}

	for i := len(path)/2 - 1; i >= 0; i-- {
		opp := len(path) - 1 - i
		path[i], path[opp] = path[opp], path[i]
	}

	return path
}
