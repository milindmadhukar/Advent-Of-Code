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
	states     map[string]int
	steps      int
	loopSize   int
}

func getState(memoryBanks []int) string {
	memoryBanksStr := utils.Map(strconv.Itoa, memoryBanks)
	return strings.Join(memoryBanksStr, ",")
}

func (d *day6) GenerateStates() {
	memoryBanks := make([]int, len(d.parsedData))
	copy(memoryBanks, d.parsedData)
	d.states[getState(memoryBanks)] = d.steps

	for {
		d.steps++

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
		if stepNo, ok := d.states[state]; ok {

			d.loopSize = d.steps - stepNo

			break
		} else {
			d.states[state] = d.steps
		}
	}
}

func (d day6) Part1() any {
	return d.steps
}

func (d day6) Part2() any {
	return d.loopSize
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

	d := day6{
		data:       data,
		startTime:  startTime,
		parsedData: parsedData,
		states:     make(map[string]int),
		loopSize:   0,
		steps:      0,
	}

	d.GenerateStates()

	return d
}

func (d day6) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
