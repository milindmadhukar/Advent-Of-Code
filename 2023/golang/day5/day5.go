package day5

import (
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day5 struct {
	data       []string
  startTune  time.Time
}

func (d day5) Part1() any {
	return 0
}

func (d day5) Part2() any {
	return 0
}

func Solve() day5 {
	data, err := utils.GetInputDataFromAOC(2023, 5)
	if err != nil {
		panic(err)
	}

  startTime := time.Now()

	// exampleFile, _ := os.ReadFile("day2/example.txt")
	// data = utils.ParseFromString(string(exampleFile))

	return day5{
		data:       data,
    startTune:  startTime,
	}
}

func (d day5) TimeTaken() time.Duration {
  return time.Since(d.startTune)
}
