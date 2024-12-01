package day4

import (
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/utils"
)

type day4 struct {
	data       []string
	parsedData [][]string
	startTime  time.Time
}

func isAnagram(s string, t string) bool {
	string1 := len(s)
	string2 := len(t)
	if string1 != string2 {
		return false
	}

	anagramMap := make(map[string]int)

	for i := 0; i < string1; i++ {
		anagramMap[string(s[i])]++
	}

	for i := 0; i < string2; i++ {
		anagramMap[string(t[i])]--
	}

	for i := 0; i < string1; i++ {
		if anagramMap[string(s[i])] != 0 {
			return false
		}
	}
	return true
}

func (d *day4) Part1() any {
	validCount := 0
	for _, line := range d.parsedData {
		unique := utils.GetUniqueElements(line)
		if len(unique) == len(line) {
			validCount++
		}
	}
	return validCount
}

func (d *day4) Part2() any {
	validCount := 0
	for _, line := range d.parsedData {
		isValid := true

		for idx1 := 0; idx1 < len(line) && isValid; idx1++ {
			for idx2 := idx1 + 1; idx2 < len(line); idx2++ {
				if isAnagram(line[idx1], line[idx2]) {
					isValid = false
					break
				}
			}
		}

		if isValid {
			validCount++
		}
	}
	return validCount
}

func Solve() *day4 {
	data, err := utils.GetInputDataFromAOC(2017, 4)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	// data = utils.GetInputDataFromFile("day4/example.txt")

	return &day4{
		data:       data,
		startTime:  startTime,
		parsedData: utils.GetSplitData(data, " "),
	}
}

func (d day4) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
