package day1

import (
	"strings"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/utils"
)

type day1 struct {
	data       []string
	parsedData []int
	startTime  time.Time
}

func (d *day1) Part1() any {
	sum := 0
	for i := 0; i < len(d.parsedData)-1; i++ {
		if d.parsedData[i] == d.parsedData[(i+1)%len(d.parsedData)] {
			sum += d.parsedData[i]
		}
	}

	return sum
}

func (d *day1) Part2() any {
	sum := 0
	for i := 0; i < len(d.parsedData)-1; i++ {
		if d.parsedData[i] == d.parsedData[(i+(len(d.parsedData)/2))%len(d.parsedData)] {
			sum += d.parsedData[i]
		}
	}

	return sum
}

func Solve() *day1 {
	data, err := utils.GetInputDataFromAOC(2017, 1)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	// exampleFileData, _ := os.ReadFile("day1/example.txt")
	// data = utils.ParseFromString(string(exampleFileData))

	data = strings.Split(data[0], "")

	parsedData := utils.StringSliceToIntegerSlice(data)

	return &day1{
		data:       data,
		parsedData: parsedData,
		startTime:  startTime,
	}
}

func (d day1) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
