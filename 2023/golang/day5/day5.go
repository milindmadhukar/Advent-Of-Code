package day5

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day5 struct {
	data_map  map[MapType][]SeedRange
	seeds     []int
	startTune time.Time
}

func findSeedLocationInSeedMap(seed int, seedMap []SeedRange) int {
	location := -1

	for _, seedRange := range seedMap {
		if seed >= seedRange.source && seed <= seedRange.source+seedRange.length-1 {
			return seedRange.destination + (seed - seedRange.source)
		}
	}

	if location == -1 {
		return seed
	}
	return location
}

func (d day5) getMapTypeFromSource(source string) MapType {
	for mapType := range d.data_map {
		if mapType.from == source {
			return mapType
		}
	}
	return MapType{}
}

func (d day5) seedToLocation(seed int) int {
	mapType := d.getMapTypeFromSource("seed")
	seedMapLocation := seed
	for i := 0; i < len(d.data_map); i++ {
		seedMap := d.data_map[mapType]
		seedMapLocation = findSeedLocationInSeedMap(seedMapLocation, seedMap)
		mapType = d.getMapTypeFromSource(mapType.to)
	}
  return seedMapLocation
}

func (d day5) Part1() any {
	lowestLocation := math.MaxInt

	for _, seed := range d.seeds {
    seedMapLocation := d.seedToLocation(seed)

		if seedMapLocation < lowestLocation {
			lowestLocation = seedMapLocation
		}
	}

	return lowestLocation
}

func (d day5) Part2() any {
	lowestLocations := math.MaxInt

	for i := 0; i < len(d.seeds); i += 2 {
		seedRangeStart := d.seeds[i]
		seedRangeLength := d.seeds[i+1]

    fmt.Println("Finding lowest location for seed range", seedRangeStart, seedRangeLength)

    for k := 0; k < seedRangeLength; k++ {
      go func(k int) {
        seedMapLocation := d.seedToLocation(seedRangeStart + k)

        if seedMapLocation < lowestLocations {
          lowestLocations = seedMapLocation
        }
      }(k)
    }

	}

  return lowestLocations

}

type MapType struct {
	from string
	to   string
}

type SeedRange struct {
	destination int
	source      int
	length      int
}

func Solve() day5 {
	rawData, err := utils.GetRawInputDataFromAOC(2023, 5)
	if err != nil {
		panic(err)
	}

	// exampleFile, _ := os.ReadFile("day5/example.txt")
	// rawData = string(exampleFile)

	startTime := time.Now()

	rawData = strings.Trim(rawData, "\n")

	data := strings.Split(rawData, "\n\n")

	seeds := utils.StringSliceToIntegerSlice(strings.Split(data[0], " ")[1:])

	category_maps := data[1:]

	var data_map = make(map[MapType][]SeedRange)

	for _, category_maps := range category_maps {
		category_data := strings.Split(category_maps, "\n")
		category_name := category_data[0][:len(category_data[0])-1]
		category_data = category_data[1:]

		from_to := strings.Split(category_name, "-to-")
		from := from_to[0]
		to := from_to[1][0:strings.Index(from_to[1], " ")]

		category := MapType{from, to}

		for _, desination_source_length := range category_data {
			destination_source_length := strings.Split(desination_source_length, " ")
			destination, _ := strconv.Atoi(destination_source_length[0])
			source, _ := strconv.Atoi(destination_source_length[1])
			length, _ := strconv.Atoi(destination_source_length[2])
			data_map[category] = append(data_map[category], SeedRange{destination, source, length})
		}
	}

	return day5{
		data_map:  data_map,
		seeds:     seeds,
		startTune: startTime,
	}
}

func (d day5) TimeTaken() time.Duration {
	return time.Since(d.startTune)
}
