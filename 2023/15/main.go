package main

import (
	aochelper "brlywk/AoC/helper"
	"fmt"
	"os"
	"slices"
	"strconv"
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

	part2 := EvaluatePart2(steps)
	fmt.Printf("Part 2: %v\n", part2)
}

// ---- Structs etc. ----------------------------

type Step struct {
	Raw         string
	RawHash     int
	Label       string
	LabelHash   int
	Operation   string
	FocalLength int
	Power       int
}

func (s Step) String() string {
	return fmt.Sprintf("{ Raw: %v | RawHash: %v | Label: %v | Hash: %v | Op: %v | FocalLength: %v } => %v",
		s.Raw, s.RawHash, s.Label, s.LabelHash, s.Operation, s.FocalLength, s.Power)
}

func (s *Step) Split() {
	l := len(s.Raw) - 1
	runes := []rune(s.Raw)

	if strings.Index(s.Raw, "=") > -1 {
		s.Operation = "="
		s.Label = string(runes[:l-1])
		fl, _ := strconv.Atoi(string(runes[l:]))
		s.FocalLength = fl

	} else {
		s.Label = string(runes[:l])
		s.Operation = "-"
	}
}

// ---- Helper ----------------------------------

func Hash(str string) int {
	currentValue := 0

	for _, char := range str {
		currentValue += int(char)
		currentValue *= 17
		currentValue %= 256
	}

	return currentValue
}

func ParseInput(fileName string) *[]Step {
	var result []Step

	byteContent, _ := os.ReadFile(fileName)
	content := string(byteContent)

	content = strings.ReplaceAll(content, "\n", "")
	steps := strings.Split(content, ",")

	for _, s := range steps {
		newStep := Step{Raw: s}
		newStep.RawHash = Hash(s)
		newStep.Split()
		newStep.LabelHash = Hash(newStep.Label)
		result = append(result, newStep)
	}

	return &result
}

// ---- Part 1 ----------------------------------

func EvaluatePart1(steps *[]Step) int {
    defer aochelper.Measure("EvaluatePart1")()
	sum := 0

	for _, step := range *steps {
		sum += step.RawHash
	}

	return sum
}

// ---- Part 2 ----------------------------------

func IndexOfStep(slice *[]Step, step Step) int {
	idx := -1

	for i, s := range *slice {
		if s.Label == step.Label {
			return i
		}
	}

	return idx
}

func SortIntoBoxes(steps *[]Step) [256][]Step {
	var result [256][]Step

	// fill array with empty slices
	for i := range result {
		result[i] = []Step{}
	}

	for _, step := range *steps {
		boxNr := step.LabelHash
		idx := IndexOfStep(&result[boxNr], step)

		// =    add to some box
		if step.Operation == "=" {
			// if found replace, otherwise add
			if idx != -1 {
				slices.Replace(result[boxNr], idx, idx+1, step)
			} else {
				result[boxNr] = append(result[boxNr], step)
			}
		}

		// -    remove from box
		if step.Operation == "-" {
			if idx != -1 {
				result[boxNr] = slices.Delete(result[boxNr], idx, idx+1)
			}
		}
	}

	return result
}

func AssignPower(boxes *[256][]Step) {
	for i := range *boxes {
		if len(boxes[i]) == 0 {
			continue
		}

		for stepIdx, step := range boxes[i] {
            stepCopy := step
			stepCopy.Power = (i + 1) * (stepIdx + 1) * step.FocalLength

            slices.Replace(boxes[i], stepIdx, stepIdx + 1, stepCopy)
		}
	}
}

func EvaluatePart2(steps *[]Step) int {
    defer aochelper.Measure("EvaluatePart2")()
	sum := 0

	boxes := SortIntoBoxes(steps)
    AssignPower(&boxes)

	for _, box := range boxes {
		if len(box) > 0 {
			// fmt.Printf("[\t%v\t]\tSteps:\n", i)
			for _, step := range box {
				// fmt.Printf("\t\t\t%v\n", step)
                sum += step.Power
			}
		}
	}

	return sum
}
