package day17

import (
	"errors"
	"fmt"
	"github.com/user3690/advent-of-code-in-go/util"
	"log"
	"strconv"
)

type heatLoss = int
type direction uint8

const (
	up direction = iota
	right
	down
	left
)

func (dir direction) TurnLeftOrRight(turn direction) direction {
	if turn != left && turn != right {
		panic("should be left or right")
	}

	switch dir {
	case up:
		return turn
	case right:
		switch turn {
		case left:
			return up
		case right:
			return down
		}
	case down:
		switch turn {
		case left:
			return right
		case right:
			return left
		}
	case left:
		switch turn {
		case left:
			return down
		case right:
			return up
		}
	}

	panic("not handled")
}

type cityBlock struct {
	state    state
	heatLoss heatLoss
}

func (cb cityBlock) move(dir direction) point {
	switch dir {
	case up:
		return point{
			x: cb.state.pos.poi.x - 1,
			y: cb.state.pos.poi.y,
		}
	case right:
		return point{
			x: cb.state.pos.poi.x,
			y: cb.state.pos.poi.y + 1,
		}
	case down:
		return point{
			x: cb.state.pos.poi.x + 1,
			y: cb.state.pos.poi.y,
		}
	case left:
		return point{
			x: cb.state.pos.poi.x,
			y: cb.state.pos.poi.y - 1,
		}
	}

	panic("panic cb turn")
}

func (cb cityBlock) turn(dir direction) direction {
	return cb.state.pos.dir.TurnLeftOrRight(dir)

}

func (cb cityBlock) goStraightOn() direction {
	return cb.state.pos.dir
}

type state struct {
	pos      position
	straight uint8
}

type position struct {
	poi point
	dir direction
}

type point struct {
	x int16 // row
	y int16 // col
}

// part 2 test1 94
// part 2 test2 71
// Part 1 674
// Part 2 773
func BothParts() {
	var (
		lines                     []string
		cityBlocks                map[point]cityBlock
		rows, cols, leastHeatLoss int
		err                       error
	)

	lines, err = util.ReadFileInLines("./2023/day17/input_test2.txt")
	if err != nil {
		log.Fatal(err)
	}

	cityBlocks, rows, cols, err = createCityBlocks(lines)
	if err != nil {
		log.Fatal(err)
	}

	leastHeatLoss, err = findRouteWithLeastHeatLoss(cityBlocks, rows, cols, 0, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(leastHeatLoss)

	leastHeatLoss, err = findRouteWithLeastHeatLoss(cityBlocks, rows, cols, 4, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(leastHeatLoss)
}

func createCityBlocks(
	lines []string,
) (
	cityBlocks map[point]cityBlock,
	rows int,
	cols int,
	err error,
) {
	var number int

	rows = len(lines)
	cols = len(lines[0])
	cityBlocks = make(map[point]cityBlock, rows*cols)

	for i, line := range lines {
		for j, char := range line {
			newPoint := point{
				x: int16(i),
				y: int16(j),
			}

			number, err = strconv.Atoi(string(char))
			if err != nil {
				return nil, 0, 0, err
			}

			cityBlocks[newPoint] = cityBlock{
				state: state{
					pos: position{
						poi: newPoint,
						dir: 0,
					},
					straight: 0,
				},
				heatLoss: number,
			}
		}
	}

	return cityBlocks, rows, cols, err
}

func findRouteWithLeastHeatLoss(
	cityBlocks map[point]cityBlock,
	rows int,
	cols int,
	minStraight, maxStraight uint8,
) (int, error) {
	targetBlock := cityBlocks[point{x: int16(rows - 1), y: int16(cols - 1)}]

	priorityQueue := util.NewPriorityQueue(
		func(itemA, itemB cityBlock) int {
			return itemA.heatLoss - itemB.heatLoss
		},
	)

	// first block below start block
	priorityQueue.Push(cityBlock{
		state: state{
			pos: position{
				poi: point{
					x: 1,
					y: 0,
				},
				dir: down,
			},
			straight: 1,
		},
	})

	// first block to the right of start block
	priorityQueue.Push(cityBlock{
		state: state{
			pos: position{
				poi: point{
					x: 0,
					y: 1,
				},
				dir: right,
			},
			straight: 1,
		},
	})

	visited := make(map[state]heatLoss)

	for !priorityQueue.IsEmpty() {
		item := priorityQueue.Pop()

		curBlockPoint := item.state.pos.poi

		if _, exists := cityBlocks[curBlockPoint]; !exists {
			continue
		}

		totalHeatLoss := cityBlocks[curBlockPoint].heatLoss + item.heatLoss
		if curBlockPoint == targetBlock.state.pos.poi {
			return totalHeatLoss, nil
		}

		if heatLossVal, exists := visited[item.state]; exists {
			if heatLossVal <= totalHeatLoss {
				continue
			}
		}

		visited[item.state] = totalHeatLoss

		if item.state.straight >= minStraight {
			newDirection := item.turn(right)
			rightBlock := cityBlocks[item.move(newDirection)]
			rightBlock.state.pos.dir = newDirection
			rightBlock.state.straight = 1
			rightBlock.heatLoss = totalHeatLoss
			priorityQueue.Push(rightBlock)

			newDirection = item.turn(left)
			leftBlock := cityBlocks[item.move(newDirection)]
			leftBlock.state.pos.dir = newDirection
			leftBlock.state.straight = 1
			leftBlock.heatLoss = totalHeatLoss
			priorityQueue.Push(leftBlock)
		}

		if item.state.straight < maxStraight {
			block := cityBlocks[item.move(item.goStraightOn())]
			block.state.pos.dir = item.goStraightOn()
			block.state.straight = item.state.straight + 1
			block.heatLoss = totalHeatLoss
			priorityQueue.Push(block)
		}
	}

	return 0, errors.New("no path found")
}
