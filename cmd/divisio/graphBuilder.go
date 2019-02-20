package main

import (
	"flag"
	"github.com/NESTLab/divisio.git/pkg/builder"
)

var numGraphsToMake = flag.Int("numGraphs", 1, "Provide the number of graphs you wish to generate")

func main() {
	flag.Parse()

	graphs := builder.GenerateGraphs(*numGraphsToMake, "graphs")
	for _, g := range graphs {
		g.PrintConnections()
	}

}
