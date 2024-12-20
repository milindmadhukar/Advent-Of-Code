package day20

import (
	"slices"
	"strings"
	"sync"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day20 struct {
	maze      [][]string
	start     Point
	end       Point
	pathSteps int
	parents   map[Point]Point
}

type Point struct {
	x, y int
}

type QueueItem struct {
	point Point
	steps int
}

var directions = []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func (d *day20) BFS(maze [][]string, secondsToCheat int) (int, map[Point]Point) {
	queue := []QueueItem{{point: d.start, steps: 0}}
	visited := make(map[Point]bool)
	next := make(map[Point]Point)

	visited[d.start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.point == d.end {
			return current.steps, next
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
			next[nextPoint] = current.point
		}
	}

	return -1, nil
}

func (d *day20) Part1() any {

	timeSaved := make(map[int]int)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for y, line := range d.maze {
		for x, char := range line {
			if x == 0 || y == 0 || x == len(d.maze[0])-1 || y == len(d.maze)-1 {
				continue
			}

			if char == "#" && char != "S" && char != "E" {
				wg.Add(1)
				go func() {
					defer wg.Done()
					var newMaze [][]string
					for _, line := range d.maze {
						newMaze = append(newMaze, slices.Clone(line))
					}
					newMaze[y][x] = "."
					steps, _ := d.BFS(newMaze, 2)
					if steps < d.pathSteps {
						mutex.Lock()
						timeSaved[d.pathSteps-steps]++
						mutex.Unlock()
					}
				}()
			}
		}
	}

	wg.Wait()

	cheatsThatSaveTimeLessThan100ps := 0

	for k, v := range timeSaved {
		if k >= 100 {
			cheatsThatSaveTimeLessThan100ps += v
		}
	}

	return cheatsThatSaveTimeLessThan100ps
}

func (d *day20) Part2() any {
	return 0
}

func Solve() *day20 {
	data, err := utils.GetInputDataFromAOC(2024, 20)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day20/example.txt")

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

	d.pathSteps, d.parents = d.BFS(d.maze)

	return &d
}
