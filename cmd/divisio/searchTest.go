package main

import (
	"encoding/json"
	"fmt"
	"github.com/NESTLab/divisio.git/pkg/graph"
	"github.com/NESTLab/divisio.git/pkg/search"
	"io/ioutil"
)

func main() {
	graphFile, err := ioutil.ReadFile("graphs/misc/example_one.json")

	if err != nil {
		panic(err)
	}

	var g2 graph.Graph
	err = json.Unmarshal(graphFile, &g2)
	if err != nil {
		panic(err)
	}

	g2.PrintConnections()

	passes := search.PostOfficeSelection(g2, search.AStarMode)
	fmt.Println(passes)

}
