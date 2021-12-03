package main

import (
	"log"
	"strconv"
	"strings"
)

func day3() {

	data := getInputData(2021, 3)
	day3Part1(data)
	day3Part2(data)

}

func getPlaces(data []string) [][]string {
	var places [][]string

	for i := 0; i < 12; i++ {
		var ls []string
		places = append(places, ls)
	}

	for _, val := range data {

		for i := 0; i < 12; i++ {
			places[i] = append(places[i], string(val[i]))
		}

	}

	return places
}

func getCommonBit(place []string, ifCommon string) string {
	number := strings.Join(place, "")
	if strings.Count(number, "0") > strings.Count(number, "1") {
		return "0"
	} else if strings.Count(number, "0") < strings.Count(number, "1") {
		return "1"
	} else {
		return ifCommon
	}
}

func flipBits(binaryNum string) string {
	var temp string
	for _, val := range binaryNum {
		num := string(val)
		if num == "0" {
			temp += "1"
		} else if num == "1" {
			temp += "0"
		}
	}
	return temp
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func getValuesWithGivenIndex(indexes []int, data []string) []string {
	var temp []string
	for _, index := range indexes {
		temp = append(temp, data[index])
	}

	return temp
}

func day3Part1(data []string) {

	places := getPlaces(data)

	var firstBitNumber string

	for _, place := range places {
		firstBitNumber += getCommonBit(place, "0")
	}

	secondBitNumber := flipBits(firstBitNumber)

	num1, err := strconv.ParseInt(firstBitNumber, 2, 64)

	if err != nil {
		log.Fatal(err)
	}

	num2, err := strconv.ParseInt(secondBitNumber, 2, 64)

	log.Println("Answer for part 1 is", num1*num2)
}

func day3Part2(data []string) {

	oxygenRating := data
	carbonDioxideRating := data

	for pos := 0; pos < 12; pos++ {

		places := getPlaces(oxygenRating)
		cBit := getCommonBit(places[pos], "1")

		var indicesToKeep []int

		for index, number := range oxygenRating {
			if string(number[pos]) == cBit {
				indicesToKeep = append(indicesToKeep, index)
			}
		}

		oxygenRating = getValuesWithGivenIndex(indicesToKeep, oxygenRating)

		if len(oxygenRating) == 1 {
			break
		}
	}

	for pos := 0; pos < 12; pos++ {

		places := getPlaces(carbonDioxideRating)
		cBit := getCommonBit(places[pos], "1")
		var indicesToKeep []int

		for index, number := range carbonDioxideRating {
			if string(number[pos]) != cBit {
				indicesToKeep = append(indicesToKeep, index)
			}
		}

		carbonDioxideRating = getValuesWithGivenIndex(indicesToKeep, carbonDioxideRating)

		if len(carbonDioxideRating) == 1 {
			break
		}
	}

	num1, err := strconv.ParseInt(oxygenRating[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	num2, err := strconv.ParseInt(carbonDioxideRating[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Answer to part 2 is", num1*num2)
}
