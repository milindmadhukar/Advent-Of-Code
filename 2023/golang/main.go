package main

import (
	"fmt"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/day1"
	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/day2"
	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/models"
)

func GetTodaysSolution() models.Solution {
	day := time.Now().Day()
	month := time.Now().Month()

	if month != 12 || day > 25 {
		panic("Not the right time to run this")
	}

	switch day {
	case 1:
		return day1.Solve()
  case 2:
    return day2.Solve()
	default:
		panic("Solution not implemented yet")
	}

}

func main() {
	var solution models.Solution

	solution = GetTodaysSolution()
   
  fmt.Println("Answer for Part 1:", solution.Part1())
  fmt.Println("Answer for Part 2:", solution.Part2())
}
