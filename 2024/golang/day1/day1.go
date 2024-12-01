package day1

import (
	"math"
	"slices"
	"strconv"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day1 struct {
	data      []string
	startTime time.Time
	leftList  []int
	rightlist []int
}

func (d day1) Part1() any {
  slices.Sort(d.leftList)
  slices.Sort(d.rightlist)

	distance := 0

	for idx := 0; idx < len(d.leftList); idx++ {
		diff := d.leftList[idx] - d.rightlist[idx]
		distance += int(math.Abs(float64(diff)))
	}

	return distance
}

func (d day1) Part2() any {
	similarity := 0

	for idx := 0; idx < len(d.leftList); idx++ {
		leftNum := d.leftList[idx]
		rightCount := utils.CountOf(d.rightlist, leftNum)
		similarity += (leftNum * rightCount)
	}

	return similarity
}

func Solve() day1 {
	data, err := utils.GetInputDataFromAOC(2024, 1)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day1/example.txt")

	startTime := time.Now()

	parsedData := utils.GetSplitData(data, "   ")
	var leftList, rightlist []int

	for _, val := range parsedData {
		lval, _ := strconv.Atoi(val[0])
		rval, _ := strconv.Atoi(val[1])
		leftList = append(leftList, lval)
		rightlist = append(rightlist, rval)
	}

	return day1{
		data:      data,
		startTime: startTime,
		leftList:  leftList,
		rightlist: rightlist,
	}
}

func (d day1) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
