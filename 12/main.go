package main

// I really, really don't like puzzles focused on matrix traversal...

import (
	"brlywk/AoC/helper"
	"fmt"
	"strconv"
	"strings"
)

const InputFile = "input.txt"

// ---- Main ------------------------------------

func main() {
	fmt.Println("Advent of Code 2023 - Day 12")
	data, _ := aochelper.NewInputData(InputFile, true)
	lines := data.GetLines()

	part1 := EvaluatePart1(&lines)
	fmt.Printf("Part 1: %v\n", part1)
}

// ---- Helper ----------------------------------

// Recursively go through the line and count how many combinations can be found
func CountCombinations(line string, nums []int) int {
	if line == "" {
		if len(nums) == 0 {
			return 1
		}

		return 0
	}

	if len(nums) == 0 {
		if strings.Contains(line, "#") {
			return 0
		}

		return 1
	}

	count := 0

	if line[0] != '#' {
		count += CountCombinations(line[1:], nums)
	}

	if line[0] != '.' {
		if nums[0] <= len(line) && !strings.Contains(line[:nums[0]], ".") && (nums[0] == len(line) || line[nums[0]] != '#') {
			if nums[0] == len(line) {
				count += CountCombinations("", nums[1:])
			} else {
				count += CountCombinations(line[nums[0]+1:], nums[1:])
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
