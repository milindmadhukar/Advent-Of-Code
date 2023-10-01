package main

import (
	"log"
	"strings"
)

type Coordinates []int
type PairOfCoordinates []Coordinates

func day5() {

	data := getDataFromFile("day5.txt")

	// data := getInputData(2021, 5)
	parsedRanges := getSplitList(data, "->")
	var parsedData []PairOfCoordinates
	for _, prange := range parsedRanges {
		var coordinates PairOfCoordinates
		for _, coordinatePair := range prange {
			pair := StringSliceToIntegerSlice(strings.Split(coordinatePair, ","))
			coordinates = append(coordinates, pair)
		}
		parsedData = append(parsedData, coordinates)
	}

	day5part1(parsedData)
	// day5part2(parsedData)

}

func day5part1(data []PairOfCoordinates) {
	log.Println(data)
	log.Println("Answer for Day 5, Part 1 is")
}

func day5part2(data []PairOfCoordinates) {
	log.Println("Answer for Day 5, Part 2 is")
}

func isHorizontalOrVertical(data []PairOfCoordinates)
