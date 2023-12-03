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

func getIndexesToCheck(rowNumber, colNumber, maxRow, maxCol int) map[int][]int {
	m := make(map[int][]int)
	minRow, minCol := 0, 0

	if rowNumber == minRow && colNumber == minCol {
		m[rowNumber] = []int{colNumber + 1}
		m[rowNumber+1] = []int{colNumber, colNumber + 1}
		return m
	}

	if rowNumber == maxRow && colNumber == minCol {
		m[rowNumber-1] = []int{colNumber, colNumber + 1}
		m[rowNumber] = []int{colNumber + 1}
		return m
	}

	if rowNumber == minRow && colNumber == maxCol {
		m[rowNumber] = []int{colNumber - 1}
		m[rowNumber+1] = []int{colNumber, colNumber - 1}
		return m
	}

	if rowNumber == maxRow && colNumber == maxCol {
		m[rowNumber-1] = []int{colNumber, colNumber - 1}
		m[rowNumber] = []int{colNumber - 1}
		return m
	}

	if rowNumber == minRow {
		m[rowNumber] = []int{colNumber - 1, colNumber + 1}
		m[rowNumber+1] = []int{colNumber - 1, colNumber, colNumber + 1}
		return m
	}

	if rowNumber == maxRow {
		m[rowNumber-1] = []int{colNumber - 1, colNumber, colNumber + 1}
		m[rowNumber] = []int{colNumber - 1, colNumber + 1}
		return m
	}

	if colNumber == minCol {
		m[rowNumber-1] = []int{colNumber, colNumber + 1}
		m[rowNumber] = []int{colNumber + 1}
		m[rowNumber+1] = []int{colNumber, colNumber + 1}
		return m
	}

	if colNumber == maxCol {
		m[rowNumber-1] = []int{colNumber - 1, colNumber}
		m[rowNumber] = []int{colNumber - 1}
		m[rowNumber+1] = []int{colNumber - 1, colNumber}
		return m
	}

	m[rowNumber-1] = []int{colNumber - 1, colNumber, colNumber + 1}
	m[rowNumber] = []int{colNumber - 1, colNumber + 1}
	m[rowNumber+1] = []int{colNumber - 1, colNumber, colNumber + 1}

	return m
}

type rowNumber struct {
	number  int
	indexes []int
	added   bool
}
