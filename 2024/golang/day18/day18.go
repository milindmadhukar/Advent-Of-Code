package day18

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day18 struct {
	data      []string
}

func (d *day18) Part1() any {
	return 0
}

func (d *day18) Part2() any {
	return 0
}

func Solve() *day18 {
	data, err := utils.GetInputDataFromAOC(2024, 18)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day18/example.txt")

	return &day18{
		data:      data,
	}
}

