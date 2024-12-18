package day11

import (
	"math"
	"strconv"
	"strings"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day11 struct {
	stones []int
}

type StonePos struct {
	stone int
	depth int
}

func Blink(count int, stones []int, limit int, cache map[StonePos]uint64) uint64 {
	if count == limit {
		return uint64(len(stones))
	}

	var sum uint64 = 0
	var res uint64

	for _, stone := range stones {
		stonePos := StonePos{stone, count}

		if v, ok := cache[stonePos]; ok {
			sum += v
			continue
		}
		if stone == 0 {
			res = Blink(count+1, []int{1}, limit, cache)
			cache[stonePos] = res
			sum += res
		} else if stoneStr := strconv.Itoa(stone); len(stoneStr)%2 == 0 {
			pow := math.Pow10(len(stoneStr) / 2)
			left := int(float64(stone) / pow)
			right := stone % int(pow)
			res = Blink(count+1, []int{left, right}, limit, cache)
			cache[stonePos] = res
			sum += res
		} else {
			res = Blink(count+1, []int{stone * 2024}, limit, cache)
			cache[stonePos] = res
			sum += res
		}
	}

	return sum
}

func (d day11) Part1() any {
	return Blink(0, d.stones, 25, make(map[StonePos]uint64))
}

func (d day11) Part2() any {
	return Blink(0, d.stones, 75, make(map[StonePos]uint64))
}

func Solve() *day11 {
	data, err := utils.GetRawInputDataFromAOC(2024, 11)
	if err != nil {
		panic(err)
	}

  // data = utils.GetRawInputDataFromFile("day11/example.txt")

	stones := utils.StringSliceToIntegerSlice(strings.Split(data, " "))

	return &day11{
		stones: stones,
	}
}
