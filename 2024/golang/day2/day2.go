package day2

import (
	"math"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day2 struct {
	data      []string
	reports   [][]int
}

func isReportValid(report []int) bool {
	isValid := true
	isAscending := false
	isDescending := false

	diff := report[0] - report[1]

	if diff < 0 {
		isAscending = true
	} else if diff > 0 {
		isDescending = true
	} else {
		// Diff is 0
		return false
	}

	for idx := 0; idx < len(report)-1; idx++ {
		diff := report[idx] - report[idx+1]
		absDiff := math.Abs(float64(diff))
		if (isAscending && diff > 0) || (isDescending && diff < 0) || (absDiff < 1 || absDiff > 3) {
			isValid = false
			break
		}
	}

	return isValid
}

func (d *day2) Part1() any {
	validCount := 0
	for _, report := range d.reports {
		if isReportValid(report) {
			validCount++
		}
	}

	return validCount
}

func (d *day2) Part2() any {
	validCount := 0

	for _, report := range d.reports {
		if isReportValid(report) {
			validCount++
		} else {
			canTolerate := false
			for idx := range report {
				newReport := make([]int, len(report))
				copy(newReport, report)
				_, newReport = utils.Pop(newReport, idx)
				if isReportValid(newReport) {
					canTolerate = true
					break
				}
			}
			if canTolerate {
				validCount++
			}
		}
	}

	return validCount
}

func Solve() *day2 {

	data, err := utils.GetInputDataFromAOC(2024, 2)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day2/example.txt")

	splitData := utils.GetSplitData(data, " ")
	var reports [][]int
	for _, line := range splitData {
		reports = append(reports, utils.StringSliceToIntegerSlice(line))
	}

	return &day2{
		data:      data,
		reports:   reports,
	}
}
