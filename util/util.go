package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Integer interface {
	int64 | int32 | int16 | int8 | int
}

func ReadFileInLines(filePath string) ([]string, error) {
	var (
		file []byte
		err  error
	)

	file, err = os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error while reading file: %w", err)
	}

	lines := strings.FieldsFunc(string(file), func(r rune) bool {
		return r == '\n'
	})

	return lines, err
}

func StrToInt64(str string) (number int64, err error) {
	number, err = strconv.ParseInt(str, 10, 64)
	if err != nil {
		return number, err
	}

	return number, err
}

func PowInt[T Integer](n T, m T) T {
	var i T

	if m == 0 {
		return 1
	}

	result := n
	for i = 2; i <= m; i++ {
		result *= n
	}

	return result
}

func FindNumbers(numbersStr string) (numbers []int64, err error) {
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

			number, err = StrToInt64(numbersStr[start : end+1])
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
