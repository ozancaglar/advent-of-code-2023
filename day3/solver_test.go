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
