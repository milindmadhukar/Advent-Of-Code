package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/solutions"
)

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  -day <number>: Run solution for a specific day")
	fmt.Println("  -all: Run all solutions")
	fmt.Println("  -help: Show this help message")
	fmt.Println("If no flags are provided, today's solution will be run.")
}

func main() {
	day := flag.String("day", "", "Specify the day to run")
	all := flag.Bool("all", false, "Run all solutions")
	help := flag.Bool("help", false, "Show help message")
	flag.Parse()

	if *help {
		printHelp()
		return
	}

	if *all {
		log.Println("Running all solutions")
		solutions.RunAllSolutions()
	} else if *day != "" {
		dayInt, err := strconv.Atoi(*day)
		if err != nil {
			log.Fatalf("Invalid day: %s", *day)
		}
		solutions.RunSolutionForDay(dayInt)
	} else {
		solutions.RunTodaysSolution()
	}
}
