package day3

import (
	"fmt"
	"log"
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
