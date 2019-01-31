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

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item
