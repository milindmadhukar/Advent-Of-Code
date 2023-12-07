package day2

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day2 struct {
	data      []string
	games     map[int]GameCubes
	startTime time.Time
}

func (d day2) Part1() any {
	sum := 0
	for gameId, game := range d.games {
		if game.redCubes[0] > 12 || game.greenCubes[0] > 13 || game.blueCubes[0] > 14 {
			continue
		}
		sum += gameId
	}

	return sum
}

func (d day2) Part2() any {
	sum := 0
	for _, game := range d.games {
		sum += game.redCubes[0] * game.greenCubes[0] * game.blueCubes[0]
	}

	return sum
}

type GameCubes struct {
	redCubes   []int
	blueCubes  []int
	greenCubes []int
}

func Solve() day2 {
	data, err := utils.GetInputDataFromAOC(2023, 2)
	if err != nil {
		panic(err)
	}

	// exampleFile, _ := os.ReadFile("day2/example.txt")
	// data = utils.ParseFromString(string(exampleFile))

	startTime := time.Now()

	var games = make(map[int]GameCubes)

	for _, line := range data {
		game := strings.Split(line, ": ")
		gameId, _ := strconv.Atoi(game[0][5:])

		cubes := game[1]
		redRegex := regexp.MustCompile(`\d+\ (red)`)
		blueRegex := regexp.MustCompile(`\d+\ (blue)`)
		greenRegex := regexp.MustCompile(`\d+\ (green)`)

		var redMatches, blueMatches, greenMatches []int

		redCubes := utils.GetSplitData(redRegex.FindAllString(cubes, -1), " ")
		for _, redCube := range redCubes {
			redCubeInt, _ := strconv.Atoi(redCube[0])
			redMatches = append(redMatches, redCubeInt)
		}

		blueCubes := utils.GetSplitData(blueRegex.FindAllString(cubes, -1), " ")
		for _, blueCube := range blueCubes {
			blueCubeInt, _ := strconv.Atoi(blueCube[0])
			blueMatches = append(blueMatches, blueCubeInt)
		}

		greenCubes := utils.GetSplitData(greenRegex.FindAllString(cubes, -1), " ")
		for _, greenCube := range greenCubes {
			greenCubeInt, _ := strconv.Atoi(greenCube[0])
			greenMatches = append(greenMatches, greenCubeInt)
		}

		slices.SortFunc(redMatches, func(a, b int) int {
			return b - a
		})
		slices.SortFunc(blueMatches, func(a, b int) int {
			return b - a
		})
		slices.SortFunc(greenMatches, func(a, b int) int {
			return b - a
		})

		games[gameId] = GameCubes{
			redCubes:   redMatches,
			blueCubes:  blueMatches,
			greenCubes: greenMatches,
		}
	}

	return day2{
		data:      data,
		games:     games,
		startTime: startTime,
	}
}

func (d day2) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
