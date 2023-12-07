package day6

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day6 struct {
	data            []string
	raceTime        []int
	distanceToCover []int
	startTime       time.Time
}

func getTotalNumberOfWays(raceTime []int, distanceToCover []int) int {
	totalNumberOfWays := 1
	for i := 0; i < len(raceTime); i++ {
		numberOfWays := 0
		raceTime := raceTime[i]
		for j := 1; j < raceTime-1; j++ {
			speed := j
			distanceCovered := (raceTime - j) * speed
			if distanceCovered > distanceToCover[i] {
				numberOfWays++
			}
		}
		if numberOfWays == 0 {
			continue
		}
		totalNumberOfWays *= numberOfWays
	}
	return totalNumberOfWays
}

func (d day6) Part1() any {
	return getTotalNumberOfWays(d.raceTime, d.distanceToCover)
}

func (d day6) Part2() any {
	newRaceTime := ""
	newDistanceToCover := ""
	for i := 0; i < len(d.raceTime); i++ {
		newRaceTime += fmt.Sprintf("%d", d.raceTime[i])
		newDistanceToCover += fmt.Sprintf("%d", d.distanceToCover[i])
	}
	raceTime, _ := strconv.Atoi(newRaceTime)
	distanceToCover, _ := strconv.Atoi(newDistanceToCover)
	return getTotalNumberOfWays([]int{raceTime}, []int{distanceToCover})
}

func Solve() day6 {
	data, err := utils.GetInputDataFromAOC(2023, 6)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	// exampleFile, _ := os.ReadFile("day6/example.txt")
	// data = utils.ParseFromString(string(exampleFile))

	raceTime := utils.StringSliceToIntegerSlice(regexp.MustCompile(`\d+`).FindAllString(data[0], -1))
	distanceToCover := utils.StringSliceToIntegerSlice(regexp.MustCompile(`\d+`).FindAllString(data[1], -1))

	return day6{
		data:            data,
		raceTime:        raceTime,
		distanceToCover: distanceToCover,
		startTime:       startTime,
	}
}

func (d day6) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
