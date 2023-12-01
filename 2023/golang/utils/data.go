package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetInputDataFromAOC(year int, day int) ([]string, error) {

  // Read the sessionkey.txt file
  sessionKey, err := os.ReadFile("sessionkey.txt")
  if err != nil {
    return nil, err
  }

  session := string(sessionKey)[:len(string(sessionKey))-1]

  url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
  req, err := http.NewRequest("GET", url, nil)
  req.AddCookie(&http.Cookie{Name: "session", Value: session})

  if err != nil {
    return nil, err
  }

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
    return nil, err
	}

	bodyString := string(bodyBytes)

  if resp.StatusCode != 200 {
    return nil, fmt.Errorf("Could not fetch input. %s", bodyString)
  }

  return ParseFromString(bodyString), nil
}
