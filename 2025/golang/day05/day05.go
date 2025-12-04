package day05

import (
	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/utils"
)

type day05 struct {
	data      []string
}

func (d *day05) Part1() any {
	return 0
}

func (d *day05) Part2() any {
	return 0
}

func Solve() *day05 {
	data, err := utils.GetInputDataFromAOC(2025, 5)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day05/example.txt")

	return &day05{
		data:      data,
	}
}

