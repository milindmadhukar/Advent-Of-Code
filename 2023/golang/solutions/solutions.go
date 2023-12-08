package solutions

import (
	"fmt"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/day1"
	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/day2"
	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/day3"
	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/day4"
	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/day5"
	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/day6"
	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/day7"
	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/day8"
	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/day9"
	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/models"
)

func GetSolution(day int) models.Solution {
	switch day {
	case 1:
		return day1.Solve()
	case 2:
		return day2.Solve()
	case 3:
		return day3.Solve()
	case 4:
		return day4.Solve()
	case 5:
		return day5.Solve()
	case 6:
		return day6.Solve()
	case 7:
		return day7.Solve()
	case 8:
		return day8.Solve()
	case 9:
		return day9.Solve()

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
