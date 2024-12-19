package day20

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day20 struct {
	data []string
}

func (d *day20) Part1() any {
	return 0
}

func (d *day20) Part2() any {
	return 0
}

func Solve() *day20 {
	data, err := utils.GetInputDataFromAOC(2024, 20)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day20/example.txt")

	return &day20{
		data: data,
	}
}
