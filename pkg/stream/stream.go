package stream

import (
	"encoding/json"
	"fmt"
	"github.com/NESTLab/divisio.git/pkg/graph"
	"io/ioutil"
	"os"
	"regexp"
)

func GraphOut(graphs map[string]*graph.Graph, directoryName string) error {
	for name, g := range graphs {
		//Convert the graph object to json format
		jsonGraph, err := json.Marshal(g)
		if err != nil {
			return err
		}
		//Build the graph name. Each graph is in its own separate folder to group all relevant files together
		dirName := fmt.Sprintf("%s/graph_%s", directoryName, name)
		graphName := fmt.Sprintf("%s/graph.json", dirName)
		os.Mkdir(dirName, os.ModeDir|0777)

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

func ResultOut(results map[string]*graph.Results, directoryName string) error {
	for name, g := range results {
		//Convert the graph object to json format
		jsonGraph, err := json.Marshal(g)
		if err != nil {
			return err
		}
		//Build the graph name. Each graph is in its own separate folder to group all relevant files together
		graphName := fmt.Sprintf("%s/graph_%s/result.json", directoryName, name)

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

func GraphIn(pathToGraphs string) (map[string]*graph.Graph, error) {
	graphJSONCheck := regexp.MustCompile(`[\w]+.json`)
	graphNumCheck := regexp.MustCompile(`graph_([\d]+)`)
	graphs := make(map[string]*graph.Graph)
	runFiles, err := ioutil.ReadDir(pathToGraphs)
	if err != nil {
		return graphs, err
	}

	for _, graphDir := range runFiles {
		if graphDir.IsDir() {
			graphNumStr := graphNumCheck.FindStringSubmatch(graphDir.Name())[1]
			graphFiles, err := ioutil.ReadDir(pathToGraphs + "/" + graphDir.Name())
			if err != nil {
				return graphs, err
			}
			for _, file := range graphFiles {
				graphMatch := graphJSONCheck.FindString(file.Name())
				if graphMatch != "" {
					fullFileName := pathToGraphs + "/" + graphDir.Name() + "/" + graphMatch
					var g graph.Graph
					graphFile, err := ioutil.ReadFile(fullFileName)
					if err != nil {
						return graphs, err
					}
					err = json.Unmarshal(graphFile, &g)
					if err != nil {
						return graphs, err
					}
					graphs[graphNumStr] = &g
				}
			}
		}

	}
	return graphs, nil
}
