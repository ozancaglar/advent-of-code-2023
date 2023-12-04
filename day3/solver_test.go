package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSymbolCol(t *testing.T) {
	assert.Equal(t, []int{3}, getSymbolCol("...*......", nil))
	assert.Equal(t, []int{}, getSymbolCol("..........", nil))
	assert.Equal(t, []int{2, 9}, getSymbolCol("..*......*", nil))
	rune := '*'
	assert.Equal(t, []int{2, 9}, getSymbolCol("..*......*", &rune))
	assert.Equal(t, []int{2, 9}, getSymbolCol("..*..$...*", &rune))
}

func TestGetDigitsInRow(t *testing.T) {
	assert.Equal(t, &rowNumber{number: 617, indexes: []int{0, 1, 2}, added: false}, getDigitsInRow("617*...617")[0])
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
