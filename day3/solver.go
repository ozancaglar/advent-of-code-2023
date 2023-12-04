package day3

import (
	"log"
	"maps"
	"strconv"
	"unicode"

	"github.com/ozancaglar/advent-of-code-2023/util"
)

type rowNumber struct {
	number  int
	indexes []int
	added   bool
}

func Solve(filename string) {
	filename = "day3/test.txt"
	scanner := util.StreamLines(filename)
	digitsInRows := make(map[int][]rowNumber)

	input := make([]string, 0)
	rowNumber := 0
	for scanner.Scan() {
		row := scanner.Text()
		input = append(input, row)
		digitsInRows[rowNumber] = getDigitsInRow(row)
		rowNumber += 1
	}

	indexesToCheck := indexesToCheckMap(input, filename)

	log.Printf("Day three, part one answer: %v", getSumForPartOne(indexesToCheck, digitsInRows))
}

func getSymbolCol(input string) (cols []int) {
	symbolColumns := make([]int, 0)
	for i, c := range input {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != 46 {
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

func indexesToCheckMap(input []string, filename string) map[int][]int {
	maxRow, err := util.LineCounter(filename)
	if err != nil {
		log.Fatal(err)
	}
	indexesToCheckMaps := make([]map[int][]int, 0)
	indexesToCheck := make(map[int][]int)

	for rowNumber, row := range input {
		symbolsInRow := getSymbolCol(row)
		for _, s := range symbolsInRow {
			indexesToCheckMaps = append(indexesToCheckMaps, getIndexesToCheck(rowNumber, s, maxRow, len(row)-1))
		}
	}

	for _, v := range indexesToCheckMaps {
		maps.Copy(indexesToCheck, v)
	}

	return indexesToCheck

}

func getSumForPartOne(indexesToCheck map[int][]int, digitsInRows map[int][]rowNumber) int {
	sum := 0

	for rowNumber := range indexesToCheck {
		for _, indexToCheckInRow := range indexesToCheck[rowNumber] {
			for _, digitInRow := range digitsInRows[rowNumber] {
				if digitInRow.added {
					continue
				}
				for _, i := range digitInRow.indexes {
					if i == indexToCheckInRow {
						sum += digitInRow.number
						digitInRow.added = true
					}
				}
			}
		}
	}
	return sum
}
