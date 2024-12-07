package day8

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day8 struct {
	data      []string
}

func (d *day8) Part1() any {
	return 0
}

func (d *day8) Part2() any {
	return 0
}

func Solve() *day8 {
	data, err := utils.GetInputDataFromAOC(2024, 8)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day8/example.txt")

	return &day8{
		data:      data,
	}
}

