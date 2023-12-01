package main

import (
	"fmt"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/models"
	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/solutions"
)

func main() {
	var solution models.Solution

	solution = solutions.GetSolution(1)

	fmt.Println("Answer for Part 1:", solution.Part1())
	fmt.Println("Answer for Part 2:", solution.Part2())
}
