package day5

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day5_1 struct {
	data_map  map[mapType][]sourceDestinationRange
	seeds     []int
	startTime time.Time
}

func sourceToDestinationInMap(source int, sdRanges []sourceDestinationRange) int {
	destination := -1
	for _, currentRange := range sdRanges {
		if source >= currentRange.source && source < currentRange.source+currentRange.length {
			return currentRange.destination + (source - currentRange.source)
		}
	}

	if destination == -1 {
		return source
	}
	return destination
}

func destinationToSourceInMap(destination int, sdRanges []sourceDestinationRange) int {
	source := -1
	for _, currentRange := range sdRanges {
		if destination >= currentRange.destination && destination <= currentRange.destination+currentRange.length-1 {
			return currentRange.source + (destination - currentRange.destination)
		}
	}

	if source == -1 {
		return destination
	}

	return source
}

func (d day5_1) getMapTypeFromSource(source string) mapType {
	for currentMapType := range d.data_map {
		if currentMapType.from == source {
			return currentMapType
		}
	}
	return mapType{}
}

func (d day5_1) getMapTypeFromDestination(destination string) mapType {
	for currentMapType := range d.data_map {
		if currentMapType.to == destination {
			return currentMapType
		}
	}
	return mapType{}
}

// NOTE: To traverse the different maps
func (d day5_1) seedToLocation(seed int) int {
	currentMapType := d.getMapTypeFromSource("seed")
	currentSource := seed
	for i := 0; i < len(d.data_map); i++ {
		mapContent := d.data_map[currentMapType]
		currentSource = sourceToDestinationInMap(currentSource, mapContent)
		currentMapType = d.getMapTypeFromSource(currentMapType.to)
	}
	return currentSource
}

// NOTE: To reversely traverse the different maps
func (d day5_1) locationToSeed(location int) int {
	currentMapType := d.getMapTypeFromDestination("location")
	currentDestination := location
	for i := 0; i < len(d.data_map); i++ {
		mapContent := d.data_map[currentMapType]
		currentDestination = destinationToSourceInMap(currentDestination, mapContent)
		if currentDestination == -1 {
			return -1
		}
		currentMapType = d.getMapTypeFromDestination(currentMapType.from)
	}
	return currentDestination
}

func (d day5_1) isSeedInRange(seed int) bool {
	for i := 0; i < len(d.seeds); i += 2 {
		seedRangeStart := d.seeds[i]
		seedRangeLength := d.seeds[i+1]

		if seed >= seedRangeStart && seed <= seedRangeStart+seedRangeLength-1 {
			return true
		}
	}
	return false
}

func (d day5_1) Part1() any {
	lowestLocation := math.MaxInt

	for _, seed := range d.seeds {
		seedLocation := d.seedToLocation(seed)

		if seedLocation < lowestLocation {
			lowestLocation = seedLocation
		}
	}

	return lowestLocation
}

func (d day5_1) Part2() any {
	currentLocation := 0
	for {
		seed := d.locationToSeed(currentLocation)
		isSeedInRange := d.isSeedInRange(seed)
		if seed != -1 && isSeedInRange {
			break
		}
		currentLocation++
	}

	return currentLocation
}

func Solve2() day5_1 {
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

	var data_map = make(map[mapType][]sourceDestinationRange)

	for _, categoryMaps := range category_maps {
		categoryData := strings.Split(categoryMaps, "\n")
		categoryName := categoryData[0][:len(categoryData[0])-1]
		categoryData = categoryData[1:]

		fromTo := strings.Split(categoryName, "-to-")
		from := fromTo[0]
		to := fromTo[1][0:strings.Index(fromTo[1], " ")]

		category := mapType{from, to}

		for _, desination_source_length := range categoryData {
			destination_source_length := strings.Split(desination_source_length, " ")
			destination, _ := strconv.Atoi(destination_source_length[0])
			source, _ := strconv.Atoi(destination_source_length[1])
			length, _ := strconv.Atoi(destination_source_length[2])
			data_map[category] = append(data_map[category], sourceDestinationRange{destination, source, length})
		}
	}

	return day5_1{
		data_map:  data_map,
		seeds:     seeds,
		startTime: startTime,
	}
}

func (d day5_1) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
