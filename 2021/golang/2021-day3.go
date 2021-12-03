package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func getData(fileName string) []string {

	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	contents := strings.Trim(string(file), " \n")

	data := strings.Split(contents, "\n")
	return data
}

func partOne(data []string) {

	places := getPlaces(data)

	var firstBitNumber string

	for _, place := range places {
		firstBitNumber += getCommonBit(place, "0")
	}

	secondBitNumber := flipBits(firstBitNumber)

	num1, err := strconv.ParseInt(firstBitNumber, 2, 64)
	num2, err := strconv.ParseInt(secondBitNumber, 2, 64)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Answer for part 1 is", num1*num2)
}

func partTwo(data []string) {

	oxygenRating := data
	carbonDioxideRating := data

	for pos := 0; pos < 12; pos++ {

		places := getPlaces(oxygenRating)
		cBit := getCommonBit(places[pos], "1")

		var indexesToKeep []int

		for index, number := range oxygenRating {
			if string(number[pos]) == cBit {
				indexesToKeep = append(indexesToKeep, index)
			}
		}

		oxygenRating = getValuesWithGivenIndex(indexesToKeep, oxygenRating)

		if len(oxygenRating) == 1 {
			break
		}
	}

	fmt.Println("CO2", carbonDioxideRating)

	for pos := 0; pos < 12; pos++ {

		places := getPlaces(carbonDioxideRating)
		cBit := getCommonBit(places[pos], "1")
		fmt.Println("Common bit is ", cBit, "at position", pos)
		var indexesToKeep []int

		for index, number := range carbonDioxideRating {
			if string(number[pos]) != cBit {
				indexesToKeep = append(indexesToKeep, index)
				fmt.Println(number)
			}
		}

		carbonDioxideRating = getValuesWithGivenIndex(indexesToKeep, carbonDioxideRating)

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

	log.Println("num1", num1, "num2", num2, "o2rating", oxygenRating, "co2rating", carbonDioxideRating)

	log.Println("Answer to part 2 is", num1*num2)
}

func main() {

	data := getData("day3.txt")
	partOne(data)
	partTwo(data)

}
