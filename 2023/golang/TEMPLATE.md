package day2

import "github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"

type day2 struct {
	data       []string
}

func (d day2) Part1() any {
	return 0
}

func (d day2) Part2() any {
	return 0
}

func Solve() day2 {
	data, err := utils.GetInputDataFromAOC(2023, 2)
	if err != nil {
		panic(err)
	}

	// exampleFile, _ := os.ReadFile("day2/example.txt")
	// data = utils.ParseFromString(string(exampleFile))

	return day2{
		data:       data,
	}
}

