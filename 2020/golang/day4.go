package main

import (
	"log"
	"strings"
)

func day4() {
	data := getRawInputData(2020, 4)
	parsedData := parseDay4Data(data)

	log.Println(parsedData)

	day4part1(parsedData)
	day4part2(parsedData)
}

func day4part1(data []map[string]string) {

}

func day4part2(data []map[string]string) {

}

func parseDay4Data(data string) []map[string]string {

	var parsedData [][]string
	perPerson := strings.Split(data, "\n\n")
	for _, person := range perPerson {
		newlineSplit := strings.Split(person, "\n")
		var personData []string
		for _, data := range newlineSplit {
			personData = append(personData, strings.Split(data, " ")...)
		}
		parsedData = append(parsedData, personData)
	}

	var personFields []map[string]string

	for _, person := range parsedData {
		fields := make(map[string]string)
		for _, keyValuePair := range person {
			keyValue := strings.Split(keyValuePair, ":")
			log.Println(keyValuePair, "Key V", keyValue)
			fields[strings.Trim(keyValue[0], " ")] = strings.Trim(keyValue[1], " ")
			personFields = append(personFields, fields)
		}
	}

	return personFields
}
