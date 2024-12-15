package day07

import (
	"strconv"
	"strings"
	"sync"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day07 struct {
	expressions []Expression
}

type Expression struct {
	result   int
	operands []int
}

func (d *day07) Compute(num1, num2 int, op string) int {
	var result int
	switch op {
	case "+":
		result = num1 + num2
	case "*":
		result = num1 * num2
	case "||":
		if num2 == 0 {
			return num1 * 10
		}
		digits := 1
		for temp := num2; temp > 0; temp /= 10 {
			digits *= 10
		}
		return num1*digits + num2
	}

	return result
}

func (d *day07) ValidExpressionsSum(operators []string) int {
	sum := 0

	var wg sync.WaitGroup
	wg.Add(len(d.expressions))

	for _, exp := range d.expressions {
		go func() {
			defer wg.Done()
			numsCount := len(exp.operands)
			operatorsLayout := utils.GeneratePermutations(operators, numsCount-1)
			for layout := range operatorsLayout {
				result := d.Compute(exp.operands[0], exp.operands[1], string(layout[0]))
				for i := 2; i < numsCount; i++ {
					result = d.Compute(result, exp.operands[i], string(layout[i-1]))
				}
				if result == exp.result {
					sum += exp.result
					return
				}
			}
		}()
	}

	wg.Wait()

	return sum
}

func (d *day07) Part1() any {
	operators := []string{"+", "*"}
	validSum := d.ValidExpressionsSum(operators)
	return validSum
}

func (d *day07) Part2() any {
	operators := []string{"+", "*", "||"}
	validSum := d.ValidExpressionsSum(operators)
	return validSum
}

func Solve() *day07 {
	data, err := utils.GetInputDataFromAOC(2024, 7)
	if err != nil {
		panic(err)
	}

	// data = utils.GetInputDataFromFile("day07/example.txt")

	var expressions []Expression

	for _, line := range data {
		var exp Expression
		split := strings.Split(line, ": ")
		exp.result, _ = strconv.Atoi(split[0])
		exp.operands = utils.StringSliceToIntegerSlice(strings.Split(split[1], " "))
		expressions = append(expressions, exp)
	}

	return &day07{
		expressions: expressions,
	}
}
