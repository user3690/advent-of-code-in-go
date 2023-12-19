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

type direction uint8

const (
	up direction = iota
	right
	down
	left
)

func (dir direction) TurnLeftOrRight(turn direction) (direction, error) {
	if turn != left && turn != right {
		return 0, errors.New("only left or right allowed")
	}

	switch dir {
	case up:
		return turn, nil
	case right:
		switch turn {
		case left:
			return up, nil
		case right:
			return down, nil
		}
	case down:
		switch turn {
		case left:
			return right, nil
		case right:
			return left, nil
		}
	case left:
		switch turn {
		case left:
			return down, nil
		case right:
			return up, nil
		}
	}

	return 0, errors.New("all cases for turning exhausted")
}

type point struct {
	x int16 // rows
	y int16 // cols
}

type position struct {
	poi point
	dir direction
}

func (pos position) Move(dir direction, steps uint8) (point, error) {
	switch dir {
	case up:
		return point{
			x: pos.poi.x - int16(steps),
			y: pos.poi.y,
		}, nil
	case right:
		return point{
			x: pos.poi.x,
			y: pos.poi.y + int16(steps),
		}, nil
	case down:
		return point{
			x: pos.poi.x + int16(steps),
			y: pos.poi.y,
		}, nil
	case left:
		return point{
			x: pos.poi.x,
			y: pos.poi.y - int16(steps),
		}, nil
	default:
		return point{}, errors.New("unknown movement direction given")
	}
}

type digInstruction struct {
	color  string
	dir    direction
	length uint8
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

func (g *lavaPit) dig(dir direction, length uint8) (position, error) {
	var (
		newPosition position
		newPoint    point
		err         error
	)

	if (g.current.dir == up && dir == down) || (g.current.dir == down && dir == up) {
		return newPosition, errors.New("cant go in opposite direction")
	}

	newPoint, err = g.current.Move(dir, length)
	if err != nil {
		return newPosition, fmt.Errorf("could not move in lavaPit: %w", err)
	}

	newPosition.poi = newPoint
	newPosition.dir = dir

	switch dir {
	case up:
		var i int16 = 0
		for i <= int16(length) {
			g.trench[g.current.poi.x-i][g.current.poi.y] = true
			i++
		}
	case right:
		var i int16 = 0
		for i <= int16(length) {
			g.trench[g.current.poi.x][g.current.poi.y+i] = true
			i++
		}
	case down:
		var i int16 = 0
		for i <= int16(length) {
			g.trench[g.current.poi.x+i][g.current.poi.y] = true
			i++
		}
	case left:
		var i int16 = 0
		for i <= int16(length) {
			g.trench[g.current.poi.x][g.current.poi.y-i] = true
			i++
		}
	}

	return newPosition, nil
}

func Part1() {
	var (
		lines        []string
		instructions []digInstruction
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

	_, err = digLoop(instructions)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("execution took %s", time.Since(start))
}

func prepareDigPlan(lines []string) (digInstructions []digInstruction, err error) {
	var (
		dir    direction
		number int
	)

	digInstructions = make([]digInstruction, len(lines))

	for i, line := range lines {
		splitLine := strings.Split(line, " ")
		dir, err = stringToDirection(splitLine[0])
		if err != nil {
			return digInstructions, err
		}

		number, err = strconv.Atoi(splitLine[1])
		if err != nil {
			return digInstructions, err
		}

		trimmed := strings.Trim(splitLine[2], "()")

		newInstruction := digInstruction{
			dir:    dir,
			length: uint8(number),
			color:  trimmed,
		}

		digInstructions[i] = newInstruction
	}

	return digInstructions, err
}

func stringToDirection(s string) (direction, error) {
	switch s {
	case "U":
		return up, nil
	case "R":
		return right, nil
	case "D":
		return down, nil
	case "L":
		return left, nil
	default:
		return up, errors.New("no possible direction given")
	}
}

func digLoop(instructions []digInstruction) (lavaPit, error) {
	var (
		newPosition position
		err         error
	)

	rows, cols, startingPoint := calculateProportions(instructions)
	fmt.Println(rows, cols, startingPoint)

	newLavaPit := newLavaPit(rows, cols, startingPoint)

	for _, instruction := range instructions {
		newPosition, err = newLavaPit.dig(instruction.dir, instruction.length)
		if err != nil {
			return newLavaPit, err
		}

		newLavaPit.current = newPosition
	}

	err = printLagoon(newLavaPit)
	if err != nil {
		return newLavaPit, err
	}

	return newLavaPit, nil
}

func calculateProportions(instructions []digInstruction) (rows int, cols int, startingPoint point) {
	var (
		width, height, minWidth, maxWidth, minHeight, maxHeight int
	)

	// find proportions of lavaPit based of instructions
	for _, instruction := range instructions {
		switch instruction.dir {
		case up:
			height -= int(instruction.length)
			if height < minHeight {
				minHeight = height
			}
		case down:
			height += int(instruction.length)
			if height > maxHeight {
				maxHeight = height
			}
		case left:
			width -= int(instruction.length)
			if width < minWidth {
				minWidth = width
			}
		case right:
			width += int(instruction.length)
			if width > maxWidth {
				maxWidth = width
			}
		}
	}

	rows = (minHeight * -1) + maxHeight
	cols = (minWidth * -1) + maxWidth

	// original start
	startingPoint.x = int16(rows - maxHeight)
	startingPoint.y = int16(cols - maxWidth)

	return rows + 1, cols + 1, startingPoint
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
