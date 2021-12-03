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

	log.Println("Answer of part 1 is:", horizontal*depth)

}

func day2part2(data [][]string) {

	// horizontal = 0
	// depth = 0
	// aim = 0

	// for instruction in data:
	//     direction = instruction[0]
	//     value = instruction[1]
	//     if direction == "forward":
	//         horizontal += value
	//         depth += value * aim

	//     elif direction == "down":
	//         aim += value
	//     elif direction == "up":
	//         aim -= value

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

	log.Println("Answer of part 2 is:", horizontal*depth)
}
