package day6

import (
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day4 struct {
	data []string
  startTime time.Time
}

func (d day4) Part1() any {
	return 0
}

func (d day4) Part2() any {
	return 0
}

func Solve() day4 {
	data, err := utils.GetInputDataFromAOC(2023, 6)
	if err != nil {
		panic(err)
	}

  startTime := time.Now()

	// exampleFile, _ := os.ReadFile("day4/example.txt")
	// data = utils.ParseFromString(string(exampleFile))

	return day4{
		data: data,
    startTime: startTime,
	}
}

func (d day4) TimeTaken() time.Duration {
  return time.Since(d.startTime)
}

