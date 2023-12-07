package day7

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPartTwoHandType(t *testing.T) {
	type handTester struct {
		hand     string
		expected int
	}
	htester := []handTester{
		// five of a kind
		{hand: "JJJJJ", expected: 1},
		{hand: "JJJJA", expected: 1},
		{hand: "JJJTT", expected: 1},
		{hand: "JJTTT", expected: 1},
		{hand: "JTTTT", expected: 1},
		{hand: "TTTTT", expected: 1},

		// four of a kind
		{hand: "JJJ3A", expected: 2},
		{hand: "JJTTA", expected: 2},
		{hand: "JTTTA", expected: 2},
		{hand: "AAAAT", expected: 2},
		{hand: "JJATT", expected: 2},

		// full house
		{hand: "JAATT", expected: 3},
		{hand: "23332", expected: 3},

		// three of a kind
		{hand: "JJ234", expected: 4},
		{hand: "J2234", expected: 4},
		{hand: "22234", expected: 4},

		// two pair
		{hand: "2244A", expected: 5},

		// pair
		{hand: "2345J", expected: 6},
		{hand: "23455", expected: 6},

		// one of a kind
		{hand: "23456", expected: 7},
	}
	for _, ht := range htester {
		if fail := assert.Equal(t, ht.expected, getPartTwoHandType(ht.hand)); !fail {
			fmt.Println(ht.hand, "failed, got:", getPartTwoHandType(ht.hand))
		}
	}
}
