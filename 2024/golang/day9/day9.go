package day9

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day9 struct {
	data      []string
}

func (d *day9) Part1() any {
	return 0
}

func (d *day9) Part2() any {
	return 0
}

func Solve() *day9 {
	data, err := utils.GetInputDataFromAOC(2024, 9)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day9/example.txt")

	return &day9{
		data:      data,
	}
}

