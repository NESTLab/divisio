package search

import "github.com/NESTLab/divisio.git/pkg/graph"

//These are different modes for how to select the Post office
//These are different modes for how to select the Post office
const (
	AStarMode = 0
)

//PostOfficeSelection takes in a graph and a chosen mode and returns a map of nodes with number of visits recorded via path planning between Tasks
func PostOfficeSelection(g graph.Graph, mode int) map[string]int {
	passes := make(map[string]int)

	//Get a list with only the task nodes in it
	taskNodes := make([]string, 1)
	for _, node := range g.Nodes {
		if node.IsTask() {
			taskNodes = append(taskNodes, node.Name)
		}
	}

	//We will reference this later to avoid finding paths from both ends
	doneNodes := make([]string, len(taskNodes))

	//Iterate over all the task nodes for starting locations
	for _, startNode := range taskNodes {
		if startNode != "" {
			//Iterate over all the task nodes for end locations
			for _, endNode := range taskNodes {
				if endNode != "" {
					//If the end node hasn't been completed
					if !contains(doneNodes, endNode) && endNode != startNode {
						var path []string
						switch mode {
						case AStarMode:
							cameFrom, _ := AStarSearch(g, startNode, endNode)
							path = ReconstructPath(cameFrom, startNode, endNode, false)
						}

						for _, passedNode := range path {
							passes[passedNode]++
						}
					}
				}
			}
			doneNodes = append(doneNodes, startNode)
		}
	}

	return passes
}

func contains(slice []string, node string) bool {
	for _, n := range slice {
		if n == node {
			return true
		}
	}
	return false
}
