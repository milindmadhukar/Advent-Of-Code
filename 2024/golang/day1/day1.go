package day1

import (
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day1 struct {
	data      []string
	startTime time.Time
}

func (d day1) Part1() any {
	return 0
}

func (d day1) Part2() any {
	return 0
}

func Solve() day1 {
	data, err := utils.GetInputDataFromAOC(2024, 1)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	// exampleFileData, _ := os.ReadFile("day1/example.txt")
	// data = utils.ParseFromString(string(exampleFile))

	return day1{
		data:      data,
		startTime: startTime,
	}
}

func (d day1) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
