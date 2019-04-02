package search

import (
	"github.com/NESTLab/divisio.git/pkg/graph"
)

//BetweennessSearch takes in a graph, and a root node, and finds the information flow from all the other [task] nodes to it
func BetweennessSearch(g graph.Graph, root string) map[string]int {
	nodeFlow := make(map[string]int)
	//Get the hierarchy for the nodes. From the start node at i=0, the slice at i=1 has the nodes 1 jump away, i=2 has two jumps away, etc
	nodeHierarchy := getNodeHierarchy(g, root)

	//Reverse the hierarchy so all the furthest nodes are at the top and we can iterate correctly
	for i := len(nodeHierarchy)/2 - 1; i >= 0; i-- {
		opp := len(nodeHierarchy) - 1 - i
		nodeHierarchy[i], nodeHierarchy[opp] = nodeHierarchy[opp], nodeHierarchy[i]
	}

	//Stores the nodes that have already been processed, to avoid duplication and to select edge weights correctly
	processed := make([]string, 0)

	//Start from the 'bottom' and work towards the start node
	for _, level := range nodeHierarchy {
		//Go through each node in that level
		for _, nodeName := range level {

			//Get the node object itself via it's name
			node := g.GetNode(nodeName)
			thisNodeFlow := nodeFlow[nodeName]

			//Add the node's rate to it's flow value. Tasks will add a positive rate, crossroads will add zero
			thisNodeFlow += node.Rate

			//Stores the sum of the outward edges (outward determined by if the toNode has already been processed) so we
			//can use the edge weights to determine the split
			var unexploredEdgeCostSum float64

			//Get the list of nodes to iterate over
			nodeEdges := g.GetEdges(nodeName)
			for _, edge := range nodeEdges {
				//We only care about the weights to the nodes we haven't visited yet
				if !levelContains(processed, edge.ToNode) && !levelContains(level, edge.ToNode) {
					unexploredEdgeCostSum += float64(edge.Weight)
				}
			}

			//Now we iterate over it again, updating the toNodes nodeFlow values with how much this spreads
			for _, edge := range nodeEdges {
				if !levelContains(processed, edge.ToNode) && !levelContains(level, edge.ToNode) {
					percentFlow := float64(1.0 - (float64(edge.Weight) / unexploredEdgeCostSum))
					nodeFlow[edge.ToNode] += int(percentFlow * float64(thisNodeFlow))
				}
			}

			processed = append(processed, nodeName)
		}
	}

	return nodeFlow
}

func getNodeHierarchy(g graph.Graph, root string) [][]string {
	nodeHierarchy := make([][]string, 0)

	level0 := make([]string, 0)
	level0 = append(level0, root)
	nodeHierarchy = append(nodeHierarchy, level0)

	for ll := 0; ll < len(g.Nodes); ll++ {
		nextLevel := make([]string, 0)
		nodeHierarchy = append(nodeHierarchy, nextLevel)
		for _, rootNode := range nodeHierarchy[ll] {
			for _, edgeNode := range g.Edges[rootNode] {
				if !hierarchyContains(nodeHierarchy, edgeNode.ToNode) {
					nodeHierarchy[ll+1] = append(nodeHierarchy[ll+1], edgeNode.ToNode)
				}
			}
		}
	}
	shrunkNodeHierarchy := append(make([][]string, 0, 0), nodeHierarchy...)
	return shrunkNodeHierarchy
}

func hierarchyContains(hierarchy [][]string, nodeToCheck string) bool {
	for _, level := range hierarchy {
		for _, node := range level {
			if node == nodeToCheck {
				return true
			}
		}
	}
	return false
}

func levelContains(level []string, nodeToCheck string) bool {
	for _, nodeName := range level {
		if nodeName == nodeToCheck {
			return true
		}
	}
	return false
}

func calcTotalCostToRoot(g graph.Graph, hierarchy [][]string, root string, start string) float64 {
	if root == start {
		return 0
	}
	var startDepth int
	for depth, level := range hierarchy {
		for _, node := range level {
			if node == start {
				startDepth = depth
			}
		}
	}

	connections := make([]string, 0, 0)

	for ii := 0; ii < len(hierarchy[startDepth-1]); ii++ {
		tmpEdge := g.GetEdge(start, hierarchy[startDepth-1][ii])
		if tmpEdge.ToNode != "" {
			connections = append(connections, hierarchy[startDepth-1][ii])
		}

	}

	var totalCost float64

	for _, node := range connections {
		totalCost += float64(g.GetEdge(start, node).Weight) + calcTotalCostToRoot(g, hierarchy, root, node)
	}
	totalCost = totalCost / float64(len(connections))
	return totalCost
}
