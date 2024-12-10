package day10

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type Point struct {
	x int
	y int
}

var deltas = []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

type day10 struct {
	data           []string
	topologicalMap [][]int
	trailheads     []Point
	visited9s      map[Point]int
}

func (d *day10) Traverse(currentPos Point, reached9 map[Point]int) {
	currentHeight := d.topologicalMap[currentPos.y][currentPos.x]
	if currentHeight == 9 {
		reached9[Point{currentPos.x, currentPos.y}] += 1
	}

	for _, delta := range deltas {
		nextPos := Point{currentPos.x + delta.x, currentPos.y + delta.y}
		if nextPos.x < 0 || nextPos.x >= len(d.topologicalMap[0]) || nextPos.y < 0 || nextPos.y >= len(d.topologicalMap) {
			continue
		}
		if d.topologicalMap[nextPos.y][nextPos.x] == currentHeight+1 {
			d.Traverse(nextPos, reached9)
		}
	}
}

func (d *day10) Part1() any {
	sum := 0
	for _, trailhead := range d.trailheads {
		visited9s := make(map[Point]int)
		d.Traverse(trailhead, visited9s)
		sum += len(visited9s)
	}
	return sum
}

func (d *day10) Part2() any {
	visited9s := make(map[Point]int)
	for _, trailhead := range d.trailheads {
		d.Traverse(trailhead, visited9s)
	}
	sum := 0
	for _, count := range visited9s {
		sum += count
	}
	return sum
}

func Solve() *day10 {
	data, err := utils.GetInputDataFromAOC(2024, 10)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day10/example.txt")

	splitData := utils.GetSplitData(data, "")

	var trailheads []Point
	for y, line := range splitData {
		for x, char := range line {
			if char == "0" {
				trailheads = append(trailheads, Point{x, y})
			}
		}
	}

	var topologicalMap [][]int
	for _, line := range splitData {
		topologicalMap = append(topologicalMap, utils.StringSliceToIntegerSlice(line))
	}

	return &day10{
		data:           data,
		topologicalMap: topologicalMap,
		trailheads:     trailheads,
	}
}
