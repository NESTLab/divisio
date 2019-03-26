package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/NESTLab/divisio.git/pkg/builder"
	"log"
	"os"
	"strconv"
	"strings"
)

var ConfigFileSaveLocation = flag.String("path", "", "The path where you wish to save the config json file")

func main() {
	flag.Parse()

	if *ConfigFileSaveLocation == "" {
		log.Fatalln("Must provide a path to the config file")
	}

	var c builder.Config
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Querying for new Graph Config")

	for {
		//Number of nodes
		c.BaseNumNodes, _ = strconv.Atoi(queryValue(reader, "Minimum number of Nodes: ", false))
		maxNodes, _ := strconv.Atoi(queryValue(reader, "Maximum number of Nodes: ", false))
		if c.BaseNumNodes < maxNodes {
			c.ExtraNumNodes = maxNodes - c.BaseNumNodes
			break
		}

	}

	for {
		//Difficulty of Task Nodes
		c.BaseNodeDiff, _ = strconv.Atoi(queryValue(reader, "Minimum difficulty of Task Nodes: ", false))
		maxDiff, _ := strconv.Atoi(queryValue(reader, "Maximum difficulty of Task Nodes: ", false))
		if c.BaseNodeDiff < maxDiff {
			c.ExtraNodeDiff = maxDiff - c.BaseNodeDiff
			break
		}

	}

	for {
		//Rate of Task Nodes
		c.BaseNodeRate, _ = strconv.Atoi(queryValue(reader, "Minimum rate of Task Nodes: ", false))
		maxRate, _ := strconv.Atoi(queryValue(reader, "Maximum rate of Task Nodes: ", false))
		if c.BaseNodeRate < maxRate {
			c.ExtraNodeRate = maxRate - c.BaseNodeRate
			break
		}

	}

	for {
		//Chance of Crossroads
		c.ChanceCrossRoads, _ = strconv.ParseFloat(queryValue(reader, "0.0 - 1.0 chance that a ndoe will be a crossroads: ", true), 64)
		if c.ChanceCrossRoads < 1.0 && c.ChanceCrossRoads >= 0.0 {
			break
		}
	}

	for {
		//Number of Edges
		c.BaseNumEdges, _ = strconv.Atoi(queryValue(reader, "Minimum number of Edges: ", false))
		maxEdges, _ := strconv.Atoi(queryValue(reader, "Maximum number of Edges: ", false))
		if c.BaseNumEdges < maxEdges {
			c.ExtraNumEdges = maxEdges - c.BaseNumEdges
			break
		}

	}

	for {
		//Number of Edges
		c.BaseEdgeWeight, _ = strconv.Atoi(queryValue(reader, "Minimum weight of Edges: ", false))
		maxWeight, _ := strconv.Atoi(queryValue(reader, "Maximum weight of Edges: ", false))
		if c.BaseEdgeWeight < maxWeight {
			c.ExtraNumEdges = maxWeight - c.BaseEdgeWeight
			break
		}

	}

	jsonConfig, err := json.Marshal(c)
	if err != nil {
		log.Fatalf("Error marshalling config json: %v", err)
	}
	fw, err := os.Create(*ConfigFileSaveLocation)
	defer fw.Close()
	if err != nil {
		log.Fatalf("Error creating config file: %v", err)
	}
	//Write the json to the file, then close the file
	fw.Write(jsonConfig)
	if err != nil {
		log.Fatalf("Error streaming out config file: %v", err)
	}

}

func queryValue(r *bufio.Reader, prompt string, isFloat bool) string {
	for {
		fmt.Print(prompt)
		text, err := r.ReadString('\n')
		if err != nil {
			log.Fatalf("Error on QueryValue with prompt: %s\n", prompt)
		}
		text = strings.Replace(text, "\n", "", -1)
		var errNum error
		if isFloat {
			_, errNum = strconv.ParseFloat(text, 64)
		} else {
			_, errNum = strconv.Atoi(text)
		}
		if errNum == nil {
			return text
		}
	}
}
