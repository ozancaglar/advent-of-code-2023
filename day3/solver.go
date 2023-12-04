package day3

import (
	"log"
	"slices"
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
	scanner := util.StreamLines(filename)
	digitsInRows := make(map[int][]*rowNumber)

	input := make([]string, 0)
	rowNumber := 0
	for scanner.Scan() {
		row := scanner.Text()
		input = append(input, row)
		digitsInRows[rowNumber] = getDigitsInRow(row)
		rowNumber += 1
	}

	maxRow, err := util.LineCounter(filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Day three, part one answer: %v", getSumForPartOne(indexesToCheckMap(input, maxRow), digitsInRows))
	log.Printf("Day three, part two answer: %v", getSumForPartTwo(input, maxRow, digitsInRows))
}

func getSymbolCol(input string, specificSymbol *rune) (cols []int) {
	symbolColumns := make([]int, 0)
	if specificSymbol != nil {
		for i, c := range input {
			if c == *specificSymbol {
				symbolColumns = append(symbolColumns, i)
			}
		}
		return symbolColumns
	}
	for i, c := range input {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != 46 {
			symbolColumns = append(symbolColumns, i)
		}
	}

	return symbolColumns
}

func getDigitsInRow(row string) []*rowNumber {
	rn := []*rowNumber{}
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
					rn = append(rn, &rowNumber{number: number, indexes: indexes})
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

func indexesToCheckMap(input []string, maxRow int) map[int][]int {
	indexesToCheckMaps := make([]map[int][]int, 0)
	indexesToCheck := make(map[int][]int)

	for rowNumber, row := range input {
		symbolsInRow := getSymbolCol(row, nil)
		for _, s := range symbolsInRow {
			indexesToCheckMaps = append(indexesToCheckMaps, getIndexesToCheck(rowNumber, s, maxRow, len(row)-1))
		}
	}

	for _, v := range indexesToCheckMaps {
		util.Merge(indexesToCheck, v)
	}

	return indexesToCheck
}

func getSumForPartOne(indexesToCheck map[int][]int, digitsInRows map[int][]*rowNumber) int {
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

// Part two

func calcGearRatio(a, b int) int {
	return a * b
}

func getSumForPartTwo(input []string, maxRow int, digitsInRows map[int][]*rowNumber) int {
	sum := 0

	for rowNum, row := range input {
		symbolsInRow := getSymbolCol(row, nil)
		for _, s := range symbolsInRow {
			gears := []*rowNumber{}
			// for one symbol '*', these are all the indexes to check
			for rowNum, indexesToCheck := range getIndexesToCheck(rowNum, s, maxRow, len(row)-1) {
				for _, indexToCheck := range indexesToCheck {
					for _, digitInRow := range digitsInRows[rowNum] {
						for _, index := range digitInRow.indexes {
							if indexToCheck == index {
								if !slices.Contains(gears, digitInRow) {
									gears = append(gears, digitInRow)
								}
								// after this, iterate through all digitsInRows and see if there are matches,
								// and if there are then multiply them and add them to the sum, otherwise reset and keep going
							}
						}
					}
				}
			}
			if len(gears) > 1 {
				gearRatio := 1
				for _, v := range gears {
					gearRatio = calcGearRatio(gearRatio, v.number)
				}
				sum += gearRatio
			}

		}
	}
	return sum
}
