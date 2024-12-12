package day13

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day13 struct {
	data      []string
}

func (d *day13) Part1() any {
	return 0
}

func (d *day13) Part2() any {
	return 0
}

func Solve() *day13 {
	data, err := utils.GetInputDataFromAOC(2024, 13)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day13/example.txt")

	return &day13{
		data:      data,
	}
}

