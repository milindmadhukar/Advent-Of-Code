package day14

import (
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

func visualize(gridSize Point, allPositions []map[Point]bool) {
	upscaleImageBy := 10
	var frames []*image.Paletted
	var delays []int

	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0, 240, 50, 255},
	}

	count := 0
	delayPerFrame := 25

	for _, positionMap := range allPositions {
		robotPositionColours := make(map[Point]color.RGBA)
		for pos := range positionMap {
			robotPositionColours[pos] = color.RGBA{0, 240, 50, 255}
		}

		upscaled := image.NewPaletted(
			image.Rect(0, 0, gridSize.x*upscaleImageBy, gridSize.y*upscaleImageBy),
			palette,
		)
		var wg sync.WaitGroup
		for y := 0; y < gridSize.y; y++ {
			wg.Add(1)
			go processRow(y, gridSize.x, upscaleImageBy, robotPositionColours, upscaled, &wg)
		}
		wg.Wait()
		frames = append(frames, upscaled)
		delays = append(delays, delayPerFrame)

		count++
	}

	outputFile, err := os.Create("day14/robots.gif")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = gif.EncodeAll(outputFile, &gif.GIF{
		Image: frames,
		Delay: delays,
	})

	if err != nil {
		panic(err)
	}
}
