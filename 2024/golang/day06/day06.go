package day06

import (
	"slices"
	"sync"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day06 struct {
	parsedData [][]string
	guardPos   Point
}

type Point struct {
	X int
	Y int
}

// Returns true if we looped.
func (d *day06) Traverse(data [][]string) (int, bool) {
	guardCurrentPos := d.guardPos

	yLimit := len(data)
	xLimit := len(data[0])

	delta := []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	deltaIdx := 0

	visited := make(map[Point]bool)
	visited[guardCurrentPos] = true

	hitOnTop := make(map[Point]bool)
	hitOnRight := make(map[Point]bool)
	hitOnBottom := make(map[Point]bool)
	hitOnLeft := make(map[Point]bool)

	for {
		nextX := guardCurrentPos.X + delta[deltaIdx].X
		nextY := guardCurrentPos.Y + delta[deltaIdx].Y

		if nextX >= xLimit || nextY >= yLimit || nextX < 0 || nextY < 0 {
			break
		} else if data[nextY][nextX] == "#" {

			switch deltaIdx {
			case 0:
				if hitOnTop[guardCurrentPos] {
					return -1, true
				}
				hitOnTop[guardCurrentPos] = true
			case 1:
				if hitOnRight[guardCurrentPos] {
					return -1, true
				}
				hitOnRight[guardCurrentPos] = true
			case 2:
				if hitOnBottom[guardCurrentPos] {
					return -1, true
				}
				hitOnBottom[guardCurrentPos] = true
			case 3:
				if hitOnLeft[guardCurrentPos] {
					return -1, true
				}
				hitOnLeft[guardCurrentPos] = true
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

func (d *day06) findGuardPos() Point {
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

func (d *day06) Part1() any {
	uniqueVisits, _ := d.Traverse(d.parsedData)
	return uniqueVisits
}

func (d *day06) Part2() any {
	count := 0
	var wg sync.WaitGroup

	yLimit := len(d.parsedData)
	xLimit := len(d.parsedData[0])
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

func Solve() *day06 {
	data, err := utils.GetInputDataFromAOC(2024, 6)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day06/example.txt")

	parsedData := utils.GetSplitData(data, "")

	d := day06{
		parsedData: parsedData,
	}
	d.guardPos = d.findGuardPos()

	return &d
}
