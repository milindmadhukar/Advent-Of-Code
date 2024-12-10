package solutions

import (
	"fmt"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day01"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day02"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day03"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day04"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day05"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day06"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day07"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day08"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day09"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day10"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/models"
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
	case 8:
		solution = day8.Solve()
	case 9:
		solution = day9.Solve()
	case 10:
		solution = day10.Solve()
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

func RunSolution(solution models.Solution, day int) {
	fmt.Println("Day", day)
	part1Time := time.Now()
	fmt.Println("Answer for Part 1:", solution.Part1())
	fmt.Println("Time taken for Part 1:", time.Since(part1Time))
	part2Time := time.Now()
	fmt.Println("Answer for Part 2:", solution.Part2())
	fmt.Println("Time taken for Part 2:", time.Since(part2Time))
}

func RunSolutionForDay(day int) {
	solution := GetSolution(day)
	if solution == nil {
		fmt.Println("No solution found for day", day)
		return
	}
	RunSolution(solution, day)
}

func RunTodaysSolution() {
	solution := GetTodaysSolution()
	RunSolution(solution, time.Now().Day())
}

func RunAllSolutions() {
	for day := 1; day <= 25; day++ {
		solution := GetSolution(day)
		if solution == nil {
			continue
		}
		RunSolution(solution, day)
	}
}
