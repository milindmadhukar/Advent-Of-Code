package solutions

import (
	"fmt"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/day1"
	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/day2"
	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/day3"
	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/day4"
	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/day5"
	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/day6"
	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/day7"
	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/models"
)

func GetSolution(day int) models.Solution {
  var solution models.Solution
	switch day {
	case 1:
		solution = day1.Solve()
	case 2:
		solution = day2.Solve()
	case 3:
		solution = day3.Solve()
	case 4:
		solution = day4.Solve()
	case 5:
		solution = day5.Solve()
	case 6:
		solution = day6.Solve()
	case 7:
		solution = day7.Solve()
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
		fmt.Println("Answer for Part 1:", solution.Part1())
		fmt.Println("Answer for Part 2:", solution.Part2())
		fmt.Println("Time taken:", solution.TimeTaken())
		fmt.Println()
	}
}
