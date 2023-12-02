package day2

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/ozancaglar/advent-of-code-2023/util"
)

type cubeCount struct {
	id          int
	cubesPerSet []cubes
}

type cubes struct {
	red   int
	blue  int
	green int
}

func Solve(filename string) {
	partOne(filename)
}

func partOne(filename string) {
	scanner := util.StreamLines(filename)
	countedCubes := make([]cubeCount, 0)

	for scanner.Scan() {
		countedCubes = append(countedCubes, countCubesInGame(scanner.Text()))
	}

	log.Printf("Day two, part one answer: %v", countIds(countedCubes, cubes{red: 12, green: 13, blue: 14}))
}

func countCubesInGame(game string) cubeCount {
	var c cubeCount
	frame := strings.Split(game, ":")
	re := regexp.MustCompile("\\d+")

	i, err := strconv.Atoi(re.FindAllString(frame[0], -1)[0])
	if err != nil {
		log.Fatal(err)
	}
	c.id = i

	j := 0
	for _, s := range strings.Split(frame[1], ";") {
		c.cubesPerSet = append(c.cubesPerSet, cubes{red: 0, blue: 0, green: 0})

		for _, furtherSplitString := range strings.Split(s, ", ") {
			cleanedString := strings.TrimSpace(furtherSplitString) // of the form 3 blue\n4 red\n...
			numberOfCubes, err := strconv.Atoi(re.FindAllString(cleanedString, -1)[0])
			if err != nil {
				log.Fatal(err)
			}

			switch strings.Split(cleanedString, " ")[1] {
			case "red":
				c.cubesPerSet[j].red += numberOfCubes
			case "blue":
				c.cubesPerSet[j].blue += numberOfCubes
			case "green":
				c.cubesPerSet[j].green += numberOfCubes
			default:
				log.Fatalf("unexpected case")
			}
		}
		j += 1
	}
	return c
}

func countIds(cc []cubeCount, totalNumberOfCubes cubes) int {
	total := 0

	for _, c := range cc {
		add := true
		for _, cubesInSet := range c.cubesPerSet {
			if cubesInSet.blue >= totalNumberOfCubes.blue || cubesInSet.red >= totalNumberOfCubes.red || cubesInSet.green >= totalNumberOfCubes.green {
				add = false
			}
		}
		if add {
			total += c.id
		}
	}

	return total
}
