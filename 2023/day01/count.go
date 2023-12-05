package day01

import (
	"fmt"
	"github.com/user3690/advent-of-code-in-go/util"
	"log"
	"os"
	"strings"
)

type wordPosition struct {
	Word  string
	Index int
}

type numberPosition struct {
	Number string
	Index  int
}

func BothParts() (numberSum int64, WordAndNumberSum int64) {
	var (
		number        int64
		relevantChars []string
		err           error
	)

	file, err := os.ReadFile("./2023/day01/input.txt")
	if err != nil {
		log.Fatalf("error while reading file: %s", err)
	}

	stringSlice := strings.FieldsFunc(string(file), func(r rune) bool {
		return r == '\n'
	})

	for _, line := range stringSlice {
		relevantChars = []string{}

		firstIndex := strings.IndexAny(line, "1234567890")
		lastIndex := strings.LastIndexAny(line, "1234567890")

		firstSingleDigit := line[firstIndex]
		lastSingleDigit := line[lastIndex]

		firstNumber := numberPosition{
			Number: fmt.Sprintf("%c", firstSingleDigit),
			Index:  firstIndex,
		}

		lastNumber := numberPosition{
			Number: fmt.Sprintf("%c", lastSingleDigit),
			Index:  lastIndex,
		}

		number, err = util.StrToInt64(firstNumber.Number + lastNumber.Number)
		numberSum += number

		var (
			wordMap   map[string]bool
			wordSlice []wordPosition
		)

		wordMap = findWords(line)
		for foundWord := range wordMap {
			firstWordIndex := strings.Index(line, foundWord)
			lastWordIndex := strings.LastIndex(line, foundWord)

			wordSlice = append(wordSlice, wordPosition{
				Word:  foundWord,
				Index: firstWordIndex,
			})

			wordSlice = append(wordSlice, wordPosition{
				Word:  foundWord,
				Index: lastWordIndex,
			})
		}

		firstWord, lastWord := findFirstAndLastWord(wordSlice)

		relevantChars = findFirstAndLastNumber(firstWord, lastWord, firstNumber, lastNumber)

		joined := strings.Join(relevantChars, "")
		number, err = util.StrToInt64(joined)
		if err != nil {
			log.Fatalf("error while converting string to int: %s", err)
		}

		WordAndNumberSum += number
	}

	return numberSum, WordAndNumberSum
}

func findWords(line string) map[string]bool {
	var (
		exists     bool
		foundWords map[string]bool
	)

	foundWords = make(map[string]bool)

	exists = strings.Contains(line, "one")
	if exists {
		foundWords["one"] = true
	}
	exists = strings.Contains(line, "two")
	if exists {
		foundWords["two"] = true
	}
	exists = strings.Contains(line, "three")
	if exists {
		foundWords["three"] = true
	}
	exists = strings.Contains(line, "four")
	if exists {
		foundWords["four"] = true
	}
	exists = strings.Contains(line, "five")
	if exists {
		foundWords["five"] = true
	}
	exists = strings.Contains(line, "six")
	if exists {
		foundWords["six"] = true
	}
	exists = strings.Contains(line, "seven")
	if exists {
		foundWords["seven"] = true
	}
	exists = strings.Contains(line, "eight")
	if exists {
		foundWords["eight"] = true
	}
	exists = strings.Contains(line, "nine")
	if exists {
		foundWords["nine"] = true
	}
	exists = strings.Contains(line, "zero")
	if exists {
		foundWords["zero"] = true
	}

	return foundWords
}

func findFirstAndLastWord(wordSlice []wordPosition) (firstWord wordPosition, lastWord wordPosition) {
	var firstIteration = true

	firstWord = wordPosition{
		Word:  "",
		Index: -1,
	}

	lastWord = wordPosition{
		Word:  "",
		Index: -1,
	}

	for _, word := range wordSlice {
		if word.Index <= firstWord.Index || firstIteration {
			firstWord = word
			firstIteration = false
		}

		if word.Index >= lastWord.Index {
			lastWord = word
		}
	}

	return firstWord, lastWord
}

func wordToInteger(word string) string {
	if "one" == word {
		return "1"
	}
	if "two" == word {
		return "2"
	}
	if "three" == word {
		return "3"
	}
	if "four" == word {
		return "4"
	}
	if "five" == word {
		return "5"
	}
	if "six" == word {
		return "6"
	}
	if "seven" == word {
		return "7"
	}
	if "eight" == word {
		return "8"
	}
	if "nine" == word {
		return "9"
	}

	return "0"
}

func findFirstAndLastNumber(
	firstWord wordPosition,
	lastWord wordPosition,
	firstNumber numberPosition,
	lastNumber numberPosition,
) []string {
	var realFirstNumber, realLastNumber string

	if firstWord.Word == "" {
		realFirstNumber = firstNumber.Number
	}

	if lastWord.Word == "" {
		realLastNumber = lastNumber.Number
	}

	if realFirstNumber != "" && realLastNumber != "" {
		return []string{realFirstNumber, realLastNumber}
	}

	if firstWord.Index < firstNumber.Index {
		realFirstNumber = wordToInteger(firstWord.Word)
	} else {
		realFirstNumber = firstNumber.Number
	}

	if lastWord.Index > lastNumber.Index {
		realLastNumber = wordToInteger(lastWord.Word)
	} else {
		realLastNumber = lastNumber.Number
	}

	return []string{realFirstNumber, realLastNumber}
}
