package main

import (
	"flag"
	"fmt"
	"github.com/NESTLab/divisio.git/pkg/builder"
	"github.com/NESTLab/divisio.git/pkg/graph"
	"github.com/NESTLab/divisio.git/pkg/search"
	"github.com/NESTLab/divisio.git/pkg/stream"
	"log"
	"os"
)

var pathToTestBatch = flag.String("testPath", "",
	"The path to the specific graph group containing the jsons for the graphs to test")
var numberToGenerate = flag.Int("numGen", 0, "The number of graphs to generate if starting sims from scratch")
var pathToGenBatch = flag.String("genPath", "",
	"The path to the root graph directory you wish to store the generated graph's json in. "+
		"A new group folder will automatically be made")
var testMode = flag.Int("mode", 0, "The search mode to run. See search/types")

func main() {
	flag.Parse()
	var generateTestGraphs bool
	if *numberToGenerate == 0 && *pathToTestBatch == "" {
		log.Fatalln("Failed to provide a number to generate, or path to load graphs from")
		os.Exit(-1)
	} else if *numberToGenerate > 0 && *pathToGenBatch == "" {
		log.Fatalf("Failed to provide a pathToTestBatch but requested %d generated graphs", *numberToGenerate)
		os.Exit(-2)
	} else if *pathToTestBatch != "" && *numberToGenerate == 0 {
		generateTestGraphs = false
	} else {
		generateTestGraphs = true
	}

	if *testMode < 0 || *testMode > 1 {
		log.Fatalln("Incorrect test mode provided")
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

	graphChan := make(chan *graph.GraphResults, 100)
	resultsChan := make(chan *graph.GraphResults, 100)

	for ww := 0; ww < 4; ww++ {
		go search.POSRoutine(graphChan, resultsChan, *testMode)
	}

	for name, g := range graphs {
		gr := new(graph.GraphResults)
		gr.GraphObj = g
		gr.Name = name
		graphChan <- gr
	}

	close(graphChan)

	var out string
	for ii := 0; ii < len(graphs); ii++ {
		gr := <-resultsChan
		out += fmt.Sprintf("%s: %v\n", gr.Name, gr.Results)
	}
	fmt.Println(out)
}
