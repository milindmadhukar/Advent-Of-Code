package main

import (
	"log"
	"strings"
)

// from utils import getInputData, getSpiltList

// if __name__ == "__main__":
//     data = getInputData(year=2020, day=2)

//     data = getSpiltList(data=data, splitChar=":")

//     rangeVals = []
//     charVals = []
//     password_vals = []

//     for line in data:
//         rangeVals.append(list(map(int, line[0][0:-2].split("-"))))
//         charVals.append(line[0][-1])
//         password_vals.append(line[1][1:])

//     # Part 1

//     valid = 0
//     for i in range(1000):
//         count = 0
//         for char in password_vals[i]:
//             if char == charVals[i]:
//                 count += 1

//         if count >= rangeVals[i][0] and count <= rangeVals[i][1]:
//             valid += 1

//     print(valid)

//     # Part 2
//     valid = 0
//     for i in range(1000):
//         lower_char = password_vals[i][(rangeVals[i][0] - 1)]
//         higher_char = password_vals[i][(rangeVals[i][1] - 1)]
//         if charVals[i] == lower_char or charVals[i] == higher_char:
//             if lower_char != higher_char:
//                 valid += 1

//     print(valid)

func day2() {

	data := getSplitList(getInputData(2020, 2), ":")

	day2part1(data)
	day2part2(data)

}

func day2part1(data [][]string) {

	var valid int

	for _, entry := range data {

		charPolicy := strings.Split(entry[0], " ")
		charLimits := StringSliceToIntegerSlice(strings.Split(charPolicy[0], "-"))
		char := strings.Trim(charPolicy[1], " ")
		password := strings.Trim(entry[1], " ")

		if count := strings.Count(password, char); count >= charLimits[0] && count <= charLimits[1] {
			valid += 1
		}
	}

	log.Println("Answer for Day 2, Part 1 is", valid)

}

func day2part2(data [][]string) {

	var valid int

	for _, entry := range data {

		charPolicy := strings.Split(entry[0], " ")
		charIndices := StringSliceToIntegerSlice(strings.Split(charPolicy[0], "-"))
		char := strings.Trim(charPolicy[1], " ")
		password := strings.Trim(entry[1], " ")

		isCharPresent := string(password[charIndices[0]-1]) == char || string(password[charIndices[1]-1]) == char
		isDifferent := password[charIndices[0]-1] != password[charIndices[1]-1]

		if isCharPresent && isDifferent {
			valid += 1
		}
	}

	log.Println("Answer for Day 2, Part 2 is", valid)

}
