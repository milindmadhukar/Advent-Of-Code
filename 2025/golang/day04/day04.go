package day04

import (
	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/utils"
)

type day04 struct {
	data      []string
}

func (d *day04) Part1() any {
	return 0
}

func (d *day04) Part2() any {
	return 0
}

func Solve() *day04 {
	data, err := utils.GetInputDataFromAOC(2025, 4)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day04/example.txt")

	return &day04{
		data:      data,
	}
}

