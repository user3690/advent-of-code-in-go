package day12

import (
	"fmt"
	"github.com/user3690/advent-of-code-in-go/util"
	"log"
	"strconv"
	"strings"
)

func BothParts() {
	var (
		groups, extendedGroups                            []int
		lines                                             []string
		number                                            int
		arrangementSum, arrangementSumPart2, arrangements uint
		cache                                             = make(map[string]uint)
		err                                               error
	)

	lines, err = util.ReadFileInLines("./2023/day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// part 1
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		splitGroups := strings.Split(splitLine[1], ",")

		groups = make([]int, 0)
		for _, group := range splitGroups {
			number, err = strconv.Atoi(group)
			if err != nil {
				log.Fatal(err)
			}

			groups = append(groups, number)
		}

		fmt.Println("found groups: ", groups)

		arrangements = calculateArrangements(&cache, splitLine[0], groups, 0)

		fmt.Println("found arrangements: ", arrangements)
		arrangementSum += arrangements

	}

	fmt.Println("sum of all arrangements: ", arrangementSum)

	// part 2
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		splitGroups := strings.Split(splitLine[1], ",")

		groups = make([]int, 0)
		for _, group := range splitGroups {
			number, err = strconv.Atoi(group)
			if err != nil {
				log.Fatal(err)
			}

			groups = append(groups, number)
		}

		var i = 0
		extendedGroups = make([]int, 0)
		for i < 5 {
			extendedGroups = append(extendedGroups, groups...)
			i++
		}

		extendedRecord := strings.Join(
			[]string{
				splitLine[0],
				splitLine[0],
				splitLine[0],
				splitLine[0],
				splitLine[0],
			},
			"?",
		)

		arrangements = calculateArrangements(&cache, extendedRecord, extendedGroups, 0)

		arrangementSumPart2 += arrangements
	}

	fmt.Println("sum of all arrangements part 2: ", arrangementSumPart2)
}

func calculateArrangements(cache *map[string]uint, record string, groups []int, curGroupLength int) (arrangements uint) {
	key := buildCacheKey(record, groups, curGroupLength)

	if value, exists := (*cache)[key]; exists {
		return value
	}

	if len(record) == 0 {
		// case we reached the end and found nothing, we can assume we only have 1 valid arrangement
		if len(groups) == 0 && curGroupLength == 0 {
			return 1
		}

		// case we found exactly one valid group at the end of the string
		if len(groups) == 1 && groups[0] == curGroupLength {
			return 1
		}

		return 0
	}

	char := record[0]

	switch char {
	case '#':
		arrangements += calculateArrangements(cache, record[1:], groups, curGroupLength+1)
	case '.':
		if curGroupLength == 0 {
			arrangements += calculateArrangements(cache, record[1:], groups, 0)
		} else if len(groups) > 0 && curGroupLength == groups[0] {
			arrangements += calculateArrangements(cache, record[1:], groups[1:], 0)
		}
	case '?':
		// we assume the ? is a #
		arrangements += calculateArrangements(cache, record[1:], groups, curGroupLength+1)

		// we assume the ? is a .
		if curGroupLength == 0 {
			arrangements += calculateArrangements(cache, record[1:], groups, 0)
		} else if len(groups) > 0 && curGroupLength == groups[0] {
			arrangements += calculateArrangements(cache, record[1:], groups[1:], 0)
		}
	}

	(*cache)[key] = arrangements

	return arrangements
}

func buildCacheKey(record string, groups []int, curGroupLength int) string {
	var key = make([]string, len(groups)+2)

	key[0] = record
	for i, group := range groups {
		key[i+1] = strconv.Itoa(group)
	}
	key[len(key)-1] = strconv.Itoa(curGroupLength)

	return strings.Join(key, ",")
}
