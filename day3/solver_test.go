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
