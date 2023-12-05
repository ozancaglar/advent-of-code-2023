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
	scanner := util.StreamLines(filename)
	re := regexp.MustCompile(`\d+`)
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
	soilToSeedMap := make(map[Range]int)
	fertilizerToSoilMap := make(map[Range]int)
	waterToFertilizerMap := make(map[Range]int)
	lightToWaterMap := make(map[Range]int)
	temperatureToLightMap := make(map[Range]int)
	humidityToTemperatureMap := make(map[Range]int)
	locationToHumidityMap := make(map[Range]int)
	var mapToPopulate *map[Range]int
	var reverseMapToPopulate *map[Range]int

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
			reverseMapToPopulate = &soilToSeedMap
			continue
		case strings.Contains(line, "soil-to-fertilizer map:"):
			i = 0
			mapToPopulate = &soilToFertilizerMap
			reverseMapToPopulate = &fertilizerToSoilMap
			continue
		case strings.Contains(line, "fertilizer-to-water map:"):
			i = 0
			mapToPopulate = &fertilizerToWater
			reverseMapToPopulate = &waterToFertilizerMap
			continue
		case strings.Contains(line, "water-to-light map:"):
			i = 0
			mapToPopulate = &waterToLight
			reverseMapToPopulate = &lightToWaterMap
			continue
		case strings.Contains(line, "light-to-temperature map:"):
			i = 0
			mapToPopulate = &lightToTemperature
			reverseMapToPopulate = &temperatureToLightMap
			continue
		case strings.Contains(line, "temperature-to-humidity map:"):
			i = 0
			mapToPopulate = &temperatureToHumidity
			reverseMapToPopulate = &humidityToTemperatureMap
			continue
		case strings.Contains(line, "humidity-to-location map:"):
			i = 0
			mapToPopulate = &humidityToLocation
			reverseMapToPopulate = &locationToHumidityMap
			continue
		}

		if i == 0 {
			if line == "" {
				i = -1
			}
			populateMap(mapToPopulate, reverseMapToPopulate, line)
		}
	}

	locations := make(map[int]int)
	for _, s := range seeds {
		locations[s] = locationForSeed(s, seedToSoilMap,
			soilToFertilizerMap,
			fertilizerToWater,
			waterToLight,
			lightToTemperature,
			temperatureToHumidity,
			humidityToLocation,
		)
	}

	log.Printf("Day five, part one answer: %v", smallestLocation(locations))

	partTwoSeeds := make([]Range, 0)
	for i := 0; i < len(seeds); i += 2 {
		if i+1 < len(seeds) {
			start := seeds[i]
			stop := seeds[i+1]

			if start != 0 && stop != 0 {
				partTwoSeeds = append(partTwoSeeds, Range{Start: start, End: start + stop - 1})
			}
		}
	}

	var partTwoAnswer int

	// Get maxInt
	maxInt := int(^uint(0) >> 1)
	for i := 1; i < maxInt; i++ {
		seed := seedForLocation(i, soilToSeedMap, fertilizerToSoilMap, waterToFertilizerMap, lightToWaterMap, temperatureToLightMap, humidityToTemperatureMap, locationToHumidityMap)
		if validSeed(seed, partTwoSeeds) {
			partTwoAnswer = i
			break
		}
	}

	log.Printf("Day five, part two answer: %v", partTwoAnswer)
}

func validSeed(seed int, seeds []Range) bool {
	for _, validSeedsSeed := range seeds {
		if seed <= validSeedsSeed.End && seed >= validSeedsSeed.Start {
			return true
		}
	}
	return false
}

func locationForSeed(seed int, seedToSoilMap,
	soilToFertilizerMap,
	fertilizerToWater,
	waterToLight,
	lightToTemperature,
	temperatureToHumidity,
	humidityToLocation map[Range]int) int {
	soil := noKeyInRangeReturnSrc(seedToSoilMap, seed)
	fertilizer := noKeyInRangeReturnSrc(soilToFertilizerMap, soil)
	water := noKeyInRangeReturnSrc(fertilizerToWater, fertilizer)
	light := noKeyInRangeReturnSrc(waterToLight, water)
	temperature := noKeyInRangeReturnSrc(lightToTemperature, light)
	humidity := noKeyInRangeReturnSrc(temperatureToHumidity, temperature)
	location := noKeyInRangeReturnSrc(humidityToLocation, humidity)

	return location
}

func seedForLocation(location int, soilToSeedMap,
	fertilizerToSoilMap,
	waterToFertilizerMap,
	lightToWaterMap,
	temperatureToLightMap,
	humidityToTemperatureMap,
	locationToHumidityMap map[Range]int) int {
	humidity := noKeyInRangeReturnSrc(locationToHumidityMap, location)
	temperature := noKeyInRangeReturnSrc(humidityToTemperatureMap, humidity)
	light := noKeyInRangeReturnSrc(temperatureToLightMap, temperature)
	water := noKeyInRangeReturnSrc(lightToWaterMap, light)
	fertilizer := noKeyInRangeReturnSrc(waterToFertilizerMap, water)
	soil := noKeyInRangeReturnSrc(fertilizerToSoilMap, fertilizer)
	return noKeyInRangeReturnSrc(soilToSeedMap, soil)
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

func populateMap(m *map[Range]int, reverseM *map[Range]int, line string) {
	splitLine := strings.Split(line, " ")
	splitLineAsInt := make([]int, 0)
	mapToPopulate := *m

	// 0 src, 1 dst, 2 range
	for _, s := range splitLine {
		splitLineAsInt = append(splitLineAsInt, util.MustParseInt(s))
	}
	//
	r := Range{Start: splitLineAsInt[1], End: splitLineAsInt[1] + splitLineAsInt[2] - 1}
	diff := splitLineAsInt[0] - splitLineAsInt[1]
	mapToPopulate[r] = diff
	reverseMToPopulate := *reverseM
	reverseRange := Range{Start: splitLineAsInt[0], End: splitLineAsInt[0] + splitLineAsInt[2] - 1}
	reverseMToPopulate[reverseRange] = diff * -1
}
