package day6

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/ozancaglar/advent-of-code-2023/util"
)

func Solve(filename string) {
	scanner := util.StreamLines(filename)
	durationOfRaces, recordDistanceInRaces := parseFile(scanner)
	partOne(durationOfRaces, recordDistanceInRaces)
	partTwo(durationOfRaces, recordDistanceInRaces)
}

func parseFile(scanner bufio.Scanner) ([]int, []int) {
	durationOfRaces := make([]int, 0)
	recordDistanceInRaces := make([]int, 0)
	i := 0
	for scanner.Scan() {
		re := regexp.MustCompile(`\d+`)
		if i%2 == 0 {
			for _, v := range re.FindAllString(scanner.Text(), -1) {
				durationOfRaces = append(durationOfRaces, util.MustParseInt(v))
			}
			i += 1
		} else {
			for _, v := range re.FindAllString(scanner.Text(), -1) {
				recordDistanceInRaces = append(recordDistanceInRaces, util.MustParseInt(v))
			}
		}
	}
	return durationOfRaces, recordDistanceInRaces
}

func calculateDistanceTravelled(buttonHoldTime int, timeInRace int) int {
	timeToTravel := timeInRace - buttonHoldTime

	return timeToTravel * buttonHoldTime
}

func partOne(durationOfRaces, recordDistanceInRaces []int) {
	numberOfTimesRecordCouldBeBeat := make([]int, 0)
	for i, d := range durationOfRaces {
		distancesGreaterThanRecord := make([]int, 0)
		for buttonHoldTime := 0; buttonHoldTime < d; buttonHoldTime++ {
			distanceTravelled := calculateDistanceTravelled(buttonHoldTime, d)
			if distanceTravelled > recordDistanceInRaces[i] {
				distancesGreaterThanRecord = append(distancesGreaterThanRecord, distanceTravelled)
			}
		}
		numberOfTimesRecordCouldBeBeat = append(numberOfTimesRecordCouldBeBeat, len(distancesGreaterThanRecord))
	}

	partOneAnswer := 1
	for _, v := range numberOfTimesRecordCouldBeBeat {
		partOneAnswer = v * partOneAnswer
	}
	log.Printf("Day six, part one answer: %v", partOneAnswer)
}

func partTwo(durationOfRaces, recordDistanceInRaces []int) {
	timeAsString := ""
	distanceAsString := ""

	for i, v := range durationOfRaces {
		timeAsString += strconv.Itoa(v)
		distanceAsString += strconv.Itoa(recordDistanceInRaces[i])
	}
	time := util.MustParseInt(timeAsString)
	distance := util.MustParseInt(distanceAsString)

	fmt.Println(time, distance)
}
