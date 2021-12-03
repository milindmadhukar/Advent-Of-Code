package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func getInputData(year, day int) []string {
	sessionKey := os.Getenv("SESSION")
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	cookie := http.Cookie{
		Name:  "session",
		Value: sessionKey,
	}
	if err != nil {
		log.Fatal(err)
	}

	req.AddCookie(&cookie)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	contents := string(body)

	var data []string
	var tmp string

	if strings.Contains(contents, "\n") {
		for _, char := range contents {
			if string(char) == "\n" {
				data = append(data, tmp)
				tmp = ""
			} else {
				tmp += string(char)
			}
		}
	} else {
		data = append(data, contents)
	}

	return data

}

func getDataFromFile(fileName string) []string {

	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	contents := strings.Trim(string(file), " \n")

	data := strings.Split(contents, "\n")
	return data
}

func StringSliceToIntegerSlice(s []string) []int {
	var ints []int

	for _, stringVal := range s {
		intVal, err := strconv.Atoi(stringVal)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, intVal)
	}

	return ints
}

func getSplitList(data []string, splitChar string) [][]string {
	var split [][]string

	for _, value := range data {
		split = append(split, strings.Split(value, splitChar))
	}

	return split
}
