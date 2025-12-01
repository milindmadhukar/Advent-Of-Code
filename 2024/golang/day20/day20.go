package day20

import (
	"fmt"
	"strings"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day20 struct {
	maze      [][]string
	start     Point
	end       Point
	pathSteps int
	parents   map[Point]Point
	nexts     map[Point]Point
}

type Point struct {
	x, y int
}

type QueueItem struct {
	point Point
	steps int
}

var directions = []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func (d *day20) BFS(maze [][]string) (int, map[Point]Point, map[Point]Point) {
	queue := []QueueItem{{point: d.start, steps: 0}}
	visited := make(map[Point]bool)
	parents := make(map[Point]Point)
	nexts := make(map[Point]Point)

	visited[d.start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.point == d.end {
			return current.steps, parents, nexts
		}

		for _, dir := range directions {
			nextPoint := Point{current.point.x + dir.x, current.point.y + dir.y}
			if nextPoint.x < 0 || nextPoint.x > len(maze[0])-1 || nextPoint.y < 0 || nextPoint.y > len(maze)-1 {
				continue
			}

			if maze[nextPoint.y][nextPoint.x] == "#" || visited[nextPoint] {
				continue
			}

			visited[nextPoint] = true
			queue = append(queue, QueueItem{point: nextPoint, steps: current.steps + 1})
			parents[nextPoint] = current.point
			nexts[current.point] = nextPoint
		}
	}

	return -1, nil, nil
}

func (d *day20) Part1() any {
	current := d.start
  secondsElapsed := 0
	for current != d.end {
		fmt.Println(current, secondsElapsed)
		current = d.nexts[current]
    secondsElapsed++
	}
	fmt.Println(current, secondsElapsed)

	return 0
}

func (d *day20) Part2() any {
	return 0
}

func Solve() *day20 {
	data, err := utils.GetInputDataFromAOC(2024, 20)
	if err != nil {
		panic(err)
	}

	data = utils.GetInputDataFromFile("day20/example.txt")

	var maze [][]string

	var start, end Point

	for y, line := range data {
		currentLine := strings.Split(line, "")
		for x, char := range currentLine {
			if char == "S" {
				start = Point{x, y}
			} else if char == "E" {
				end = Point{x, y}
			}
		}
		maze = append(maze, currentLine)
	}

	d := day20{
		maze:  maze,
		start: start,
		end:   end,
	}

	d.pathSteps, d.parents, d.nexts = d.BFS(d.maze)

	return &d
}
