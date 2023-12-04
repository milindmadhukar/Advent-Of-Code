package day4

import (
	"regexp"
	"strings"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day4 struct {
	data         []string
	startTime    time.Time
	scratchCards map[int]*scratchCard
}

func checkNumInSlice(num int, slice []int) bool {
	for _, n := range slice {
		if n == num {
			return true
		}
	}
	return false
}

func (d day4) Part1() any {
	sum := 0

	for _, scratchCard := range d.scratchCards {
		currentScore := 0
		for _, num := range scratchCard.current {
			if checkNumInSlice(num, scratchCard.winnings) {
				if currentScore == 0 {
					currentScore = 1
				} else {
					currentScore *= 2
				}
			}
		}
		sum += currentScore
	}

	return sum
}

func (d day4) Part2() any {
	sum := 0

	for id := 1; id <= len(d.scratchCards); id++ {
		scratchCard := d.scratchCards[id]
		matchingNumbersCount := 0
		for _, num := range scratchCard.current {
			if checkNumInSlice(num, scratchCard.winnings) {
				matchingNumbersCount++
			}
		}

	matchingLoop:
		for i := 1; i <= matchingNumbersCount; i++ {
			_, ok := d.scratchCards[id+i]
			if !ok {
				break matchingLoop
			}
			d.scratchCards[id+i].count += d.scratchCards[id].count
		}
	}

	for _, scratchCard := range d.scratchCards {
		sum += scratchCard.count
	}

	return sum
}

type scratchCard struct {
	winnings []int
	current  []int
	count    int
}

func Solve() day4 {
	data, err := utils.GetInputDataFromAOC(2023, 4)
	if err != nil {
		panic(err)
	}

	// exampleFile, _ := os.ReadFile("day4/example.txt")
	// data = utils.ParseFromString(string(exampleFile))

	var scratchCards = make(map[int]*scratchCard)

	for idx, line := range data {
		line = line[9:]
		winnings_and_current := strings.Split(line, " | ")
		winningNumsStr := regexp.MustCompile(`\d+`).FindAllString(winnings_and_current[0], -1)
		currentNumStr := regexp.MustCompile(`\d+`).FindAllString(winnings_and_current[1], -1)
		winning_nums := utils.StringSliceToIntegerSlice(winningNumsStr)
		current_num := utils.StringSliceToIntegerSlice(currentNumStr)

		scratchCards[idx+1] = &scratchCard{
			winnings: winning_nums,
			current:  current_num,
			count:    1,
		}
	}

	startTime := time.Now()

	return day4{
		data:         data,
		startTime:    startTime,
		scratchCards: scratchCards,
	}
}

func (d day4) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
