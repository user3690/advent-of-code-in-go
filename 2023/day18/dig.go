package day18

import (
	"errors"
	"fmt"
	"github.com/user3690/advent-of-code-in-go/util"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x int // rows
	y int // cols
}

type position struct {
	poi point
	dir util.Direction
}

func (pos position) Move(dir util.Direction, steps uint32) (point, error) {
	switch dir {
	case util.Up:
		return point{
			x: pos.poi.x - int(steps),
			y: pos.poi.y,
		}, nil
	case util.Right:
		return point{
			x: pos.poi.x,
			y: pos.poi.y + int(steps),
		}, nil
	case util.Down:
		return point{
			x: pos.poi.x + int(steps),
			y: pos.poi.y,
		}, nil
	case util.Left:
		return point{
			x: pos.poi.x,
			y: pos.poi.y - int(steps),
		}, nil
	default:
		return point{}, errors.New("unknown movement direction given")
	}
}

type digInstruction struct {
	dir    util.Direction
	length uint32
}

type trench [][]bool
type lavaPit struct {
	current position
	trench  trench
}

func newLavaPit(rows int, cols int, startingPoint point) lavaPit {
	trenchMap := make(trench, rows)
	for i := range trenchMap {
		trenchMap[i] = make([]bool, cols)
	}

	trenchMap[startingPoint.x][startingPoint.y] = true

	startingPosition := position{
		poi: startingPoint,
		dir: 0,
	}

	return lavaPit{
		current: startingPosition,
		trench:  trenchMap,
	}
}

