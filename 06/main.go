package main

import (
	"log"
	"regexp"
	"strconv"

	"brlywk/AoC/helper"
)

const InputFile = "input.txt"

type Race struct {
	Time         int
	BestDistance int
}

// ---- Main ------------------------------------

func main() {
	log.Println("Advent of Code 2023 - Day 6")

	content, err := aochelper.ReadFile(InputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	lines := aochelper.ConvertContentToSlice(&content)
	races := ParseInput(&lines)
	part1 := EvaluatePart1(&races)

	log.Printf("Part 1: %v\n", part1)
}

// ---- Helper ----------------------------------

func ParseInput(lines *[]string) []Race {
	var result []Race

	re := regexp.MustCompile(`\d+`)

	timeMatches := re.FindAllString((*lines)[0], -1)
	distMatches := re.FindAllString((*lines)[1], -1)

	for i, t := range timeMatches {
		time, _ := strconv.Atoi(t)
		dist, _ := strconv.Atoi(distMatches[i])

		newRace := Race{
			Time:         time,
			BestDistance: dist,
		}

		result = append(result, newRace)
	}

	return result
}

// ---- Part 1 ----------------------------------

func TimePressedToDistance(total int, pressed int) int {
	if pressed == 0 {
		return 0
	}

	travelTime := total - pressed
	return travelTime * pressed
}

func CalculateWaysToBeatRecord(race *Race) int {
	r := *race
	var result []int

	for t := 1; t < r.Time; t++ {
		dist := TimePressedToDistance(r.Time, t)

		if dist > r.BestDistance {
			result = append(result, t)
		}
	}

	return len(result)
}

func EvaluatePart1(races *[]Race) int {
	result := 1

	for _, r := range *races {
		result *= CalculateWaysToBeatRecord(&r)
	}

	return result
}
