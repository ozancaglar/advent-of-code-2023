package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSymbolCol(t *testing.T) {
	assert.Equal(t, []int{3}, getSymbolCol("...*......"))
	assert.Equal(t, []int{}, getSymbolCol(".........."))
	assert.Equal(t, []int{2, 9}, getSymbolCol("..*......*"))
}

func TestGetDigitsInRow(t *testing.T) {
	assert.Equal(t, []rowNumber{}, getDigitsInRow("...*......"))
	assert.Equal(t, []rowNumber{{number: 35, indexes: []int{2, 3}, added: false}, {number: 633, indexes: []int{6, 7, 8}, added: false}}, getDigitsInRow("..35..633."))
	assert.Equal(t, []rowNumber{{number: 617, indexes: []int{0, 1, 2}, added: false}}, getDigitsInRow("617*......"))
	assert.Equal(t, []rowNumber{{number: 617, indexes: []int{0, 1, 2}, added: false}, {number: 617, indexes: []int{7, 8, 9}, added: false}}, getDigitsInRow("617*...617"))
}

func TestGetIndexesToCheck(t *testing.T) {
	assert.Equal(t, map[int][]int{0: {1}, 1: {0, 1}}, getIndexesToCheck(0, 0, 9, 9))
	assert.Equal(t, map[int][]int{0: {8}, 1: {9, 8}}, getIndexesToCheck(0, 9, 9, 9))
	assert.Equal(t, map[int][]int{0: {0, 2}, 1: {0, 1, 2}}, getIndexesToCheck(0, 1, 9, 9))
	assert.Equal(t, map[int][]int{8: {0, 1}, 9: {1}}, getIndexesToCheck(9, 0, 9, 9))
	assert.Equal(t, map[int][]int{8: {0, 1, 2}, 9: {0, 2}}, getIndexesToCheck(9, 1, 9, 9))
	assert.Equal(t, map[int][]int{8: {9, 8}, 9: {8}}, getIndexesToCheck(9, 9, 9, 9))
	assert.Equal(t, map[int][]int{0: {0, 1, 2}, 1: {0, 2}, 2: {0, 1, 2}}, getIndexesToCheck(1, 1, 9, 9))
}

// func TestSumForPartOne(t *testing.T) {
// 	assert.Equal(t, 4361, sumForPartOne())
// }
