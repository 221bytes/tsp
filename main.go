package main

import (
	"os"

	"github.com/221bytes/tsp/algo"
	"github.com/221bytes/tsp/plan"
	"github.com/221bytes/tsp/visualisation"
)

func main() {
	var m *plan.Map
	if len(os.Args) == 2 {
		m = plan.CreateMap(os.Args[1])
	} else {
		m = plan.CreateMap("test/map.txt")
	}
	path := algo.AStar(m)
	x := len(path)

	for _, n := range path {
		m.Grid[n.X][n.Y] = x
		x--
	}
	m.DisplayMap()
	visualisation.Run(m)
}
