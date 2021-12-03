// from utils import getInputData

// data = getInputData(year=2021, typecast=int)
// print(
//     "Answer to part 1:",
//     len([data[i] for i in range(0, len(data) - 1) if data[i + 1] > data[i]]),
// )

// print(
//     "Answer to part 2:",
//     len(
//         [
//             i
//             for i in range(0, len(data) - 1)
//             if data[i] + data[i + 1] + data[i + 2]
//             < data[i + 1] + data[i + 2] + data[i + 3]
//         ]
//     ),
// )

package main

import (
	"log"
)

func day1() {
	data := StringSliceToIntegerSlice(getInputData(2021, 1))
	day1part1(data)
	day1part2(data)
}

func day1part1(data []int) {
	ctr := 0
	for i := 0; i < len(data)-1; i++ {
		if data[i] < data[i+1] {
			ctr += 1
		}
	}
	log.Println("Answer for part 1", ctr)
}

func day1part2(data []int) {

	ctr := 0
	for i := 0; i < len(data)-3; i++ {
		if (data[i] + data[i+1] + data[i+2]) < (data[i+1] + data[i+2] + data[i+3]) {
			ctr++
		}
	}
	log.Println("Answer for part 2", ctr)
}
