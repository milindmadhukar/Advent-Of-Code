package day5

import (
	"math"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

type day5 struct {
	data_map  map[mapType][]sourceDestinationRange
	seeds     []int
	startTime time.Time
}

type mapRange struct {
	start    int
	end      int
	negative bool
}

type sourceDestinationRange struct {
	destination int
	source      int
	length      int
}

type mapType struct {
	from string
	to   string
}

func (d day5) getMapTypeFromSource(source string) mapType {
	for currentMapType := range d.data_map {
		if currentMapType.from == source {
			return currentMapType
		}
	}
	return mapType{}
}

func getCompleteMapRange(sdRanges []sourceDestinationRange) []mapRange {
	var mapRanges []mapRange

	for _, sdRange := range sdRanges {
		mapRanges = append(mapRanges, mapRange{sdRange.source, sdRange.source + sdRange.length - 1, false})
	}

	start := sdRanges[0].source

	if start > 0 {
		mapRanges = append(mapRanges, mapRange{0, start - 1, true})
	}

	for i := 0; i < len(sdRanges)-1; i++ {
		currentRange := sdRanges[i]
		nextRange := sdRanges[i+1]
		if currentRange.source+currentRange.length != nextRange.source {
			mapRanges = append(mapRanges, mapRange{currentRange.source + currentRange.length, nextRange.source - 1, true})
		}
	}

	end := sdRanges[len(sdRanges)-1].source

	if end < math.MaxInt {
		mapRanges = append(mapRanges, mapRange{end + 1, math.MaxInt, true})
	}

	slices.SortFunc(mapRanges, func(a mapRange, b mapRange) int {
		return a.start - b.start
	})

	return mapRanges
}

func intersection(a mapRange, b mapRange) mapRange {
	var intersection mapRange

	if a.start <= b.end && b.start <= a.end {
		intersection.start = max(a.start, b.start)
		intersection.end = min(a.end, b.end)
	}

	return intersection
}

func (d day5) Walk1(currentValue int, currentMapType mapType) int {
	emptyMap := mapType{}
	if currentMapType == emptyMap {
		return currentValue
	}

	sdRanges := d.data_map[currentMapType]

	for _, currentRange := range sdRanges {
		if currentValue >= currentRange.source && currentValue < currentRange.source+currentRange.length {
			newValue := currentRange.destination + (currentValue - currentRange.source)
			return d.Walk1(newValue, d.getMapTypeFromSource(currentMapType.to))
		}
	}

	return d.Walk1(currentValue, d.getMapTypeFromSource(currentMapType.to))
}

func (d day5) Walk2(currentMapRanges []mapRange, currentMapType mapType) []mapRange {
	emptyMap := mapType{}
	if currentMapType == emptyMap {
		return currentMapRanges
	}

	var newMapRange []mapRange

	completeRange := getCompleteMapRange(d.data_map[currentMapType])

	orignalRangesIdx := -1
	for _, currentRange := range completeRange {
		if !currentRange.negative {
			orignalRangesIdx++
		}
		for _, currentMapRange := range currentMapRanges {
			rangeIntersection := intersection(currentRange, currentMapRange)
			emptyMap := mapRange{}
			if rangeIntersection == emptyMap {
				continue
			}
			if currentRange.negative {
				newMapRange = append(newMapRange, rangeIntersection)
			} else {
        sdRange := d.data_map[currentMapType][orignalRangesIdx]
        start := sdRange.destination + (rangeIntersection.start - currentRange.start)
        end := sdRange.destination + (rangeIntersection.end - currentRange.start)
				newMapRange = append(newMapRange, mapRange{
					start:    start,
					end:      end,
					negative: false,
				})
			}
		}
	}

	return d.Walk2(newMapRange, d.getMapTypeFromSource(currentMapType.to))
}

func (d day5) Part1() any {
	lowestLocation := math.MaxInt

	for _, seed := range d.seeds {
		location := d.Walk1(seed, d.getMapTypeFromSource("seed"))
		if location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func (d day5) Part2() any {
	var seedRanges []mapRange
	for i := 0; i < len(d.seeds); i += 2 {
		startSeed := d.seeds[i]
		rangeLength := d.seeds[i+1]
		seedRanges = append(seedRanges, mapRange{
			start: startSeed,
			end:   startSeed + rangeLength - 1,
		})
	}

	finalRanges := d.Walk2(seedRanges, d.getMapTypeFromSource("seed"))

	slices.SortFunc(finalRanges, func(a mapRange, b mapRange) int {
		return a.start - b.start
	})

	return finalRanges[0].start
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

		slices.SortFunc(data_map[category], func(a, b sourceDestinationRange) int {
			return a.source - b.source
		})
	}

	return day5{
		data_map:  data_map,
		seeds:     seeds,
		startTime: startTime,
	}
}

func (d day5) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
