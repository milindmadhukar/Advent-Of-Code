package day2

import (
	"math"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day2 struct {
	data      []string
	timeTaken time.Duration
	reports   [][]int
}

func isReportValid(report []int) bool {
	isValid := true
	var isAscending = false
	var isDescending = false

	diff := report[0] - report[1]

	if diff == 0 {
		return false
	}

	if diff < 0 {
		isDescending = true
	}
	if diff > 0 {
		isAscending = true
	}

	for idx := 0; idx < len(report)-1; idx++ {
		diff := report[idx] - report[idx+1]
		if isAscending && diff < 0 {
			isValid = false
			break
		}

		if isDescending && diff > 0 {
			isValid = false
			break
		}

		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			isValid = false
			break
		}
	}

	return isValid
}

func (d *day2) Part1() any {

	validCount := 0

	for _, report := range d.reports {
		isValid := isReportValid(report)
		if isValid {
			validCount++
		}

	}

	return validCount
}

func (d *day2) Part2() any {
	validCount := 0

	for _, report := range d.reports {
		isValid := isReportValid(report)

		if isValid {
			validCount++
		} else {
			canTolerate := false

			for idx := range report {
				newReport := make([]int, len(report))
				copy(newReport, report)

				newReport = append(newReport[:idx], newReport[idx+1:]...)
				isValid := isReportValid(newReport)

				if isValid {
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

	startTime := time.Now()

	splitData := utils.GetSplitData(data, " ")

	var reports [][]int

	for _, line := range splitData {
		reports = append(reports, utils.StringSliceToIntegerSlice(line))
	}

	endTime := time.Now()

	return &day2{
		data:      data,
		timeTaken: endTime.Sub(startTime),
		reports:   reports,
	}
}

func (d day2) TimeTaken() time.Duration {
	return d.timeTaken
}
