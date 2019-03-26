package main

import (
	"encoding/json"
	"flag"
	"github.com/NESTLab/divisio.git/pkg/builder"
	"io/ioutil"
	"log"
	"strings"
)

var numGraphsToMake = flag.Int("numGraphs", 1, "Provide the number of graphs you wish to generate")
var pathToConfigFile = flag.String("config", "", "The path to the config json containing parameters on how to make graphs")

func main() {
	flag.Parse()

	if *numGraphsToMake <= 0 {
		log.Fatalln("Must give positive number of graphs to generate")
	}
	if *pathToConfigFile == "" {
		log.Fatalln("Must provide a valid path to config json file")
	}

	confFile, err := ioutil.ReadFile(*pathToConfigFile)
	if err != nil {
		log.Fatalf("Error reading in config file: %v", err)
	}

	var c builder.Config
	err = json.Unmarshal(confFile, &c)

	dirSlice := strings.Split(*pathToConfigFile, "/")
	dirSlice = append(dirSlice[:len(dirSlice)-1], dirSlice[len(dirSlice):]...)

	saveLocation := strings.Join(dirSlice, "/")

	c.GenerateGraphs(*numGraphsToMake, saveLocation)
}
