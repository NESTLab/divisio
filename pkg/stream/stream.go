package stream

import (
	"encoding/json"
	"fmt"
	"github.com/NESTLab/divisio.git/pkg/graph"
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
