package stream

import (
	"encoding/json"
	"fmt"
	"github.com/NESTLab/divisio.git/pkg/graph"
	"io/ioutil"
	"os"
)

func Out(graphs []*graph.Graph, directoryName string) error {
	for ii, g := range graphs {
		//Convert the graph object to json format
		jsonGraph, err := json.Marshal(g)
		if err != nil {
			return err
		}
		//Build the graph name. Each graph is named for the group it's in to avoid any duplicates
		graphName := fmt.Sprintf("%s/graph_%d.json", directoryName, ii+1)
		//Create the file
		fw, err := os.Create(graphName)
		if err != nil {
			return err
		}
		//Write the json to the file, then close the file
		fw.Write(jsonGraph)
		fw.Close()
	}
	return nil
}

func In(pathToGraphs string) (map[string]*graph.Graph, error) {
	graphs := make(map[string]*graph.Graph)
	files, err := ioutil.ReadDir(pathToGraphs)
	if err != nil {
		return graphs, err
	}

	for _, filename := range files {
		graphFile, err := ioutil.ReadFile(pathToGraphs + "/" + filename.Name())
		if err != nil {
			return graphs, err
		}
		var g graph.Graph
		err = json.Unmarshal(graphFile, &g)
		if err != nil {
			return graphs, err
		}
		graphs[filename.Name()] = &g
	}
	return graphs, nil
}
