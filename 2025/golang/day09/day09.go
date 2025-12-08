package day09

import (
	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/utils"
)

type day09 struct {
	data      []string
}

func (d *day09) Part1() any {
	return 0
}

func (d *day09) Part2() any {
	return 0
}

func Solve() *day09 {
	data, err := utils.GetInputDataFromAOC(2025, 9)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day09/example.txt")

	return &day09{
		data:      data,
	}
}

