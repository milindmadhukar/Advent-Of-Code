package main

import "log"

func day1() {

	data := StringSliceToIntegerSlice(getInputData(2020, 1))
	day1part1(data)
	day1part2(data)
}

func day1part1(data []int) {

	for _, i := range data {
		for _, j := range data {
			if i+j == 2020 {
				log.Println("Answer for Day 1, Part is", i*j)
				return
			}
		}
	}

}

func day1part2(data []int) {

	for _, i := range data {
		for _, j := range data {
			for _, k := range data {
				if i+j+k == 2020 {
					log.Println("Answer for Day 1, Part is", i*j*k)
					return
				}
			}
		}
	}

}
