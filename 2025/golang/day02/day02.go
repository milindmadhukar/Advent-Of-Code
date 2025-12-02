package day02

import (
	"strconv"
	"strings"

	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/utils"
)

type day02 struct {
	data       string
	idSequence [][]uint64
}

func (d *day02) Part1() any {
	var invalidIdSum uint64 = 0

	for _, sequence := range d.idSequence {

		for currentId := range utils.GenerateRange(sequence[0], sequence[1]) {
			currentIdStr := strconv.FormatUint(currentId, 10)
			if len(currentIdStr)%2 != 0 {
				continue
			}
			currentIdStrLen := len(currentIdStr)
			if currentIdStr[:currentIdStrLen/2] == currentIdStr[currentIdStrLen/2:] {
				invalidIdSum += currentId
			}
		}
	}

	return invalidIdSum
}

func (d *day02) Part2() any {
	var invalidIdSum uint64 = 0

	for _, sequence := range d.idSequence {
		for currentId := sequence[0]; currentId <= sequence[1]; currentId++ {
			currentIdStr := strconv.FormatUint(currentId, 10)
			isInvalid := false
			for subStrLen := 1; subStrLen <= len(currentIdStr)/2; subStrLen++ {
				subStr := currentIdStr[:subStrLen]
				tempStr := currentIdStr
				for strings.HasPrefix(tempStr, subStr) {
					tempStr = strings.TrimPrefix(tempStr, subStr)
				}
				if tempStr == "" {
					isInvalid = true
					break
				}
			}

			if isInvalid {
				invalidIdSum += currentId
			}
		}
	}

	return invalidIdSum
}

func Solve() *day02 {
	data, err := utils.GetRawInputDataFromAOC(2025, 2)
	if err != nil {
		panic(err)
	}

	// data = utils.GetRawInputDataFromFile("day02/example.txt")

	splitData := strings.Split(data, ",")

	idSequenceInts := make([][]uint64, len(splitData))
	for i, r := range splitData {
		nums := strings.Split(r, "-")
		idSequenceInts[i] = utils.StringSliceToUint64Slice(nums)
	}

	return &day02{
		data:       data,
		idSequence: idSequenceInts,
	}
}
