package day7

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day7 struct {
	data      []string
}

func (d *day7) Part1() any {
	return 0
}

func (d *day7) Part2() any {
	return 0
}

func Solve() *day7 {
	data, err := utils.GetInputDataFromAOC(2024, 7)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day7/example.txt")

	return &day7{
		data:      data,
	}
}

