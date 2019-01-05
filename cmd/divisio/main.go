package main

import (
	"github.com/NESTLab/divisio.git/pkg/builder"
	"time"
)

func main() {
	g := builder.GraphBuilderRand(time.Now().Unix())

	g.PrintConnections()

}
