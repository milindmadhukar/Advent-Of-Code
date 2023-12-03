package solutions

import (
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/day1"
	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/day2"
	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/day3"
	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/day4"
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
