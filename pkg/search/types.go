package search

import "github.com/NESTLab/divisio.git/pkg/graph"

// An Item is something we manage in a priority queue.
type Item struct {
	Node     *graph.Node // The value of the item; arbitrary.
	priority int         // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item
