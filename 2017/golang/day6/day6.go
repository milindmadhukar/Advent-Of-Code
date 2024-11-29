package day6

import (
	"strconv"
	"strings"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/utils"
)

type day6 struct {
	data       string
	startTime  time.Time
	parsedData []int
}

func getState(memoryBanks []int) string {
	memoryBanksStr := utils.Map(strconv.Itoa, memoryBanks)
	return strings.Join(memoryBanksStr, ",")
}

func (d day6) Part1() any {

	states := make(map[string]bool)
	steps := 0

	memoryBanks := make([]int, len(d.parsedData))
	copy(memoryBanks, d.parsedData)
	states[getState(memoryBanks)] = true

	for {
		steps++

		maxBlockIdx := 0
		maxBlockCount := memoryBanks[0]

		for idx, blockCount := range memoryBanks {
			if blockCount > maxBlockCount {
				maxBlockCount = blockCount
				maxBlockIdx = idx
			}
		}

		count := maxBlockCount
		memoryBanks[maxBlockIdx] = 0
		maxBlockIdx = (maxBlockIdx + 1) % len(memoryBanks)

		for {
			if count == 0 {
				break
			}
			memoryBanks[maxBlockIdx]++
			count--
			maxBlockIdx = (maxBlockIdx + 1) % len(memoryBanks)
		}

		state := getState(memoryBanks)
		if _, ok := states[state]; ok {
			break
		} else {
			states[state] = true
		}
	}

	return steps
}

func (d day6) Part2() any {
	return 0
}

func Solve() day6 {
	data, err := utils.GetRawInputDataFromAOC(2017, 6)
	if err != nil {
		panic(err)
	}

	// exampleFileData, _ := os.ReadFile("day6/example.txt")
	// data = strings.Trim(string(exampleFileData), "\t\n ")

	startTime := time.Now()

	parsedData := utils.StringSliceToIntegerSlice(strings.Split(data, "\t"))

	return day6{
		data:       data,
		startTime:  startTime,
		parsedData: parsedData,
	}
}

func (d day6) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
