package algo

import (
	"container/heap"
	"fmt"

	"github.com/221bytes/tsp/plan"
)

func getPath(cameFrom map[plan.Point]plan.Point, current *plan.Item, start plan.Point) (path []plan.Point) {
	path = []plan.Point{}
	fmt.Printf("%s", "Path has been found\n")
	for !current.P.IsEqual(start) {
		current.P = cameFrom[current.P]
		path = append(path, current.P)
	}
	return
}

func AStar(m *plan.Map) []plan.Point {
	openQueue := &plan.PriorityQueue{}
	heap.Init(openQueue)
	cost_so_far := make(map[plan.Point]int)
	cameFrom := make(map[plan.Point]plan.Point)
	start := m.Start
	end := m.End
	heap.Push(openQueue, &plan.Item{P: start, Priority: 0})
	cost_so_far[start] = 0
	cameFrom[start] = start
	for {
		if openQueue.Len() == 0 {
			fmt.Printf("%s\n", "no path")
			for _, n := range cameFrom {
				m.Grid[n.X][n.Y] = 'x'
			}
			m.Grid[start.X][start.Y] = 'S'
			m.Grid[end.X][end.Y] = 'E'
			return nil
		}
		current := heap.Pop(openQueue).(*plan.Item)

		if current.P.IsEqual(end) {
			return getPath(cameFrom, current, start)
		}

		for _, neighborPoint := range m.GetNeighbors(current.P) {
			cost := plan.AddPoint(&neighborPoint, &end)
			if _, ok := cost_so_far[neighborPoint]; ok == false || 100000-cost < cost_so_far[neighborPoint] {
				cost_so_far[neighborPoint] = 100000 - cost
				neighbor := &plan.Item{P: neighborPoint, Priority: 100000 - cost}
				heap.Push(openQueue, neighbor)
				cameFrom[neighborPoint] = current.P
			}
		}
	}
}
