package stream

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/NESTLab/divisio.git/pkg/graph"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
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

func VizOut(results map[string]*graph.Results, directoryName string) error {
	for graphName, gr := range results {
		output := "graph G {\n\toverlap=\"scale\"\n"
		for _, node := range gr.GraphObj.Nodes {
			if node.Rate > 0 {
				output += fmt.Sprintf("\t\"%s\" [shape=\"box\" label=\"%s\\nR=%d\\nD=%d\\nF=%d\"];\n", node.Name, node.Name, node.Rate, node.Difficulty, results[graphName].Results[node.Name])
			} else {
				output += fmt.Sprintf("\t\"%s\" [shape=\"circle\" label=\"%s\\nF=%d\"];\n", node.Name, node.Name, results[graphName].Results[node.Name])
			}

		}
		for originNode, edges := range gr.GraphObj.Edges {
			for _, edge := range edges {
				fromNum, err := strconv.Atoi(originNode)
				if err != nil {
					return err
				}
				toNum, err := strconv.Atoi(edge.ToNode)
				if err != nil {
					return err
				}
				if fromNum < toNum {
					output += fmt.Sprintf("\t\"%s\" -- \"%s\" [label=\"%d\"];\n", originNode, edge.ToNode, edge.Weight)
				}
			}
		}
		output += "\n}"
		//Build the graph name. Each graph is in its own separate folder to group all relevant files together
		graphPath := fmt.Sprintf("%s/graph_%s/visual.gv", directoryName, graphName)
		visualName := fmt.Sprintf("%s/graph_%s/image.png", directoryName, graphName)

		//Create the file
		fw, err := os.Create(graphPath)
		if err != nil {
			return err
		}
		//Write the json to the file, then close the file
		fw.Write([]byte(output))
		fw.Close()

		cmd := exec.Command("neato", "-Tpng", graphPath, "-o", visualName)
		var out bytes.Buffer
		cmd.Stderr = &out
		err = cmd.Run()
		if err != nil {
			fmt.Printf("%q\n", out.String())
			return err
		}

	}
	return nil
}

func GraphIn(pathToGraphs string) (map[string]*graph.Graph, error) {
	graphJSONCheck := regexp.MustCompile(`graph.json`)
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
