package day2

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/ozancaglar/advent-of-code-2023/util"
)

type cubeCount struct {
	id    int
	cubes cubes
}

type cubes struct {
	red   int
	green int
	blue  int
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

	for _, s := range strings.Split(frame[1], ";") {
		for _, furtherSplitString := range strings.Split(s, ", ") {
			cleanedString := strings.TrimSpace(furtherSplitString) // of the form 3 blue\n4 red\n...
			numberOfCubes, err := strconv.Atoi(string(cleanedString[0]))
			if err != nil {
				log.Fatal(err)
			}
			switch strings.Split(cleanedString, " ")[1] {
			case "red":
				c.cubes.red += numberOfCubes
			case "green":
				c.cubes.green += numberOfCubes
			case "blue":
				c.cubes.blue += numberOfCubes
			default:
				log.Fatalf("unexpected case")
			}

		}
	}
	return c
}

func countIds(cc []cubeCount, desiredNumberOfCubes cubes) int {
	total := 0

	for _, c := range cc {
		if c.cubes.blue >= desiredNumberOfCubes.blue && c.cubes.red >= desiredNumberOfCubes.red && c.cubes.green >= desiredNumberOfCubes.green {
			total += c.id
		}
	}

	return total
}
