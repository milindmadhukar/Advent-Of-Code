package day12

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"slices"
	"sync"
	"time"
)

func randomColor() color.RGBA {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
}

func assignColors(clusters []Cluster) map[Point]color.RGBA {
	clusterColors := make(map[Point]color.RGBA)
	adjacencyMap := make(map[Point]map[Point]bool)

	for _, cluster := range clusters {
		clusterColor := randomColor()
		for plantPos := range cluster.plants {
			presentLeft := cluster.plants[Point{plantPos.x - 1, plantPos.y}]
			presentRight := cluster.plants[Point{plantPos.x + 1, plantPos.y}]
			presentTop := cluster.plants[Point{plantPos.x, plantPos.y - 1}]
			presentBottom := cluster.plants[Point{plantPos.x, plantPos.y + 1}]
			if presentLeft && presentRight && presentTop && presentBottom {
				clusterColors[plantPos] = color.RGBA{0, 0, 0, 255}
			} else {
				clusterColors[plantPos] = clusterColor
			}
		}
		for plantPos := range cluster.plants {
			if _, exists := adjacencyMap[plantPos]; !exists {
				adjacencyMap[plantPos] = make(map[Point]bool)
			}
			for _, delta := range deltas {
				neighbor := Point{plantPos.x + delta.x, plantPos.y + delta.y}
				if _, neighborExists := adjacencyMap[neighbor]; neighborExists {
					adjacencyMap[plantPos][neighbor] = true
					adjacencyMap[neighbor][plantPos] = true
				}
			}
		}
	}
	return clusterColors
}

func findClusters(x, y int, plantType string, garden [][]string, xMax, yMax int, pointsFound map[Point]bool) {
	if x < 0 || y < 0 || x >= xMax || y >= yMax || plantType == "." {
		return
	}
	if garden[y][x] == plantType {
		pointsFound[Point{x, y}] = true
		garden[y][x] = "."
		for _, delta := range deltas {
			findClusters(x+delta.x, y+delta.y, plantType, garden, xMax, yMax, pointsFound)
		}
	}
}

func processRow(y int, xMax int, upscaledImageBy int, clusterColors map[Point]color.RGBA, upscaled *image.RGBA, wg *sync.WaitGroup) {
	defer wg.Done()

	for x := 0; x < xMax; x++ {
		for i := 0; i < upscaledImageBy; i++ {
			for j := 0; j < upscaledImageBy; j++ {
				upscaled.Set(x*upscaledImageBy+i, y*upscaledImageBy+j, clusterColors[Point{x, y}])
			}
		}
	}
}

func visualize(gardenOrignal [][]string) {
  var garden [][]string

  for _, row := range gardenOrignal {
    garden = append(garden, slices.Clone(row))
  }

	yMax := len(garden)
	xMax := len(garden[0])
	var clusters []Cluster
	upscalePointsBy := 1

	for y, row := range garden {
		for x, cell := range row {
			if cell != "." {
				cluster := Cluster{plantType: cell}
				cluster.plants = make(map[Point]bool)
				findClusters(x, y, cell, garden, xMax, yMax, cluster.plants)
				expandedPlants := make(map[Point]bool)
				for plantPos := range cluster.plants {
					for i := 0; i < upscalePointsBy; i++ {
						for j := 0; j < upscalePointsBy; j++ {
							expandedPlants[Point{plantPos.x*upscalePointsBy + i, plantPos.y*upscalePointsBy + j}] = true
						}
					}
				}
				cluster.plants = expandedPlants
				cluster.area = len(cluster.plants)
				clusters = append(clusters, cluster)
			}
		}
	}

	clusterColors := assignColors(clusters)

	xMax = xMax * upscalePointsBy
	yMax = yMax * upscalePointsBy

	upscaleImageBy := 1000
	upscaled := image.NewRGBA(image.Rect(0, 0, xMax*upscaleImageBy, yMax*upscaleImageBy))

	var wg sync.WaitGroup
	for y := 0; y < yMax; y++ {
		wg.Add(1)
		go processRow(y, xMax, upscaleImageBy, clusterColors, upscaled, &wg)
	}

	wg.Wait()

	outputFile, err := os.Create("day12/garden.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	png.Encode(outputFile, upscaled)
}
