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

func StringSliceToInt64Slice(s []string) []int64 {
	return Map(MustClosure(func(str string) (int64, error) {
		return strconv.ParseInt(str, 10, 64)
	}), s)
}

func StringSliceToInt32Slice(s []string) []int32 {
	return Map(MustClosure(func(str string) (int32, error) {
		val, err := strconv.ParseInt(str, 10, 32)
		return int32(val), err
	}), s)
}

func StringSliceToUint64Slice(s []string) []uint64 {
	return Map(MustClosure(func(str string) (uint64, error) {
		return strconv.ParseUint(str, 10, 64)
	}), s)
}
