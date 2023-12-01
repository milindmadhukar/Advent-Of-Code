package day1

import (
	"slices"
	"strconv"
	"unicode"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day1 struct {
	data       []string
	parsedData [][]number
}

type number struct {
	num    int
	isWord bool
}

var names_to_number = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func getNumber(tmp string) (int, bool) {
	i := 0
	for {
		if tmp[i:] == "" {
			break
		}
		if v, ok := names_to_number[tmp[i:]]; ok {
			return v, true
		} else {
      i += 1
    }
	}
	return -1, false
}

func SplitIntoNumbers(data []string) [][]number {
	var numList [][]number

	for _, line := range data {
		tmp := ""
		var digits []number
		for _, char := range line {
			if unicode.IsLetter(char) {
				tmp += string(char)
				v, ok := getNumber(tmp)
				if ok {
					digits = append(digits, number{
						num:    v,
						isWord: true,
					})
					tmp = ""
				} else {
					continue
				}
			}

			if unicode.IsDigit(char) {
				num, _ := strconv.Atoi(string(char))
				digits = append(digits, number{
					num:    num,
					isWord: false,
				})
				tmp = ""
			}
		}
		numList = append(numList, digits)
	}

	return numList
}

func (d day1) Part1() any {
	sum := 0
	for _, line := range d.parsedData {
		num := 0
		for _, digit := range line {
			if !digit.isWord {
				num += digit.num
				break
			}
		}
		num *= 10

		var lineCopy []number
    copy(lineCopy, line)
		slices.Reverse(lineCopy)

		for _, digit := range lineCopy {
			if !digit.isWord {
				num += digit.num
				break
			}
		}
		sum += num
	}
	return sum
}

func (d day1) Part2() any {
	sum := 0

	for _, digits := range d.parsedData {
    if len(digits) < 1 {
      continue
    }
		num := digits[0].num
		num *= 10
		num += digits[len(digits)-1].num

		sum += num
	}

	return sum
}

func Solve() day1 {
	data, err := utils.GetInputDataFromAOC(2023, 1)
	// exampleFile, _ := os.ReadFile("day1/example.txt")
	// data = utils.ParseFromString(string(exampleFile))

	if err != nil {
		panic(err)
	}

	return day1{
		data:       data,
		parsedData: SplitIntoNumbers(data),
	}
}
