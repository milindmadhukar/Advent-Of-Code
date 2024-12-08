package day8

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day8 struct {
	data     []string
	antennas map[byte][]Point
}

type Point struct {
	x int
	y int
}

func (d *day8) Part1() any {
	antiNodes := make(map[Point]bool)
	for _, locations := range d.antennas {
		for pair := range utils.GenerateCombinations(locations, 2) {
			yDiff := pair[0].y - pair[1].y
			xDiff := pair[0].x - pair[1].x

			antiNode1 := Point{pair[0].x + xDiff, pair[0].y + yDiff}
			antiNode2 := Point{pair[1].x - xDiff, pair[1].y - yDiff}

			if antiNode1.x >= 0 && antiNode1.y >= 0 && antiNode1.x < len(d.data[0]) && antiNode1.y < len(d.data) {
				antiNodes[antiNode1] = true
			}
			if antiNode2.x >= 0 && antiNode2.y >= 0 && antiNode2.x < len(d.data[0]) && antiNode2.y < len(d.data) {
				antiNodes[antiNode2] = true
			}
		}
	}

	return len(antiNodes)
}

func (d *day8) Part2() any {
	antiNodes := make(map[Point]bool)
	for _, locations := range d.antennas {
		for pair := range utils.GenerateCombinations(locations, 2) {
			yDiff := pair[0].y - pair[1].y
			xDiff := pair[0].x - pair[1].x
      var antiNode Point

			dist := 0
			for {
				antiNode = Point{pair[0].x + (xDiff * dist), pair[0].y + (yDiff * dist)}
				if antiNode.x >= 0 && antiNode.y >= 0 && antiNode.x < len(d.data[0]) && antiNode.y < len(d.data) {
					antiNodes[antiNode] = true
					dist++
				} else {
					break
				}
			}

			dist = 0
			for {
				antiNode = Point{pair[1].x - (xDiff * dist), pair[1].y - (yDiff * dist)}
				if antiNode.x >= 0 && antiNode.y >= 0 && antiNode.x < len(d.data[0]) && antiNode.y < len(d.data) {
					antiNodes[antiNode] = true
					dist++
				} else {
					break
				}
			}
		}
	}

	return len(antiNodes)
}

func Solve() *day8 {
	data, err := utils.GetInputDataFromAOC(2024, 8)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day8/example.txt")

	antennas := make(map[byte][]Point)

	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] == '.' {
				continue
			}

			antennas[data[y][x]] = append(antennas[data[y][x]], Point{x, y})
		}
	}

	return &day8{
		data:     data,
		antennas: antennas,
	}
}
