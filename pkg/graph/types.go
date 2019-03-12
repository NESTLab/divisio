package graph

import "sync"

//Node is each hub of the graph
//If the node is a 'task' then it will have some positive difficulty and rate
//if it is a 'crossroad' then it will have the difficulty and rate set to zero
//the crossings value will be incremented during the testing phase where the most popular nodes are found
type Node struct {
	Name       string   `json:"name"`
	Difficulty int      `json:"difficulty"`
	Rate       int      `json:"rate"`
	Pos        Location `json:"pos"`
}

//Edge holds a reference to the connected Node, and the weight of the edge itself
type Edge struct {
	ToNode string `json:"to_node"`
	Weight int    `json:"weight"`
}

//Graph holds all of the nodes and connections
//It has a list of node references, to allow direct edits
//It has a map of nodes to edges, that way you can find all the connections of each node
type Graph struct {
	Nodes []*Node           `json:"nodes"`
	Edges map[string][]Edge `json:"edges"`
	lock  sync.RWMutex
}

//Location stores x and y coordinates in a separate
type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}

//GraphResults holds a graph ref, its name, and the eventual results from the post office selection
type GraphResults struct {
	GraphObj *Graph
	Name     string
	Results  map[string]int
}
