package search

import (
	"github.com/NESTLab/divisio.git/pkg/graph"
	"gonum.org/v1/gonum/mat"
	"log"
	"strconv"
)

//These are different modes for how to select the Post office
//These are different modes for how to select the Post office
const (
	AStarMode       = 0
	BetweennessMode = 1
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
		//Ignore empty strings from slice array lengthening
		if startNode != "" {

			switch mode {
			case AStarMode:
				for _, endNode := range taskNodes {
					if endNode != "" {
						if !contains(doneNodes, endNode) && endNode != startNode {
							cameFrom, _ := AStarSearch(g, startNode, endNode)
							//Stores the path from task node to task node
							var path []string
							path = AStarReconstructPath(cameFrom, startNode, endNode, false)
							for _, passedNode := range path {
								passes[passedNode]++
							}
						}
					}
				}
				doneNodes = append(doneNodes, startNode)
			case BetweennessMode:
				nodePasses := BetweennessSearch(g, startNode)
				for name, rank := range nodePasses {
					passes[name] += rank
				}
			}

		}
	}

	return passes
}

func CalculateLaplacian(g graph.Graph) *mat.Dense {
	matSize := len(g.Nodes)
	lapData := make([]float64, 0)
	//Go through each row, as NewDense takes a row-major slice
	for node, edges := range g.Edges {
		rowData := make([]float64, matSize)
		numEdges := len(edges)
		for _, edge := range edges {
			edgeNum, err := strconv.Atoi(edge.ToNode)
			if err != nil {
				log.Fatalln(err)
			}
			rowData[edgeNum] = -1
		}
		nodeNum, err := strconv.Atoi(node)
		if err != nil {
			log.Fatalln(err)
		}
		rowData[nodeNum] = float64(numEdges)
		lapData = append(lapData, rowData...)
	}

	return mat.NewDense(matSize, matSize, lapData)
}

//contains is just a simple check to see if a node is within a slice of nodes already
func contains(slice []string, node string) bool {
	for _, n := range slice {
		if n == node {
			return true
		}
	}
	return false
}

func POSRoutine(graphs <-chan *graph.Results, result chan<- *graph.Results, mode int) {
	for g := range graphs {

		g.Results = PostOfficeSelection(*g.GraphObj, mode)
		g.Laplacian = CalculateLaplacian(*g.GraphObj)

		result <- g
	}
}
