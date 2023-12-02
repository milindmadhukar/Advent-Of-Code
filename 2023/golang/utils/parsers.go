package utils

import "strings"

func ParseFromString(data string) []string {
  data = strings.Trim(data, " ")
  data = strings.Trim(data, "\n")
  return strings.Split(data, "\n")
}

func GetSplitData(data []string, splitter string) [][]string {
  var splitData [][]string

  for _, line := range data {
    splitData = append(splitData, strings.Split(line, splitter))
  }

  return splitData
}
