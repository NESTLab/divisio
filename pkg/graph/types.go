package graph

import "sync"

//Node is each hub of the graph
//If the node is a 'task' then it will have some positive difficulty and rate
//if it is a 'crossroad' then it will have the difficulty and rate set to zero
//the crossings value will be incremented during the testing phase where the most popular nodes are found
type Node struct {
	name       string
	difficulty int
	rate       int
	crossings  int
}

//Edge holds a reference to the connected Node, and the weight of the edge itself
type Edge struct {
	toNode *Node
	weight int
}

//Graph holds all of the nodes and connections
//It has a list of node references, to allow direct edits
//It has a map of nodes to edges, that way you can find all the connections of each node
type Graph struct {
	nodes []*Node
	edges map[Node][]Edge
	lock  sync.RWMutex
}
