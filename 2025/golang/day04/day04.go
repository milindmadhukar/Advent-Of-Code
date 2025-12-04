package day04

import (
	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/utils"
)

type day04 struct {
	data [][]string
}

var directions = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func (d *day04) countAdjacentPaperRolls(idxX, idxY int) int {
	count := 0
	line := d.data[idxY]
	
	for _, dir := range directions {
		newX := idxX + dir[0]
		newY := idxY + dir[1]
		
		if newX >= 0 && newX < len(line) && newY >= 0 && newY < len(d.data) {
			if d.data[newY][newX] == "@" {
				count++
			}
		}
	}
	
	return count
}

func (d *day04) Part1() any {
	validCount := 0
	
	for idxY, line := range d.data {
		for idxX, char := range line {
			if char == "@" && d.countAdjacentPaperRolls(idxX, idxY) < 4 {
				validCount++
			}
		}
	}
	
	return validCount
}

func (d *day04) Part2() any {
	validCount := 0
	
	for {
		pointsToRemove := make([][2]int, 0)
		
		for idxY, line := range d.data {
			for idxX, char := range line {
				if char == "@" && d.countAdjacentPaperRolls(idxX, idxY) < 4 {
					validCount++
					pointsToRemove = append(pointsToRemove, [2]int{idxX, idxY})
				}
			}
		}
		
		if len(pointsToRemove) == 0 {
			break
		}
		
		for _, point := range pointsToRemove {
			d.data[point[1]][point[0]] = "."
		}
	}
	
	return validCount
}

func Solve() *day04 {
	data, err := utils.GetInputDataFromAOC(2025, 4)
	if err != nil {
		panic(err)
	}
	// data = utils.GetInputDataFromFile("day04/example.txt")
	splitData := utils.GetSplitData(data, "")
	
	return &day04{
		data: splitData,
	}
}
