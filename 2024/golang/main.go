package main

import (
	"fmt"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/solutions"
)

func main() {

	solution := solutions.GetSolution(1)

	fmt.Println("Answer for Part 1:", solution.Part1())
	fmt.Println("Answer for Part 2:", solution.Part2())
	fmt.Println("Time taken:", solution.TimeTaken())

	// solutions.RunAllSolutions()
}
