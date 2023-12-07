package day07

import (
	"fmt"
	"github.com/user3690/advent-of-code-in-go/util"
	"log"
	"sort"
	"strconv"
	"strings"
)

type handType = uint8

const (
	highCard handType = iota
	onePair
	twoPair
	threeOfKind
	fullHouse
	fourOfKind
	fiveOfKind
)

var cardValues = map[rune]uint8{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

var cardValuesWithActiveJoker = map[rune]uint8{
	'J': 0,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

type handInfo struct {
	hand     string
	handType handType
	rank     uint16
	bid      uint16
}

func Part1() uint64 {
	var (
		hands []handInfo
	)

	lines, err := util.ReadFileInLines("./2023/day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	hands, err = prepareData(lines, false)
	if err != nil {
		log.Fatal(err)
	}

	return calculateTotalWinnings(hands, cardValues)
}

func prepareData(lines []string, withJoker bool) ([]handInfo, error) {
	var (
		hands       []handInfo
		newHandInfo handInfo
		number      uint64
		err         error
	)

	hands = make([]handInfo, len(lines))
	for i, line := range lines {
		splitLine := strings.Split(line, " ")
		number, err = strconv.ParseUint(splitLine[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error while parsing number: %w", err)
		}

		if withJoker {
			newHandInfo = handInfo{
				hand:     splitLine[0],
				handType: determineHandTypeWithJoker(splitLine[0]),
				rank:     1,
				bid:      uint16(number),
			}
		} else {
			newHandInfo = handInfo{
				hand:     splitLine[0],
				handType: determineHandType(splitLine[0]),
				rank:     1,
				bid:      uint16(number),
			}
		}

		hands[i] = newHandInfo
	}

	return hands, err
}

func determineHandType(hand string) handType {
	var (
		cardCount = make(map[rune]uint8)
		countList = []uint8{0, 0, 0, 0, 0}
		i         uint8
	)
	for _, letter := range hand {
		cardCount[letter] += 1
	}

	for _, count := range cardCount {
		countList[i] = count
		i++
	}

	sort.SliceStable(countList, func(i, j int) bool {
		return countList[i] > countList[j]
	})

	handTypeCode := fmt.Sprintf(
		"%d%d%d%d%d",
		countList[0],
		countList[1],
		countList[2],
		countList[3],
		countList[4],
	)

	switch handTypeCode {
	case "11111":
		return highCard
	case "21110":
		return onePair
	case "22100":
		return twoPair
	case "31100":
		return threeOfKind
	case "32000":
		return fullHouse
	case "41000":
		return fourOfKind
	case "50000":
		return fiveOfKind
	default:
		return 0
	}
}

func calculateTotalWinnings(hands []handInfo, currentCardValues map[rune]uint8) uint64 {
	var (
		matchPlayed   = make(map[string]bool)
		winner        handInfo
		totalWinnings uint64
	)

	for i, hand1 := range hands {
		for j, hand2 := range hands {
			if hand1 == hand2 || matchPlayed[hand1.hand+hand2.hand] || matchPlayed[hand2.hand+hand1.hand] {
				continue
			}

			if hand1.handType > hand2.handType {
				hand1.rank += 1
				hands[i] = hand1
			} else if hand1.handType < hand2.handType {
				hand2.rank += 1
				hands[j] = hand2
			} else {
				winner, _ = determineWinnerByCardValue(hand1, hand2, currentCardValues)

				if winner == hand1 {
					hand1.rank += 1
					hands[i] = hand1
				} else {
					hand2.rank += 1
					hands[j] = hand2
				}
			}

			matchPlayed[hand1.hand+hand2.hand] = true
		}
	}

	for _, hand := range hands {
		totalWinnings += uint64(hand.rank) * uint64(hand.bid)
	}

	return totalWinnings
}

func determineWinnerByCardValue(
	hand1 handInfo,
	hand2 handInfo,
	currentCardValues map[rune]uint8,
) (handInfo, handInfo) {
	var i int
	for i < len(hand1.hand) {
		char1 := hand1.hand[i]
		char2 := hand2.hand[i]

		if currentCardValues[rune(char1)] > currentCardValues[rune(char2)] {
			return hand1, hand2
		}

		if currentCardValues[rune(char1)] < currentCardValues[rune(char2)] {
			return hand2, hand1
		}

		i++
	}

	// should throw error here
	return hand1, hand2
}

func Part2() uint64 {
	var (
		hands []handInfo
	)

	lines, err := util.ReadFileInLines("./2023/day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	hands, err = prepareData(lines, true)
	if err != nil {
		log.Fatal(err)
	}

	return calculateTotalWinnings(hands, cardValuesWithActiveJoker)
}

func determineHandTypeWithJoker(hand string) handType {
	var (
		cardCount = make(map[rune]uint8)
		countList = []uint8{0, 0, 0, 0, 0}
		i         uint8
	)

	for _, card := range hand {
		cardCount[card] += 1
	}

	if _, exists := cardCount['J']; exists {
		var highestCount uint8
		var highestValueCard rune
		for card, count := range cardCount {
			if card == 'J' {
				continue
			}

			if highestCount < count {
				highestCount = count
				highestValueCard = card
			} else if highestCount == count && cardValuesWithActiveJoker[highestValueCard] < cardValuesWithActiveJoker[card] {
				highestValueCard = card
			}
		}

		newHand := strings.Replace(hand, "J", string(highestValueCard), -1)
		cardCount = make(map[rune]uint8)
		for _, card := range newHand {
			cardCount[card] += 1
		}
	}

	for _, count := range cardCount {
		countList[i] = count
		i++
	}

	sort.SliceStable(countList, func(i, j int) bool {
		return countList[i] > countList[j]
	})

	handTypeCode := fmt.Sprintf(
		"%d%d%d%d%d",
		countList[0],
		countList[1],
		countList[2],
		countList[3],
		countList[4],
	)

	switch handTypeCode {
	case "11111":
		return highCard
	case "21110":
		return onePair
	case "22100":
		return twoPair
	case "31100":
		return threeOfKind
	case "32000":
		return fullHouse
	case "41000":
		return fourOfKind
	case "50000":
		return fiveOfKind
	default:
		return 0
	}
}
