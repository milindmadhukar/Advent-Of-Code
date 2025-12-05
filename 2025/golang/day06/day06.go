package day06

import (
	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/utils"
)

type day06 struct {
	data      []string
}

func (d *day06) Part1() any {
	return 0
}

func (d *day06) Part2() any {
	return 0
}

func Solve() *day06 {
	data, err := utils.GetInputDataFromAOC(2025, 6)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day06/example.txt")

	return &day06{
		data:      data,
	}
}

