package day10

import (
	"os"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day10 struct {
	data      []string
	startTime time.Time
}

func (d day10) Part1() any {
	return 0
}

func (d day10) Part2() any {
	return 0
}

func Solve() day10 {
	data, err := utils.GetInputDataFromAOC(2023, 10)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	exampleFile, _ := os.ReadFile("day10/example.txt")
	data = utils.ParseFromString(string(exampleFile))

	return day10{
		data:      data,
		startTime: startTime,
	}
}

func (d day10) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}

