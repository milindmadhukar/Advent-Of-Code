package day16

import (
	"fmt"
	"math"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day16 struct {
	maze       [][]string
	start, end Point
}

type Point struct {
	x, y int
}

var deltas = []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func popSmallest(pq map[Point]int) Point {
	var min = math.MaxInt
	var minPoint Point

	for k, v := range pq {
		if v < min {
			min = v
			minPoint = k
		}
	}

	delete(pq, minPoint)
	return minPoint
}

func (d *day16) Part1() any {
	visited := make(map[Point]bool)
	dist := make(map[Point]int)

	for y, row := range d.maze {
		for x := range row {
			dist[Point{x, y}] = math.MaxInt
		}
	}

	visited[d.start] = true
	dist[d.start] = 0

	pq := make(map[Point]int)
	nodeDeltaIdx := make(map[Point]int)

	pq[d.start] = 0
	nodeDeltaIdx[d.start] = 1

	var current Point

	for len(pq) > 0 {
		current = popSmallest(pq)

		if current == d.end {
			break
		}

		for idx, delta := range deltas {
			neighbor := Point{current.x + delta.x, current.y + delta.y}
			if d.maze[neighbor.y][neighbor.x] == "#" {
				continue
			}

			if _, ok := visited[neighbor]; ok {
				continue
			}

			diff := utils.Abs(idx - nodeDeltaIdx[current])
			newDist := dist[current] + (diff * 1000) + 1
			if newDist < dist[neighbor] {
				dist[neighbor] = newDist
				nodeDeltaIdx[neighbor] = idx
				pq[neighbor] = dist[neighbor]
			}
		}
	}

	return dist[d.end]
}

func (d *day16) Part2() any {
	return 0
}

func Solve() *day16 {
	data, err := utils.GetInputDataFromAOC(2024, 16)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day16/example.txt")

	maze := utils.GetSplitData(data, "")
	var start, end Point

	for y, row := range maze {
		for x, col := range row {
			if col == "S" {
				start = Point{x, y}
			} else if col == "E" {
				end = Point{x, y}
			}

		}
	}
  
  fmt.Println(start, end)

	return &day16{
		maze:  maze,
		start: start,
		end:   end,
	}
}
