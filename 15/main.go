package main

import (
	"fmt"
	"os"
	"strings"
)

const TestFile = "input_test.txt"
const InputFile = "input.txt"

// ---- Main ------------------------------------

func main() {
	fmt.Println("Advent of Code 2023 - Day 15")
	steps := ParseInput(InputFile)

	part1 := EvaluatePart1(steps)
	fmt.Printf("Part 1: %v\n", part1)
}

// ---- Structs etc. ----------------------------

type Step struct {
	Raw  string
	Hash int
}

func (s Step) String() string {
	return fmt.Sprintf("{ String: %v | Hash: %v }", s.Raw, s.Hash)
}

func (s *Step) CalculateHash() {
	currentValue := 0

	for _, char := range s.Raw {
		currentValue += int(char)
		currentValue *= 17
		currentValue %= 256
	}

	s.Hash = currentValue
}

// ---- Helper ----------------------------------

func ParseInput(fileName string) *[]Step {
	var result []Step

	byteContent, _ := os.ReadFile(fileName)
	content := string(byteContent)

	content = strings.ReplaceAll(content, "\n", "")
	steps := strings.Split(content, ",")

	for _, s := range steps {
		newStep := Step{Raw: s}
		newStep.CalculateHash()
		result = append(result, newStep)
	}

	return &result
}

// ---- Part 1 ----------------------------------

func EvaluatePart1(steps *[]Step) int {
	sum := 0

	for _, step := range *steps {
		sum += step.Hash
	}

	return sum
}
