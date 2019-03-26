package main

import (
	"flag"
	"github.com/NESTLab/divisio.git/pkg/graph"
	"github.com/NESTLab/divisio.git/pkg/search"
	"github.com/NESTLab/divisio.git/pkg/stream"
	"log"
)

var pathToTestBatch = flag.String("testPath", "",
	"The path to the specific graph group containing the jsons for the graphs to test")
var testMode = flag.Int("mode", 0, "The search mode to run. See search/types")

func main() {
	flag.Parse()
	if *pathToTestBatch == "" {
		log.Fatalln("No path to test batch provided")
	}
	if *testMode < 0 || *testMode > 1 {
		log.Fatalln("Incorrect test mode provided")
	}

	graphs := make(map[string]*graph.Graph)

	streamGraphs, err := stream.GraphIn(*pathToTestBatch)
	if err != nil {
		log.Fatal(err)
	}
	graphs = streamGraphs

	graphChan := make(chan *graph.Results, 100)
	resultsChan := make(chan *graph.Results, 100)

	for ww := 0; ww < 4; ww++ {
		go search.POSRoutine(graphChan, resultsChan, *testMode)
	}

	for name, g := range graphs {
		gr := new(graph.Results)
		gr.GraphObj = g
		gr.Name = name
		graphChan <- gr
	}

	close(graphChan)

	results := make(map[string]*graph.Results)
	for ii := 0; ii < len(graphs); ii++ {
		gr := <-resultsChan
		results[gr.Name] = gr
	}

	close(resultsChan)

	err = stream.ResultOut(results, *pathToTestBatch)
	if err != nil {
		log.Fatalln(err)
	}
}
