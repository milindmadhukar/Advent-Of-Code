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
	var ints []int

	for _, stringVal := range s {
    stringVal = strings.Trim(stringVal, " ")
		intVal, err := strconv.Atoi(stringVal)
		if err != nil {
			panic(err)
		}
		ints = append(ints, intVal)
	}

	return ints
}
