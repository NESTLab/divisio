package search

import "github.com/NESTLab/divisio.git/pkg/graph"

func BetweennessSearch(g graph.Graph, start string) map[string]int {
	nodeFlow := make(map[string]int)
	return nodeFlow
}

func GetNodeHierarchy(g graph.Graph, start string) [][]string {
	nodeHierarchy := make([][]string, 0)

	level0 := make([]string, 0)
	level0 = append(level0, start)
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

	return nodeHierarchy
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
