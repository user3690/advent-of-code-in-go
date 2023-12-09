package day09

import (
	"fmt"
	"github.com/user3690/advent-of-code-in-go/util"
	"log"
	"strconv"
	"strings"
)

func BothParts() (int64, int64) {
	var (
		lines                                  []string
		histories                              [][]int32
		sumOfAllNextValues, sumOfAllPrevValues int64
		nextVal, prevVal                       int32
		err                                    error
	)

	lines, err = util.ReadFileInLines("./2023/day09/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	histories, err = prepareData(lines)
	if err != nil {
		log.Fatal(err)
	}

	for _, history := range histories {
		nextVal, prevVal = findNextValueInHistory(history)
		sumOfAllNextValues += int64(nextVal)
		sumOfAllPrevValues += int64(prevVal)
	}

	return sumOfAllNextValues, sumOfAllPrevValues
}

func prepareData(lines []string) ([][]int32, error) {
	var (
		history   []int32
		histories = make([][]int32, len(lines))
		number    int64
		err       error
	)

	for i, line := range lines {
		lineSplit := strings.Split(line, " ")
		history = make([]int32, len(lineSplit))
		for j, numberStr := range lineSplit {
			number, err = strconv.ParseInt(numberStr, 10, 32)
			if err != nil {
				return nil, fmt.Errorf("error while parsing number: %w", err)
			}

			history[j] = int32(number)
		}

		histories[i] = history
	}

	return histories, err
}

func findNextValueInHistory(history []int32) (int32, int32) {
	var historyMatrix [][]int32

	historyMatrix = append(historyMatrix, history)

	historyMatrix = calculateSequenceOfDifferences(historyMatrix, history)

	nextValue := calculateNextValue(historyMatrix)
	prevValue := calculatePrevValue(historyMatrix)

	return nextValue, prevValue
}

func calculateSequenceOfDifferences(historyMatrix [][]int32, history []int32) [][]int32 {
	var (
		isNonZeroValue                    = false
		i                                 int
		curNumber, nextNumber, difference int32
		differenceSequence                = make([]int32, len(history)-1)
	)

	for i < len(history)-1 {
		curNumber = history[i]
		nextNumber = history[i+1]

		difference = nextNumber - curNumber

		differenceSequence[i] = difference

		if difference != 0 {
			isNonZeroValue = true
		}

		i++
	}

	historyMatrix = append(historyMatrix, differenceSequence)

	if isNonZeroValue {
		historyMatrix = calculateSequenceOfDifferences(historyMatrix, differenceSequence)
	}

	return historyMatrix
}

func calculateNextValue(historyMatrix [][]int32) int32 {
	var (
		curLastNumber, prevLastNumber, newLastNumber int32
	)

	i := len(historyMatrix) - 2
	for i > 0 {
		curLastNumber = historyMatrix[i][len(historyMatrix[i])-1]
		prevLastNumber = historyMatrix[i-1][len(historyMatrix[i-1])-1]

		newLastNumber = curLastNumber + prevLastNumber
		historyMatrix[i-1] = append(historyMatrix[i-1], newLastNumber)

		i--
	}

	return newLastNumber
}

func calculatePrevValue(historyMatrix [][]int32) int32 {
	var (
		curFirstNumber, prevFirstNumber, newFirstNumber int32
	)

	i := len(historyMatrix) - 2
	for i > 0 {
		curFirstNumber = historyMatrix[i][0]
		prevFirstNumber = historyMatrix[i-1][0]

		if prevFirstNumber < curFirstNumber {
			newFirstNumber = prevFirstNumber - curFirstNumber
		} else {
			newFirstNumber = prevFirstNumber - curFirstNumber
		}

		var newSlice = make([]int32, 1)
		newSlice[0] = newFirstNumber
		newSlice = append(newSlice, historyMatrix[i-1]...)

		historyMatrix[i-1] = newSlice

		i--
	}

	return newFirstNumber
}
