package day10

import (
	"fmt"
	"github.com/user3690/advent-of-code-in-go/util"
	"log"
)

const (
	vertical      = '|'
	horizontal    = '-'
	northEastBend = 'L'
	northWestBend = 'J'
	southEastBend = 'F'
	SouthWestBend = '7'
	ground        = '.'
	startingPos   = 'S'
)

type position struct {
	row, col uint16
}

type tile struct {
	pipeType   rune
	pos        position
	discovered bool
}

type pipeMap = map[position]tile

func BothParts() {
	var (
		lines        []string
		playingField pipeMap
		startTile    tile
		steps        uint64
		err          error
	)

	lines, err = util.ReadFileInLines("./2023/day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	playingField, startTile = prepareData(lines)

	steps, err = walkLoop(playingField, startTile)
	if err != nil {
		log.Fatal(err)
	}

	insideArea := countInsideTiles(playingField, len(lines), len(lines[0]))
	fmt.Printf(
		"total steps in loop: %d, farthest point: %d, inside area is %d tiles large\n",
		steps,
		steps/2,
		insideArea,
	)

	return
}

func prepareData(lines []string) (pipeMap, tile) {
	var (
		newPipeMap  = make(pipeMap)
		startTile   tile
		newPosition position
	)

	for row, line := range lines {
		for col, letter := range line {
			newPosition = position{col: uint16(col), row: uint16(row)}
			newPipeMap[newPosition] = tile{
				pipeType: letter,
				pos:      newPosition,
			}

			if letter == startingPos {
				startTile = tile{
					pipeType:   letter,
					pos:        newPosition,
					discovered: true,
				}

				newPipeMap[startTile.pos] = startTile
			}
		}
	}

	return newPipeMap, startTile
}

func walkLoop(playingField pipeMap, startTile tile) (uint64, error) {
	var (
		steps             uint64
		nextTile, curTile tile
		err               error
	)

	neighbors := findWalkableNeighborsForStartPos(playingField, startTile)
	if len(neighbors) != 2 {
		return 0, fmt.Errorf("there should exactly be two walkable neighbor tiles: %v", neighbors)
	}

	nextTile = neighbors[0]
	steps++

	for {
		nextTile.discovered = true
		playingField[nextTile.pos] = nextTile
		steps++

		curTile = nextTile
		neighbors = findWalkableNeighbor(playingField, curTile)

		if !neighbors[0].discovered {
			nextTile = neighbors[0]

			continue
		}

		if !neighbors[1].discovered {
			nextTile = neighbors[1]

			continue
		}

		for _, neighbor := range neighbors {
			if neighbor.pipeType == startingPos {
				break
			}
		}

		break
	}

	return steps, err
}

func findWalkableNeighbor(playingField pipeMap, curTile tile) []tile {
	var neighbors []tile

	switch curTile.pipeType {
	case horizontal:
		rightWalkableTile := playingField[position{col: curTile.pos.col + 1, row: curTile.pos.row}]
		leftWalkableTile := playingField[position{col: curTile.pos.col - 1, row: curTile.pos.row}]
		neighbors = append(neighbors, rightWalkableTile, leftWalkableTile)
	case vertical:
		upperWalkableTile := playingField[position{col: curTile.pos.col, row: curTile.pos.row - 1}]
		lowerWalkableTile := playingField[position{col: curTile.pos.col, row: curTile.pos.row + 1}]
		neighbors = append(neighbors, upperWalkableTile, lowerWalkableTile)
	case northEastBend:
		upperWalkableTile := playingField[position{col: curTile.pos.col, row: curTile.pos.row - 1}]
		rightWalkableTile := playingField[position{col: curTile.pos.col + 1, row: curTile.pos.row}]
		neighbors = append(neighbors, upperWalkableTile, rightWalkableTile)
	case northWestBend:
		upperWalkableTile := playingField[position{col: curTile.pos.col, row: curTile.pos.row - 1}]
		leftWalkableTile := playingField[position{col: curTile.pos.col - 1, row: curTile.pos.row}]
		neighbors = append(neighbors, upperWalkableTile, leftWalkableTile)
	case southEastBend:
		lowerWalkableTile := playingField[position{col: curTile.pos.col, row: curTile.pos.row + 1}]
		rightWalkableTile := playingField[position{col: curTile.pos.col + 1, row: curTile.pos.row}]
		neighbors = append(neighbors, lowerWalkableTile, rightWalkableTile)
	case SouthWestBend:
		lowerWalkableTile := playingField[position{col: curTile.pos.col, row: curTile.pos.row + 1}]
		leftWalkableTile := playingField[position{col: curTile.pos.col - 1, row: curTile.pos.row}]
		neighbors = append(neighbors, lowerWalkableTile, leftWalkableTile)
	}

	return neighbors
}

func findWalkableNeighborsForStartPos(playingField pipeMap, startTile tile) []tile {
	var neighbors []tile

	// get upper tile
	walkableTile := playingField[position{col: startTile.pos.col, row: startTile.pos.row - 1}]
	if walkableTile.pipeType == vertical ||
		walkableTile.pipeType == SouthWestBend ||
		walkableTile.pipeType == southEastBend {
		neighbors = append(neighbors, walkableTile)
	}

	// get right tile
	walkableTile = playingField[position{col: startTile.pos.col + 1, row: startTile.pos.row}]
	if walkableTile.pipeType == horizontal ||
		walkableTile.pipeType == northWestBend ||
		walkableTile.pipeType == SouthWestBend {
		neighbors = append(neighbors, walkableTile)
	}

	// get lower tile
	walkableTile = playingField[position{col: startTile.pos.col, row: startTile.pos.row + 1}]
	if walkableTile.pipeType == vertical ||
		walkableTile.pipeType == northEastBend ||
		walkableTile.pipeType == northWestBend {
		neighbors = append(neighbors, walkableTile)
	}

	// get left tile
	walkableTile = playingField[position{col: startTile.pos.col - 1, row: startTile.pos.row}]
	if walkableTile.pipeType == horizontal ||
		walkableTile.pipeType == northEastBend ||
		walkableTile.pipeType == southEastBend {
		neighbors = append(neighbors, walkableTile)
	}

	return neighbors
}

// countInsideTiles for better visuals, I convert the letters '|-LJ7F' => '│─└┘┐┌'
// at every vertical, northWestBend, northEastBend I assume that I enter an inside area and at the next occurrence
// of these symbols, I assume that I left the inside of the maze
func countInsideTiles(playingField pipeMap, rows int, cols int) uint16 {
	var (
		i, j        int
		insideCount uint16
		newTile     tile
		inside      bool
		convert     = map[rune]rune{
			ground:        ground,
			vertical:      '│',
			horizontal:    horizontal,
			northEastBend: '└',
			northWestBend: '┘',
			SouthWestBend: '┐',
			southEastBend: '┌',
			startingPos:   'X',
		}
	)

	for i < rows {
		j = 0
		inside = false
		for j < cols {
			newTile = playingField[position{row: uint16(i), col: uint16(j)}]
			if newTile.discovered &&
				(newTile.pipeType == vertical ||
					newTile.pipeType == startingPos ||
					newTile.pipeType == northWestBend ||
					newTile.pipeType == northEastBend) {
				inside = !inside
			}

			if (inside && newTile.pipeType == ground && j != cols-1) || (inside && !newTile.discovered) {
				insideCount++
				fmt.Print(string('$'))
				j++
				continue
			}

			if newTile.discovered {
				fmt.Print(string(convert[newTile.pipeType]))
			} else {
				fmt.Print(string(ground))
			}

			j++
		}

		fmt.Print("\n")
		i++
	}

	return insideCount
}
