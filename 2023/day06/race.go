package day06

import (
	"fmt"
	"github.com/user3690/advent-of-code-in-go/util"
	"log"
	"strconv"
	"strings"
	"time"
)

type racePart1 struct {
	time     uint16
	distance uint16
}

type racePart2 struct {
	time     uint64
	distance uint64
}

func Part1() uint32 {
	var (
		lines []string
		races []racePart1
		err   error
	)

	start := time.Now()

	lines, err = util.ReadFileInLines("./2023/day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	races, err = prepareDataPart1(lines)
	if err != nil {
		log.Fatal(err)
	}

	wins := calculateCountOfWins(races)

	log.Printf("execution took %v", time.Since(start).Milliseconds())

	return wins
}

func prepareDataPart1(lines []string) ([]racePart1, error) {
	var races []racePart1

	for _, line := range lines {
		if strings.Contains(line, "Time") {
			lineSplit := strings.Split(line, ":")
			numbers, err := util.FindNumbers(lineSplit[1])
			if err != nil {
				return races, fmt.Errorf("error while casting numbers %w", err)
			}

			races = make([]racePart1, len(numbers))
			for i, number := range numbers {
				races[i] = racePart1{
					time:     uint16(number),
					distance: 0,
				}
			}
		}

		if strings.Contains(line, "Distance") {
			lineSplit := strings.Split(line, ":")
			numbers, err := util.FindNumbers(lineSplit[1])
			if err != nil {
				return races, fmt.Errorf("error while casting numbers %w", err)
			}

			for i, number := range numbers {
				newRace := races[i]
				newRace.distance = uint16(number)
				races[i] = newRace
			}
		}
	}

	return races, nil
}

func calculateCountOfWins(races []racePart1) uint32 {
	var (
		winsPerRace, holdButtonTime uint16
		wins                        uint32 = 1
	)

	for _, raceHighScore := range races {
		winsPerRace = 0
		holdButtonTime = 0
		for holdButtonTime <= raceHighScore.time {
			speed := holdButtonTime
			travelTime := raceHighScore.time - holdButtonTime
			distance := speed * travelTime

			if distance > raceHighScore.distance {
				winsPerRace++
			}

			holdButtonTime++
		}

		wins *= uint32(winsPerRace)
	}

	return wins
}

func Part2() uint64 {
	var (
		lines []string
		race  racePart2
		err   error
	)

	start := time.Now()

	lines, err = util.ReadFileInLines("./2023/day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	race, err = prepareDataPart2(lines)
	if err != nil {
		log.Fatal(err)
	}

	wins := calculateAllWinPossibilities(race)

	log.Printf("execution took %v", time.Since(start))

	return wins
}

func prepareDataPart2(lines []string) (racePart2, error) {
	var newRace racePart2

	lineSplit := strings.Split(lines[0], ":")
	number, err := strconv.ParseUint(strings.Replace(lineSplit[1], " ", "", -1), 10, 64)
	if err != nil {
		return newRace, fmt.Errorf("error while casting numbers %w", err)
	}

	newRace.time = number

	lineSplit = strings.Split(lines[1], ":")
	number, err = strconv.ParseUint(strings.Replace(lineSplit[1], " ", "", -1), 10, 64)
	if err != nil {
		return newRace, fmt.Errorf("error while casting numbers %w", err)
	}

	newRace.distance = number

	return newRace, nil
}

func calculateAllWinPossibilities(race racePart2) uint64 {
	var holdButtonTime, wins uint64

	for holdButtonTime <= race.time {
		speed := holdButtonTime
		travelTime := race.time - holdButtonTime
		distance := speed * travelTime

		if distance > race.distance {
			wins++
		}

		holdButtonTime++
	}

	return wins
}
