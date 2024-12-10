package day11

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day11 struct {
	data      []string
}

func (d *day11) Part1() any {
	return 0
}

func (d *day11) Part2() any {
	return 0
}

func Solve() *day11 {
	data, err := utils.GetInputDataFromAOC(2024, 11)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day11/example.txt")

	return &day11{
		data:      data,
	}
}

