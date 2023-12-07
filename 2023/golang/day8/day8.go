package day8

import (
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day8 struct {
	data       []string
  startTune  time.Time
}

func (d day8) Part1() any {
	return 0
}

func (d day8) Part2() any {
	return 0
}

func Solve() day8 {
	data, err := utils.GetInputDataFromAOC(2023, 8)
	if err != nil {
		panic(err)
	}

  startTime := time.Now()

	// exampleFile, _ := os.ReadFile("day2/example.txt")
	// data = utils.ParseFromString(string(exampleFile))

	return day8{
		data:       data,
    startTune:  startTime,
	}
}

func (d day8) TimeTaken() time.Duration {
  return time.Since(d.startTune)
}

