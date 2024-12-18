package day17

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day17 struct {
	regA, regB, regC   int
	programs           []int
	instructionPointer int
	outputs            []string
}

func (d *day17) ComboOperandValue(operand int) int {
	switch operand {
	case 0:
		return operand
	case 1:
		return operand
	case 2:
		return operand
	case 3:
		return operand
	case 4:
		return d.regA
	case 5:
		return d.regB
	case 6:
		return d.regC
	default:
		panic("Invalid operand")
	}
}

func (d *day17) ExecuteInstruction(opcode int, operand int) {
	switch opcode {
	// adv
	case 0:
		d.regA = int(float64(d.regA) / math.Pow(2, float64(d.ComboOperandValue(operand))))
	// bxl
	case 1:
		d.regB ^= operand
	// bst
	case 2:
		d.regB = d.ComboOperandValue(operand) % 8
	// jnz
	case 3:
		if d.regA != 0 {
			d.instructionPointer = operand
		}
	// bxc
	case 4:
		d.regB ^= d.regC
	// out
	case 5:
		d.outputs = append(d.outputs, strconv.Itoa(d.ComboOperandValue(operand)%8))
	// bdv
	case 6:
		d.regB = int(float64(d.regA) / math.Pow(2, float64(d.ComboOperandValue(operand))))
	// cdv
	case 7:
		d.regC = int(float64(d.regA) / math.Pow(2, float64(d.ComboOperandValue(operand))))
	}
}

func (d *day17) Part1() any {
	d.instructionPointer = 0
	for d.instructionPointer < len(d.programs) {
		instruction := d.programs[d.instructionPointer]
		operand := d.programs[d.instructionPointer+1]
		d.ExecuteInstruction(instruction, operand)
		if instruction == 3 {
			if d.regA == 0 {
				d.instructionPointer += 2
			}
			continue
		}
		d.instructionPointer += 2
	}

	return strings.Join(d.outputs, ",")
}

func (d *day17) Find(programs []int, answer int) int {
	if len(programs) == 0 {
		return answer
	}

	for d.regB = range utils.GenerateRange(8) {
		d.regA = (answer << 3) | d.regB
		d.regB = 0
		d.regC = 0
		d.outputs = nil

		for d.instructionPointer = range utils.GenerateRange(0, len(d.programs)-2, 2) {
			instruction := d.programs[d.instructionPointer]
			operand := d.programs[d.instructionPointer+1]
			if instruction == 0 && operand == 3 {
				continue
			}

			d.ExecuteInstruction(instruction, operand)

			if instruction == 5 && d.regB%8 == programs[len(programs)-1] {
				subAnswer := d.Find(programs[:len(programs)-1], d.regA)
				if subAnswer != -1 {
					return subAnswer
				}
			}
		}
	}

	return -1
}

func (d *day17) Part2() any {
	// This solution assumes there is only one jump in the end 3 0
	// That there is exists a left shift by 3 on A register and happens only once
	// There is a single output in the program
	return d.Find(d.programs, 0)
}

func Solve() *day17 {
	data, err := utils.GetRawInputDataFromAOC(2024, 17)
	if err != nil {
		panic(err)
	}

	/*
		fileData, _ := os.ReadFile("day17/example.txt")
		data = string(fileData)
		data = strings.Trim(data, " ")
		data = strings.Trim(data, "\n")
	*/

	splitData := strings.Split(data, "\n")

	var regA, regB, regC int
	fmt.Sscanf(splitData[0], "Register A: %d", &regA)
	fmt.Sscanf(splitData[1], "Register B: %d", &regB)
	fmt.Sscanf(splitData[2], "Register C: %d", &regC)
	programsStr := strings.Split(splitData[4], " ")
	programs := utils.StringSliceToIntegerSlice(strings.Split(programsStr[1], ","))

	return &day17{
		regA,
		regB,
		regC,
		programs,
		0,
		make([]string, 0),
	}
}
