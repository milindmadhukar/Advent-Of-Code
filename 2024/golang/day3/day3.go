package day3

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day3 struct {
	data      []string
}

func (d *day3) Part1() any {
	return 0
}

func (d *day3) Part2() any {
	return 0
}

func Solve() *day3 {
	data, err := utils.GetInputDataFromAOC(2024, 3)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day3/example.txt")

	return &day3{
		data:      data,
	}
}
