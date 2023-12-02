package day1

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/ozancaglar/advent-of-code-2023/util"
)

var (
	wordToDigit = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	wordCorrectionMap = map[string]string{
		"oneight":   "oneeight",
		"twone":     "twoone",
		"sevenine":  "sevennine",
		"nineight":  "nineeight",
		"eightwo":   "eighttwo",
		"eighthree": "eightthree",
	}
)

func Solve(filename string) {
	partOne(filename)
	partTwo(filename)
}

func partOne(filename string) {
	scanner := util.StreamLines(filename)

	re := regexp.MustCompile("[1-9]+")
	total := 0
	for scanner.Scan() {
		matches := re.FindAllString(scanner.Text(), -1)
		joinedMatches := strings.Join(matches, "")

		total += addPoints(joinedMatches)
	}

	log.Printf("Day one, part one answer: %v", total)
}

func partTwo(filename string) {
	scanner := util.StreamLines(filename)

	re := regexp.MustCompile("(?:one|two|three|four|five|six|seven|eight|nine|zero|[0-9])")

	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		for k, v := range wordCorrectionMap {
			text = strings.Replace(text, k, v, -1)
		}
		matches := re.FindAllString(text, -1)
		for i, match := range matches {
			if len(match) == 1 {
				continue
			}
			matches[i] = wordToDigit[match]
		}
		joinedMatches := strings.Join(matches, "")
		total += addPoints(joinedMatches)
	}

	log.Printf("Day one, part two answer: %v", total)
}

func addPoints(joinedMatches string) int {
	i, err := strconv.Atoi(string(joinedMatches[0]) + string(joinedMatches[len(joinedMatches)-1]))
	if err != nil {
		log.Fatal(err)
	}
	return i
}
