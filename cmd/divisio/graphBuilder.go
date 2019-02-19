package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/NESTLab/divisio.git/pkg/builder"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

var numGraphsToMake = flag.Int("numGraphs", 1, "Provide the number of graphs you wish to generate")

func main() {
	flag.Parse()

	//store the contents of 'graphs' as a slice of strings
	files, err := ioutil.ReadDir("graphs")
	if err != nil {
		log.Fatal(err)
	}

	//regex matcher for the directory names. Group () around the number is important for later
	var directoryCheck = regexp.MustCompile(`[\w]+([0-9]+)`)
	var maxDirNum int

	//Iterate over all the names of the files in the 'graphs' directory
	for _, f := range files {
		if f.IsDir() {
			//Pull out a slice of matches by checking the name to the directoryCheck regex
			//index 0 will be the full match, index 1 will be only the number
			val := directoryCheck.FindStringSubmatch(f.Name())
			if val != nil {
				ii, err := strconv.Atoi(val[1])
				if err != nil {
					log.Fatal(err)
				}
				//We're trying to find the largest number here
				if ii > maxDirNum {
					maxDirNum = ii
				}
			}

		}
	}
	//As this is human forward, we'll be nice and index by 1
	dirNum := maxDirNum + 1

	//Build the string for the path to the file
	dirName := fmt.Sprintf("graphs/group_%d", dirNum)

	//Make the file with free permissions
	os.Mkdir(dirName, os.ModeDir|0777)

	//Each loop creates a new graph
	for ii := 0; ii < *numGraphsToMake; ii++ {
		//Build the graph with time as the seed. Since this executes in under a second, in order to get true randomness
		//we add the index*1000 of the loop to the time.
		g := builder.GraphBuilderRand(time.Now().Unix() + int64(1000*ii))

		//Convert the graph object to json format
		g2, err := json.Marshal(g)
		if err != nil {
			log.Fatal(err)
		}
		//Build the graph name. Each graph is named for the group it's in to avoid any duplicates
		graphName := fmt.Sprintf("%s/group_%d_graph_%d.json", dirName, dirNum, ii+1)
		//Create the file
		fw, err := os.Create(graphName)
		if err != nil {
			log.Fatal(err)
		}
		//Write the json to the file, then close the file
		fw.Write(g2)
		fw.Close()
	}

}
