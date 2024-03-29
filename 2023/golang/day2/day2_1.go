package day2

import (
	"strconv"
	"strings"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day2_1 struct {
	data map[int]Game
  startTime time.Time
}

func (d day2_1) Part1() any {
	sum := 0

	for gameId, game := range d.data {
		redMaxCount := 12
		greenMaxCount := 13
		blueMaxCount := 14

		valid := true
	gameLoop:
		for _, cubeSet := range game {
			for _, cube := range cubeSet {
				if (cube.colour == "red" && cube.quantity > redMaxCount) || (cube.colour == "blue" && cube.quantity > blueMaxCount) || (cube.colour == "green" && cube.quantity > greenMaxCount) {
					valid = false
					break gameLoop
				}
			}
		}
		if valid {
			sum += gameId
		}
	}

	return sum
}

func (d day2_1) Part2() any {
	sum := 0

	for _, game := range d.data {
		var minRed, minGreen, minBlue int
		for _, cubeSet := range game {
			for _, cube := range cubeSet {
				if cube.colour == "red" && cube.quantity > minRed {
					minRed = cube.quantity
				} else if cube.colour == "green" && cube.quantity > minGreen {
					minGreen = cube.quantity
				} else if cube.colour == "blue" && cube.quantity > minBlue {
					minBlue = cube.quantity
				}
			}
		}
		power := minRed * minBlue * minGreen

		sum += power
	}

	return sum
}

type Cube struct {
	quantity int
	colour   string
}

type CubeSet []Cube

type Game []CubeSet

func Solve2() day2_1 {
	rawData, err := utils.GetInputDataFromAOC(2023, 2)
	if err != nil {
		panic(err)
	}


  startTime := time.Now()

	var data = make(map[int]Game)

	gamesToCubes := utils.GetSplitData(rawData, ": ")
	for _, gameToCube := range gamesToCubes {
		gameId, _ := strconv.Atoi(gameToCube[0][5:])
		gameCubeSetsSlice := strings.Split(gameToCube[1], "; ")
		var game Game
		for _, gameCubeSet := range gameCubeSetsSlice {
			cubesSlice := strings.Split(gameCubeSet, ", ")
			var cubeSet CubeSet
			for _, cube := range cubesSlice {
				cubeSlice := strings.Split(cube, " ")
				cubeAmount, _ := strconv.Atoi(cubeSlice[0])
				cubeColour := cubeSlice[1]
				cubeSet = append(cubeSet, Cube{
					quantity: cubeAmount,
					colour:   cubeColour,
				})
			}
			game = append(game, cubeSet)
		}
		data[gameId] = game
	}

	return day2_1{
		data: data,
    startTime: startTime,
	}
}


func (d day2_1) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
