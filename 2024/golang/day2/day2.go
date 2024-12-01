package day2

import (
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day2 struct {
	data      []string
  timeTaken time.Duration
	startTime time.Time
}

func (d *day2) Part1() any {
	return 0
}

func (d *day2) Part2() any {
	return 0
}

func Solve() *day2 {
	data, err := utils.GetInputDataFromAOC(2024, 1)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	// data = utils.GetInputDataFromFile("day2/example.txt")

  endTime := time.Now()

	return &day2{
		data:      data,
		startTime: startTime,
    timeTaken: endTime.Sub(startTime),
	}
}

func (d day2) TimeTaken() time.Duration {
	return d.timeTaken
}



