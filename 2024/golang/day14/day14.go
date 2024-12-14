package day14

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day14 struct {
	data []string
}

func (d *day14) Part1() any {
	return 0
}

func (d *day14) Part2() any {
	return 0
}

func Solve() *day14 {
	data, err := utils.GetInputDataFromAOC(2024, 14)
	if err != nil {
		panic(err)
	}

	data = utils.GetInputDataFromFile("day14/example.txt")

	return &day14{
		data: data,
	}
}
