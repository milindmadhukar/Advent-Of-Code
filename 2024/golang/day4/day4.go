package day4

import (
	"strings"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day4 struct {
	data       []string
	parsedData [][]string
}

func (d *day4) Part1() any {
	count := 0

	// Horizontals
	for _, row := range d.parsedData {
		s := strings.Join(row, "")
		count += strings.Count(s, "XMAS") + strings.Count(s, "SAMX")
	}

	transposed := utils.Transpose(d.parsedData)
	// Verticals
	for _, row := range transposed {
		s := strings.Join(row, "")
		count += strings.Count(s, "XMAS") + strings.Count(s, "SAMX")
	}

	// Diagonals
	diagonals := utils.Diagonals(d.parsedData)
	for _, row := range diagonals {
		s := strings.Join(row, "")
		count += strings.Count(s, "XMAS") + strings.Count(s, "SAMX")
	}

	return count
}

func isValidSubgrid(grid [][]string) bool {
	return grid[1][1] == "A" &&
		((grid[0][0] == "M" && grid[2][2] == "S") || (grid[0][0] == "S" && grid[2][2] == "M")) &&
		((grid[0][2] == "M" && grid[2][0] == "S") || (grid[0][2] == "S" && grid[2][0] == "M"))
}

func (d *day4) Part2() any {
	count := 0

	for i := 0; i < len(d.parsedData)-2; i++ {
		for j := 0; j < len(d.parsedData[i])-2; j++ {
			var threeByThree [][]string
			for k := 0; k < 3; k++ {
				threeByThree = append(threeByThree, d.parsedData[i+k][j:j+3])
			}

			if isValidSubgrid(threeByThree) {
				count++
			}
		}
	}

	return count
}

func Solve() *day4 {
	data, err := utils.GetInputDataFromAOC(2024, 4)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day4/example.txt")

	parsedData := utils.GetSplitData(data, "")

	return &day4{
		data:       data,
		parsedData: parsedData,
	}
}
