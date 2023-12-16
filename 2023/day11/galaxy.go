package day11

import (
	"fmt"
	"github.com/user3690/advent-of-code-in-go/util"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func BothParts() {
	var (
		minDistanceSum                  int
		lines                           []string
		galaxyPhoto                     [][]int
		emptyRows, emptyCols, galaxyIds []int
		galaxyMap                       = make(map[int][]int)
		err                             error
	)

	lines, err = util.ReadFileInLines("./2023/day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	galaxyPhoto = generateGalaxyPhoto(lines)
	emptyRows = findEmptyRows(galaxyPhoto)
	emptyCols = findEmptyRows(transpose(galaxyPhoto))

	for y, line := range galaxyPhoto {
		for x, cell := range line {
			if cell > 0 {
				galaxyIds = append(galaxyIds, cell)
				galaxyMap[cell] = []int{x, y}
			}
		}
	}

	minDistanceSum = calculateMinDistanceSum(galaxyMap, galaxyIds, emptyRows, emptyCols, 2)
	fmt.Println("Part 1: ", minDistanceSum)

	minDistanceSum = calculateMinDistanceSum(galaxyMap, galaxyIds, emptyRows, emptyCols, 1_000_000)
	fmt.Println("Part 2: ", minDistanceSum)
}

func generateGalaxyPhoto(lines []string) [][]int {
	galaxyPhoto := make([][]int, len(lines))
	id := 1

	for i, line := range lines {
		galaxyPhoto[i] = make([]int, len(line))
		for j, ch := range line {
			if ch == '#' {
				galaxyPhoto[i][j] = id
				id++
			}
		}
	}

	return galaxyPhoto
}

// mirror photo
func transpose(galaxyPhoto [][]int) [][]int {
	result := make([][]int, len(galaxyPhoto[0]))

	for i := 0; i < len(galaxyPhoto[0]); i++ {
		result[i] = make([]int, len(galaxyPhoto))
		for j := 0; j < len(galaxyPhoto); j++ {
			result[i][j] = galaxyPhoto[j][i]
		}
	}

	return result
}

func findEmptyRows(galaxyPhoto [][]int) []int {
	var emptyRows []int

	for i, line := range galaxyPhoto {
		diff := 0
		for _, galaxyId := range line {
			diff += galaxyId
		}

		if diff == 0 {
			emptyRows = append(emptyRows, i)
		}
	}

	return emptyRows
}

func calculateMinDistanceSum(
	galaxyMap map[int][]int,
	galaxyIds []int,
	emptyRows []int,
	emptyCols []int,
	expansion int,
) int {
	var minDistanceSum int

	for i := 0; i < len(galaxyIds); i++ {
		for j := i + 1; j < len(galaxyIds); j++ {
			galaxy1 := galaxyIds[i]
			galaxy2 := galaxyIds[j]
			g1Coords := []int{galaxyMap[galaxy1][0], galaxyMap[galaxy1][1]}
			g2Coords := []int{galaxyMap[galaxy2][0], galaxyMap[galaxy2][1]}

			galaxy1AddedExpansion := make([]int, 2)
			galaxy2AddedExpansion := make([]int, 2)
			for _, emptyCol := range emptyCols {
				if emptyCol < g1Coords[0] {
					galaxy1AddedExpansion[0] += expansion - 1
				}

				if emptyCol < g2Coords[0] {
					galaxy2AddedExpansion[0] += expansion - 1
				}
			}

			for _, emptyRow := range emptyRows {
				if emptyRow < g1Coords[1] {
					galaxy1AddedExpansion[1] += expansion - 1
				}

				if emptyRow < g2Coords[1] {
					galaxy2AddedExpansion[1] += expansion - 1
				}
			}

			g1Coords[0] += galaxy1AddedExpansion[0]
			g1Coords[1] += galaxy1AddedExpansion[1]
			g2Coords[0] += galaxy2AddedExpansion[0]
			g2Coords[1] += galaxy2AddedExpansion[1]

			x := math.Abs(float64(g2Coords[0] - g1Coords[0]))
			y := math.Abs(float64(g2Coords[1] - g1Coords[1]))
			minDistance := int(x + y)

			minDistanceSum += minDistance
		}
	}

	return minDistanceSum
}

// old code

type position struct {
	row int
	col int
}

type galaxyPair struct {
	pos1 position
	pos2 position
}

func expandPhoto(lines []string) ([]string, error) {
	var (
		colsOfPhoto          = len(lines[0])
		emptyRows, emptyCols []int
		file                 *os.File
		err                  error
	)

	emptyRows = oldFindEmptyRows(lines)
	fmt.Println(emptyRows)

	emptyCols = findEmptyCols(lines)
	fmt.Println(emptyCols)

	var newLines []string
	var newString string
	for i, row := range lines {
		if slices.Contains(emptyRows, i) {
			newString = ""
			var x int
			for x < len(emptyCols)+colsOfPhoto {
				newString += string('.')
				x++
			}
			newString += string('\n')
			newLines = append(newLines, newString)
		}

		newString = ""
		for j, col := range row {
			if slices.Contains(emptyCols, j) {
				newString += string('.')
			}

			newString += string(col)
		}

		newString += string('\n')
		newLines = append(newLines, newString)
	}

	file, err = os.Create("./2023/day11/test.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}

	for _, line := range newLines {
		_, err = fmt.Fprint(file, line)
		if err != nil {
			return nil, fmt.Errorf("failed to write to file: %w", err)
		}
	}

	err = file.Close()
	if err != nil {
		if err != nil {
			return nil, fmt.Errorf("failed to close file: %w", err)
		}
	}

	return newLines, nil
}

func oldFindEmptyRows(lines []string) (emptyRows []int) {
	var (
		i          int
		isEmptyCol bool
	)

	for _, rows := range lines {
		isEmptyCol = true
		for _, cols := range rows {
			if cols == '#' {
				isEmptyCol = false
			}
		}

		i++
		if isEmptyCol {
			emptyRows = append(emptyRows, i)
		}
	}

	return emptyRows
}

func findEmptyCols(lines []string) (emptyCols []int) {
	var (
		i, j       int
		isEmptyRow bool
	)

	for j < len(lines[0]) {
		isEmptyRow = true
		i = 0
		for i < len(lines) {
			letter := lines[i][j]

			if letter == '#' {
				isEmptyRow = false
			}

			i++
		}

		if isEmptyRow {
			emptyCols = append(emptyCols, j)
		}

		j++
	}

	return emptyCols
}

func findGalaxies(lines []string) []position {
	var (
		line       string
		xIdx, yIdx int
		galaxies   []position
	)

	for yIdx, line = range lines {
		xIdx = 0
		for {
			xIdx = strings.Index(line, "#")
			line = strings.Replace(line, "#", ".", 1)

			if xIdx == -1 {
				break
			}

			newPosition := position{col: yIdx, row: xIdx}
			galaxies = append(galaxies, newPosition)
		}
	}

	return galaxies
}

func createPairs(galaxies []position) map[string]galaxyPair {
	var (
		pairs = make(map[string]galaxyPair)
	)

	for _, position1 := range galaxies {
		for _, position2 := range galaxies {
			if position1 == position2 {
				continue
			}

			x1Coordinates := []int{position1.row, position2.row}
			y1Coordinates := []int{position1.col, position2.col}
			x2Coordinates := []int{position2.row, position1.row}
			y2Coordinates := []int{position2.col, position1.col}
			coordinates1Str := numbersToString(append(x1Coordinates, y1Coordinates...))
			coordinates2Str := numbersToString(append(x2Coordinates, y2Coordinates...))

			_, exists1 := pairs[coordinates1Str]
			_, exists2 := pairs[coordinates2Str]

			if exists1 || exists2 {
				continue
			}

			pairs[coordinates1Str] = galaxyPair{
				pos1: position1,
				pos2: position2,
			}
		}
	}

	return pairs
}

func numbersToString(numbers []int) string {
	var (
		numberStr string
	)

	for _, number := range numbers {
		numberStr += strconv.FormatInt(int64(number), 10)
	}

	return numberStr
}

func calculateMinDistance(pair galaxyPair) int {
	x := math.Abs(float64(pair.pos2.row - pair.pos1.row))
	y := math.Abs(float64(pair.pos2.col - pair.pos1.col))

	return int(x + y)
}
