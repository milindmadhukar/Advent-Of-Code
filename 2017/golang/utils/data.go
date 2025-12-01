package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetRawInputDataFromAOC(year, day int) (string, error) {
	// Read the sessionkey.txt file
	sessionKey, err := os.ReadFile("sessionkey.txt")
	if err != nil {
		return "", err
	}

	session := string(sessionKey)[:len(string(sessionKey))-1]

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	bodyString := string(bodyBytes)

	bodyString = strings.Trim(bodyString, "\n")

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Could not fetch input. %s", bodyString)
	}

	return bodyString, err
}

func GetInputDataFromAOC(year int, day int) ([]string, error) {
	bodyString, err := GetRawInputDataFromAOC(year, day)
	if err != nil {
		return nil, err
	}
	return ParseFromString(bodyString), nil
}

func GetInputDataFromFile(fileName string) []string {

	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	contents := strings.Trim(string(file), "\n")

	data := strings.Split(contents, "\n")
	return data
}
