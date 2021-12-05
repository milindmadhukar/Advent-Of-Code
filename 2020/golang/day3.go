package main

import "log"

func day3() {

	data := getInputData(2020, 3)
	day3part1(data)
	day3part2(data)

}

func getTreesEncountered(data []string, xStep, yStep int) int {

	var x, y, count int

	for i := 0; i < (len(data)-1)/yStep; i++ {
		x += xStep
		y += yStep

		if data[y][x%len(data[y])] == '#' {
			count += 1
		}
	}

	return count
}

func day3part1(data []string) {
	log.Println("Answer for Day 3, Part 1 is", getTreesEncountered(data, 3, 1))
}

func day3part2(data []string) {

	answer := getTreesEncountered(data, 1, 1) * getTreesEncountered(data, 3, 1) * getTreesEncountered(data, 5, 1) * getTreesEncountered(data, 7, 1) * getTreesEncountered(data, 1, 2)

	log.Println("Answer for Day 3, Part 2 is", answer)

}
