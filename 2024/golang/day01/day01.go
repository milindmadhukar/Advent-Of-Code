package day01

import (
	"slices"
	"strconv"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day01 struct {
	leftList  []int
	rightlist []int
}

func (d day01) Part1() any {
	slices.Sort(d.leftList)
	slices.Sort(d.rightlist)

	distance := 0

	for idx := 0; idx < len(d.leftList); idx++ {
		diff := d.leftList[idx] - d.rightlist[idx]
		distance += utils.Abs(diff)
	}

	/*
		return utils.Sum(
			utils.Map(func(pair models.Pair[int, int]) int {
				return utils.Abs(pair.First - pair.Second)
			},
				utils.Zip(d.leftList, d.rightlist),
			),
		)
	*/

	return distance
}

func (d day01) Part2() any {
	similarity := 0

	countsRight := utils.CountOfAll(d.rightlist)
	for idx := 0; idx < len(d.leftList); idx++ {
		leftNum := d.leftList[idx]
		rightCount := countsRight[leftNum]
		similarity += (leftNum * rightCount)
	}

	return similarity
}

func Solve() day01 {
	data, err := utils.GetInputDataFromAOC(2024, 1)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day01/example.txt")

	parsedData := utils.GetSplitData(data, "   ")
	var leftList, rightlist []int

	for _, val := range parsedData {
		lval, _ := strconv.Atoi(val[0])
		rval, _ := strconv.Atoi(val[1])
		leftList = append(leftList, lval)
		rightlist = append(rightlist, rval)
	}

	return day01{
		leftList:  leftList,
		rightlist: rightlist,
	}
}
