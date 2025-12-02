package day01

import (
	"fmt"
	"strconv"

	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/utils"
)

type day01 struct {
	data []string
}

func (d *day01) Part1() any {
	location := 50
	zeroCount := 0

	for _, line := range d.data {
		direction := string(line[0])
		amt, _ := strconv.Atoi(line[1:])

		if direction == "R" {
			location = (location + amt) % 100
		} else if direction == "L" {
			location = (location - amt + 100) % 100
		} else {
			fmt.Println("Invalid direction:", direction)
		}

		if location == 0 {
			zeroCount++
		}
	}

	return zeroCount
}

func (d *day01) Part2() any {
	location := 50
	touchedZeroCount := 0

	for _, line := range d.data {
		direction := string(line[0])
		amt, _ := strconv.Atoi(line[1:])

		for range amt {
			if direction == "R" {
				location++
				if location == 100 {
					location = 0
				}
			} else {
				location--
				if location == -1 {
					location = 99
				}
			}

			if location == 0 {
				touchedZeroCount++
			}
		}
	}

	return touchedZeroCount
}

func Solve() *day01 {
	data, err := utils.GetInputDataFromAOC(2025, 1)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day01/example.txt")

	return &day01{
		data: data,
	}
}
