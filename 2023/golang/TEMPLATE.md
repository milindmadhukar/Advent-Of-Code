package day1

import "github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"

type day1 struct {
	data       []string
}

func (d day1) Part1() any {
	return 0
}

func (d day1) Part2() any {
	return 0
}

func Solve() day1 {
	data, err := utils.GetInputDataFromAOC(2023, 1)
	if err != nil {
		panic(err)
	}

	// exampleFile, _ := os.ReadFile("day1/example.txt")
	// data = utils.ParseFromString(string(exampleFile))

	return day1{
		data:       data,
	}
}

