package utils

import (
	"strconv"
	"strings"
)

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

func StringSliceToIntegerSlice(s []string) []int {
  return Map(MustClosure(strconv.Atoi), s)
}
