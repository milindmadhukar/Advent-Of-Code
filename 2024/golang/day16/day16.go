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

type PointDir struct {
	p      Point
	dirIdx int
}

var deltas = []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func (d *day16) Adjacents(current PointDir) func(yield func(PointDir, int) bool) {
	return func(yield func(PointDir, int) bool) {
		newDir1 := (current.dirIdx + 1) % 4
		if !yield(PointDir{p: current.p, dirIdx: newDir1}, 1000) {
			return
		}

		newDir2 := (current.dirIdx - 1 + 4) % 4
		if !yield(PointDir{p: current.p, dirIdx: newDir2}, 1000) {
			return
		}

		newPoint := Point{x: current.p.x + deltas[current.dirIdx].x, y: current.p.y + deltas[current.dirIdx].y}
		if d.maze[newPoint.y][newPoint.x] != "#" {
			if !yield(PointDir{p: newPoint, dirIdx: current.dirIdx}, 1) {
				return
			}
		}
	}
}

func popSmallest(pq map[PointDir]int) PointDir {
	var min = math.MaxInt
	var minPoint PointDir

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
	dist := make(map[PointDir]int)
	for y, row := range d.maze {
		for x := range row {
			for i := 0; i < 4; i++ {
				dist[PointDir{p: Point{x, y}, dirIdx: i}] = math.MaxInt
			}
		}
	}
	pq := make(map[PointDir]int)

	start := PointDir{p: d.start, dirIdx: 1}

	dist[start] = 0
	pq[start] = 0

	var current PointDir

	for len(pq) > 0 {
		current = popSmallest(pq)
		if current.p == d.end {
			break
		}

		for adjacent, distDelta := range d.Adjacents(current) {
			newDist := dist[current] + distDelta
			if newDist < dist[adjacent] {
				dist[adjacent] = newDist
				pq[adjacent] = dist[adjacent]
			}
		}
	}

	return dist[current]
}

func (d *day16) Part2() any {
	dist := make(map[PointDir]int)
	for y, row := range d.maze {
		for x := range row {
			for i := 0; i < 4; i++ {
				dist[PointDir{p: Point{x, y}, dirIdx: i}] = math.MaxInt
			}
		}
	}
	pq := make(map[PointDir]int)

	from := make(map[PointDir][]PointDir)

	start := PointDir{p: d.start, dirIdx: 1}

	dist[start] = 0
	pq[start] = 0

	var current PointDir

	for len(pq) > 0 {
		current = popSmallest(pq)

		for adjacent, distDelta := range d.Adjacents(current) {
			newDist := dist[current] + distDelta
			if newDist < dist[adjacent] {
				dist[adjacent] = newDist
				pq[adjacent] = dist[adjacent]
				from[adjacent] = []PointDir{current}
			} else if newDist <= dist[adjacent] {
				from[adjacent] = append(from[adjacent], current)
			}
		}
	}

  for k, v := range from {
    from[k] = utils.GetUniqueElements(v)
  }

	stack := []PointDir{{p: d.end, dirIdx: 2}}
	concernedNodes := make(map[PointDir]bool)
	concernedNodes[PointDir{p: d.end, dirIdx: 2}] = true

	for len(stack) > 0 {
		current, stack = utils.Pop(stack, len(stack)-1)

		for _, prev := range from[current] {
			if _, ok := concernedNodes[prev]; !ok {
				concernedNodes[prev] = true
				stack = append(stack, prev)
			}
		}
	}

	points := make(map[Point]bool)
	for point := range concernedNodes {
		points[point.p] = true
	}

	return len(points)
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
