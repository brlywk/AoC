package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14

	INPUT_FILE = "input.txt"
	TEST_FILE  = "input_test.txt"
)

type GameInfo struct {
	Id       int
	MaxRed   int
	MaxBlue  int
	MaxGreen int
	Possible bool
}

func main() {
	const fileToParse string = INPUT_FILE

	content, err := ReadFile(fileToParse)
	if err != nil {
		log.Fatalf("Unable to read file %v. Exiting", fileToParse)
	}

	parsedGames := ParseInput(content)

	part1 := EvaluateGames(parsedGames)
	part2 := EvaluateGamesPart2(parsedGames)

	log.Printf("Part 1\tSum of valid game's IDs: %v", part1)
	log.Printf("Part 2\tSum of power games: %v", part2)
}

// Read the content of a file and return as string
func ReadFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Parse the games and create Slice of GameInfos
func ParseInput(content string) []GameInfo {
	lines := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")

	var result []GameInfo = []GameInfo{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		result = append(result, ParseLine(line))
	}

	return result
}

// Parse a single line (i.e. game) and return info about game
func ParseLine(line string) GameInfo {
	game := strings.Split(line, ":")
	id, _ := strconv.Atoi(strings.Split(game[0], " ")[1])

	result := GameInfo{
		Id:       id,
		MaxRed:   0,
		MaxBlue:  0,
		MaxGreen: 0,
		Possible: true,
	}

	rounds := strings.Split(game[1], ";")

	for _, round := range rounds {
		colors := strings.Split(strings.Trim(round, " "), ",")

		for _, color := range colors {
			numAndColor := strings.Split(strings.Trim(color, " "), " ")
			n, _ := strconv.Atoi(numAndColor[0])
			c := numAndColor[1]

			if c == "red" && n > result.MaxRed {
				result.MaxRed = n
			}
			if c == "blue" && n > result.MaxBlue {
				result.MaxBlue = n
			}
			if c == "green" && n > result.MaxGreen {
				result.MaxGreen = n
			}
		}
	}

	if result.MaxRed > MAX_RED || result.MaxBlue > MAX_BLUE || result.MaxGreen > MAX_GREEN {
		result.Possible = false
	}

	return result
}

// Check all games parsed and sum up possible games
func EvaluateGames(games []GameInfo) int {
	sumValid := 0

	for _, game := range games {
		if game.Possible {
			sumValid += game.Id
		}
	}

	return sumValid
}

// Modified evaluation for Part 2
func EvaluateGamesPart2(games []GameInfo) int {
	sumPower := 0

	for _, game := range games {
		power := game.MaxRed * game.MaxBlue * game.MaxGreen
		sumPower += power
	}

	return sumPower
}
