package day4

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/ozancaglar/advent-of-code-2023/util"
)

func Solve(filename string) {
	scanner := util.StreamLines(filename)
	re := regexp.MustCompile(`\d+`)
	cards := []string{}

	winningNumbers := make(map[int][]int)

	j := 0
	for scanner.Scan() {
		splitString := strings.Split(scanner.Text(), "|")
		cards = append(cards, splitString[1])
		winningNumbersInRow := re.FindAllString(splitString[0], -1)[1:]
		winningNumbersInRowAsInt := make([]int, 0)
		for _, n := range winningNumbersInRow {
			i, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			winningNumbersInRowAsInt = append(winningNumbersInRowAsInt, i)
		}
		winningNumbers[j] = winningNumbersInRowAsInt
		j += 1
	}

	parsedCards := parseCards(cards, re)
	log.Printf("Day four, part one answer: %v", sumForPartOne(parsedCards, winningNumbers))
	log.Printf("Day four, part two answer: %v", partTwo(parsedCards))
}

type card struct {
	id                     int
	numberOfWinningNumbers int
	numbers                []int
}

func parseCards(cards []string, re *regexp.Regexp) []*card {
	parsedCards := make([]*card, 0)
	for cardNumber, c := range cards {
		numbers := make([]int, 0)
		for _, n := range re.FindAllString(c, -1) {
			i, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, i)
		}
		parsedCards = append(parsedCards, util.ToPtr(card{id: cardNumber + 1, numbers: numbers}))
	}
	return parsedCards
}

func sumForPartOne(parsedCards []*card, winningNumbers map[int][]int) int {
	sum := 0
	for i, c := range parsedCards {
		for _, v := range winningNumbers[i] {
			for _, number := range c.numbers {
				if v == number {
					c.numberOfWinningNumbers += 1
				}
			}
		}
		if c.numberOfWinningNumbers > 0 {
			sum += int(math.Pow(2, float64(c.numberOfWinningNumbers)-1))
		}
	}
	return sum
}

func partTwo(parsedCards []*card) int {
	sum := 0
	remainingCalculations := make(map[int]int)
	for i := 1; i <= len(parsedCards)-1; i++ {
		remainingCalculations[i] = 1
	}

	for i, p := range parsedCards {
		if i == 0 {
			sum += 1
			for j := i + p.numberOfWinningNumbers; j > i; j-- {
				if j <= len(parsedCards)-1 {
					remainingCalculations[j] += 1
				}
			}
		}

		if i > 0 {
			for remainingCalculations[i] != 0 {
				sum += 1
				for j := i + p.numberOfWinningNumbers; j > i; j-- {
					if j <= len(parsedCards)-1 {
						remainingCalculations[j] += 1
					}
				}
				remainingCalculations[i] -= 1
			}
		}
	}

	return sum
}
