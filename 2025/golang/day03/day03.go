package day03

import (
	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/utils"
)

type day03 struct {
	data      []string
}

func (d *day03) Part1() any {
	return 0
}

func (d *day03) Part2() any {
	return 0
}

func Solve() *day03 {
	data, err := utils.GetInputDataFromAOC(2025, 3)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day03/example.txt")

	return &day03{
		data:      data,
	}
}

