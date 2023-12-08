package day9

import (
	"os"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day9 struct {
	data      []string
	startTime time.Time
}

func (d day9) Part1() any {
	return 0
}

func (d day9) Part2() any {
	return 0
}

func Solve() day9 {
	data, err := utils.GetInputDataFromAOC(2023, 9)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	exampleFile, _ := os.ReadFile("day9/example.txt")
	data = utils.ParseFromString(string(exampleFile))

	return day9{
		data:      data,
		startTime: startTime,
	}
}

func (d day9) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}

