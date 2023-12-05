package day05

import (
	"log"
	"os"
	"strings"
	"unicode"
)

func Part1() int64 {
	var (
		file []byte
		err  error
	)

	file, err = os.ReadFile("./2023/day05/input_test.txt")
	if err != nil {
		log.Fatalf("error while reading file: %s", err)
	}

	lines := strings.FieldsFunc(string(file), func(r rune) bool {
		return r == '\n'
	})

	prepareData(lines)

	return 0
}

func prepareData(lines []string) {
	for _, line := range lines {
		for _, letter := range line {
			if unicode.IsLetter(letter) {

			}
		}

	}
}
