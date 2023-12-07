package day7

import (
	"log"
	"slices"
	"sort"
	"strings"

	"github.com/ozancaglar/advent-of-code-2023/util"
	"golang.org/x/exp/maps"
)

type hand struct {
	hand string
	bid  int
}

var (
	cardRankings = map[rune]int{
		'A': 1,
		'K': 2,
		'Q': 3,
		'J': 4,
		'T': 5,
		'9': 6,
		'8': 7,
		'7': 8,
		'6': 9,
		'5': 10,
		'4': 11,
		'3': 12,
		'2': 13,
	}
)

func Solve(filename string) {
	highCards := make([]hand, 0)
	onePairs := make([]hand, 0)
	twoPairs := make([]hand, 0)
	threeOfAKinds := make([]hand, 0)
	fullHouses := make([]hand, 0)
	fourOfAKinds := make([]hand, 0)
	fiveOfAKinds := make([]hand, 0)
	scanner := util.StreamLines("day7/input.txt")
	for scanner.Scan() {
		splitHandAndBid := strings.Split(scanner.Text(), " ")
		switch getHandType(splitHandAndBid[0]) {
		case 1:
			fiveOfAKinds = append(fiveOfAKinds, hand{hand: splitHandAndBid[0], bid: util.MustParseInt(splitHandAndBid[1])})
		case 2:
			fourOfAKinds = append(fourOfAKinds, hand{hand: splitHandAndBid[0], bid: util.MustParseInt(splitHandAndBid[1])})
		case 3:
			fullHouses = append(fullHouses, hand{hand: splitHandAndBid[0], bid: util.MustParseInt(splitHandAndBid[1])})
		case 4:
			threeOfAKinds = append(threeOfAKinds, hand{hand: splitHandAndBid[0], bid: util.MustParseInt(splitHandAndBid[1])})
		case 5:
			twoPairs = append(twoPairs, hand{hand: splitHandAndBid[0], bid: util.MustParseInt(splitHandAndBid[1])})
		case 6:
			onePairs = append(onePairs, hand{hand: splitHandAndBid[0], bid: util.MustParseInt(splitHandAndBid[1])})
		case 7:
			highCards = append(highCards, hand{hand: splitHandAndBid[0], bid: util.MustParseInt(splitHandAndBid[1])})
		}
	}

	allHands := []*[]hand{&highCards, &onePairs, &twoPairs, &threeOfAKinds, &fullHouses, &fourOfAKinds, &fiveOfAKinds}
	for _, h := range allHands {
		if len(*h) < 1 {
			continue
		}
		rankHands(*h)
	}
	sum := 0
	i := 1
	for _, h := range allHands {
		for _, hand := range *h {
			sum += hand.bid * i
			i += 1
		}
	}

	log.Printf("Day seven, part one answer: %v", sum)
}

func rankHands(hands []hand) []hand {
	sort.Slice(hands, func(i, j int) bool {
		for k := 0; k < len(hands[i].hand); k++ {
			rankI, _ := cardRankings[rune(hands[i].hand[k])]
			rankJ, _ := cardRankings[rune(hands[j].hand[k])]

			if rankI != rankJ {
				// Compare the first character based on rank
				if k == 0 {
					return rankI > rankJ
				}
				// Compare previous character indexes based on rank
				if k > 0 {
					rankII, _ := cardRankings[rune(hands[i].hand[k-1])]
					rankJJ, _ := cardRankings[rune(hands[j].hand[k-1])]
					if rankII == rankJJ {
						return rankI > rankJ
					}
					return false
				}
			}
		}

		return len(hands[i].hand) > len(hands[j].hand)
	})
	return hands
}

func getHandType(hand string) int {
	m := make(map[rune]int)
	for _, c := range hand {
		_, ok := m[c]
		if !ok {
			m[c] = 1
		} else {
			m[c] += 1
		}
	}

	var threeCheck, twoCheck bool
	switch slices.Max(maps.Values(m)) {
	case 5: // five of a kind
		return 1
	case 4: // four of a kind
		return 2
	case 3:
		threeCheck = true
		// check values and see if three of a kind or full house
	case 2:
		twoCheck = true
		// check values and see if one pair or two pair
	case 1: // high card
		return 7
	default:
		log.Fatal("unreachable condition")
	}
	if threeCheck {
		if slices.Contains(maps.Values(m), 2) {
			return 3 // full house
		} else {
			return 4 // 3 of a kind
		}
	}

	if twoCheck {
		numberOfTwos := 0
		for _, v := range m {
			if v == 2 {
				numberOfTwos += 1
			}
		}
		if numberOfTwos > 1 {
			return 5 // two pair
		}

		return 6 // one pair
	}
	return 0
}
