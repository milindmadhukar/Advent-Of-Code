package day10

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day10 struct {
	data      []string
}

func (d *day10) Part1() any {
	return 0
}

func (d *day10) Part2() any {
	return 0
}

func Solve() *day10 {
	data, err := utils.GetInputDataFromAOC(2024, 10)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day10/example.txt")

	return &day10{
		data:      data,
	}
}

