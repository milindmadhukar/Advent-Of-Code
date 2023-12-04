package day3

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day3 struct {
	data []string
  startTime time.Time
}

var symbols string = `+-*/@&$#=%`

func leftCheck(numIdx []int, line string) bool {
	if numIdx[0] != 0 {
		if strings.Contains(symbols, string(line[numIdx[0]-1])) {
			return true
		}
	}
	return false
}

func rightCheck(numIdx []int, line string) bool {
	if numIdx[1] != len(line) {
		if strings.Contains(symbols, string(line[numIdx[1]])) {
			return true
		}
	}
	return false
}

func verticalCheck(numIdx []int, line string) bool {
	if strings.Contains(symbols, string(line[numIdx[0]])) || strings.Contains(symbols, string(line[numIdx[1]-1])) {
		return true
	}

	idxRange := make([]int, 2)

	if numIdx[0] == 0 {
		idxRange[0] = 0
	} else {
		idxRange[0] = numIdx[0] - 1
	}

	if numIdx[1] == len(line) {
		idxRange[1] = len(line)
	} else {
		idxRange[1] = numIdx[1] + 1
	}

	for _, char := range line[idxRange[0]:idxRange[1]] {
		if strings.Contains(symbols, string(char)) {
			return true
		}
	}

	return false
}

func (d day3) Part1() any {
	sum := 0
	r := regexp.MustCompile(`\d+`)

	for lineIdx, line := range d.data {
		numberIndexes := r.FindAllStringIndex(line, -1)
		for _, numIdx := range numberIndexes {
			num, _ := strconv.Atoi(line[numIdx[0]:numIdx[1]])
			var left, right, top, bottom bool
			left = leftCheck(numIdx, line)
			right = rightCheck(numIdx, line)

			if lineIdx != 0 {
				top = verticalCheck(numIdx, d.data[lineIdx-1])
			}

			if lineIdx != len(d.data)-1 {
				bottom = verticalCheck(numIdx, d.data[lineIdx+1])
			}

			if left || right || top || bottom {
				sum += num
			}
		}
	}

	return sum
}

type GearPart struct {
	startIdx int
	endIdx   int
	value    int
}

func (d day3) Part2() any {
	sum := 0

	for lineIdx, line := range d.data {

		gears := regexp.MustCompile(`\*`).FindAllStringIndex(line, -1)
		for _, gear := range gears {
			gearIdx := gear[0]

			var partsOnTop, partsOnBottom, partsOnSameLine [][]int

			partsOnSameLine = regexp.MustCompile(`\d+`).FindAllStringIndex(line, -1)
			if lineIdx != 0 {
				partsOnTop = regexp.MustCompile(`\d+`).FindAllStringIndex(d.data[lineIdx-1], -1)
			}
			if lineIdx != len(d.data)-1 {
				partsOnBottom = regexp.MustCompile(`\d+`).FindAllStringIndex(d.data[lineIdx+1], -1)
			}

			var gearParts []GearPart

			// Checking left and right
			for _, part := range partsOnSameLine {
				gearPart := GearPart{
					startIdx: part[0],
					endIdx:   part[1],
				}
				value, _ := strconv.Atoi(line[part[0]:part[1]])
				gearPart.value = value

				if part[1] == gearIdx {
					gearParts = append(gearParts, gearPart)
				}

				if part[0]-1 == gearIdx {
					gearParts = append(gearParts, gearPart)
				}
			}

			// Checking top and bottom

			for _, part := range partsOnTop {
				gearPart := GearPart{
					startIdx: part[0],
					endIdx:   part[1],
				}
				value, _ := strconv.Atoi(d.data[lineIdx-1][part[0]:part[1]])
				gearPart.value = value

				if part[1] >= gearIdx && part[0]-1 <= gearIdx {
					gearParts = append(gearParts, gearPart)
				}
			}

			for _, part := range partsOnBottom {
				gearPart := GearPart{
					startIdx: part[0],
					endIdx:   part[1],
				}
				value, _ := strconv.Atoi(d.data[lineIdx+1][part[0]:part[1]])
				gearPart.value = value

				if part[1] >= gearIdx && part[0]-1 <= gearIdx {
					gearParts = append(gearParts, gearPart)
				}
			}

			if len(gearParts) > 2 {
				panic("More than 2 adjacents")
			}

      if len(gearParts) != 2 {
        continue
      }

			sum += gearParts[0].value * gearParts[1].value
		}
	}

	return sum
}

func Solve() day3 {
	data, err := utils.GetInputDataFromAOC(2023, 3)
	if err != nil {
		panic(err)
	}

  startTime := time.Now()

	// exampleFile, _ := os.ReadFile("day3/example.txt")
	// data = utils.ParseFromString(string(exampleFile))

	return day3{
		data: data,
    startTime: startTime,
	}
}

func (d day3) TimeTaken() time.Duration {
  return time.Since(d.startTime)
}
