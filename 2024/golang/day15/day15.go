package day15

import (
	"fmt"
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

func (d *day15) PrintGridPart1() {
	for y := 0; y < d.gridSize.y; y++ {
		for x := 0; x < d.gridSize.x; x++ {
			if wallFound := d.walls[Point{x, y}]; wallFound {
				fmt.Print("#")
			} else if boxFound := d.boxes[Point{x, y}]; boxFound {
				fmt.Print("O")
			} else if d.robot.x == x && d.robot.y == y {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *day15) PrintGridPart2(boxes map[Point]string) {
	for y := 0; y < d.gridSize.y; y++ {
		for x := 0; x < 2*d.gridSize.x; x++ {
			if wallFound := d.walls[Point{x, y}]; wallFound {
				fmt.Print("#")
			} else if boxFound := boxes[Point{x, y}]; boxFound != "" {
				fmt.Print(boxFound)
			} else if d.robot.x == x && d.robot.y == y {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()

}

func (d *day15) MoveBoxesPart1(dx, dy int) {
	if wallFound := d.walls[Point{d.robot.x + dx, d.robot.y + dy}]; wallFound {
		return
	}

	var boxesToMove []Point

	tempPos := d.robot
	for {
		tempPos = Point{tempPos.x + dx, tempPos.y + dy}
		if boxFound := d.boxes[tempPos]; boxFound {
			if wallFound := d.walls[Point{tempPos.x + dx, tempPos.y + dy}]; wallFound {
				return
			}
			boxesToMove = append(boxesToMove, tempPos)
		} else {
			break
		}
	}

	for _, box := range boxesToMove {
		delete(d.boxes, box)
	}

	for _, box := range boxesToMove {
		d.boxes[Point{box.x + dx, box.y + dy}] = true
	}

	d.robot = Point{d.robot.x + dx, d.robot.y + dy}
}

func (d *day15) Part1() any {
	boxesOriginal := make(map[Point]bool)
	for k, v := range d.boxes {
		boxesOriginal[k] = v
	}

	var robotOrignal Point = d.robot
	fmt.Println("Original Robot", robotOrignal)

	for _, move := range d.moves {
		switch move {
		case "<":
			d.MoveBoxesPart1(-1, 0)
		case ">":
			d.MoveBoxesPart1(1, 0)
		case "^":
			d.MoveBoxesPart1(0, -1)
		case "v":
			d.MoveBoxesPart1(0, 1)
		}

		// fmt.Println("Current Move", move)
		// d.PrintGrid(d.gridSize)
	}

	d.robot = robotOrignal
	d.boxes = boxesOriginal

	sum := 0
	for box, isPresent := range d.boxes {
		if isPresent {
			sum += (100 * box.y) + box.x
		}
	}

	return sum
}

func (d *day15) MoveBoxesPart2(dx, dy int, boxes map[Point]string) {
	if wallFound := d.walls[Point{d.robot.x + dx, d.robot.y + dy}]; wallFound {
		return
	}

	boxesToMove := make(map[Point]string)
	tempPos := d.robot

	if dy == 0 {
		for {
			tempPos = Point{tempPos.x + dx, tempPos.y + dy}
			if boxVal := boxes[tempPos]; boxVal != "" {
				if wallFound := d.walls[Point{tempPos.x + dx, tempPos.y + dy}]; wallFound {
					return
				}
				boxesToMove[tempPos] = boxVal
			} else {
				break
			}
		}
	} else {
		for {
			// Check phase
			for box := range boxesToMove {
				if wallFound := d.walls[Point{box.x + dx, box.y + dy}]; wallFound {
					return
				}
			}

			// Add phase
			if len(boxesToMove) == 0 {
				if boxFound := boxes[Point{tempPos.x + dx, tempPos.y + dy}]; boxFound != "" {
					if boxFound == "[" {
						boxesToMove[Point{tempPos.x, tempPos.y + dy}] = "["
						boxesToMove[Point{tempPos.x + +1, tempPos.y + dy}] = "]"
					} else if boxFound == "]" {
						boxesToMove[Point{tempPos.x - 1, tempPos.y + dy}] = "["
						boxesToMove[Point{tempPos.x, tempPos.y + dy}] = "]"
					}
				} else {
					break
				}
			} else {
				nextBoxes := make(map[Point]string)
				for k, v := range boxesToMove {
					nextBoxes[k] = v
				}
				oldLen := len(nextBoxes)
				for box := range boxesToMove {
					if boxFound := boxes[Point{box.x + dx, box.y + dy}]; boxFound != "" {
						if boxFound == "[" {
							nextBoxes[Point{box.x + dx, box.y + dy}] = "["
							nextBoxes[Point{box.x + 1, box.y + dy}] = "]"
						} else if boxFound == "]" {
							nextBoxes[Point{box.x - 1, box.y + dy}] = "["
							nextBoxes[Point{box.x, box.y + dy}] = "]"
						}
					}
				}
				if len(nextBoxes) == oldLen {
					break
				} else {
					for k, v := range nextBoxes {
						boxesToMove[k] = v
					}
				}
			}
		}
	}

	for box := range boxesToMove {
		delete(boxes, box)
	}

	for box, boxVal := range boxesToMove {
		boxes[Point{box.x + dx, box.y + dy}] = boxVal
	}

	d.robot = Point{d.robot.x + dx, d.robot.y + dy}
}

func (d *day15) Part2() any {
	walls := make(map[Point]bool)
	boxes := make(map[Point]string)

	for k, v := range d.walls {
		walls[Point{2 * k.x, k.y}] = v
		walls[Point{2*k.x + 1, k.y}] = v
	}

	for k := range d.boxes {
		boxes[Point{2 * k.x, k.y}] = "["
		boxes[Point{2*k.x + 1, k.y}] = "]"
	}

	d.robot = Point{2 * d.robot.x, d.robot.y}

	for p := range d.walls {
		delete(d.walls, p)
	}
	for p := range d.boxes {
		delete(d.boxes, p)
	}

	d.walls = walls

	for _, move := range d.moves {
		switch move {
		case "<":
			d.MoveBoxesPart2(-1, 0, boxes)
		case ">":
			d.MoveBoxesPart2(1, 0, boxes)
		case "^":
			d.MoveBoxesPart2(0, -1, boxes)
		case "v":
			d.MoveBoxesPart2(0, 1, boxes)
		}

		// fmt.Println("Current Move", move)
		// d.PrintGridPart2(boxes)
	}

	sum := 0
	for box, isPresent := range boxes {
		if isPresent == "[" {
			sum += (100 * box.y) + box.x
		}
	}

	return sum
}

func Solve() *day15 {
	data, err := utils.GetRawInputDataFromAOC(2024, 15)
	if err != nil {
		panic(err)
	}

	/*
		fileData, _ := os.ReadFile("day15/example.txt")
		data = string(fileData)
		data = strings.Trim(data, " ")
		data = strings.Trim(data, "\n")
	*/

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
