package day18

import (
	"fmt"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day18 struct {
	bytePositions         []Point
	gridSize              Point
	nanoSecondsToSimulate int
}

type Point struct {
	x, y int
}

var directions = []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

type QueueItem struct {
	point Point
	steps int
}

func (d *day18) BFS(memoryMap map[Point]bool) int {
	target := Point{d.gridSize.x - 1, d.gridSize.y - 1}
	visited := make(map[Point]bool)

	queue := []QueueItem{{point: Point{0, 0}, steps: 0}}
	visited[Point{0, 0}] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.point == target {
			return current.steps
		}

		for _, dir := range directions {
			next := Point{current.point.x + dir.x, current.point.y + dir.y}

			if next.x < 0 || next.x > d.gridSize.x-1 || next.y < 0 || next.y > d.gridSize.y-1 {
				continue
			}
			if memoryMap[next] || visited[next] {
				continue
			}

			visited[next] = true
			queue = append(queue, QueueItem{point: next, steps: current.steps + 1})
		}
	}

	return -1
}

func (d *day18) Part1() any {
	memoryMap := make(map[Point]bool)
	for i := 0; i < d.nanoSecondsToSimulate; i++ {
		memoryMap[d.bytePositions[i]] = true
	}

	return d.BFS(memoryMap)
}

func (d *day18) Part2() any {
	secondsSimulated := 0
	for {
		memoryMap := make(map[Point]bool)
		for i := 0; i < secondsSimulated; i++ {
			memoryMap[d.bytePositions[i]] = true
		}

		result := d.BFS(memoryMap)
		if result == -1 {
			bytePos := d.bytePositions[secondsSimulated-1]
			return fmt.Sprintf("%d,%d", bytePos.x, bytePos.y)
		}

		secondsSimulated++
	}
}

func Solve() *day18 {
	data, err := utils.GetInputDataFromAOC(2024, 18)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day18/example.txt")

	var bytePositions []Point
	for _, line := range data {
		var pos Point
		fmt.Sscanf(line, "%d,%d", &pos.x, &pos.y)
		bytePositions = append(bytePositions, pos)
	}

	// gridSize := Point{7, 7}
	// nanoSecondsToSimulate := 12

	gridSize := Point{71, 71}
	nanoSecondsToSimulate := 1024

	return &day18{
		bytePositions:         bytePositions,
		gridSize:              gridSize,
		nanoSecondsToSimulate: nanoSecondsToSimulate,
	}
}
