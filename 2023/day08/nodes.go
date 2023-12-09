package day08

import (
	"github.com/user3690/advent-of-code-in-go/util"
	"log"
	"strings"
)

type instruction uint8
type node struct {
	name  string
	left  string
	right string
}

const (
	left instruction = iota
	right
)

func BothParts() (uint64, int) {
	var (
		lines        []string
		instructions []instruction
		nodes        map[string]node
		steps        uint64
		stepsAsGhost []int
		err          error
	)

	lines, err = util.ReadFileInLines("./2023/day08/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	instructions, nodes = prepareData(lines)
	if err != nil {
		log.Fatal(err)
	}

	steps = countSteps(instructions, nodes)
	stepsAsGhost = countStepsAsGhost(instructions, nodes)

	overallSteps := util.LCM(stepsAsGhost[0], stepsAsGhost[1], stepsAsGhost...)

	return steps, overallSteps
}

func prepareData(lines []string) ([]instruction, map[string]node) {
	var (
		instructions = make([]instruction, len(lines[0]))
		nodes        = make(map[string]node)
	)

	for i, letter := range lines[0] {
		if letter == 'R' {
			instructions[i] = right
		} else {
			instructions[i] = left
		}
	}

	for _, line := range lines[1:] {
		newNode := node{}
		splitLine := strings.Split(line, "=")
		newNode.name = strings.TrimSpace(splitLine[0])

		replacer := strings.NewReplacer("(", "", ")", "", ",", "")
		nodeNames := replacer.Replace(splitLine[1])

		splitNodeNames := strings.Split(strings.TrimSpace(nodeNames), " ")

		newNode.left = splitNodeNames[0]
		newNode.right = splitNodeNames[1]

		nodes[newNode.name] = newNode
	}

	return instructions, nodes
}

func countSteps(instructions []instruction, nodes map[string]node) (steps uint64) {
	var nextNode, end node

	nextNode = nodes["AAA"]
	end = nodes["ZZZ"]

	for i := 0; i < len(instructions); i++ {
		if nextNode == end {
			break
		}

		if instructions[i] == left {
			nextNode = nodes[nextNode.left]
		} else {
			nextNode = nodes[nextNode.right]
		}

		steps++

		// reset loop
		if i == len(instructions)-1 {
			i = -1
		}
	}

	return steps
}

// countStepsAsGhost the assumption is that the first found ending, is also the correct ending.
// The input file is constructed that way, otherwise this shouldn't work.
func countStepsAsGhost(instructions []instruction, nodes map[string]node) []int {
	var (
		curNodes []node
		allSteps []int
		steps    int
	)

	for _, curNode := range nodes {
		index := strings.LastIndexAny(curNode.name, "A")
		if index == 2 {
			curNodes = append(curNodes, curNode)
		}
	}

	allSteps = make([]int, len(curNodes))

	for i, curNode := range curNodes {
		steps = 0
		for j := 0; j < len(instructions); j++ {
			if curNode.name[2] == 'Z' {
				break
			}

			if instructions[j] == left {
				curNode = nodes[curNode.left]
			} else {
				curNode = nodes[curNode.right]
			}

			steps++

			// reset loop
			if j == len(instructions)-1 {
				j = -1
			}
		}

		allSteps[i] = steps
	}

	return allSteps
}
