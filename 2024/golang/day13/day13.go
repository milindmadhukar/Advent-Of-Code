package day13

import (
	"fmt"
	"strings"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day13 struct {
	clawMachines []ClawMachine
}

func CramersRule(a1, b1, c1, a2, b2, c2 int) (float64, float64, error) {
	det := float64((a1 * b2) - (a2 * b1))

	if det == 0 {
		return 0, 0, fmt.Errorf("Determinant is zero")
	}

	detX := float64((c1 * b2) - (c2 * b1))
	detY := float64((a1 * c2) - (a2 * c1))

	if detX/det != float64(int(detX/det)) || detY/det != float64(int(detY/det)) {
		return 0, 0, fmt.Errorf("Determinant is not a whole number")
	}

	return detX / det, detY / det, nil
}

func (d *day13) Part1() any {
	sum := 0
	for _, clawMachine := range d.clawMachines {
		aPresses, bPresses, err := CramersRule(clawMachine.aPressDelta.x, clawMachine.bPressDelta.x, clawMachine.prizePos.x, clawMachine.aPressDelta.y, clawMachine.bPressDelta.y, clawMachine.prizePos.y)
		if err != nil || (aPresses > 100 || bPresses > 100) {
			continue
		}

		sum += (int(aPresses) * 3) + int(bPresses)
	}

	return sum
}

func (d *day13) Part2() any {
	sum := 0
	for _, clawMachine := range d.clawMachines {
		aPresses, bPresses, err := CramersRule(clawMachine.aPressDelta.x, clawMachine.bPressDelta.x, clawMachine.prizePos.x+10000000000000, clawMachine.aPressDelta.y, clawMachine.bPressDelta.y, clawMachine.prizePos.y+10000000000000)
		if err != nil {
			continue
		}
		sum += (int(aPresses) * 3) + int(bPresses)
	}

	return sum
}

type Point struct {
	x, y int
}

type ClawMachine struct {
	aPressDelta Point
	bPressDelta Point
	prizePos    Point
}

func Solve() *day13 {
	data, err := utils.GetRawInputDataFromAOC(2024, 13)
	if err != nil {
		panic(err)
	}

	// data = utils.GetRawInputDataFromFile("day13/example.txt")

	clawMachinesStr := strings.Split(data, "\n\n")

	var clawMachines []ClawMachine
	for _, clawMachineStr := range clawMachinesStr {

		lines := strings.Split(clawMachineStr, "\n")

		var clawMachine ClawMachine
		var x, y int
		fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &x, &y)
		clawMachine.aPressDelta = Point{x, y}

		fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &x, &y)
		clawMachine.bPressDelta = Point{x, y}

		fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &x, &y)
		clawMachine.prizePos = Point{x, y}

		clawMachines = append(clawMachines, clawMachine)
	}

	return &day13{
		clawMachines: clawMachines,
	}
}
