package day17

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day17 struct {
	data []string
}

func (d *day17) Part1() any {
	return 0
}

func (d *day17) Part2() any {
	return 0
}

func Solve() *day17 {
	data, err := utils.GetInputDataFromAOC(2024, 17)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day17/example.txt")

	return &day17{
		data: data,
	}
}
