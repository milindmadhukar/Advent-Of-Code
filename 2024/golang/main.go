package main

import (
	"fmt"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/solutions"
)

func main() {

	solution := solutions.GetTodaysSolution()

	part1Time := time.Now()
	fmt.Println("Answer for Part 1:", solution.Part1())
	fmt.Println("Time taken for Part 1:", time.Since(part1Time))
	part2Time := time.Now()
	fmt.Println("Answer for Part 2:", solution.Part2())
	fmt.Println("Time taken for Part 2:", time.Since(part2Time))
}
