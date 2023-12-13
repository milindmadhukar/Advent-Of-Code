package day10

import (
	"fmt"
	"math"
	"os"
	"slices"
	"time"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day10 struct {
	data      []string
	startTime time.Time
	start     *pipe
	vertices  map[point]*pipe
	tiles     map[point]*tile
}

func (d day10) Part1() any {

	// Using d.start node, calculate the distances to each vertex and store it in its distance field
	d.start.distance = 0

	var queue = []*pipe{d.start}
	var visited = make(map[point]bool)

	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:]

		for _, adj := range currentVertex.adjacents {
			if !visited[adj.pos] {
				adj.distance = currentVertex.distance + 1
				queue = append(queue, adj)
				visited[adj.pos] = true
			}
		}
	}

	// Find the vertex with the maximum distances
	maxDistance := 0

	for _, v := range d.vertices {
		if v.distance > maxDistance {
			maxDistance = v.distance
		}
	}

	return maxDistance
}

func (d day10) Part2() any {
	count := 0

	var visitedTiles = make(map[point]bool)

	for _, v := range d.tiles {
		var newVisitedTiles = make(map[point]bool)
		if !visitedTiles[v.pos] {
			bfs2(v, d.tiles, newVisitedTiles, d.data)
		}

		didReachEdge := false
		for k, v := range newVisitedTiles {
			if v && (k.x == 0 || k.y == 0 || k.x == len(d.data[0])-1 || k.y == len(d.data)-1) {
				didReachEdge = true
				break
			}
		}

		if !didReachEdge {
			count += len(newVisitedTiles)

			for k, v := range newVisitedTiles {
				fmt.Println(k, v)
			}
		}

		for k, v := range newVisitedTiles {
			visitedTiles[k] = v
		}
	}

	return count
}

type pipe struct {
	pos       point
	adjacents []*pipe
	pipeType  rune
	distance  int
}

type tile struct {
	pos point
}

type point struct {
	x, y int
}

var directions = []point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func bfs1(start *pipe, pipes map[point]*pipe, data []string) {
	var visitedPipes = make(map[point]bool)
	visitedPipes[start.pos] = true
	queue := []*pipe{start}

	for len(queue) > 0 {
		currentPipe := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		for _, direction := range directions {
			nextPoint := point{currentPipe.pos.x + direction.x, currentPipe.pos.y + direction.y}
			if nextPoint.x < 0 || nextPoint.y < 0 || nextPoint.x >= len(data[0]) || nextPoint.y >= len(data) {
				continue
			}

			currentPipeType := rune(data[nextPoint.y][nextPoint.x])

			if currentPipeType == '.' {
				continue
			}

			// when going up
			if direction.x == 0 && direction.y == -1 && !slices.Contains([]rune{'|', '7', 'F'}, currentPipeType) {
				// Make it a dot
				continue
			}

			// when going down
			if direction.x == 0 && direction.y == 1 && !slices.Contains([]rune{'|', 'J', 'L'}, currentPipeType) {
				continue
			}

			// when going left
			if direction.x == -1 && direction.y == 0 && !slices.Contains([]rune{'-', 'L', 'F'}, currentPipeType) {
				continue
			}

			// when going right
			if direction.x == 1 && direction.y == 0 && !slices.Contains([]rune{'-', '7', 'J'}, currentPipeType) {
				continue
			}

			newPipe := &pipe{
				pos:      nextPoint,
				pipeType: currentPipeType,
				distance: math.MaxInt,
			}

			if !visitedPipes[nextPoint] {
				visitedPipes[nextPoint] = true
				pipes[nextPoint] = newPipe
				currentPipe.adjacents = append(currentPipe.adjacents, newPipe)
				newPipe.adjacents = append(newPipe.adjacents, currentPipe)
				queue = append(queue, newPipe)
			}
		}
	}

	for _, line := range data {
		fmt.Println(line)
	}
}

func bfs2(start *tile, tiles map[point]*tile, visitedTiles map[point]bool, data []string) {
	visitedTiles[start.pos] = true
	queue := []*tile{start}

	for len(queue) > 0 {
		currentTile := queue[0]
		queue = queue[1:]

		for _, direction := range directions {
			nextPoint := point{currentTile.pos.x + direction.x, currentTile.pos.y + direction.y}
			if nextPoint.x < 0 || nextPoint.y < 0 || nextPoint.x >= len(data[0]) || nextPoint.y >= len(data) {
				continue
			}

			currentTileType := rune(data[nextPoint.y][nextPoint.x])

			if currentTileType != '.' {
				continue
			}

			newTile := &tile{
				pos: nextPoint,
			}

			if !visitedTiles[nextPoint] {
				visitedTiles[nextPoint] = true
				tiles[nextPoint] = newTile
				queue = append(queue, newTile)
			}
		}
	}
}

func findMainLoop(start *pipe) map[point]*pipe {
	var mainLoop = make(map[point]*pipe)
	// Walk through adjacents of start until you reach start again. Longest loop is the main loop

	var visited = make(map[point]bool)
	visited[start.pos] = true
	queue := []*pipe{start}

	for len(queue) > 0 {
		currentPipe := queue[0]
		queue = queue[1:]

		for _, adj := range currentPipe.adjacents {
			if !visited[adj.pos] {
				visited[adj.pos] = true
				queue = append(queue, adj)
			}
		}
	}

	return mainLoop
}

func Solve() day10 {
	data, err := utils.GetInputDataFromAOC(2023, 10)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	exampleFile, _ := os.ReadFile("day10/example.txt")
	data = utils.ParseFromString(string(exampleFile))

	var start *pipe

	// Finding S
	for i, row := range data {
		for j, char := range row {
			if char == 'S' {
				start = &pipe{
					pos:      point{j, i},
					pipeType: rune(char),
				}
			}
		}
	}

	var pipes = make(map[point]*pipe)
	pipes[start.pos] = start

	bfs1(start, pipes, data)

	// Find all the tiles
	var tiles = make(map[point]*tile)
	for i, row := range data {
		for j, char := range row {
			if char == '.' {
				tiles[point{j, i}] = &tile{
					pos: point{j, i},
				}
			}
		}
	}

	pointHash := func(p point) string {
		return fmt.Sprintf("%d,%d", p.x, p.y)
	}

	pipeGraph := graph.New(pointHash, graph.Directed())

	for _, v := range pipes {
		pipeGraph.AddVertex(v.pos,
			graph.VertexAttribute(
				"label",
				fmt.Sprintf("%s - (%d,%d)", string(v.pipeType), v.pos.x, v.pos.y),
			),
		)
	}

	for _, v := range pipes {
		for _, adj := range v.adjacents {
			pipeGraph.AddEdge(pointHash(v.pos), pointHash(adj.pos))
		}
	}

	file, _ := os.Create("day10/graph.gv")
	_ = draw.DOT(pipeGraph, file)

	return day10{
		data:      data,
		start:     start,
		vertices:  pipes,
		tiles:     tiles,
		startTime: startTime,
	}
}

func (d day10) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
