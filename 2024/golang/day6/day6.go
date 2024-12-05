package day6

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day6 struct {
	data      []string
}

func (d *day6) Part1() any {
	return 0
}

func (d *day6) Part2() any {
	return 0
}

func Solve() *day6 {
	data, err := utils.GetInputDataFromAOC(2024, 6)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day6/example.txt")

	return &day6{
		data:      data,
	}
}

