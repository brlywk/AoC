package main

import (
	"brlywk/AoC/helper"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const InputFile = "input.txt"

// ---- Constructs ------------------------------

type History struct {
	Raw          []int
	Extras       [][]int
	Extrapolated []int
	Next         int
}

func (h History) String() string {
	return fmt.Sprintf("{\n\tRaw: %v\n\tExtras: %v\nExtrapolated: %v\nNext: %v\n}\n\n", h.Raw, h.Extras, h.Extrapolated, h.Next)
}

// Creates Extras (additional history lines) from the Raw field of the
// calling history;
//
// Does nothing if Raw field is empty or does not exist
func (h *History) CreateExtras() {
	if h.Raw == nil || len(h.Raw) == 0 {
		return
	}

	extras := [][]int{}

	currentExtra := &h.Raw

	for !checkAllZero(currentExtra) {
		tmpExtra := []int{}

		for i := 0; i < len(*currentExtra)-1; i++ {
			diff := (*currentExtra)[i+1] - (*currentExtra)[i]
			tmpExtra = append(tmpExtra, diff)
		}

		extras = append(extras, tmpExtra)
		currentExtra = &tmpExtra
	}

	h.Extras = extras
}

// Create extrapolation and Next expected value in History
func (h *History) CreateExtrapolation() {
	addValues := []int{}

	prev := 0

	for i := len(h.Extras) - 1; i >= 0; i-- {
		curr := h.Extras[i]
		lastValue := curr[len(curr)-1]

		extrapolatedValue := prev + lastValue
		addValues = append(addValues, extrapolatedValue)

		prev = extrapolatedValue
	}

	slices.Reverse(addValues)
	h.Extrapolated = addValues

	h.Next = h.Raw[len(h.Raw)-1] + addValues[0]
}

// Part 2
type PreHistory struct {
	Ref           *History
	Leftrapolated []int
	Prev          int
}

func (p PreHistory) String() string {
	return fmt.Sprintf("{\n\tHistory: %v\n\tLeft: %v\nPrev: %v\n}\n\n", p.Ref, p.Leftrapolated, p.Prev)
}

func (p *PreHistory) CreateLeftrapolation() {
	addValues := []int{}
	extras := p.Ref.Extras

	prev := 0

	for i := len(extras) - 1; i >= 0; i-- {
		curr := extras[i]

		firstValue := curr[0]
		leftrapolatedValue := firstValue - prev
		addValues = append(addValues, leftrapolatedValue)
		prev = leftrapolatedValue
	}

	slices.Reverse(addValues)
	p.Leftrapolated = addValues

	p.Prev = p.Ref.Raw[0] - addValues[0]
}

// ---- Main ------------------------------------

func main() {
	fmt.Println("Advent of Code 2023 - Day 9")

	data, err := aochelper.NewInputData(InputFile, true)
	if err != nil {
		aochelper.PrintAndQuit("Unable to read input file: %v", err)
	}

	lines := data.GetLines()
	histories := ParseInput(&lines)

	part1 := EvaluatePart1(&histories)
	fmt.Printf("Part 1: %v\n", part1)

	preHistories := CreatePrehistoricSlice(&histories)
	part2 := EvaluatePart2(&preHistories)
	fmt.Printf("Part 2: %v\n", part2)
}

// ---- Helper ----------------------------------

// Checks if all values in a slice are 0
func checkAllZero(slice *[]int) bool {
	for _, v := range *slice {
		if v != 0 {
			return false
		}
	}

	return true
}

func ParseInput(lines *[]string) []History {
	result := []History{}

	for _, line := range *lines {
		nhStrSlice := strings.Fields(line)
		nhSlice := aochelper.MapSlice(nhStrSlice, func(e string) int {
			n, _ := strconv.Atoi(e)
			return n
		})

		nh := History{Raw: nhSlice}
		nh.CreateExtras()
		nh.CreateExtrapolation()
		result = append(result, nh)
	}

	return result
}

// ---- Part 1 ----------------------------------

func EvaluatePart1(data *[]History) int {
	sum := 0

	for _, h := range *data {
		sum += h.Next
	}

	return sum
}

// ---- Part 2 ----------------------------------

func CreatePrehistoricSlice(data *[]History) []PreHistory {
	result := []PreHistory{}

	for _, h := range *data {
		pre := PreHistory{Ref: &h}
		pre.CreateLeftrapolation()

		result = append(result, pre)
	}

	return result
}

func EvaluatePart2(preData *[]PreHistory) int {
	sum := 0

	for _, p := range *preData {
		sum += p.Prev
	}

	return sum
}
