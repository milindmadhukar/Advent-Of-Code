package day05

import (
	"slices"
	"strconv"
	"strings"

	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/utils"
)

type day05 struct {
	data             string
	ingredientRanges [][2]uint64
	ingredients      []uint64
}

func (d *day05) Part1() any {
	validIngredients := 0

	for _, ingredient := range d.ingredients {
		for _, ingredientRange := range d.ingredientRanges {
			if ingredient >= ingredientRange[0] && ingredient <= ingredientRange[1] {
				validIngredients++
				break
			}
		}
	}

	return validIngredients
}

func (d *day05) Part2() any {
	type SweepingLineItem struct {
		position uint64
		value    int
	}
	sweepingLineItems := make([]SweepingLineItem, len(d.ingredientRanges)*2)
	for idx, ingredientRange := range d.ingredientRanges {
		sweepingLineItems[2*idx] = SweepingLineItem{ingredientRange[0], 1}
		sweepingLineItems[2*idx+1] = SweepingLineItem{ingredientRange[1] + 1, -1}
	}
	slices.SortFunc(sweepingLineItems, func(a, b SweepingLineItem) int {
		if a.position != b.position {
			return int(a.position) - int(b.position)
		}
		return b.value - a.value
	})

	var ingredientCount uint64
	k := 0
	var rangeStart uint64

	for _, sweepingLineItem := range sweepingLineItems {
		if k == 0 && sweepingLineItem.value > 0 {
			rangeStart = sweepingLineItem.position
		}
		k += sweepingLineItem.value
		if k == 0 {
			ingredientCount += sweepingLineItem.position - rangeStart
		}
	}
	return ingredientCount
}

func Solve() *day05 {
	data, err := utils.GetRawInputDataFromAOC(2025, 5)
	if err != nil {
		panic(err)
	}

	// data = utils.GetRawInputDataFromFile("day05/example.txt")

	splitData := strings.Split(data, "\n\n")
	ingredientRanges := utils.Map(
		func(val string) [2]uint64 {
			bounds := strings.Split(val, "-")
			start, _ := strconv.ParseInt(bounds[0], 10, 64)
			end, _ := strconv.ParseInt(bounds[1], 10, 64)
			return [2]uint64{uint64(start), uint64(end)}
		},
		strings.Split(splitData[0], "\n"))

	ingredients := utils.StringSliceToUint64Slice(strings.Split(splitData[1], "\n"))

	return &day05{
		data:             data,
		ingredientRanges: ingredientRanges,
		ingredients:      ingredients,
	}
}
