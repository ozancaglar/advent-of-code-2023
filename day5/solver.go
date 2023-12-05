package day5

import (
	"log"
	"regexp"
	"slices"
	"strings"

	"github.com/ozancaglar/advent-of-code-2023/util"
)

func Solve(filename string) {
	scanner := util.StreamLines("./day5/input.txt")
	re := regexp.MustCompile("\\d+")
	seeds := make([]int, 0)
	i := -1
	seedToSoilMap := make(map[int]int)
	soilToFertilizerMap := make(map[int]int)
	fertilizerToWater := make(map[int]int)
	waterToLight := make(map[int]int)
	lightToTemperature := make(map[int]int)
	temperatureToHumidity := make(map[int]int)
	humidityToLocation := make(map[int]int)
	var mapToPopulate *map[int]int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "seeds: ") {
			for _, s := range re.FindAllString(line, -1) {
				seeds = append(seeds, util.MustParseInt(s))
			}
		}

		switch {
		case line == "":
			continue
		case strings.Contains(line, "seed-to-soil map:"):
			i = 0
			mapToPopulate = &seedToSoilMap
			continue
		case strings.Contains(line, "soil-to-fertilizer map:"):
			i = 0
			mapToPopulate = &soilToFertilizerMap
			continue
		case strings.Contains(line, "fertilizer-to-water map:"):
			i = 0
			mapToPopulate = &fertilizerToWater
			continue
		case strings.Contains(line, "water-to-light map:"):
			i = 0
			mapToPopulate = &waterToLight
			continue
		case strings.Contains(line, "light-to-temperature map:"):
			i = 0
			mapToPopulate = &lightToTemperature
			continue
		case strings.Contains(line, "temperature-to-humidity map:"):
			i = 0
			mapToPopulate = &temperatureToHumidity
			continue
		case strings.Contains(line, "humidity-to-location map:"):
			i = 0
			mapToPopulate = &humidityToLocation
			continue
		}

		if i == 0 {
			if line == "" {
				i = -1
			}
			populateMap(mapToPopulate, line)
		}
	}
	locations := make(map[int]int)

	for _, s := range seeds {
		fertilizer := noKeyReturnSrc(soilToFertilizerMap, s)
		water := noKeyReturnSrc(fertilizerToWater, fertilizer)
		light := noKeyReturnSrc(waterToLight, water)
		temperature := noKeyReturnSrc(lightToTemperature, light)
		humidity := noKeyReturnSrc(temperatureToHumidity, temperature)
		locations[s] = noKeyReturnSrc(humidityToLocation, humidity)
	}

	log.Printf("Day five, part one answer: %v", smallestLocation(locations))
}

func smallestLocation(locations map[int]int) int {
	locationsSlice := make([]int, 0)
	for _, v := range locations {
		locationsSlice = append(locationsSlice, v)
	}

	return slices.Min(locationsSlice)
}

func noKeyReturnSrc(m map[int]int, k int) int {
	v, ok := m[k]
	if !ok {
		return k
	}

	return v
}

func populateMap(m *map[int]int, line string) {
	splitLine := strings.Split(line, " ")
	splitLineAsInt := make([]int, 0)
	mapToPopulate := *m

	// 0 src, 1 dst, 2 range
	for _, s := range splitLine {
		splitLineAsInt = append(splitLineAsInt, util.MustParseInt(s))
	}

	for i, j := splitLineAsInt[0], splitLineAsInt[1]; i < splitLineAsInt[0]+splitLineAsInt[2]; i, j = i+1, j+1 {
		mapToPopulate[j] = i
	}
}
