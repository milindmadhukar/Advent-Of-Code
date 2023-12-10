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
	start     *vertex
	vertices  map[point]*vertex
}

func (d day10) Part1() any {

	// Using d.start node, calculate the distances to each vertex and store it in its distance field
	dijkstra(d.start, d.vertices)

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
	return 0
}

type vertex struct {
	pos       point
	adjacents []*vertex
	pipe      rune
	distance  int
}

type point struct {
	x, y int
}

func bfs(start *vertex, vertices map[point]*vertex, data []string) {
	directions := []point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var visited = make(map[point]bool)
	visited[start.pos] = true
	queue := []*vertex{start}

	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:]

		for _, direction := range directions {
			nextPoint := point{currentVertex.pos.x + direction.x, currentVertex.pos.y + direction.y}
			if nextPoint.x < 0 || nextPoint.y < 0 || nextPoint.x >= len(data[0]) || nextPoint.y >= len(data) {
				continue
			}

			currentPipe := rune(data[nextPoint.y][nextPoint.x])

			if data[nextPoint.y][nextPoint.x] == '.' {
				continue
			}

			// when going up
			if direction.x == 0 && direction.y == -1 && !slices.Contains([]rune{'|', '7', 'F'}, currentPipe) {
				continue
			}

			// when going down
			if direction.x == 0 && direction.y == 1 && !slices.Contains([]rune{'|', 'J', 'L'}, currentPipe) {
				continue
			}

			// when going left
			if direction.x == -1 && direction.y == 0 && !slices.Contains([]rune{'-', 'L', 'F'}, currentPipe) {
				continue
			}

			// when going right
			if direction.x == 1 && direction.y == 0 && !slices.Contains([]rune{'-', '7', 'J'}, currentPipe) {
				continue
			}

			newVertex := &vertex{
				pos:      nextPoint,
				pipe:     currentPipe,
				distance: math.MaxInt,
			}

			if !visited[nextPoint] {
				visited[nextPoint] = true
				vertices[nextPoint] = newVertex
				currentVertex.adjacents = append(currentVertex.adjacents, newVertex)
				queue = append(queue, newVertex)
			}
		}
	}
}

func Solve() day10 {
	data, err := utils.GetInputDataFromAOC(2023, 10)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	exampleFile, _ := os.ReadFile("day10/example2.txt")
	data = utils.ParseFromString(string(exampleFile))

	var start *vertex

	// Finding S
	for i, row := range data {
		for j, char := range row {
			if char == 'S' {
				start = &vertex{
					pos:  point{j, i},
					pipe: rune(char),
				}
			}
		}
	}

	var vertices = make(map[point]*vertex)
	vertices[start.pos] = start

	bfs(start, vertices, data)

	pointHash := func(p point) string {
		return fmt.Sprintf("%d,%d", p.x, p.y)
	}

	g := graph.New(pointHash, graph.Directed())

	for _, v := range vertices {
		g.AddVertex(v.pos, graph.VertexAttribute("label", string(v.pipe)))
	}

	for _, v := range vertices {
		for _, adj := range v.adjacents {
			g.AddEdge(pointHash(v.pos), pointHash(adj.pos))
		}
	}

	file, _ := os.Create("day10/graph.gv")
	_ = draw.DOT(g, file)

	fmt.Println(len(vertices))

	return day10{
		data:      data,
		start:     start,
		vertices:  vertices,
		startTime: startTime,
	}
}

func (d day10) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