func Part1() {
	var (
		lines        []string
		instructions []digInstruction
		volume       uint
		err          error
	)

	start := time.Now()

	lines, err = util.ReadFileInLines("./2023/day18/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	instructions, err = prepareDigPlan(lines)
	if err != nil {
		log.Fatal(err)
	}

	volume, err = calculateAreaPart1(instructions)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("volume of pit is %d", volume)
	log.Printf("execution Part 1 took %s", time.Since(start))
}

func prepareDigPlan(lines []string) (digInstructions []digInstruction, err error) {
	var (
		dir    util.Direction
		length int
	)

	digInstructions = make([]digInstruction, len(lines))

	for i, line := range lines {
		splitLine := strings.Split(line, " ")
		dir, err = stringToDirection(splitLine[0])
		if err != nil {
			return digInstructions, err
		}

		length, err = strconv.Atoi(splitLine[1])
		if err != nil {
			return digInstructions, err
		}

		newInstruction := digInstruction{
			dir:    dir,
			length: uint32(length),
		}

		digInstructions[i] = newInstruction
	}

	return digInstructions, err
}

func stringToDirection(s string) (util.Direction, error) {
	switch s {
	case "U":
		return util.Up, nil
	case "R":
		return util.Right, nil
	case "D":
		return util.Down, nil
	case "L":
		return util.Left, nil
	default:
		return util.Up, errors.New("no possible direction given")
	}
}

func calculateAreaPart1(instructions []digInstruction) (volume uint, err error) {
	var (
		emptyLavaPit lavaPit
		insidePoint  point
	)

	rows, cols, startingPoint := calculateProportions(instructions)
	emptyLavaPit = newLavaPit(rows, cols, startingPoint)

	for _, instruction := range instructions {
		emptyLavaPit, err = digTrench(emptyLavaPit, instruction.dir, instruction.length)
		if err != nil {
			return volume, err
		}
	}

	insidePoint, err = findPointInsideTrench(emptyLavaPit)
	if err != nil {
		return volume, err
	}

	emptyLavaPit = digInsideAreaOfLavaPit(emptyLavaPit, insidePoint)

	err = printLagoon(emptyLavaPit)
	if err != nil {
		return volume, err
	}

	return countArea(emptyLavaPit), nil
}

func calculateProportions(instructions []digInstruction) (rows int, cols int, startingPoint point) {
	var (
		width, height, minWidth, maxWidth, minHeight, maxHeight int
	)

	// find proportions of lavaPit based of instructions
	for _, instruction := range instructions {
		switch instruction.dir {
		case util.Up:
			height -= int(instruction.length)
			if height < minHeight {
				minHeight = height
			}
		case util.Down:
			height += int(instruction.length)
			if height > maxHeight {
				maxHeight = height
			}
		case util.Left:
			width -= int(instruction.length)
			if width < minWidth {
				minWidth = width
			}
		case util.Right:
			width += int(instruction.length)
			if width > maxWidth {
				maxWidth = width
			}
		}
	}

	rows = (minHeight * -1) + maxHeight
	cols = (minWidth * -1) + maxWidth

	// original start
	startingPoint.x = rows - maxHeight
	startingPoint.y = cols - maxWidth

	return rows + 1, cols + 1, startingPoint
}

func digTrench(freshLavaPit lavaPit, dir util.Direction, length uint32) (lavaPit, error) {
	var (
		newPosition position
		newPoint    point
		err         error
	)

	if (freshLavaPit.current.dir == util.Up && dir == util.Down) || (freshLavaPit.current.dir == util.Down && dir == util.Up) {
		return freshLavaPit, errors.New("cant go in opposite direction")
	}

	newPoint, err = freshLavaPit.current.Move(dir, length)
	if err != nil {
		return freshLavaPit, fmt.Errorf("could not move in lavaPit: %w", err)
	}

	newPosition.poi = newPoint
	newPosition.dir = dir

	switch dir {
	case util.Up:
		var i = 0
		for i <= int(length) {
			freshLavaPit.trench[freshLavaPit.current.poi.x-i][freshLavaPit.current.poi.y] = true
			i++
		}
	case util.Right:
		var i = 0
		for i <= int(length) {
			freshLavaPit.trench[freshLavaPit.current.poi.x][freshLavaPit.current.poi.y+i] = true
			i++
		}
	case util.Down:
		var i = 0
		for i <= int(length) {
			freshLavaPit.trench[freshLavaPit.current.poi.x+i][freshLavaPit.current.poi.y] = true
			i++
		}
	case util.Left:
		var i = 0
		for i <= int(length) {
			freshLavaPit.trench[freshLavaPit.current.poi.x][freshLavaPit.current.poi.y-i] = true
			i++
		}
	}

	freshLavaPit.current = newPosition

	return freshLavaPit, nil
}

func findPointInsideTrench(trenchedLavaPit lavaPit) (newPoint point, err error) {
	for i, row := range trenchedLavaPit.trench {
		for j, isTrench := range row {
			if j > 0 && j < len(row)-1 && !row[j-1] && isTrench && !row[j+1] {
				newPoint = point{x: i, y: j + 1}

				return newPoint, nil
			}
		}
	}

	return newPoint, errors.New("no inside point found")
}

func digInsideAreaOfLavaPit(trenchedLavaPit lavaPit, insidePoint point) lavaPit {
	queue := util.FiFoQueue[point]{}
	diggedMap := make(map[point]struct{})

	// dig trench in inside point
	trenchedLavaPit.trench[insidePoint.x][insidePoint.y] = true
	for _, neighbour := range findNoTrenchNeighbours(trenchedLavaPit, insidePoint) {
		queue.Push(neighbour)
	}

	for !queue.IsEmpty() {
		currentPoint := queue.Pop()

		// first we mark as trenched
		trenchedLavaPit.trench[currentPoint.x][currentPoint.y] = true

		// find not trenched neighbours
		for _, neighbour := range findNoTrenchNeighbours(trenchedLavaPit, currentPoint) {
			if _, exists := diggedMap[neighbour]; !exists {
				queue.Push(neighbour)

				diggedMap[neighbour] = struct{}{}
			}
		}
	}

	return trenchedLavaPit
}

func findNoTrenchNeighbours(trenchedLavaPit lavaPit, currentPoint point) (noTrenchNeighbours []point) {
	var isTrench bool

	maxHeight := len(trenchedLavaPit.trench)
	maxWidth := len(trenchedLavaPit.trench[0])

	// upper
	if currentPoint.x-1 >= 0 {
		isTrench = trenchedLavaPit.trench[currentPoint.x-1][currentPoint.y]
		if !isTrench {
			noTrenchNeighbours = append(noTrenchNeighbours, point{x: currentPoint.x - 1, y: currentPoint.y})
		}
	}

	// right
	if int(currentPoint.y+1) < maxWidth {
		isTrench = trenchedLavaPit.trench[currentPoint.x][currentPoint.y+1]
		if !isTrench {
			noTrenchNeighbours = append(noTrenchNeighbours, point{x: currentPoint.x, y: currentPoint.y + 1})
		}
	}

	// lower
	if int(currentPoint.x+1) < maxHeight {
		isTrench = trenchedLavaPit.trench[currentPoint.x+1][currentPoint.y]
		if !isTrench {
			noTrenchNeighbours = append(noTrenchNeighbours, point{x: currentPoint.x + 1, y: currentPoint.y})
		}
	}

	// left
	if currentPoint.y-1 >= 0 {
		isTrench = trenchedLavaPit.trench[currentPoint.x][currentPoint.y-1]
		if !isTrench {
			noTrenchNeighbours = append(noTrenchNeighbours, point{x: currentPoint.x, y: currentPoint.y - 1})
		}
	}

	return noTrenchNeighbours
}

func countArea(trenchedLavaPit lavaPit) (volume uint) {
	for _, row := range trenchedLavaPit.trench {
		for _, isTrenched := range row {
			if isTrenched {
				volume += 1
			}
		}
	}

	return volume
}

func printLagoon(trenchedLavaPit lavaPit) error {
	file, err := os.Create("./2023/day18/lavapit.txt")
	if err != nil {
		return err
	}

	for _, row := range trenchedLavaPit.trench {
		for _, col := range row {
			if col {
				_, err = fmt.Fprint(file, "#")
				if err != nil {
					return err
				}

				continue
			}

			_, err = fmt.Fprint(file, ".")
			if err != nil {
				return err
			}
		}

		_, err = fmt.Fprint(file, "\n")
		if err != nil {
			return err
		}
	}

	file.Close()

	return nil
}

func Part2() {
	var (
		lines        []string
		instructions []digInstruction
		points       []util.Point
		volume       int
		err          error
	)

	start := time.Now()

	lines, err = util.ReadFileInLines("./2023/day18/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	instructions, err = prepareDigPlanPart2(lines)
	if err != nil {
		log.Fatal(err)
	}

	newPoint := util.Point{}
	perimeter := 0
	for _, instruction := range instructions {
		newPoint, err = newPoint.Move(instruction.dir, instruction.length)

		perimeter += int(instruction.length)
		points = append(points, newPoint)
	}

	volume = calculateAreaPart2(perimeter, points)

	log.Printf("volume of pit is %d", volume)
	log.Printf("execution Part 2 took %s", time.Since(start))
}

func calculateAreaPart2(perimeter int, points []util.Point) int {
	n := len(points)
	points = append(points, util.Point{X: points[0].X, Y: points[0].Y})
	points = append(points, util.Point{X: points[1].X, Y: points[1].Y})
	area := 0

	for i := 1; i <= n; i++ {
		area += points[i].Y * (points[i+1].X - points[i-1].X)
	}

	return area/2 + perimeter/2 + 1
}

func prepareDigPlanPart2(lines []string) (digInstructions []digInstruction, err error) {
	var (
		dir    util.Direction
		length int64
	)

	digInstructions = make([]digInstruction, len(lines))

	for i, line := range lines {
		splitLine := strings.Split(line, " ")

		trimmed := strings.Trim(splitLine[2], "#()")
		length, err = strconv.ParseInt(trimmed[0:5], 16, 64)
		if err != nil {
			return digInstructions, err
		}

		switch rune(trimmed[5]) {
		case '0':
			dir = util.Right
		case '1':
			dir = util.Down
		case '2':
			dir = util.Left
		case '3':
			dir = util.Up
		default:
			return nil, errors.New("no valid direction")
		}

		newInstruction := digInstruction{
			dir:    dir,
			length: uint32(length),
		}

		digInstructions[i] = newInstruction
	}

	return digInstructions, err
}
