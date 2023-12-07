package day7

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/milindmadhukar/Advent-Of-Code/2023/golang/utils"
)

func getUniqueCardsFromHand(hand string) []string {
	var uniqueCards []string
	for _, card := range hand {
		if !slices.Contains(uniqueCards, string(card)) {
			uniqueCards = append(uniqueCards, string(card))
		}
	}
	return uniqueCards
}

func getHandType(hand string) int {
	var handType = map[string]int{
		"highCard":     1,
		"onePair":      2,
		"twoPair":      3,
		"threeOfAKind": 4,
		"fullHouse":    5,
		"fourOfAKind":  6,
		"fiveOfAKind":  7,
	}

	uniqueCards := getUniqueCardsFromHand(hand)

	if len(uniqueCards) == 1 {
		return handType["fiveOfAKind"]
	}

	if len(uniqueCards) == 2 {
		for _, card := range uniqueCards {
			if strings.Count(hand, card) == 4 {
				return handType["fourOfAKind"]
			} else if strings.Count(hand, card) == 3 {
				return handType["fullHouse"]
			}
		}
	}

	if len(uniqueCards) == 3 {
		for _, card := range uniqueCards {
			if strings.Count(hand, card) == 3 {
				return handType["threeOfAKind"]
			} else if strings.Count(hand, card) == 2 {
				return handType["twoPair"]
			}
		}
	}

	if len(uniqueCards) == 4 {
		for _, card := range uniqueCards {
			if strings.Count(hand, card) == 2 {
				return handType["onePair"]
			}
		}
	}

	if len(uniqueCards) == 5 {
		return handType["highCard"]
	}

	return -1
}

func compareHands(hand1, hand2 string, cardValues map[string]int) int {
	for idx := 0; idx < len(hand1); idx++ {
		if hand1[idx] != hand2[idx] {
			return cardValues[string(hand1[idx])] - cardValues[string(hand2[idx])]
		}
	}
	return 0
}

func fillJokersValue(hand string) string {
	if !strings.Contains(hand, "J") {
		return hand
	}

	if hand == "JJJJJ" {
		return "AAAAA"
	}

	uniqueCards := getUniqueCardsFromHand(hand)
	maxCount := 0
	maxCard := ""

	for _, card := range uniqueCards {
		if count := strings.Count(hand, card); count > maxCount && card != "J" {
			maxCount = count
			maxCard = card
		}
	}

	return strings.ReplaceAll(hand, "J", maxCard)
}

type day7 struct {
	data      [][]string
	startTime time.Time
}

func (d day7) Part1() any {
	sum := 0

	var cardValues = map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	slices.SortFunc(d.data, func(hand1, hand2 []string) int {
		handType1 := getHandType(hand1[0])
		handType2 := getHandType(hand2[0])
		if handType1 != handType2 {
			return getHandType(hand1[0]) - getHandType(hand2[0])
		} else {
			return compareHands(hand1[0], hand2[0], cardValues)
		}
	})

	for idx, hand := range d.data {
		handBet, _ := strconv.Atoi(hand[1])
		sum += handBet * (idx + 1)
	}

	return sum
}

func (d day7) Part2() any {
	sum := 0

	var cardValues = map[string]int{
		"J": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	slices.SortFunc(d.data, func(hand1, hand2 []string) int {
		hand1WithJokerValues := fillJokersValue(hand1[0])
		hand2WithJokerValues := fillJokersValue(hand2[0])
		handType1 := getHandType(hand1WithJokerValues)
		handType2 := getHandType(hand2WithJokerValues)
		if handType1 != handType2 {
			return handType1 - handType2
		} else {
			return compareHands(hand1[0], hand2[0], cardValues)
		}
	})

	for idx, hand := range d.data {
		fmt.Println(hand)
		handBet, _ := strconv.Atoi(hand[1])
		sum += handBet * (idx + 1)
	}

	return sum
}

func Solve() day7 {
	rawData, err := utils.GetInputDataFromAOC(2023, 7)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	// exampleFile, err := os.ReadFile("day7/example.txt")
	// rawData = utils.ParseFromString(string(exampleFile))

	data := utils.GetSplitData(rawData, " ")

	return day7{
		data:      data,
		startTime: startTime,
	}
}

func (d day7) TimeTaken() time.Duration {
	return time.Since(d.startTime)
}
