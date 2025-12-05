package solutions

import (
	"fmt"
	"log"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/day01"
	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/day02"
	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/day03"
	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/day04"
	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/day05"
	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/day06"
	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/models"
)

func GetSolution(day int) (models.Solution, time.Time) {
	var solution models.Solution
	startTime := time.Now()
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
	}
	return solution, startTime
}

func RunSolution(solution models.Solution, startTime time.Time, day int) {
	fmt.Println("Day", day)
	part1Time := time.Now()
	fmt.Println("Answer for Part 1:", solution.Part1())
	fmt.Println("Time taken for Part 1:", time.Since(part1Time))
	part2Time := time.Now()
	fmt.Println("Answer for Part 2:", solution.Part2())
	fmt.Println("Time taken for Part 2:", time.Since(part2Time))
	fmt.Println("Total time taken:", time.Since(startTime))
}

func RunSolutionForDay(day int) {
	solution, startTime := GetSolution(day)
	if solution == nil {
		log.Fatal("No solution found for Day ", day)
		return
	}
	RunSolution(solution, startTime, day)
}

func RunTodaysSolution() {
	day := time.Now().Day()
	month := time.Now().Month()

	if month != 12 || day > 25 {
		panic("Not the right time to run this\nTry running --help")
	}

	solution, startTime := GetSolution(day)

	if solution == nil {
		log.Fatal("No solution found for Day ", day)
		return
	}

	RunSolution(solution, startTime, time.Now().Day())
}

func RunAllSolutions() {
	for day := 1; day <= 25; day++ {
		solution, startTime := GetSolution(day)
		if solution == nil {
			continue
		}
		RunSolution(solution, startTime, day)
	}
}
