package day4

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/ozancaglar/advent-of-code-2023/util"
)

func Solve(filename string) {
	scanner := util.StreamLines(filename)
	re := regexp.MustCompile("\\d+")
	cards := []string{}
	parsedCards := make([]*card, 0)
	winningNumbers := make(map[int]bool)

	for scanner.Scan() {
		splitString := strings.Split(scanner.Text(), "|")
		cards = append(cards, splitString[1])
		for _, n := range re.FindAllString(splitString[0], -1) {
			i, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			winningNumbers[i] = true
		}
	}

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
	sum := 0
	for _, c := range parsedCards {
		for _, n := range c.numbers {
			_, ok := winningNumbers[n]
			if ok {
				c.numberOfWinningNumbers += 1
			}
		}
		fmt.Println(c)
		fmt.Println(c.numberOfWinningNumbers)
		if c.numberOfWinningNumbers > 0 {
			sum += int(math.Pow(2, float64(c.numberOfWinningNumbers)-1))
		}
	}
	fmt.Println(sum)
}

type card struct {
	id                     int
	numberOfWinningNumbers int
	numbers                []int
}
