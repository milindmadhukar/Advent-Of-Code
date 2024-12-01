package day5

import (
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/utils"
)

type day5 struct {
	data       []string
	startTime  time.Time
	parsedData []int
}

func (d *day5) Part1() any {
	ptr := 0
	steps := 0
  offsets := make([]int, len(d.parsedData))
  copy(offsets, d.parsedData)
	for ptr >= 0 && ptr < len(offsets) {
		offset := offsets[ptr]
		offsets[ptr]++
		ptr += offset
		steps++
	}

	return steps
}

func (d *day5) Part2() any {
	ptr := 0
	steps := 0
  offsets := make([]int, len(d.parsedData))
  copy(offsets, d.parsedData)
	for ptr >= 0 && ptr < len(offsets) {
		offset := offsets[ptr]
		if offset >= 3 {
			offsets[ptr]--
		} else {
			offsets[ptr]++
		}
		ptr += offset
		steps++
	}

	return steps
}

func Solve() *day5 {
	data, err := utils.GetInputDataFromAOC(2017, 5)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	return &day5{
		data:       data,
		startTime:  startTime,
		parsedData: utils.StringSliceToIntegerSlice(data),
	}
}

func (d day5) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
