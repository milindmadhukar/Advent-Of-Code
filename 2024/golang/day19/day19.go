package day19

import (
	"fmt"
	"slices"
	"strings"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day19 struct {
	towels          []string
	desiredPatterns []string
}

func (d *day19) Towels(pattern string, count int, cache map[string]int) int {
	if val, ok := cache[pattern]; ok {
		return val
	}

	if len(pattern) == 0 {
		return count
	}

	for _, towel := range d.towels {
		if strings.HasPrefix(pattern, towel) {
			if val := d.Towels(pattern[len(towel):], count+1, cache); val != 0 {
				cache[pattern] = val
				return val
			}
		}
	}

	return 0
}

func (d *day19) Part1() any {
	possibleCount := 0

	counts := make(map[string]int)
	cache := make(map[string]int)

	for idx, pattern := range d.desiredPatterns {
		fmt.Println("Current", idx+1)
		if val := d.Towels(pattern, 0, cache); val != -1 {
			possibleCount++
			counts[pattern] = val
		}
	}

	fmt.Println(counts)

	return possibleCount
}

func (d *day19) Part2() any {
	return 0
}

func Solve() *day19 {
	data, err := utils.GetRawInputDataFromAOC(2024, 19)
	if err != nil {
		panic(err)
	}

	// data = utils.GetRawInputDataFromFile("day19/example.txt")

	splitData := strings.Split(data, "\n\n")
	towels := strings.Split(splitData[0], ", ")
	desiredPatterns := strings.Split(splitData[1], "\n")

	slices.SortFunc(towels, func(a, b string) int {
		return len(a) - len(b)
	})

	return &day19{
		towels:          towels,
		desiredPatterns: desiredPatterns,
	}
}
