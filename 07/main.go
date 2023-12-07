package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"brlywk/AoC/helper"
)

const InputFile = "input.txt"

type HandType uint

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var Cards = [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var Cards2 = [...]string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

type Hand struct {
	Raw    string
	RawMap map[string]int
	Bid    int
	Rank   int
	Type   HandType
}

func (h Hand) String() string {
	return fmt.Sprintf("Hand: %v\tMap: %v\tBid: %v\tRank: %v\tType: %v\n", h.Raw, h.RawMap, h.Bid, h.Rank, h.RawMap)
}

func (h *Hand) Compare(otherHand *Hand, part2 bool) int {
	if h.Type != otherHand.Type {
		if h.Type > otherHand.Type {
			return 1
		} else {
			return -1
		}
	}

	slice := Cards[:]
	if part2 {
		slice = Cards2[:]
	}

	for i, c := range (*h).Raw {
		cStr := string(c)
		ocStr := string((*otherHand).Raw[i])

		cIdx := slices.Index(slice, cStr)
		ocIdx := slices.Index(slice, ocStr)

		if cIdx == ocIdx {
			continue
		} else {
			if cIdx > ocIdx {
				return 1
			} else {
				return -1
			}
		}
	}

	return 0
}

// ---- Main ------------------------------------

func main() {
	log.Println("Advent of Code 2023 - Day 7")

	data, err := aochelper.NewInputData(InputFile, true)
	if err != nil {
		log.Fatalf("Unable to read input file: %v", err)
	}
	lines := data.GetLines()

	handsPart1 := ParseInput(&lines, false)
	handsPart1 = PreparePart1(&handsPart1)
	part1 := EvaluatePart1(&handsPart1)
	log.Printf("Part 1: %v\n", part1)

	handsPart2 := ParseInput(&lines, true)
	handsPart2 = PreparePart2(&handsPart2)
	part2 := EvaluatePart2(&handsPart2)
	log.Printf("Part 2: %v\n", part2)
}

// ---- Helper ----------------------------------

// For AoC we can assume that all cards in the Raw string exist...
func GetCardValue(card *string, cards *[13]string) int {
	for i, c := range *cards {
		if c == *card {
			// Let's calculate the actual card value, i.e. 2 = 0 + 2
			return i + 2
		}
	}

	return -1
}

func StringToMap(str *string) map[string]int {
	result := map[string]int{}

	for _, c := range *str {
		result[string(c)]++
	}

	return result
}

func IsFiveOfAKind(m *map[string]int) bool {
	return len(*m) == 1
}

func IsFourOfAKind(m *map[string]int) bool {
	if len(*m) != 2 {
		return false
	}

	hasFour := false
	for _, v := range *m {
		if v == 4 {
			hasFour = true
		}
	}

	return hasFour
}

func IsFullHouse(m *map[string]int) bool {
	if len(*m) != 2 {
		return false
	}

	hasThree := false
	hasTwo := false
	for _, v := range *m {
		if v == 3 {
			hasThree = true
		}

		if v == 2 {
			hasTwo = true
		}
	}

	return hasThree && hasTwo
}

func IsThreeOfAKind(m *map[string]int) bool {
	if len(*m) != 3 {
		return false
	}

	hasThree := false
	for _, v := range *m {
		if v == 3 {
			hasThree = true
		}
	}

	return hasThree
}

func IsTwoPair(m *map[string]int) bool {
	return len(*m) == 3 && !IsThreeOfAKind(m)
}

func IsOnePair(m *map[string]int) bool {
	return len(*m) == 4
}

func FillJokers(rawMap *map[string]int) {
	m := *rawMap
	_, hasJoker := m["J"]
	if !hasJoker {
		return
	}

	jokerCount := m["J"]
	delete(m, "J")

	highestKeyCount := 0
	keyToFill := ""

	for k, v := range m {
		if v > highestKeyCount {
			highestKeyCount = v
			keyToFill = k
		}
	}

	m[keyToFill] += jokerCount
}

func GetHandType(rawMap *map[string]int, part2 bool) HandType {
	if part2 {
		FillJokers(rawMap)
	}

	if IsFiveOfAKind(rawMap) {
		return 6
	}

	if IsFourOfAKind(rawMap) {
		return 5
	}

	if IsFullHouse(rawMap) {
		return 4
	}

	if IsThreeOfAKind(rawMap) {
		return 3
	}

	if IsTwoPair(rawMap) {
		return 2
	}

	if IsOnePair(rawMap) {
		return 1
	}

	return 0
}

func ParseInput(lines *[]string, part2 bool) []Hand {
	var result []Hand

	for _, line := range *lines {
		split := strings.Split(line, " ")
		raw := split[0]
		rawMap := StringToMap(&raw)
		bid, _ := strconv.Atoi(split[1])

		hand := Hand{
			Raw:    raw,
			RawMap: rawMap,
			Bid:    bid,
			Rank:   0,
			Type:   GetHandType(&rawMap, part2),
		}

		result = append(result, hand)
	}

	return result
}

func SortHands(hands *[]Hand, part2 bool) {
	slices.SortStableFunc(*hands, func(h1, h2 Hand) int {
		return h1.Compare(&h2, part2)
	})

	slices.Reverse(*hands)
}

func AssignRank(hands *[]Hand) {
	maxRank := len(*hands)

	for i := range *hands {
		(*hands)[i].Rank = maxRank
		maxRank--
	}
}

// ---- Part 1 ----------------------------------

func PreparePart1(hands *[]Hand) []Hand {
	handCopy := *hands

	SortHands(&handCopy, false)
	AssignRank(&handCopy)

	return handCopy
}

func EvaluatePart1(hands *[]Hand) int {
	sum := 0

	for _, hand := range *hands {
		sum += hand.Bid * hand.Rank
	}

	return sum
}

// ---- Part 2 ----------------------------------

func PreparePart2(hands *[]Hand) []Hand {
	handCopy := *hands

	SortHands(&handCopy, true)
	AssignRank(&handCopy)

	return handCopy
}

func EvaluatePart2(hands *[]Hand) int {
	sum := 0

	for _, hand := range *hands {
		sum += hand.Bid * hand.Rank
	}

	return sum
}
