package day14

import (
	"fmt"
	"slices"
	"strings"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day14 struct {
	data     []string
	robots   []Robot
	gridSize Point
}

type Point struct {
	x, y int
}

type Robot struct {
	position Point
	velocity Point
}

func (d *day14) Part1() any {
	robots := slices.Clone(d.robots)

	secondsToSimulate := 100

	for i := 0; i < secondsToSimulate; i++ {
		for idx := range robots {
			robots[idx].position.x = (robots[idx].position.x + robots[idx].velocity.x) % d.gridSize.x
			robots[idx].position.y = (robots[idx].position.y + robots[idx].velocity.y) % d.gridSize.y

			if robots[idx].position.x < 0 {
				robots[idx].position.x = d.gridSize.x + robots[idx].position.x
			}

			if robots[idx].position.y < 0 {
				robots[idx].position.y = d.gridSize.y + robots[idx].position.y
			}
		}
	}

	midX := d.gridSize.x / 2
	midY := d.gridSize.y / 2
	var firstQuadrant, secondQuadrant, thirdQuadrant, fourthQuadrant int

	for _, robot := range robots {
		if robot.position.x > midX && robot.position.y < midY {
			firstQuadrant++
		} else if robot.position.x < midX && robot.position.y < midY {
			secondQuadrant++
		} else if robot.position.x < midX && robot.position.y > midY {
			thirdQuadrant++
		} else if robot.position.x > midX && robot.position.y > midY {
			fourthQuadrant++
		}
	}

	return firstQuadrant * secondQuadrant * thirdQuadrant * fourthQuadrant
}

func (d *day14) Part2() any {

	robots := slices.Clone(d.robots)

	var allPositions []map[Point]bool

	count := 0
	for {
		positions := make(map[Point]bool)
		for idx := range robots {
			robots[idx].position.x = (robots[idx].position.x + robots[idx].velocity.x) % d.gridSize.x
			robots[idx].position.y = (robots[idx].position.y + robots[idx].velocity.y) % d.gridSize.y

			if robots[idx].position.x < 0 {
				robots[idx].position.x = d.gridSize.x + robots[idx].position.x
			}

			if robots[idx].position.y < 0 {
				robots[idx].position.y = d.gridSize.y + robots[idx].position.y
			}
			positions[robots[idx].position] = true
		}

		allPositions = append(allPositions, positions)

		count++

		if len(positions) == len(robots) {
			break
		}
	}

	visualize(d.gridSize, allPositions[count-6:])

	return count
}

func Solve() *day14 {
	data, err := utils.GetInputDataFromAOC(2024, 14)
	if err != nil {
		panic(err)
	}

	var robots []Robot

	for _, line := range data {
		split := strings.Split(line, " ")
		var robot Robot
		fmt.Sscanf(split[0], "p=%d,%d", &robot.position.x, &robot.position.y)
		fmt.Sscanf(split[1], "v=%d,%d", &robot.velocity.x, &robot.velocity.y)
		robots = append(robots, robot)
	}

	// gridSize := Point{11, 7}
	gridSize := Point{101, 103}

	return &day14{
		data:     data,
		robots:   robots,
		gridSize: gridSize,
	}
}
