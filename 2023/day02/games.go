package day02

import (
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type color = int

const (
	unknown color = iota
	red
	green
	blue
)

type game struct {
	Id      int64
	Results []result
}

type result struct {
	Red   int64
	Green int64
	Blue  int64
}

func BothParts() (partOne int64, partTwo int64) {
	var (
		file  []byte
		games []game
		err   error
	)

	file, err = os.ReadFile("./2023/day02/input.txt")
	if err != nil {
		log.Fatalf("error while reading file: %s", err)
	}

	lines := strings.FieldsFunc(string(file), func(r rune) bool {
		return r == '\n'
	})

	games, err = prepareData(lines)
	if err != nil {
		log.Fatalf("error while preparing data: %s", err)
	}

	for _, playedGame := range games {
		fmt.Printf("%+v\n", playedGame)
	}

	return sumIdsOfPossibleGames(games), calculateComputePower(games)
}

func prepareData(lines []string) (games []game, err error) {
	var (
		gameId, cubeCount int64
		cubeColor         int
	)

	for _, line := range lines {
		var results []result
		splits := strings.Split(line, ":")

		gameIdString := splits[0]
		gameIdSplit := strings.Split(gameIdString, " ")
		gameId, err = strconv.ParseInt(strings.TrimSpace(gameIdSplit[1]), 10, 64)
		if err != nil {
			err = fmt.Errorf("error while parsing id: %w", err)

			return nil, err
		}

		allGamesPlayed := splits[1]
		allGamesSplitted := strings.Split(allGamesPlayed, ";")

		for _, gameResult := range allGamesSplitted {
			cubes := strings.Split(gameResult, ",")

			newResult := result{}
			for _, cubeResult := range cubes {
				cubeColor, cubeCount, err = detectCubeColorAndCount(cubeResult)
				if err != nil {
					err = fmt.Errorf("error detecting cube color and count: %w", err)

					return nil, err
				}

				if cubeColor == red {
					newResult.Red = cubeCount
				}

				if cubeColor == blue {
					newResult.Blue = cubeCount
				}

				if cubeColor == green {
					newResult.Green = cubeCount
				}
			}

			results = append(results, newResult)
		}

		newGame := game{
			Id:      gameId,
			Results: results,
		}

		games = append(games, newGame)
	}

	return games, err
}

func detectCubeColorAndCount(cubes string) (color color, count int64, err error) {
	if strings.Contains(cubes, "green") {
		count, err = strconv.ParseInt(strings.Split(cubes, " ")[1], 10, 64)
		if err != nil {
			err = fmt.Errorf("error while parsing id: %w", err)

			return unknown, 0, err
		}

		return green, count, err
	}

	if strings.Contains(cubes, "red") {
		count, err = strconv.ParseInt(strings.Split(cubes, " ")[1], 10, 64)
		if err != nil {
			err = fmt.Errorf("error while parsing id: %w", err)

			return unknown, 0, err
		}

		return red, count, err
	}

	if strings.Contains(cubes, "blue") {
		count, err = strconv.ParseInt(strings.Split(cubes, " ")[1], 10, 64)
		if err != nil {
			err = fmt.Errorf("error while parsing id: %w", err)

			return unknown, 0, err
		}

		return blue, count, err
	}

	return unknown, 0, errors.New(fmt.Sprintf("no valid color found: %s", cubes))
}

func sumIdsOfPossibleGames(games []game) int64 {
	var sum int64

	for _, playedGame := range games {
		if isPossibleGame(playedGame.Results) {
			sum += playedGame.Id
		}
	}

	return sum
}

func isPossibleGame(results []result) bool {
	for _, resultSet := range results {
		if resultSet.Red > 12 {
			return false
		}

		if resultSet.Green > 13 {
			return false
		}

		if resultSet.Blue > 14 {
			return false
		}
	}

	return true
}

func calculateComputePower(games []game) int64 {
	var computePower int64

	for _, playedGame := range games {
		var reds, greens, blues []int64
		for _, resultSet := range playedGame.Results {
			reds = append(reds, resultSet.Red)
			greens = append(greens, resultSet.Green)
			blues = append(blues, resultSet.Blue)
		}

		maxRed := slices.Max(reds)
		maxGreen := slices.Max(greens)
		maxBlue := slices.Max(blues)

		computePower += maxRed * maxGreen * maxBlue
	}

	return computePower
}
