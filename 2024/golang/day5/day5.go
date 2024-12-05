package day5

import (
	"slices"
	"strings"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day5 struct {
	data              string
	pageOrderingRules [][]int
	pagesToProduce    [][]int
	validLines        [][]int
	invalidLines      [][]int
}

func (d *day5) ValidateLines() {
	pagesBefore := make(map[int][]int)

	for _, line := range d.pageOrderingRules {
		pagesBefore[line[0]] = append(pagesBefore[line[0]], line[1])
	}

	for _, line := range d.pagesToProduce {
		currentLine := slices.Clone(line)
		slices.Reverse(currentLine)
		idx := 0
		valid := true
		for idx < len(currentLine) && valid {
			shouldNotContain := pagesBefore[currentLine[idx]]
			left := currentLine[idx+1:]
			for _, val := range left {
				if utils.Contains(shouldNotContain, val) {
					valid = false
					break
				}
			}
			idx++
		}

		if valid {
			d.validLines = append(d.validLines, line)
		} else {
      // Fix the line
			for {
				idx := 0
				nowValid := true
				for idx < len(currentLine) {
					shouldNotContain := pagesBefore[currentLine[idx]]
					left := currentLine[idx+1:]
					for _, val := range left {
						if utils.Contains(shouldNotContain, val) {
							nowValid = false
							invalidPageIdx := utils.IndexOf(currentLine, val)
							invalidPage, currentLine := utils.Pop(currentLine, invalidPageIdx)
							currentLine = utils.Insert(currentLine, idx, invalidPage)
						}
					}
					idx++
				}

				if nowValid {
					break
				}
			}
      
			d.invalidLines = append(d.invalidLines, currentLine)
		}
	}
}

func (d *day5) Part1() any {
	sum := 0
	for _, validLine := range d.validLines {
		midIdx := len(validLine) / 2
		sum += validLine[midIdx]
	}

	return sum
}

func (d *day5) Part2() any {
	sum := 0
	for _, validLine := range d.invalidLines {
		midIdx := len(validLine) / 2
		sum += validLine[midIdx]
	}

	return sum
}

func Solve() *day5 {
	data, err := utils.GetRawInputDataFromAOC(2024, 5)
	if err != nil {
		panic(err)
	}

	/*
		fileData, _ := os.ReadFile("day5/example.txt")
		data = string(fileData)
		data = strings.Trim(data, " ")
		data = strings.Trim(data, "\n")
	*/

	splitData := strings.Split(data, "\n\n")

	top := utils.GetSplitData(strings.Split(splitData[0], "\n"), "|")
	bottom := utils.GetSplitData(strings.Split(splitData[1], "\n"), ",")

	var pageOrderingRules [][]int
	for _, line := range top {
		pageOrderingRules = append(pageOrderingRules, utils.StringSliceToIntegerSlice(line))
	}

	var pagesToProduce [][]int
	for _, line := range bottom {
		pagesToProduce = append(pagesToProduce, utils.StringSliceToIntegerSlice(line))
	}

	d := day5{
		data:              data,
		pageOrderingRules: pageOrderingRules,
		pagesToProduce:    pagesToProduce,
	}

	d.ValidateLines()

	return &d
}
