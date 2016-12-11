package plan

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Map struct {
	Start Point
	End   Point
	Grid  [][]int
	tree  map[Location][]Location
}

func CreateMap(fName string) *Map {
	file, err := os.Open(fName)
	if err != nil {
		log.Fatal(err)
	}
	m := &Map{}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	x := 0
	for scanner.Scan() {
		s := scanner.Text()
		m.Grid = append(m.Grid, make([]int, len(s)))
		y := 0
		for _, value := range s {
			myInt := int(value)
			m.Grid[x][y] = myInt
			y++
		}
		x++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	m.Start = Point{X: 2, Y: 2}
	m.End = Point{X: 1, Y: 5}
	return m
}

func (m *Map) DisplayMap() {
	for x := 0; x < len(m.Grid); x++ {
		for y := 0; y < len(m.Grid[0]); y++ {
			switch m.Grid[x][y] {
			case 119:
				fmt.Printf("%s", "w  ")
				break
			case 32:
				fmt.Printf("%s", "   ")
				break
			default:
				fmt.Printf("%2d ", m.Grid[x][y])

			}
		}
		fmt.Printf("\n")
	}
}

func (m *Map) GetNeighbors(current Point) (neighbors []Point) {
	if current.Y+1 < len(m.Grid[current.X]) && m.Grid[current.X][current.Y+1] == ' ' {
		neighbors = append(neighbors, Point{X: current.X, Y: current.Y + 1})
	}
	if current.Y-1 > 0 && m.Grid[current.X][current.Y-1] == ' ' {
		neighbors = append(neighbors, Point{X: current.X, Y: current.Y - 1})
	}
	if m.Grid[current.X+1][current.Y] == ' ' {
		neighbors = append(neighbors, Point{X: current.X + 1, Y: current.Y})
	}
	if current.X > 0 && m.Grid[current.X-1][current.Y] == ' ' {
		neighbors = append(neighbors, Point{X: current.X - 1, Y: current.Y})
	}
	return
}

func (m *Map) initTree() {
	m.tree[Location{lat: 34.978476, long: 135.964855}] =
		[]Location{Location{lat: 34.978584, long: 135.963530},
			Location{lat: 34.979322, long: 135.965021},
			Location{lat: 34.979322, long: 135.964260},
		}
	m.tree[Location{lat: 34.978584, long: 135.963530}] =
		[]Location{Location{lat: 34.978584, long: 135.963530},
			Location{lat: 34.979322, long: 135.965021},
			Location{lat: 34.979322, long: 135.964260},
		}
}
