package day16

import (
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day16 struct {
	data      []string
}

func (d *day16) Part1() any {
	return 0
}

func (d *day16) Part2() any {
	return 0
}

func Solve() *day16 {
	data, err := utils.GetInputDataFromAOC(2024, 16)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day16/example.txt")

	return &day16{
		data:      data,
	}
}

