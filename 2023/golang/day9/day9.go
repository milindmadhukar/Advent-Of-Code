package day9

import (
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day9 struct {
	data      [][]string
	sequences [][]int
	startTime time.Time
}

func allZeroes(arr []int) bool {
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}

func nextValue(series []int) int {
	if allZeroes(series) {
		return series[len(series)-1]
	}

	var differences []int
	for i := 0; i < len(series)-1; i++ {
		differences = append(differences, series[i+1]-series[i])
	}

	return series[len(series)-1] + nextValue(differences)
}

func previousValue(series []int) int {
	if allZeroes(series) {
		return series[0]
	}

	var differences []int
	for i := 0; i < len(series)-1; i++ {
		differences = append(differences, series[i+1]-series[i])
	}

	return series[0] - previousValue(differences)
}

func (d day9) Part1() any {
	sum := 0
	for _, series := range d.sequences {
		sum += nextValue(series)
	}
	return sum
}

func (d day9) Part2() any {
	sum := 0
	for _, series := range d.sequences {
		sum += previousValue(series)
	}
	return sum
}

func Solve() day9 {
	rawData, err := utils.GetInputDataFromAOC(2023, 9)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	// exampleFile, _ := os.ReadFile("day9/example.txt")
	// rawData = utils.ParseFromString(string(exampleFile))

	data := utils.GetSplitData(rawData, " ")

	var sequences [][]int
	for _, line := range data {
		sequences = append(sequences, utils.StringSliceToIntegerSlice(line))
	}

	return day9{
		data:      data,
		sequences: sequences,
		startTime: startTime,
	}
}

func (d day9) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
