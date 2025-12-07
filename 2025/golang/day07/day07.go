package day07

import (
	"fmt"
	"strings"

	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/utils"
)

type day07 struct {
	data   []string
	source Point
}

type Point struct {
	x, y int
}

func (d *day07) Part1() any {
	fmt.Println(d.source)

	splitterHitCount := 0

	beamsSources := make(map[Point]bool, 0)
	beamsSources[d.source] = true
	for len(beamsSources) > 0 {
		newBeamsSources := make(map[Point]bool, 0)
		for beamSource := range beamsSources {
			nextPoint := Point{x: beamSource.x, y: beamSource.y + 1}
			if nextPoint.y >= len(d.data) {
				continue
			}
			nextChar := d.data[nextPoint.y][nextPoint.x]
			switch nextChar {
			case '^':
				splitterHitCount++
				newBeamsSources[Point{x: nextPoint.x - 1, y: nextPoint.y}] = true
				newBeamsSources[Point{x: nextPoint.x + 1, y: nextPoint.y}] = true
			case '.':
				newBeamsSources[nextPoint] = true
			}
		}
		beamsSources = newBeamsSources
	}

	return splitterHitCount
}

type PathNode struct {
	point  Point
	left   *PathNode
	right  *PathNode
	isLeaf bool
}

func (d *day07) Part2() any {
	cache := make(map[Point]int)

	var countPaths func(p Point) int
	countPaths = func(p Point) int {
		if count, exists := cache[p]; exists {
			return count
		}

		nextY := p.y + 1
		if nextY >= len(d.data) {
			cache[p] = 1
			return 1
		}

		nextChar := d.data[nextY][p.x]
		var count int

		switch nextChar {
		case '^':
			count = countPaths(Point{x: p.x - 1, y: nextY}) +
				countPaths(Point{x: p.x + 1, y: nextY})
		case '.':
			count = countPaths(Point{x: p.x, y: nextY})
		default:
			count = 0
		}

		cache[p] = count
		return count
	}

	return countPaths(d.source)
}

func Solve() *day07 {
	data, err := utils.GetInputDataFromAOC(2025, 7)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day07/example.txt")

	source := Point{x: strings.Index(data[0], "S"), y: 0}

	return &day07{
		data:   data,
		source: source,
	}
}
