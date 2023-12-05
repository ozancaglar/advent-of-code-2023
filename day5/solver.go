package day5

import (
	"log"
	"regexp"
	"slices"
	"strings"

	"github.com/ozancaglar/advent-of-code-2023/util"
)

type Range struct {
	Start int
	End   int
}

func (r *Range) inRange(i int) bool {
	return r.Start <= i && i <= r.End
}

func Solve(filename string) {
	scanner := util.StreamLines("./day5/input.txt")
	re := regexp.MustCompile("\\d+")
	seeds := make([]int, 0)
	i := -1

	// the idea is these maps return the difference you need to add to the seed to get to the soil from a seed in a range so it's more efficient
	seedToSoilMap := make(map[Range]int)
	soilToFertilizerMap := make(map[Range]int)
	fertilizerToWater := make(map[Range]int)
	waterToLight := make(map[Range]int)
	lightToTemperature := make(map[Range]int)
	temperatureToHumidity := make(map[Range]int)
	humidityToLocation := make(map[Range]int)
	var mapToPopulate *map[Range]int

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

	log.Printf("Day five, part one answer: %v", smallestLocation(locationsForSeeds(seeds, seedToSoilMap,
		soilToFertilizerMap,
		fertilizerToWater,
		waterToLight,
		lightToTemperature,
		temperatureToHumidity,
		humidityToLocation)))
}

func locationsForSeeds(seeds []int, seedToSoilMap,
	soilToFertilizerMap,
	fertilizerToWater,
	waterToLight,
	lightToTemperature,
	temperatureToHumidity,
	humidityToLocation map[Range]int) map[int]int {
	locations := make(map[int]int)
	for _, s := range seeds {
		soil := noKeyInRangeReturnSrc(seedToSoilMap, s)
		fertilizer := noKeyInRangeReturnSrc(soilToFertilizerMap, soil)
		water := noKeyInRangeReturnSrc(fertilizerToWater, fertilizer)
		light := noKeyInRangeReturnSrc(waterToLight, water)
		temperature := noKeyInRangeReturnSrc(lightToTemperature, light)
		humidity := noKeyInRangeReturnSrc(temperatureToHumidity, temperature)
		locations[s] = noKeyInRangeReturnSrc(humidityToLocation, humidity)
	}
	return locations
}

func noKeyInRangeReturnSrc(m map[Range]int, key int) int {
	for k, v := range m {
		if k.inRange(key) {
			return key + v
		}
	}
	return key
}

func smallestLocation(locations map[int]int) int {
	locationsSlice := make([]int, 0)
	for _, v := range locations {
		locationsSlice = append(locationsSlice, v)
	}

	return slices.Min(locationsSlice)
}

func populateMap(m *map[Range]int, line string) {
	splitLine := strings.Split(line, " ")
	splitLineAsInt := make([]int, 0)
	mapToPopulate := *m

	// 0 src, 1 dst, 2 range
	for _, s := range splitLine {
		splitLineAsInt = append(splitLineAsInt, util.MustParseInt(s))
	}

	r := Range{Start: splitLineAsInt[1], End: splitLineAsInt[1] + splitLineAsInt[2] - 1}
	diff := splitLineAsInt[0] - splitLineAsInt[1]
	mapToPopulate[r] = diff
}
