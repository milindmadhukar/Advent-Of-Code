package day12

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day12 struct {
	data      []string
}

func (d *day12) Part1() any {
	return 0
}

func (d *day12) Part2() any {
	return 0
}

func Solve() *day12 {
	data, err := utils.GetInputDataFromAOC(2024, 12)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day12/example.txt")

	return &day12{
		data:      data,
	}
}

