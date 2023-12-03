package day3

import (
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/ozancaglar/advent-of-code-2023/util"
)

func Solve(filename string) {
	scanner := util.StreamLines(filename)
	input := make([]string, 0)
	sum := 0

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	for i, row := range input {
		fmt.Println(i, row)
	}

	log.Printf("Day three, part one answer: %v", sum)
}

func getSymbolCol(input string) (cols []int) {
	symbolColumns := make([]int, 0)
	for i, c := range input {
		if !unicode.IsLetter(c) && c != 46 {
			symbolColumns = append(symbolColumns, i)
		}
	}

	return symbolColumns
}

func getDigitsInRow(row string) []rowNumber {
	rn := make([]rowNumber, 0)
	toContinue := 0
	for i, c := range row {
		if toContinue != 0 {
			toContinue -= 1
			continue
		}

		if unicode.IsNumber(c) {
			for j, k := range row[i:] {
				endOfRow := unicode.IsNumber(k) && i+j == len(row)-1
				if !unicode.IsNumber(k) || endOfRow {
					endIndex := i + j
					if endOfRow {
						endIndex += 1
					}
					number, err := strconv.Atoi(row[i:endIndex])
					if err != nil {
						log.Fatal(err)
					}

					indexes := make([]int, 0)
					for k := i; k < endIndex; k++ {
						indexes = append(indexes, k)
					}
					rn = append(rn, rowNumber{number: number, indexes: indexes})
					toContinue = j
					break
				}
			}
		}
	}
	return rn
}

type rowNumber struct {
	number  int
	indexes []int
	added   bool
}
