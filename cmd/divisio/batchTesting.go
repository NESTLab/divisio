package main

import (
	"flag"
	"fmt"
	"github.com/NESTLab/divisio.git/pkg/builder"
	"github.com/NESTLab/divisio.git/pkg/graph"
	"github.com/NESTLab/divisio.git/pkg/search"
	"github.com/NESTLab/divisio.git/pkg/stream"
	"log"
)

var pathToTestBatch = flag.String("testPath", "",
	"The path to the specific graph group containing the jsons for the graphs to test")
var numberToGenerate = flag.Int("numGen", 0, "The number of graphs to generate if starting sims from scratch")
var pathToGenBatch = flag.String("genPath", "",
	"The path to the root graph directory you wish to store the generated graph's json in. "+
		"A new group folder will automatically be made")

func main() {
	flag.Parse()
	var generateTestGraphs bool
	if *numberToGenerate == 0 && *pathToTestBatch == "" {
		log.Fatalln("Failed to provide a number to generate, or path to load graphs from")
	}
	if *numberToGenerate > 0 && *pathToGenBatch == "" {
		log.Fatalf("Failed to provide a pathToTestBatch but requested %d generated graphs", *numberToGenerate)
	} else {
		generateTestGraphs = true
	}

	graphs := make(map[string]*graph.Graph)

	if generateTestGraphs {
		graphs = builder.GenerateGraphs(*numberToGenerate, *pathToGenBatch)

	} else {
		streamGraphs, err := stream.In(*pathToTestBatch)
		if err != nil {
			log.Fatal(err)
		}
		graphs = streamGraphs
	}

	var output string
	for name, g := range graphs {
		passes := search.PostOfficeSelection(*g, search.AStarMode)
		output = fmt.Sprintf("%s%s: %v\n", output, name, passes)
		fmt.Printf("%s:####################", name)
		g.PrintConnections()
	}
	fmt.Println(output)
}
