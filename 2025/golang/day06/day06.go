package day06

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/utils"
)

type day06 struct {
	data          []string
	numbersStrCol [][]string
	numbersCol    [][]int
	operators     []string
}

func (d *day06) Part1() any {
	var grandTotal uint64 = 0

	for colIdx, operator := range d.operators {
		if operator == "+" {
			grandTotal += uint64(utils.Sum(d.numbersCol[colIdx]))
		} else if operator == "*" {
			grandTotal += uint64(utils.Product(d.numbersCol[colIdx]))
		}
	}

	return grandTotal
}

func (d *day06) Part2() any {
	var grandTotal uint64 = 0

	for operatorIdx, operator := range d.operators {
		numbersStr := utils.Transpose(utils.GetSplitData(d.numbersStrCol[operatorIdx], ""))
		numbersToOperate := make([]int, len(numbersStr))
		for numberStrIdx, numberStrParts := range numbersStr {
			currentNumberStr := ""
			for _, part := range numberStrParts {
				if part != " " {
					currentNumberStr += part
				}
			}
			number, _ := strconv.Atoi(currentNumberStr)
			numbersToOperate[numberStrIdx] = number
		}

		if operator == "+" {
			grandTotal += uint64(utils.Sum(numbersToOperate))
		} else if operator == "*" {
			grandTotal += uint64(utils.Product(numbersToOperate))
		}
	}

	return grandTotal
}

func Solve() *day06 {
	data, err := utils.GetInputDataFromAOC(2025, 6)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day06/example.txt")

	operatorsLine := data[len(data)-1]
	operators := regexp.MustCompile(`[+*]`).FindAllString(operatorsLine, -1)
	operatorIndexRanges := regexp.MustCompile(`[+*]`).FindAllStringIndex(operatorsLine, -1)
	operatorIndexes := utils.Map(
		func(r []int) int {
			return r[0]
		}, operatorIndexRanges)

	numbersList := data[:len(data)-1]
	numbersStr := make([][]string, len(numbersList))

	idxStart := 0
	for operatorStrIdx := 1; operatorStrIdx <= len(operatorIndexRanges); operatorStrIdx++ {
		var idxEnd int
		if operatorStrIdx < len(operatorIndexRanges) {
			idxEnd = operatorIndexes[operatorStrIdx] - 1
		}
		for lineIdx, line := range numbersList {
			var numberStr string
			if operatorStrIdx == len(operatorIndexRanges) {
				numberStr = line[idxStart:]
			} else {
				numberStr = line[idxStart:idxEnd]
			}
			numbersStr[lineIdx] = append(numbersStr[lineIdx], numberStr)
		}
		if operatorStrIdx < len(operatorIndexRanges) {
			idxStart = operatorIndexes[operatorStrIdx]
		}
	}

	numbersStrCol := utils.Transpose(numbersStr)
	numbersCol := make([][]int, len(numbersStrCol))

	for colIdx, numbersStr := range numbersStrCol {
		numbersCol[colIdx] = utils.Map(func(s string) int {
			num, _ := strconv.Atoi(strings.TrimSpace(s))
			return num
		}, numbersStr)
	}

	return &day06{
		data:          data,
		numbersCol:    numbersCol,
		numbersStrCol: numbersStrCol,
		operators:     operators,
	}
}
