package day02

import (
	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/utils"
)

type day02 struct {
	data      []string
}

func (d *day02) Part1() any {
	return 0
}

func (d *day02) Part2() any {
	return 0
}

func Solve() *day02 {
	data, err := utils.GetInputDataFromAOC(2025, 1)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day02/example.txt")

	return &day02{
		data:      data,
	}
}

