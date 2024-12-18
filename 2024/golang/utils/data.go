package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GetRawInputDataFromAOC(year, day int) (string, error) {
	cacheFileName := filepath.Join("cache", fmt.Sprintf("aoc_%d_day%d.txt", year, day))

	if cacheData, err := os.ReadFile(cacheFileName); err == nil {
		return string(cacheData), nil
	}

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

	// Save to cache file
	if err := os.MkdirAll(filepath.Dir(cacheFileName), 0755); err != nil {
		return "", fmt.Errorf("Failed to create cache directory: %v", err)
	}
	if err := os.WriteFile(cacheFileName, []byte(bodyString), 0644); err != nil {
		return "", fmt.Errorf("Failed to write cache file: %v", err)
	}

	return bodyString, nil
}

func GetInputDataFromAOC(year int, day int) ([]string, error) {
	bodyString, err := GetRawInputDataFromAOC(year, day)
	if err != nil {
		return nil, err
	}
	return ParseFromString(bodyString), nil
}

func GetInputDataFromFile(fileName string) []string {
	contents := GetRawInputDataFromFile(fileName)

	data := strings.Split(contents, "\n")
	return data
}

func GetRawInputDataFromFile(fileName string) string {
	fileData, _ := os.ReadFile("day09/example.txt")
	data := string(fileData)
	data = strings.Trim(data, " ")
	data = strings.Trim(data, "\n")

	return data
}
