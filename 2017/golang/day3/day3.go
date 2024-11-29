package day3

import (
	"math"
	"strconv"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2017/golang/utils"
)

type day3 struct {
	data                      []string
	num                       int
	startTime                 time.Time
	allocatedValues           map[Point]int
	coordinates               map[int]Point
	firstValueBiggerThanInput int
}

type Point struct {
	x int
	y int
}

func (p Point) manhattanDistance() int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}

func (d *day3) GenerateMemoryBlock() {
	dx := []int{1, 0, -1, 0, 1, -1, -1, 1}
	dy := []int{0, 1, 0, -1, 1, 1, -1, -1}

	d.allocatedValues[Point{0, 0}] = 1

	i := 2
	opCount := 1
	dirIdx := 0
	currentX := 0
	currentY := 0
	run := true

	for run {
		for k := 0; k < 2 && run; k++ {
			for j := 0; j < opCount; j++ {
				currentX += dx[dirIdx]
				currentY += dy[dirIdx]
				p := Point{
					x: currentX,
					y: currentY,
				}

				d.coordinates[i] = p

				sum := 0
				for idx := range 8 {
					newPoint := Point{
						x: p.x + dx[idx],
						y: p.y + dy[idx],
					}

					v, ok := d.allocatedValues[newPoint]
					if ok {
						sum += v
					}
				}

				d.allocatedValues[p] = sum

				if sum > d.num && d.firstValueBiggerThanInput == -1 {
					d.firstValueBiggerThanInput = sum
				}

				if i == d.num {
					run = false
					break
				}
				i++
			}
			dirIdx = (dirIdx + 1) % 4
		}
		opCount++
	}
}

func (d day3) Part1() any {
	return d.coordinates[d.num].manhattanDistance()
}

func (d day3) Part2() any {
	return d.firstValueBiggerThanInput
}

func Solve() day3 {

	data, err := utils.GetInputDataFromAOC(2017, 3)
	if err != nil {
		panic(err)
	}
	num, _ := strconv.Atoi(data[0])

	startTime := time.Now()

	d := day3{
		data:                      data,
		num:                       num,
		startTime:                 startTime,
		allocatedValues:           make(map[Point]int),
		coordinates:               make(map[int]Point),
		firstValueBiggerThanInput: -1,
	}
	d.GenerateMemoryBlock()

	return d
}

func (d day3) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
