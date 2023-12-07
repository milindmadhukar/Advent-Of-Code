package day7

import (
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

func fillJokersValue(hand string, cardValues map[string]int) string {
	// Joker can become any card
	if !strings.Contains(hand, "J") {
		return hand
	}

	if hand == "JJJJJ" {
		return "AAAAA"
	}

	uniqueCards := getUniqueCardsFromHand(hand)
	if len(uniqueCards) == 2 {
		// Get the card which is not J
		var card string
		for _, c := range uniqueCards {
			if c != "J" {
				card = c
			}
		}
		return strings.ReplaceAll(hand, "J", card)
	}

	if len(uniqueCards) == 3 {
		var card1, card2 string
		for _, c := range uniqueCards {
			if c != "J" {
				if card1 == "" {
					card1 = c
				} else {
					card2 = c
				}
			}
		}

		card1Count := strings.Count(hand, card1)
		card2Count := strings.Count(hand, card2)
		JCount := strings.Count(hand, "J")

		if card1Count == 3 && JCount == 1 {
			return strings.ReplaceAll(hand, "J", card1)
		} else if card2Count == 3 && JCount == 1 {
			return strings.ReplaceAll(hand, "J", card2)
		}

		if card1Count == 2 && JCount == 2 {
			return strings.ReplaceAll(hand, "J", card1)
		} else if card2Count == 2 && JCount == 2 {
			return strings.ReplaceAll(hand, "J", card2)
		}

		if card1Count == card2Count {
			var higherCard string
			if cardValues[card1] > cardValues[card2] {
				higherCard = card1
			} else {
				higherCard = card2
			}
			return strings.ReplaceAll(hand, "J", higherCard)
		}
	}

	//1233J, 123JJ

	if len(uniqueCards) == 4 {
		var card1, card2, card3 string
		for _, c := range uniqueCards {
			if c != "J" {
				if card1 == "" {
					card1 = c
				} else if card2 == "" {
					card2 = c
				} else {
					card3 = c
				}
			}
		}
		if strings.Count(hand, card1) == 2 {
			return strings.ReplaceAll(hand, "J", card1)
		}
		if strings.Count(hand, card2) == 2 {
			return strings.ReplaceAll(hand, "J", card2)
		}
		if strings.Count(hand, card3) == 2 {
			return strings.ReplaceAll(hand, "J", card3)
		}

		var higherCard string
		if cardValues[card1] > cardValues[card2] && cardValues[card1] > cardValues[card3] {
			higherCard = card1
		} else if cardValues[card2] > cardValues[card1] && cardValues[card2] > cardValues[card3] {
			higherCard = card2
		} else {
			higherCard = card3
		}
		return strings.ReplaceAll(hand, "J", higherCard)
	}

	if len(uniqueCards) == 5 {
		var card1, card2, card3, card4 string
		for _, c := range uniqueCards {
			if c != "J" {
				if card1 == "" {
					card1 = c
				} else if card2 == "" {
					card2 = c
				} else if card3 == "" {
					card3 = c
				} else {
					card4 = c
				}
			}
		}

		var higherCard string
		if cardValues[card1] > cardValues[card2] && cardValues[card1] > cardValues[card3] && cardValues[card1] > cardValues[card4] {
			higherCard = card1
		} else if cardValues[card2] > cardValues[card1] && cardValues[card2] > cardValues[card3] && cardValues[card2] > cardValues[card4] {
			higherCard = card2
		} else if cardValues[card3] > cardValues[card1] && cardValues[card3] > cardValues[card2] && cardValues[card3] > cardValues[card4] {
			higherCard = card3
		} else {
			higherCard = card4
		}
		return strings.ReplaceAll(hand, "J", higherCard)
	}

	return hand
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
		hand1WithJokerValues := fillJokersValue(hand1[0], cardValues)
		hand2WithJokerValues := fillJokersValue(hand2[0], cardValues)
		handType1 := getHandType(hand1WithJokerValues)
		handType2 := getHandType(hand2WithJokerValues)
		if handType1 != handType2 {
			return handType1 - handType2
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

func Solve() day7 {
	rawData, err := utils.GetInputDataFromAOC(2023, 7)
	if err != nil {
		panic(err)
	}

	startTime := time.Now()

	// exampleFile, _ := os.ReadFile("day7/example.txt")
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
