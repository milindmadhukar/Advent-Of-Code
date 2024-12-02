package solutions

import (
	"fmt"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day1"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day2"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/models"
)

func GetSolution(day int) models.Solution {
	var solution models.Solution
	switch day {
	case 1:
		solution = day1.Solve()
	case 2:
		solution = day2.Solve()
	}
	return solution
}

func GetTodaysSolution() models.Solution {
	day := time.Now().Day()
	month := time.Now().Month()

	if month != 12 || day > 25 {
		panic("Not the right time to run this")
	}

	return GetSolution(day)
}

func RunAllSolutions() {
	for i := 1; i <= 25; i++ {
		solution := GetSolution(i)
		if solution == nil {
			continue
		}
		fmt.Println("Day", i)
    part1Time := time.Now()
		fmt.Println("Answer for Part 1:", solution.Part1())
    fmt.Println("Time taken for Part 1:", time.Since(part1Time))
    part2Time := time.Now()
		fmt.Println("Answer for Part 2:", solution.Part2())
    fmt.Println("Time taken for Part 2:", time.Since(part2Time))
		fmt.Println()
	}
}
