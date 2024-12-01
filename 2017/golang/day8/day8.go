package day8

import (
	"math"
	"strconv"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/utils"
)

type day8 struct {
	data         []string
	startTime    time.Time
	instructions []Instruction
}

type Instruction struct {
	Register  string
	Operation string
	Value     int
	Cond      Condition
}

type Condition struct {
	Lhs   string
	Relop string
	Rhs   int
}

func (d day8) ConditionResult(cond Condition, registerValues map[string]int) bool {
	var conditionResult bool
	conditionRegister := registerValues[cond.Lhs]
	conditionValue := cond.Rhs

	switch cond.Relop {
	case ">":
		conditionResult = (conditionRegister > conditionValue)
	case "<":
		conditionResult = (conditionRegister < conditionValue)
	case ">=":
		conditionResult = (conditionRegister >= conditionValue)
	case "<=":
		conditionResult = (conditionRegister <= conditionValue)
	case "==":
		conditionResult = (conditionRegister == conditionValue)
	case "!=":
		conditionResult = (conditionRegister != conditionValue)
	}
	return conditionResult
}

func (d *day8) Part1() any {

	registerValues := make(map[string]int)

	for _, instruction := range d.instructions {
		if d.ConditionResult(instruction.Cond, registerValues) {
			if instruction.Operation == "inc" {
				registerValues[instruction.Register] += instruction.Value
			} else if instruction.Operation == "dec" {
				registerValues[instruction.Register] -= instruction.Value
			}
		}
	}

	maxVal := math.MinInt
	for _, regVal := range registerValues {
		if regVal > maxVal {
			maxVal = regVal
		}
	}

	return maxVal
}

func (d *day8) Part2() any {

	registerValues := make(map[string]int)
	maxVal := math.MinInt

	for _, instruction := range d.instructions {
		if d.ConditionResult(instruction.Cond, registerValues) {
			if instruction.Operation == "inc" {
				registerValues[instruction.Register] += instruction.Value
			} else if instruction.Operation == "dec" {
				registerValues[instruction.Register] -= instruction.Value
			}

			if registerValues[instruction.Register] > maxVal {
				maxVal = registerValues[instruction.Register]
			}
		}
	}

	return maxVal
}

func Solve() *day8 {
	data, err := utils.GetInputDataFromAOC(2017, 8)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	// data = utils.GetInputDataFromFile("day8/example.txt")

	splitData := utils.GetSplitData(data, " ")

	var instructions []Instruction
	registerValues := make(map[string]int)

	for _, line := range splitData {
		ins := Instruction{}
		ins.Register = line[0]
		registerValues[ins.Register] = 0
		ins.Operation = line[1]
		ins.Value, _ = strconv.Atoi(line[2])
		ins.Cond.Lhs = line[4]
		ins.Cond.Relop = line[5]
		ins.Cond.Rhs, _ = strconv.Atoi(line[6])
		instructions = append(instructions, ins)
	}

	return &day8{
		data:         data,
		startTime:    startTime,
		instructions: instructions,
	}
}

func (d day8) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
