package day6

import (
	"slices"
	"sync"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day6 struct {
	data       []string
	parsedData [][]string
	guardPos   Point
}

type Point struct {
	X int
	Y int
}

// Returns true if we looped.
func (d *day6) Traverse(data [][]string) (int, bool) {
	guardCurrentPos := d.guardPos

	yLimit := len(data)
	xLimit := len(data[0])

	delta := []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	deltaIdx := 0

	visited := make(map[Point]bool)
	visited[guardCurrentPos] = true

	var hitOnTop []Point
	var hitOnRight []Point
	var hitOnBottom []Point
	var hitOnLeft []Point

	for {
		nextX := guardCurrentPos.X + delta[deltaIdx].X
		nextY := guardCurrentPos.Y + delta[deltaIdx].Y

		if nextX >= xLimit || nextY >= yLimit || nextX < 0 || nextY < 0 {
			break
		} else if data[nextY][nextX] == "#" {

			switch deltaIdx {
			case 0:
				if utils.Contains(hitOnTop, guardCurrentPos) {
					return -1, true
				}
				hitOnTop = append(hitOnTop, guardCurrentPos)
			case 1:
				if utils.Contains(hitOnRight, guardCurrentPos) {
					return -1, true
				}
				hitOnRight = append(hitOnRight, guardCurrentPos)
			case 2:
				if utils.Contains(hitOnBottom, guardCurrentPos) {
					return -1, true
				}
				hitOnBottom = append(hitOnBottom, guardCurrentPos)
			case 3:
				if utils.Contains(hitOnLeft, guardCurrentPos) {
					return -1, true
				}
				hitOnLeft = append(hitOnLeft, guardCurrentPos)
			}
			deltaIdx = (deltaIdx + 1) % 4
		} else {
			guardCurrentPos.X = nextX
			guardCurrentPos.Y = nextY
			visited[guardCurrentPos] = true
		}
	}

	uniqueVisits := len(visited)
	return uniqueVisits, false
}

func (d *day6) findGuardPos() Point {
	for y := 0; y < len(d.parsedData); y++ {
		line := d.parsedData[y]
		for x := 0; x < len(line); x++ {
			if d.parsedData[y][x] == "^" {
				return Point{X: x, Y: y}
			}
		}
	}
	return Point{}
}
func (d *day6) Part1() any {
	uniqueVisits, _ := d.Traverse(d.parsedData)
	return uniqueVisits
}

func (d *day6) Part2() any {
	count := 0
	var wg sync.WaitGroup

	yLimit := len(d.data)
	xLimit := len(d.data[0])
	wg.Add(xLimit * yLimit)

	for y := 0; y < yLimit; y++ {
		for x := 0; x < xLimit; x++ {
			go func() {
				defer wg.Done()
				if d.parsedData[y][x] == "#" || d.parsedData[y][x] == "^" {
					return
				}
				var newData [][]string
				for _, line := range d.parsedData {
					newData = append(newData, slices.Clone(line))
				}
				newData[y][x] = "#"
				if _, looped := d.Traverse(newData); looped {
					count++
				}
			}()
		}
	}

	wg.Wait()

	return count
}

func Solve() *day6 {
	data, err := utils.GetInputDataFromAOC(2024, 6)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day6/example.txt")

	parsedData := utils.GetSplitData(data, "")

	d := day6{
		data:       data,
		parsedData: parsedData,
	}
	d.guardPos = d.findGuardPos()

	return &d
}
