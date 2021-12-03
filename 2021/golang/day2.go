package main

import (
	"log"
	"strconv"
)

func day2() {

	data := getSplitList(getInputData(2021, 2), " ")
	day2part1(data)
	day2part2(data)

}

func day2part1(data [][]string) {

	var horizontal, depth int

	for _, instruction := range data {
		direction := instruction[0]
		value, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}

	log.Println("Answer for Day 2, Part 1 is:", horizontal*depth)

}

func day2part2(data [][]string) {
	var horizontal, depth, aim int

	for _, instruction := range data {
		direction := instruction[0]
		value, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "forward":
			horizontal += value
			depth += value * aim
		case "down":
			aim += value

		case "up":
			aim -= value
		}
	}

	log.Println("Answer for Day 2, Part 2 is:", horizontal*depth)
}
