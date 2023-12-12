package main

// I really, really don't like puzzles focused on matrix traversal...

import (
	"brlywk/AoC/helper"
	"fmt"
	"strconv"
	"strings"
)

const InputFile = "input.txt"

// const InputFile = "input_test.txt"

// ---- Main ------------------------------------

func main() {
	fmt.Println("Advent of Code 2023 - Day 12")
	data, _ := aochelper.NewInputData(InputFile, true)
	lines := data.GetLines()

	part1 := EvaluatePart1(&lines)
	fmt.Printf("Part 1: %v\n", part1)

	// We could cache already checked combinations for a line...
	// or just brute force it...
	unfoldedLines := UnfoldPatterns(&lines)
	part2 := EvaluatePart1(unfoldedLines)
	fmt.Printf("Part 2: %v\n", part2)
}

// ---- Helper ----------------------------------

// Recursively go through the line and count how many combinations can be found
func CountCombinations(pattern string, nums []int) int {
	if pattern == "" {
		if len(nums) == 0 {
			return 1
		}

		return 0
	}

	if len(nums) == 0 {
		if strings.Contains(pattern, "#") {
			return 0
		}

		return 1
	}

	count := 0

	if pattern[0] != '#' {
		count += CountCombinations(pattern[1:], nums)
	}

	if pattern[0] != '.' {
		if nums[0] <= len(pattern) && !strings.Contains(pattern[:nums[0]], ".") && (nums[0] == len(pattern) || pattern[nums[0]] != '#') {
			if nums[0] == len(pattern) {
				count += CountCombinations("", nums[1:])
			} else {
				count += CountCombinations(pattern[nums[0]+1:], nums[1:])
			}
		}

	}

	return count
}

// ---- Part 1 ----------------------------------

func EvaluatePart1(lines *[]string) int {
	sum := 0

	for _, line := range *lines {
		split := strings.Fields(line)
		pattern := split[0]
		numStr := strings.Split(split[1], ",")
		nums := aochelper.MapSlice(numStr, func(str string) int {
			n, _ := strconv.Atoi(str)
			return n
		})

		sum += CountCombinations(pattern, nums)
	}

	return sum
}

// ---- Part 2 ----------------------------------

func UnfoldPatterns(lines *[]string) *[]string {
	unfoldedLines := []string{}

	for _, line := range *lines {
		var newLine strings.Builder
		split := strings.Fields(line)

		newPattern := strings.Repeat(split[0]+"?", 5)
		newLine.WriteString(newPattern[:len(newPattern)-1])
		newLine.WriteString(" ")
		newNums := strings.Repeat(split[1]+",", 5)
		newLine.WriteString(newNums[:len(newNums)-1])

		unfoldedLines = append(unfoldedLines, newLine.String())
	}

	return &unfoldedLines
}
