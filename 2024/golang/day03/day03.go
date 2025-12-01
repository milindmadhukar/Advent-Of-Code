package day03

import (
	"fmt"
	"regexp"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day03 struct {
	data string
}

func (d *day03) Part1() any {
	sum := 0
	var num1, num2 int
	reg := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := reg.FindAllString(d.data, -1)
	for _, match := range matches {
		fmt.Sscanf(match, "mul(%d,%d)", &num1, &num2)
		sum += num1 * num2
	}

	return sum
}

func (d *day03) Part2() any {
	sum := 0
	var num1, num2 int
	reg := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don\'t\(\)`)
	matches := reg.FindAllString(d.data, -1)
	isMulEnabled := true
	for _, match := range matches {
		if match == "do()" {
			isMulEnabled = true
		} else if match == "don't()" {
			isMulEnabled = false
		} else if isMulEnabled {
			fmt.Sscanf(match, "mul(%d,%d)", &num1, &num2)
			sum += num1 * num2
		}
	}
	return sum
}

func Solve() *day03 {
	data, err := utils.GetRawInputDataFromAOC(2024, 3)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day03/example.txt")

	return &day03{
		data: data,
	}
}
