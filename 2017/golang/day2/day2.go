package day2

import (
	"math"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/utils"
)

type day2 struct {
	data       []string
	startTime  time.Time
	parsedData [][]int
}

func (d *day2) Part1() any {

	checksum := 0

	for _, line := range d.parsedData {
		smallest := math.MaxInt
		largest := math.MinInt
		for _, num := range line {
			if num > largest {
				largest = num
			}
			if num < smallest {
				smallest = num
			}
		}
		checksum += (largest - smallest)
	}

	return checksum
}

func (d *day2) Part2() any {
	checksum := 0

	for _, line := range d.parsedData {
		found := false
		for idx1, num1 := range line {
			for idx2, num2 := range line {
				if idx1 == idx2 {
					continue
				}
				if num1%num2 == 0 {
					found = true
					checksum += num1 / num2
					break
				}
			}
			if found {
				break
			}
		}
	}

	return checksum
}

func Solve() *day2 {
	data, err := utils.GetInputDataFromAOC(2017, 2)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	splitData := utils.GetSplitData(data, "\t")

	var parsedData [][]int
	for _, val := range splitData {
		parsedData = append(parsedData, utils.StringSliceToIntegerSlice(val))
	}

	return &day2{
		data:       data,
		startTime:  startTime,
		parsedData: parsedData,
	}
}

func (d day2) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
