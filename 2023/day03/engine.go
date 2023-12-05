package day03

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type symbolPosition struct {
	symbol   string
	Position int64
}

type numberPosition struct {
	Value    int64
	Position int64
	Length   int64
}

func BothParts() (partOne int64, partTwo int64) {
	var (
		file                              []byte
		err                               error
		searchMap                         map[int64]string
		validPartNumbers, validGearRatios []int64
	)

	file, err = os.ReadFile("./2023/day03/input.txt")
	if err != nil {
		log.Fatalf("error while reading file: %s", err)
	}

	lines := strings.FieldsFunc(string(file), func(r rune) bool {
		return r == '\n'
	})

	searchMap = make(map[int64]string)
	for i, line := range lines {
		searchMap[int64(i+1)] = line
	}

	for n := range searchMap {
		curLine, prevLine, nextLine := getLines(n, searchMap)

		posSymbols := findPositionOfSymbols(curLine)
		curPosNumbers := findPositionOfNumbers(curLine)
		prevPosNumbers := findPositionOfNumbers(prevLine)
		nextPosNumbers := findPositionOfNumbers(nextLine)

		validPartNumbers = append(
			validPartNumbers,
			findValidPartNumbers(posSymbols, curPosNumbers, prevPosNumbers, nextPosNumbers)...,
		)

		validGearRatios = append(
			validGearRatios,
			findValidGearRatios(posSymbols, curPosNumbers, prevPosNumbers, nextPosNumbers),
		)
	}

	for _, partNumber := range validPartNumbers {
		partOne += partNumber
	}

	for _, sumGearRatio := range validGearRatios {
		partTwo += sumGearRatio
	}

	return partOne, partTwo
}

func getLines(lineNumber int64, searchMap map[int64]string) (curLine string, prevLine string, nextLine string) {
	curLine = searchMap[lineNumber]

	if _, exists := searchMap[lineNumber-1]; exists {
		prevLine = searchMap[lineNumber-1]
	}

	if _, exists := searchMap[lineNumber+1]; exists {
		nextLine = searchMap[lineNumber+1]
	}

	return curLine, prevLine, nextLine
}

func findPositionOfSymbols(line string) []symbolPosition {
	var symbolPositions []symbolPosition
	for {
		index := strings.IndexAny(line, "*=/%@-+$#&")
		if index == -1 {
			break
		}

		symbol := fmt.Sprintf("%c", line[index])

		line = strings.Replace(line, symbol, ".", 1)

		newSymbolPosition := symbolPosition{
			symbol:   symbol,
			Position: int64(index),
		}

		symbolPositions = append(symbolPositions, newSymbolPosition)
	}

	return symbolPositions
}

func findPositionOfNumbers(line string) (numberPos []numberPosition) {
	var (
		conv int64
		err  error
	)

	expr, err := regexp.Compile("\\d+")
	if err != nil {
		log.Fatal(err)
	}

	numbers := expr.FindAllString(line, -1)
	indexForNumbers := expr.FindAllStringIndex(line, -1)

	for i, found := range numbers {
		conv, err = strconv.ParseInt(found, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		newNumberPos := numberPosition{
			Value:    conv,
			Position: int64(indexForNumbers[i][0]),
			Length:   int64(indexForNumbers[i][1] - indexForNumbers[i][0]),
		}

		numberPos = append(numberPos, newNumberPos)
	}

	return numberPos
}

func findValidPartNumbers(
	symbolPositions []symbolPosition,
	curNumberPos []numberPosition,
	prevNumberPos []numberPosition,
	nextNumberPos []numberPosition,
) []int64 {
	var partNumbers []int64

	for _, symbol := range symbolPositions {
		for _, numberPos := range curNumberPos {
			// number right of symbol
			if symbol.Position+1 == numberPos.Position {
				partNumbers = append(partNumbers, numberPos.Value)
			}

			// number left of symbol
			if symbol.Position-1 == (numberPos.Position + numberPos.Length - 1) {
				partNumbers = append(partNumbers, numberPos.Value)
			}
		}

		// numbers above symbol
		for _, numberPos := range prevNumberPos {
			if symbol.Position >= numberPos.Position-1 && symbol.Position <= (numberPos.Position+numberPos.Length) {
				partNumbers = append(partNumbers, numberPos.Value)
			}
		}

		// numbers under symbol
		for _, numberPos := range nextNumberPos {
			if symbol.Position >= numberPos.Position-1 && symbol.Position <= (numberPos.Position+numberPos.Length) {
				partNumbers = append(partNumbers, numberPos.Value)
			}
		}
	}

	return partNumbers
}

func findValidGearRatios(
	symbolPositions []symbolPosition,
	curNumberPos []numberPosition,
	prevNumberPos []numberPosition,
	nextNumberPos []numberPosition,
) int64 {
	var gearRatios []int64
	var sumGearRatios int64

	for _, symbol := range symbolPositions {
		gearRatios = []int64{}
		if symbol.symbol != "*" {
			continue
		}

		for _, numberPos := range curNumberPos {
			// number right of symbol
			if symbol.Position+1 == numberPos.Position {
				gearRatios = append(gearRatios, numberPos.Value)
			}

			// number left of symbol
			if symbol.Position-1 == (numberPos.Position + numberPos.Length - 1) {
				gearRatios = append(gearRatios, numberPos.Value)
			}
		}

		// numbers above symbol
		for _, numberPos := range prevNumberPos {
			if symbol.Position >= numberPos.Position-1 && symbol.Position <= (numberPos.Position+numberPos.Length) {
				gearRatios = append(gearRatios, numberPos.Value)
			}
		}

		// numbers under symbol
		for _, numberPos := range nextNumberPos {
			if symbol.Position >= numberPos.Position-1 && symbol.Position <= (numberPos.Position+numberPos.Length) {
				gearRatios = append(gearRatios, numberPos.Value)
			}
		}

		if len(gearRatios) != 2 {
			continue
		}

		sumGearRatios += gearRatios[0] * gearRatios[1]
	}

	return sumGearRatios
}
