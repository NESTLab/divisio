package main

import "github.com/NESTLab/divisio.git/pkg/graph"

func main() {
	var g graph.Graph

	nA := graph.MakeNode("A", 0, 0)
	nB := graph.MakeNode("B", 0, 0)
	nC := graph.MakeNode("C", 0, 0)
	nD := graph.MakeNode("D", 0, 0)
	nE := graph.MakeNode("E", 0, 0)
	nF := graph.MakeNode("F", 0, 0)

	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

	g.AddEdge(&nA, &nB, 100)
	g.AddEdge(&nA, &nC, 50)
	g.AddEdge(&nB, &nE, 25)
	g.AddEdge(&nC, &nE, 12)
	g.AddEdge(&nE, &nF, 6)
	g.AddEdge(&nD, &nA, 3)

	g.PrintConnections()

}
