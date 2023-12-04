package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	InputFile = "input.txt"
	TestFile  = "input_test.txt"
)

type GameInfo struct {
	Id             int
	WinningNumbers []int
	PlayerNumbers  []int
	Points         int
}

func (g GameInfo) String() string {
	return fmt.Sprintf("{ Id: %v, WinningNumbers: %v, PlayerNumbers: %v, Points: %v}",
		g.Id, g.WinningNumbers, g.PlayerNumbers, g.Points)
}

func main() {
	log.Println("Advent of Code 2023 - Day 4")

	content, err := ReadFile(InputFile)
	if err != nil {
		log.Fatalf("Unable to read input file: %v", err)
	}
	lines := ParseInput(content)
	games := ParseGames(lines)
	points := EvaluatePart1(games)

	log.Printf("Day 4 - Part 1: %v", points)
}

// ---- Helper ----------------------------------

// Count how many winning numbers a player has
// not a general function, as we know len(playerNums) > len(winNums)
func CountSameElements(winNums []int, playerNums []int) int {
	// once more: no error handling to keep it short
	count := 0

	for _, pnum := range playerNums {
		// we could iterate again, but let's use something built-in
		if slices.Index(winNums, pnum) > -1 {
			count++
		}
	}
	return count
}

// ---- Part 1 ----------------------------------

func ReadFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func ParseInput(content string) []string {
	var result []string
	lines := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if len(line) > 0 {
			result = append(result, line)
		}
	}

	return result
}

func ParseGames(games []string) []GameInfo {
	var result []GameInfo

	for _, game := range games {
		// no error handling in here, so no check if len(...) > 1
		// let's assume input is well formatted and has no error
		idAndNums := strings.Split(game, ":")
		gameIdStr := strings.Split(idAndNums[0], " ")[1]
		gameId, _ := strconv.Atoi(gameIdStr)
		winAndPlay := strings.Split(idAndNums[1], "|")

		winNumsStr := strings.Fields(winAndPlay[0])
		var winNums []int
		for _, wnumstr := range winNumsStr {
			wnumint, _ := strconv.Atoi(wnumstr)
			winNums = append(winNums, wnumint)
		}

		playerNumsStr := strings.Fields(winAndPlay[1])
		var playerNums []int
		for _, pnumstr := range playerNumsStr {
			pnumint, _ := strconv.Atoi(pnumstr)
			playerNums = append(playerNums, pnumint)
		}

		countMatches := CountSameElements(winNums, playerNums)
		points := math.Pow(2, float64(countMatches)-1)

		gameInfo := GameInfo{
			Id:             gameId,
			WinningNumbers: winNums,
			PlayerNumbers:  playerNums,
			Points:         int(points),
		}

		result = append(result, gameInfo)
	}

	return result
}

func EvaluatePart1(games []GameInfo) int {
	sum := 0

	for _, game := range games {
		sum += game.Points
	}

	return sum
}
