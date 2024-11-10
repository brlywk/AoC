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

	// content, err := aochelper.ReadFile(InputFile)
	// if err != nil {
	// 	log.Fatalf("Error reading input file: %v", err)
	// }
	// lines := aochelper.ConvertContentToSlice(&content)
	data, err := aochelper.NewInputData(InputFile, true)
	if err != nil {
		log.Fatalf("Error creating data: %v", err)
	}
	lines := data.GetLines()

	racesPart1 := ParseInputPart1(&lines)
	part1 := EvaluatePart1(&racesPart1)
	log.Printf("Part 1: %v\n", part1)

	racePart2 := ParseInputPart2(&lines)
	part2 := EvaluatePart2(&racePart2)
	log.Printf("Part 2: %v\n", part2)
}

// ---- Helper ----------------------------------

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

// ---- Part 1 ----------------------------------

func ParseInputPart1(lines *[]string) []Race {
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

func EvaluatePart1(races *[]Race) int {
	result := 1

	for _, r := range *races {
		result *= CalculateWaysToBeatRecord(&r)
	}

	return result
}

// ---- Part 2 ----------------------------------

func ParseInputPart2(lines *[]string) Race {
	re := regexp.MustCompile(`\d+`)

	timeMatches := re.FindAllString((*lines)[0], -1)
	distMatches := re.FindAllString((*lines)[1], -1)

	timeStr := ""
	distStr := ""

	for i := 0; i < len(timeMatches); i++ {
		timeStr += timeMatches[i]
		distStr += distMatches[i]
	}

	time, _ := strconv.Atoi(timeStr)
	dist, _ := strconv.Atoi(distStr)

	race := Race{
		Time:         time,
		BestDistance: dist,
	}

	return race
}

func EvaluatePart2(race *Race) int {
	return CalculateWaysToBeatRecord(race)
}
