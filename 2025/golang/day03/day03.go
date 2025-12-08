package day03

import (
	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/utils"
)

type day03 struct {
	batteries [][]string
}

func GreedySubsequenceSelector(bank []int, subsequenceLen int) int {
	currentLen := len(bank)
	countCanBeSkipped := currentLen - subsequenceLen
	if countCanBeSkipped < 0 || subsequenceLen == 0 {
		return 0
	}

	maxNum := bank[0]
	maxIdx := 0
	for i := 0; i <= countCanBeSkipped; i++ {
		if bank[i] > maxNum {
			maxNum = bank[i]
			maxIdx = i
		}
		if maxNum == 9 {
			break
		}
	}

	nextDigits := GreedySubsequenceSelector(bank[maxIdx+1:], subsequenceLen-1)
	multiplier := 1
	for range subsequenceLen - 1 {
		multiplier *= 10
	}

	return maxNum*multiplier + nextDigits
}

func (d *day03) Part1() any {
	joltageSum := 0
	for _, bank := range d.batteries {
		bankInt := utils.StringSliceToIntegerSlice(bank)
		maxNum := GreedySubsequenceSelector(bankInt, 2)
		joltageSum += maxNum
	}
	return joltageSum
}

func (d *day03) Part2() any {
	joltageSum := 0
	for _, bank := range d.batteries {
		bankInt := utils.StringSliceToIntegerSlice(bank)
		maxNum := GreedySubsequenceSelector(bankInt, 12)
		joltageSum += maxNum
	}
	return joltageSum
}

func Solve() *day03 {
	data, err := utils.GetInputDataFromAOC(2025, 3)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day03/example.txt")

	batteries := utils.GetSplitData(data, "")
	return &day03{
		batteries: batteries,
	}
}
