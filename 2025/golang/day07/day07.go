package day07

import (
	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/utils"
)

type day07 struct {
	data      []string
}

func (d *day07) Part1() any {
	return 0
}

func (d *day07) Part2() any {
	return 0
}

func Solve() *day07 {
	data, err := utils.GetInputDataFromAOC(2025, 7)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day07/example.txt")

	return &day07{
		data:      data,
	}
}

