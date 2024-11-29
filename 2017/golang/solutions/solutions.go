package solutions

import (
	"fmt"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/day1"
	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/day2"
	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/models"
)

func GetSolution(day int) models.Solution {
	switch day {
	case 1:
		return day1.Solve()
	case 2:
		return day2.Solve()

	default:
		panic("Solution not implemented yet")
	}
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
		fmt.Println("Day", i)
		fmt.Println("Answer for Part 1:", solution.Part1())
		fmt.Println("Answer for Part 2:", solution.Part2())
		fmt.Println("Time taken:", solution.TimeTaken())
		fmt.Println()
	}
}
