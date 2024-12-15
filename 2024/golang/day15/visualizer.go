package day15

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"sync"
)

func processRow(y int, xMax int, upscaledImageBy int, clusterColors map[Point]color.RGBA, upscaled *image.Paletted, wg *sync.WaitGroup) {
	defer wg.Done()

	for x := 0; x < xMax; x++ {
		for i := 0; i < upscaledImageBy; i++ {
			for j := 0; j < upscaledImageBy; j++ {
				upscaled.Set(x*upscaledImageBy+i, y*upscaledImageBy+j, clusterColors[Point{x, y}])
			}
		}
	}
}

func visualizePart1(gridSize Point, walls map[Point]bool, boxes map[Point]bool, robot Point) *image.Paletted {
	upscaleImageBy := 60

	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff},
		color.RGBA{0x00, 0x00, 0xff, 0xff},
	}

	colorMap := make(map[Point]color.RGBA)

	for y := 0; y < gridSize.y; y++ {
		for x := 0; x < gridSize.x; x++ {
			if wallFound := walls[Point{x, y}]; wallFound {
				colorMap[Point{x, y}] = color.RGBA{0xff, 0, 0, 255}
			} else if boxFound := boxes[Point{x, y}]; boxFound {
				colorMap[Point{x, y}] = color.RGBA{0, 0xff, 0, 255}
			} else if robot.x == x && robot.y == y {
				colorMap[Point{x, y}] = color.RGBA{0, 0, 0xff, 255}
			} else {
				colorMap[Point{x, y}] = color.RGBA{0, 0, 0, 255}
			}
		}
	}

	upscaled := image.NewPaletted(
		image.Rect(0, 0, gridSize.x*upscaleImageBy, gridSize.y*upscaleImageBy),
		palette,
	)
	var wg sync.WaitGroup
	for y := 0; y < gridSize.y; y++ {
		wg.Add(1)
		go processRow(y, gridSize.x, upscaleImageBy, colorMap, upscaled, &wg)
	}
	wg.Wait()

	return upscaled
}

func visualizePart2(gridSize Point, walls map[Point]bool, boxes map[Point]string, robot Point) *image.Paletted {
	upscaleImageBy := 60

	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x90, 0xee, 0x90, 0xff},
		color.RGBA{0x06, 0x40, 0x2b, 0xff},
		color.RGBA{0x00, 0x00, 0xff, 0xff},
	}

	colorMap := make(map[Point]color.RGBA)

	for y := 0; y < gridSize.y; y++ {
		for x := 0; x < gridSize.x; x++ {
			if wallFound := walls[Point{x, y}]; wallFound {
				colorMap[Point{x, y}] = color.RGBA{0xff, 0, 0, 255}
			} else if boxVal := boxes[Point{x, y}]; boxVal != "" {
				if boxVal == "[" {
					colorMap[Point{x, y}] = color.RGBA{0x06, 0x40, 0x2b, 0xff}
				} else {
					colorMap[Point{x, y}] = color.RGBA{0x90, 0xee, 0x90, 0xff}
				}
			} else if robot.x == x && robot.y == y {
				colorMap[Point{x, y}] = color.RGBA{0, 0, 0xff, 255}
			} else {
				colorMap[Point{x, y}] = color.RGBA{0, 0, 0, 255}
			}
		}
	}

	upscaled := image.NewPaletted(
		image.Rect(0, 0, gridSize.x*upscaleImageBy, gridSize.y*upscaleImageBy),
		palette,
	)
	var wg sync.WaitGroup
	for y := 0; y < gridSize.y; y++ {
		wg.Add(1)
		go processRow(y, gridSize.x, upscaleImageBy, colorMap, upscaled, &wg)
	}
	wg.Wait()

	return upscaled
}

func EncodeGif(frames []*image.Paletted, delay int, filename string) {
	outputFile, err := os.Create(fmt.Sprintf("day15/%s.gif", filename))
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	delays := make([]int, len(frames))
	for i := range delays {
		delays[i] = delay
	}

	delays[len(delays)-1] = delay * 5

	err = gif.EncodeAll(outputFile, &gif.GIF{
		Image:     frames,
		Delay:     delays,
		LoopCount: -1,
	})

	if err != nil {
		panic(err)
	}
}
