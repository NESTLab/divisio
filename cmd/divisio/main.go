package main

import (
	"encoding/json"
	"fmt"
	"github.com/NESTLab/divisio.git/pkg/builder"
	"github.com/NESTLab/divisio.git/pkg/graph"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	g := builder.GraphBuilderRand(time.Now().Unix())

	g.PrintConnections()

	graphMarshall, _ := json.Marshal(g)

	fmt.Println(string(graphMarshall))

	fileWriter, _ := os.Create("Output.json")
	defer fileWriter.Close()

	fileWriter.Write(graphMarshall)
	fileWriter.Close()

	graphFile, err := ioutil.ReadFile("Output.json")

	if err != nil {
		panic(err)
	}

	var g2 graph.Graph
	err = json.Unmarshal(graphFile, &g2)
	if err != nil {
		panic(err)
	}

	g2.PrintConnections()

}
