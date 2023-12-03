package day3

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day3 struct {
	data []string
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

func (d day3) Part2() any {
	sum := 0

	// To check in all eight directions

	for lineIdx, line := range d.data {

		gears := regexp.MustCompile(`\*`).FindAllStringIndex(line, -1)
	gearLoop:
		for _, gear := range gears {
			gearIdx := gear[0]

      fmt.Println("Line: ", line, "Gear: ", gearIdx)

			var partsOnTop, partsOnBottom, partsOnSameLine [][]int

			partsOnSameLine = regexp.MustCompile(`\d+`).FindAllStringIndex(line, -1)
			if lineIdx != 0 {
				partsOnTop = regexp.MustCompile(`\d+`).FindAllStringIndex(d.data[lineIdx-1], -1)
			}
			if lineIdx != len(d.data)-1 {
				partsOnBottom = regexp.MustCompile(`\d+`).FindAllStringIndex(d.data[lineIdx+1], -1)
			}

      fmt.Println("Checking for horizontally on same",)

			for idx := 0; idx < len(partsOnSameLine)-1; idx++ {
				part1idx := partsOnSameLine[idx]
				part2idx := partsOnSameLine[idx+1]
				if part1idx[1] == gearIdx && part2idx[0]-1 == gearIdx {
					num1, _ := strconv.Atoi(line[part1idx[0]:part1idx[1]])
					num2, _ := strconv.Atoi(line[part2idx[0]:part2idx[1]])
					sum += num1 * num2
					fmt.Println("Gear", num1, num2)
					continue gearLoop
				}
			}

      fmt.Println("Checking for horizontally on top")

			for idx := 0; idx < len(partsOnTop)-1; idx++ {
				part1idx := partsOnTop[idx]
				part2idx := partsOnTop[idx+1]
				if part1idx[1] == gearIdx && part2idx[0]-1 == gearIdx {
					num1, _ := strconv.Atoi(d.data[lineIdx-1][part1idx[0]:part1idx[1]])
					num2, _ := strconv.Atoi(d.data[lineIdx-1][part2idx[0]:part2idx[1]])
					sum += num1 * num2
					fmt.Println("Gear", num1, num2)
					continue gearLoop
				}
			}

      fmt.Println("Checking for horizontally on bottom")

			for idx := 0; idx < len(partsOnBottom)-1; idx++ {
				part1idx := partsOnBottom[idx]
				part2idx := partsOnBottom[idx+1]
				if part1idx[1] == gearIdx && part2idx[0]-1 == gearIdx {
					num1, _ := strconv.Atoi(d.data[lineIdx+1][part1idx[0]:part1idx[1]])
					num2, _ := strconv.Atoi(d.data[lineIdx+1][part2idx[0]:part2idx[1]])
					sum += num1 * num2
					fmt.Println("Gear", num1, num2)
					continue gearLoop
				}
			}

			// Checking vertically
			if lineIdx == 0 || lineIdx == len(d.data)-1 {
				continue gearLoop
			}

			fmt.Println("Checking for vertically")

			var partOnTop []int
			var partOnBottom []int

			// find part on top
			for _, part := range partsOnTop {
        if part[1] >= gearIdx && part[0] - 1 <= gearIdx {
          partOnTop = part
        }
			}

			// find part on bottom
			for _, part := range partsOnBottom {
        if part[1] >= gearIdx && part[0] - 1 <= gearIdx {
          partOnBottom = part
        }
			}

			if len(partOnTop) == 0 || len(partOnBottom) == 0 {
				continue gearLoop
			}

			num1, _ := strconv.Atoi(d.data[lineIdx-1][partOnTop[0]:partOnTop[1]])
			num2, _ := strconv.Atoi(d.data[lineIdx+1][partOnBottom[0]:partOnBottom[1]])
			fmt.Println("Gear", num1, num2)
			sum += num1 * num2
		}
	}

	return sum
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

func Solve() day3 {
	data, err := utils.GetInputDataFromAOC(2023, 3)
	if err != nil {
		panic(err)
	}

	return day3{
		data: data,
	}
}
