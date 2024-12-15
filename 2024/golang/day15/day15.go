package day15

import (
	"os"
	"strings"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day15 struct {
	walls    map[Point]bool
	boxes    map[Point]bool
	robot    Point
	moves    []string
	gridSize Point
}

type Point struct {
	x, y int
}

func (d *day15) PrintGrid() {
	for y := 0; y < d.gridSize.y; y++ {
		for x := 0; x < d.gridSize.x; x++ {
			if wallFound := d.walls[Point{x, y}]; wallFound {
				print("#")
			} else if boxFound := d.boxes[Point{x, y}]; boxFound {
				print("O")
			} else if d.robot.x == x && d.robot.y == y {
				print("@")
			} else {
				print(".")
			}
		}
		println()
	}
	println()
}

func (d *day15) MoveBoxes(dx, dy int, boxes map[Point]bool) {
	if wallFound := d.walls[Point{d.robot.x + dx, d.robot.y + dy}]; wallFound {
		return
	}

	var boxesToMove []Point
	hitAWall := false

	tempPos := d.robot
	for {
		tempPos = Point{tempPos.x + dx, tempPos.y + dy}
		if boxFound := boxes[tempPos]; boxFound {
			if wallFound := d.walls[Point{tempPos.x + dx, tempPos.y + dy}]; wallFound {
				hitAWall = true
				break
			}
			boxesToMove = append(boxesToMove, tempPos)
		} else {
			break
		}
	}

	if hitAWall {
		return
	}

	for _, box := range boxesToMove {
		delete(boxes, box)
	}

	for _, box := range boxesToMove {
		boxes[Point{box.x + dx, box.y + dy}] = true
	}

	d.robot = Point{d.robot.x + dx, d.robot.y + dy}
}

func (d *day15) Part1() any {
	boxes := make(map[Point]bool)
	for k, v := range d.boxes {
		boxes[k] = v
	}

	for _, move := range d.moves {
		switch move {
		case "<":
			d.MoveBoxes(-1, 0, boxes)
		case ">":
			d.MoveBoxes(1, 0, boxes)
		case "^":
			d.MoveBoxes(0, -1, boxes)
		case "v":
			d.MoveBoxes(0, 1, boxes)
		}
	}

	sum := 0
	for box, isPresent := range boxes {
		if isPresent {
			sum += (100 * box.y) + box.x
		}
	}

	return sum
}

func (d *day15) Part2() any {
	return 0
}

func Solve() *day15 {
	data, err := utils.GetRawInputDataFromAOC(2024, 15)
	if err != nil {
		panic(err)
	}

	fileData, _ := os.ReadFile("day15/example.txt")
	data = string(fileData)
	data = strings.Trim(data, " ")
	data = strings.Trim(data, "\n")

	splitData := strings.Split(data, "\n\n")
	grid := utils.GetSplitData(strings.Split(splitData[0], "\n"), "")
	moves := strings.Split(splitData[1], "")
	walls := make(map[Point]bool)
	boxes := make(map[Point]bool)
	var robot Point

	for y, row := range grid {
		for x, cell := range row {
			if cell == "#" {
				walls[Point{x, y}] = true
			} else if cell == "O" {
				boxes[Point{x, y}] = true
			} else if cell == "@" {
				robot = Point{x, y}
			}
		}
	}

	maxY := len(grid)
	maxX := len(grid[0])

	gridSize := Point{maxX, maxY}

	return &day15{
		walls:    walls,
		boxes:    boxes,
		robot:    robot,
		moves:    moves,
		gridSize: gridSize,
	}
}
