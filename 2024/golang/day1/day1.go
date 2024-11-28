package day1

import (
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day8 struct {
	data      []string
	startTime time.Time
}

func (d day8) Part1() any {
	return 0
}

func (d day8) Part2() any {
	return 0
}

func Solve() day8 {
	data, err := utils.GetInputDataFromAOC(2024, 1)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	// exampleFile, _ := os.ReadFile("day8/example.txt")
	// data = utils.ParseFromString(string(exampleFile))

	return day8{
		data:      data,
		startTime: startTime,
	}
}

func (d day8) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
