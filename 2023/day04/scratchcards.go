package day04

import (
	"github.com/user3690/advent-of-code-in-go/util"
	"log"
	"os"
	"slices"
	"strings"
	"unicode"
)

type scratchCard struct {
	Id             int64
	WinningNumbers []int64
	Numbers        []int64
}

func BothParts() (points int64, wonCards int64) {
	var (
		file                    []byte
		cardNumber              int64
		winningNumbers, numbers []int64
		scratchCards            []scratchCard
		err                     error
	)

	file, err = os.ReadFile("./2023/day04/input.txt")
	if err != nil {
		log.Fatalf("error while reading file: %s", err)
	}

	lines := strings.FieldsFunc(string(file), func(r rune) bool {
		return r == '\n'
	})

	for _, line := range lines {
		cardNumber, winningNumbers, numbers, err = prepareData(line)
		if err != nil {
			log.Fatalf("error while parsing line: %s", err)
		}

		newScratchCard := scratchCard{
			Id:             cardNumber,
			WinningNumbers: winningNumbers,
			Numbers:        numbers,
		}

		scratchCards = append(scratchCards, newScratchCard)
	}

	for _, card := range scratchCards {
		points += util.PowInt(2, countWinningNumbers(card)) / 2
	}

	// Part2
	var scratchCardMap = make(map[int64]scratchCard)

	for _, card := range scratchCards {
		scratchCardMap[card.Id] = card
	}

	wonCards = countWonScratchCards(scratchCardMap)

	return points, wonCards
}

func prepareData(line string) (cardNumber int64, winningNumbers []int64, numbers []int64, err error) {
	firstSplit := strings.Split(line, ":")
	cardNumber, err = util.StrToInt64(strings.Trim(strings.SplitAfter(firstSplit[0], "Card")[1], " "))
	if err != nil {
		return cardNumber, winningNumbers, numbers, err
	}

	secondSplit := strings.Split(firstSplit[1], "|")

	winningNumbers, err = findNumbers(secondSplit[0])
	if err != nil {
		return cardNumber, winningNumbers, numbers, err
	}

	numbers, err = findNumbers(secondSplit[1])
	if err != nil {
		return cardNumber, winningNumbers, numbers, err
	}

	return cardNumber, winningNumbers, numbers, err
}

func findNumbers(numbersStr string) (numbers []int64, err error) {
	var (
		start  = -1
		end    = -1
		number int64
	)

	numbersStr = numbersStr + "x"

	for idx, char := range numbersStr {
		if unicode.IsNumber(char) {
			if start == -1 {
				start = idx
			}

			end = idx
		} else {
			if start == -1 {
				continue
			}

			number, err = util.StrToInt64(numbersStr[start : end+1])
			if err != nil {
				return numbers, err
			}

			numbers = append(numbers, number)
			start = -1
			end = -1
		}
	}

	return numbers, err
}

func countWinningNumbers(card scratchCard) int64 {
	var count int64
	for _, number := range card.Numbers {
		if slices.Contains(card.WinningNumbers, number) {
			count++
		}
	}

	return count
}

func countWonScratchCards(scratchCardMap map[int64]scratchCard) int64 {
	var count int64
	for _, card := range scratchCardMap {
		count += giveCardsOfWinningNumbers(card, scratchCardMap) + 1
	}

	return count
}

func giveCardsOfWinningNumbers(card scratchCard, scratchCardMap map[int64]scratchCard) int64 {
	var (
		count, maxId int64
		cards        []scratchCard
	)

	for _, number := range card.Numbers {
		if slices.Contains(card.WinningNumbers, number) {
			count++
		}
	}

	maxId = count + card.Id
	for maxId > card.Id {
		if _, exists := scratchCardMap[maxId]; exists {
			cards = append(cards, scratchCardMap[maxId])
		}

		maxId--
	}

	for _, newCard := range cards {
		count += giveCardsOfWinningNumbers(newCard, scratchCardMap)
	}

	return count
}
