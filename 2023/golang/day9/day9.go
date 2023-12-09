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

func nextValue(series []int, last int) int {
	if allZeroes(series) {
		return last
	}

	var differences []int
	for i := 0; i < len(series)-1; i++ {
		differences = append(differences, series[i+1]-series[i])
	}

	last = differences[len(differences)-1]

	return series[len(series)-1] + nextValue(differences, last)
}

func previousValue(series []int, first int) int {
	if allZeroes(series) {
		return first
	}

	var differences []int
	for i := 0; i < len(series)-1; i++ {
		differences = append(differences, series[i+1]-series[i])
	}

	first = differences[0]

	return series[0] - previousValue(differences, first)
}

func (d day9) Part1() any {
	sum := 0
	for _, series := range d.sequences {
		sum += nextValue(series, series[len(series)-1])
	}
	return sum
}

func (d day9) Part2() any {
	sum := 0
	for _, series := range d.sequences {
		sum += previousValue(series, series[0])
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
