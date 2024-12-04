package day5

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day5 struct {
	data      []string
}

func (d *day5) Part1() any {
	return 0
}

func (d *day5) Part2() any {
	return 0
}

func Solve() *day5 {
	data, err := utils.GetInputDataFromAOC(2024, 5)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day5/example.txt")

	return &day5{
		data:      data,
	}
}

