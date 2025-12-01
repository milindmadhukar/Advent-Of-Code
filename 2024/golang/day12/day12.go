package day12

import (
	"math"

	"github.com/milindmadhukar/Advent-Of-Code/2024/golang/utils"
)

type day12 struct {
	garden   [][]string
	clusters []Cluster
}

type Point struct {
	x, y int
}

type Cluster struct {
	plantType string
	plants    map[Point]bool
	area      int
}

var deltas = []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func (d *day12) Part1() any {
	sum := 0
	for _, cluster := range d.clusters {
		perimeter := 0
		for plantPos := range cluster.plants {
			for _, delta := range deltas {
				if _, ok := cluster.plants[Point{plantPos.x + delta.x, plantPos.y + delta.y}]; !ok {
					perimeter++
				}
			}
		}
		sum += cluster.area * perimeter
	}
	return sum
}

func (d *day12) Part2() any {
	sum := 0
	for _, cluster := range d.clusters {
		sum += cluster.area * calculateSides(cluster.plants)
	}
	return sum
}

// idea from https://github.com/SHA65536/AdventOfCodeGo/blob/42536cbc645bba04f1c0f1cd1eefbf83ef53b9f5/2024/day12/day12.go#L81
// I did not come up with this scanning approach.
// credit to https://github.com/Karitham/aoc/blob/cec0847f73a507b6b843d9b49c517c466a48bf65/2024/src/12.py#L51 for cleaning the code so that I can understand lol.
func calculateSides(plants map[Point]bool) int {
	minX, maxX := math.MaxInt32, math.MinInt32
	minY, maxY := math.MaxInt32, math.MinInt32
	for point := range plants {
		if point.x < minX {
			minX = point.x
		}
		if point.x > maxX {
			maxX = point.x
		}
		if point.y < minY {
			minY = point.y
		}
		if point.y > maxY {
			maxY = point.y
		}
	}

	rows := make([][]Point, maxX-minX+1)
	for i := range rows {
		rows[i] = make([]Point, maxY-minY+1)
		for j := range rows[i] {
			rows[i][j] = Point{x: i + minX, y: j + minY}
		}
	}

	cols := make([][]Point, maxY-minY+1)
	for i := range cols {
		cols[i] = make([]Point, maxX-minX+1)
		for j := range cols[i] {
			cols[i][j] = Point{x: j + minX, y: i + minY}
		}
	}

	scan := func(points []Point, dx, dy int) int {
		count := 0
		bulk := false
		for _, pos := range points {
			_, inCluster := plants[pos]
			if !inCluster {
				bulk = false
				continue
			}
			_, nextInCluster := plants[Point{x: pos.x + dx, y: pos.y + dy}]
			if !nextInCluster && !bulk {
				count++
			}
			bulk = !nextInCluster
		}
		return count
	}

	up := 0
	for _, row := range rows {
		up += scan(row, -1, 0)
	}

	down := 0
	for _, row := range rows {
		down += scan(row, 1, 0)
	}

	left := 0
	for _, col := range cols {
		left += scan(col, 0, -1)
	}

	right := 0
	for _, col := range cols {
		right += scan(col, 0, 1)
	}

	return up + down + left + right
}

func (d *day12) findClusters(x, y int, plantType string, xMax, yMax int, pointsFound map[Point]bool) {
	if x < 0 || y < 0 || x >= xMax || y >= yMax || plantType == "." {
		return
	}

	if d.garden[y][x] == plantType {
		pointsFound[Point{x, y}] = true
		d.garden[y][x] = "."

		for _, delta := range deltas {
			d.findClusters(x+delta.x, y+delta.y, plantType, xMax, yMax, pointsFound)
		}
	}
}

func Solve() *day12 {
	data, err := utils.GetInputDataFromAOC(2024, 12)
	if err != nil {
		panic(err)
	}

	data = utils.GetInputDataFromFile("day12/example.txt")

	garden := utils.GetSplitData(data, "")

	visualize(garden)

	yMax := len(garden)
	xMax := len(garden[0])

	d := day12{
		garden: garden,
	}

	var clusters []Cluster
	for y, row := range garden {
		for x, cell := range row {
			if cell != "." {
				cluster := Cluster{plantType: cell}
				cluster.plants = make(map[Point]bool)
				d.findClusters(x, y, cell, xMax, yMax, cluster.plants)
				cluster.area = len(cluster.plants)
				clusters = append(clusters, cluster)
			}
		}
	}

	d.clusters = clusters

	return &d
}
