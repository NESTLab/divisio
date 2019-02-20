package builder

import (
	"fmt"
	"github.com/NESTLab/divisio.git/pkg/graph"
	"github.com/NESTLab/divisio.git/pkg/stream"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

//GraphBuilderRand takes in an instance of rand.Rand and returns a reference to a graph.
//The random quantities and their ranges are as follows:
//numNodes: 10 - 14
//Chance the node is a Task: 0.3
//number of edges 2-4x number of nodes
//edge weight: 50-100
func GraphBuilderRand(randObj *rand.Rand) *graph.Graph {
	var g graph.Graph

	numNodes := randObj.Intn(5) + 10

	for ii := 0; ii < numNodes; ii++ {
		name := fmt.Sprintf("%d", ii)

		//30% chance of being a task. ~2-5 task nodes per graph
		if randObj.Intn(100) > 30 {
			g.AddNodeRand(randObj, name, true)
		} else {
			g.AddNodeRand(randObj, name, false)
		}
	}

	edgeNumFactor := randObj.Intn(2) + 2 //2-4x number of nodes

	for ii := 0; ii < edgeNumFactor*numNodes; ii++ {
		randWeight := randObj.Intn(50) + 50
		g.AddEdgeRand(randObj, randWeight)
	}

	fmt.Println("New graph built")
	fmt.Println("Number of Nodes: " + fmt.Sprintf("%d", numNodes))
	fmt.Println("Number of Edges: " + fmt.Sprintf("%d", edgeNumFactor*numNodes))

	return &g
}

//GenerateGraphs generates numGraphsToMake of random graphs, and stores them at path
func GenerateGraphs(numGraphsToMake int, path string) []*graph.Graph {
	graphs := make([]*graph.Graph, numGraphsToMake, numGraphsToMake)
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

	//Make the random generator
	randSeed := rand.NewSource(time.Now().Unix())
	randObj := rand.New(randSeed)

	//Each loop creates a new graph
	for ii := 0; ii < numGraphsToMake; ii++ {
		//Build the graph
		g := GraphBuilderRand(randObj)

		//Since the slice was pre-generated to the correct size, we don't need to append it
		graphs[ii] = g
	}

	stream.Out(graphs, path)
	return graphs
}
