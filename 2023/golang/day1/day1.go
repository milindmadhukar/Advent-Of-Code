package day1

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day1 struct {
	data []string
}

func reverseString(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func (d day1) Part1() any {
	sum := 0
	r := regexp.MustCompile(`(\d)`)
	for _, line := range d.data {
		digits := r.FindAllString(line, -1)
		first, _ := strconv.Atoi(digits[0])
		last, _ := strconv.Atoi(digits[len(digits)-1])
		sum += (first * 10) + last
	}
	return sum
}

func (d day1) Part2() any {

	var names_to_number = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	sum := 0

  // NOTE: Finding first number by using the normal regex but reversing the regex and input to find first from last.

	digitRegexPattern := `one|two|three|four|five|six|seven|eight|nine`
	digitRegexPatternReverse := reverseString(digitRegexPattern)

  digitRegexPattern = fmt.Sprintf(`(\d|%s)`, digitRegexPattern)
  digitRegexPatternReverse = fmt.Sprintf(`(\d|%s)`, digitRegexPatternReverse)

	digitRegex := regexp.MustCompile(digitRegexPattern)
	digitRegexReverse := regexp.MustCompile(digitRegexPatternReverse)

	for _, line := range d.data {
		var first, last int
		var err error
		firstStr := digitRegex.FindString(line)
		lastStr := digitRegexReverse.FindString(reverseString(line))
		first, err = strconv.Atoi(firstStr)
		if err != nil {
			first = names_to_number[firstStr]
		}
		last, err = strconv.Atoi(lastStr)
    if err != nil {
      last = names_to_number[reverseString(lastStr)]
    }
    sum += (first*10) + last
	}

	return sum
}

func Solve() day1 {
	data, err := utils.GetInputDataFromAOC(2023, 1)
	// exampleFile, _ := os.ReadFile("day1/input.txt")
	// data = utils.ParseFromString(string(exampleFile))

	if err != nil {
		panic(err)
	}

	return day1{
		data: data,
	}
}
