package day7

import (
	"strconv"
	"strings"
	"sync"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day7 struct {
	data        []string
	expressions []Expression
}

type Expression struct {
	result   int
	operands []int
}

func (d *day7) Compute(num1, num2 int, op string) int {
	var result int
	switch op {
	case "+":
		result = num1 + num2
	case "*":
		result = num1 * num2
	case "||":
		result, _ = strconv.Atoi(strconv.Itoa(num1) + strconv.Itoa(num2))
	}

	return result
}

func (d *day7) ValidExpressionsSum(operators []string) int {
	sum := 0

	var wg sync.WaitGroup
	wg.Add(len(d.expressions))

	for _, exp := range d.expressions {
		go func() {
			defer wg.Done()
			numsCount := len(exp.operands)
			operatorsLayout := utils.Permutations(operators, numsCount-1)
			for layout := range operatorsLayout {
				result := d.Compute(exp.operands[0], exp.operands[1], string(layout[0]))
				for i := 2; i < numsCount; i++ {
					result = d.Compute(result, exp.operands[i], string(layout[i-1]))
					if result > exp.result {
						break
					}
				}
				if result == exp.result {
					sum += exp.result
					break
				}
			}
		}()
	}

	wg.Wait()

	return sum
}

func (d *day7) Part1() any {
	operators := []string{"+", "*"}
	validSum := d.ValidExpressionsSum(operators)
	return validSum
}

func (d *day7) Part2() any {
	operators := []string{"+", "*", "||"}
	validSum := d.ValidExpressionsSum(operators)
	return validSum
}

func Solve() *day7 {
	data, err := utils.GetInputDataFromAOC(2024, 7)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day7/example.txt")

	var expressions []Expression

	for _, line := range data {
		var exp Expression
		split := strings.Split(line, ": ")
		exp.result, _ = strconv.Atoi(split[0])
		exp.operands = utils.StringSliceToIntegerSlice(strings.Split(split[1], " "))
		expressions = append(expressions, exp)
	}

	return &day7{
		data:        data,
		expressions: expressions,
	}
}
