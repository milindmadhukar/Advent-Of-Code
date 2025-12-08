package day08

import (
	"container/heap"
	"math"

	"slices"

	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/utils"
)

type Coordinate struct {
	X           int
	Y           int
	Z           int
	Connections []Coordinate
}

type CoordPair struct {
	A        Coordinate
	B        Coordinate
	Distance float64
}

type CoordPairHeap []CoordPair

func (h CoordPairHeap) Len() int           { return len(h) }
func (h CoordPairHeap) Less(i, j int) bool { return h[i].Distance < h[j].Distance }
func (h CoordPairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *CoordPairHeap) Push(x any) {
	*h = append(*h, x.(CoordPair))
}
func (h *CoordPairHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type day08 struct {
	data          []string
	isExampleData bool
	splitData     [][]string
	coords        []Coordinate
	minHeap       *CoordPairHeap
}

func isPointPresentInGroup(coord Coordinate, group []Coordinate) bool {
	for _, gCoord := range group {
		if coord.X == gCoord.X && coord.Y == gCoord.Y && coord.Z == gCoord.Z {
			return true
		}
	}
	return false
}

func manageGroup(coordPair CoordPair, groups [][]Coordinate) [][]Coordinate {
	coordPair.A.Connections = append(coordPair.A.Connections, coordPair.B)
	coordPair.B.Connections = append(coordPair.B.Connections, coordPair.A)

	aGroupIdx := -1
	bGroupIdx := -1

	for groupIdx, group := range groups {
		if isPointPresentInGroup(coordPair.A, group) {
			aGroupIdx = groupIdx
		}
		if isPointPresentInGroup(coordPair.B, group) {
			bGroupIdx = groupIdx
		}
	}

	if aGroupIdx != -1 && bGroupIdx != -1 {
		if aGroupIdx == bGroupIdx {
		} else {
			groups[aGroupIdx] = append(groups[aGroupIdx], groups[bGroupIdx]...)
			groups = slices.Delete(groups, bGroupIdx, bGroupIdx+1)
		}
	} else if aGroupIdx != -1 {
		groups[aGroupIdx] = append(groups[aGroupIdx], coordPair.B)
	} else if bGroupIdx != -1 {
		groups[bGroupIdx] = append(groups[bGroupIdx], coordPair.A)
	} else {
		groups = append(groups, []Coordinate{coordPair.A, coordPair.B})
	}

	return groups
}

func (d *day08) Part1() any {
	minHeap := &CoordPairHeap{}
	*minHeap = append(*minHeap, (*d.minHeap)...)
	heap.Init(minHeap)

	var timesToJoin int
	if d.isExampleData {
		timesToJoin = 10
	} else {
		timesToJoin = 1000
	}

	groups := make([][]Coordinate, 0)
	for range timesToJoin {
		coordPair := heap.Pop(minHeap).(CoordPair)
		groups = manageGroup(coordPair, groups)
	}

	slices.SortFunc(groups, func(a, b []Coordinate) int {
		return len(b) - len(a)
	})

	return len(groups[0]) * len(groups[1]) * len(groups[2])
}

func (d *day08) Part2() any {
	minHeap := &CoordPairHeap{}
	*minHeap = append(*minHeap, (*d.minHeap)...)
	heap.Init(minHeap)

	groups := make([][]Coordinate, 0)
	var currentCoordPair CoordPair
	for len(groups) < 1 || len(groups[0]) < len(d.coords) {
		currentCoordPair = heap.Pop(minHeap).(CoordPair)
		groups = manageGroup(currentCoordPair, groups)
	}

	return currentCoordPair.A.X * currentCoordPair.B.X
}

func Solve() *day08 {
	data, err := utils.GetInputDataFromAOC(2025, 8)
	if err != nil {
		panic(err)
	}

	exampleData := false
	if exampleData {
		data = utils.GetInputDataFromFile("day08/example.txt")
	}

	splitData := utils.GetSplitData(data, ",")
	var rowInt []int
	coords := make([]Coordinate, len(splitData))
	for i, row := range splitData {
		rowInt = utils.StringSliceToIntegerSlice(row)
		coords[i] = Coordinate{
			X: rowInt[0],
			Y: rowInt[1],
			Z: rowInt[2],
		}
	}

	minHeap := &CoordPairHeap{}
	heap.Init(minHeap)

	for coords := range utils.GenerateCombinations(coords, 2) {
		coordPair := CoordPair{coords[0], coords[1], 0}
		a := coordPair.A.X - coordPair.B.X
		b := coordPair.A.Y - coordPair.B.Y
		c := coordPair.A.Z - coordPair.B.Z
		coordPair.Distance = float64(math.Sqrt(float64(a*a + b*b + c*c)))

		heap.Push(minHeap, coordPair)
	}

	return &day08{
		data:          data,
		isExampleData: exampleData,
		splitData:     splitData,
		coords:        coords,
		minHeap:       minHeap,
	}
}
