package day19

import (
	"strings"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day19 struct {
	towels          []string
	desiredPatterns []string
	counts          map[string]int
	cache           map[string]int
}

func (d *day19) Towels(pattern string) int {
	if val, ok := d.cache[pattern]; ok {
		return val
	}

	if len(pattern) == 0 {
		return 1
	}

	possibilites := 0

	for _, towel := range d.towels {
		if strings.HasPrefix(pattern, towel) {
			possibilites += d.Towels(pattern[len(towel):])
		}
	}
	d.cache[pattern] = possibilites

	return possibilites
}

func (d *day19) Part1() any {
	return len(d.counts)
}

func (d *day19) Part2() any {
	sum := 0
	for _, val := range d.counts {
		sum += val
	}
	return sum
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

	d := day19{
		towels: towels,
		counts: make(map[string]int),
		cache:  make(map[string]int),
	}

	for _, pattern := range desiredPatterns {
		if val := d.Towels(pattern); val > 0 {
			d.counts[pattern] = val
		}
	}

	return &d
}
