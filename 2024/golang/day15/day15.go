package day15

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day15 struct {
	data []string
}

func (d *day15) Part1() any {
	return 0
}

func (d *day15) Part2() any {
	return 0
}

func Solve() *day15 {
	data, err := utils.GetInputDataFromAOC(2024, 15)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day15/example.txt")

	return &day15{
		data: data,
	}
}
