package solutions

import (
	"fmt"
	"log"
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
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day11"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day12"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day13"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day14"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day15"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day16"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day17"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day18"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day19"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/day20"
	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/models"
)

func GetSolution(day int) models.Solution {
	var solution models.Solution
	switch day {
	case 1:
		solution = day01.Solve()
	case 2:
		solution = day02.Solve()
	case 3:
		solution = day03.Solve()
	case 4:
		solution = day04.Solve()
	case 5:
		solution = day05.Solve()
	case 6:
		solution = day06.Solve()
	case 7:
		solution = day07.Solve()
	case 8:
		solution = day08.Solve()
	case 9:
		solution = day09.Solve()
	case 10:
		solution = day10.Solve()
	case 11:
		solution = day11.Solve()
	case 12:
		solution = day12.Solve()
	case 13:
		solution = day13.Solve()
	case 14:
		solution = day14.Solve()
	case 15:
		solution = day15.Solve()
	case 16:
		solution = day16.Solve()
	case 17:
		solution = day17.Solve()
	case 18:
		solution = day18.Solve()
	case 19:
		solution = day19.Solve()
	case 20:
		solution = day20.Solve()
	}
	return solution
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
		log.Fatal("No solution found for Day ", day)
		return
	}
	RunSolution(solution, day)
}

func RunTodaysSolution() {
	day := time.Now().Day()
	month := time.Now().Month()

	if month != 12 || day > 25 {
		panic("Not the right time to run this\nTry running --help")
	}

	solution := GetSolution(day)

	if solution == nil {
		log.Fatal("No solution found for Day ", day)
		return
	}

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
