package day6

import (
	"bufio"
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

func calculateDistanceTravelled(buttonHoldTime int, raceDuration int) int {
	timeToTravel := raceDuration - buttonHoldTime

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
	distanceToBeat := util.MustParseInt(distanceAsString)
	startedBeatingRecord := 0
	stoppedBeatingRecord := 0
	for i := 1; i < time+10000; i += 10000 {
		if startedBeatingRecord != 0 && stoppedBeatingRecord != 0 {
			break
		}

		if startedBeatingRecord == 0 {
			startedBeatingRecord = findStartRecord(i, time, distanceToBeat)
		}

		if startedBeatingRecord != 0 && stoppedBeatingRecord == 0 {
			stoppedBeatingRecord = findStopRecord(startedBeatingRecord, time, distanceToBeat)
		}
	}

	log.Printf("Day six, part two answer: %v", stoppedBeatingRecord-startedBeatingRecord)
}

func findStartRecord(timeToTry, raceTimeDuration, distanceToBeat int) int {
	distanceTravelled := calculateDistanceTravelled(timeToTry, raceTimeDuration)
	if distanceTravelled > distanceToBeat {
		for timeInLinearSearchRange := timeToTry - 10000; timeInLinearSearchRange < timeToTry; timeInLinearSearchRange++ {
			if calculateDistanceTravelled(timeInLinearSearchRange, raceTimeDuration) > distanceToBeat {
				return timeInLinearSearchRange
			}
		}
	}
	return 0
}

func findStopRecord(startedBeatingRecord, raceTimeDuration, distanceToBeat int) int {
	for k := startedBeatingRecord; k < raceTimeDuration+10000; k += 10000 {
		distanceTravelled := calculateDistanceTravelled(k, raceTimeDuration)
		if distanceTravelled < distanceToBeat {
			for timeInLinearSearchRange := k - 10000; timeInLinearSearchRange < k; timeInLinearSearchRange++ {
				if calculateDistanceTravelled(timeInLinearSearchRange, raceTimeDuration) < distanceToBeat {
					return timeInLinearSearchRange
				}
			}
		}
	}
	return 0
}
