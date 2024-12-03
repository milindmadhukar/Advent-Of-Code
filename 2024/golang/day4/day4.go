package day4

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day4 struct {
	data      []string
}

func (d *day4) Part1() any {
	return 0
}

func (d *day4) Part2() any {
	return 0
}

func Solve() *day4 {
	data, err := utils.GetInputDataFromAOC(2024, 4)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day4/example.txt")

	return &day4{
		data:      data,
	}
}

