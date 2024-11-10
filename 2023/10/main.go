package main

// I really, really don't like puzzles focused on matrix traversal...

import (
	"brlywk/AoC/helper"
	"fmt"
	"math"
)

const InputFile = "input.txt"

// ---- Constructs ------------------------------

type Tile struct {
	Label string
	Row   int
	Col   int
}

func (t Tile) String() string {
	return fmt.Sprintf("'%v' > %v:%v", t.Label, t.Row, t.Col)
}

// Check if two tiles are the same
func (tile *Tile) Equals(otherTile *Tile) bool {
	return tile.Row == otherTile.Row && tile.Col == otherTile.Col
}

// Checks if the current tile is withing the bounds of our map
func (t *Tile) InBounds(matrix *[][]string) bool {
	m := *matrix

	// All lines have same 'width' so any line is fine for this check
	return t.Col >= 0 && t.Col < len(m[0]) && t.Row >= 0 && t.Row < len(m)
}

// Get all tiles adjacent to this tile
func (t *Tile) GetAdjacent(matrix *[][]string) []*Tile {
	result := []*Tile{}

	// check all surrounding indices
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if math.Abs(float64(i)) == math.Abs(float64(j)) {
				continue
			}

			tmpT := Tile{Col: t.Col + j, Row: t.Row + i}

			// check if adjacent tile is within matrix
			if tmpT.InBounds(matrix) {
				tmpT.Label = (*matrix)[tmpT.Row][tmpT.Col]

				if !tmpT.Equals(t) && t.CanGoTo(&tmpT) {
					result = append(result, &tmpT)
				}
			}
		}
	}

	return result
}

// Can other tile be reached from this tile?
//
// Does not check the actual bounds, rather only looks at the label
// and if the other tile is in a place that would be reachable
func (tile *Tile) CanGoTo(otherTile *Tile) bool {
	t := *tile
	o := *otherTile

	switch t.Label {

	case "|":
		return t.Col == o.Col && (t.Row-1 == o.Row || t.Row+1 == o.Row)
	case "-":
		return t.Row == o.Row && (t.Col-1 == o.Col || t.Col+1 == o.Col)
	case "L":
		return (t.Col == o.Col && t.Row-1 == o.Row) || (t.Row == o.Row && t.Col+1 == o.Col)
	case "J":
		return (t.Col == o.Col && t.Row-1 == o.Row) || (t.Row == o.Row && t.Col-1 == o.Col)
	case "7":
		return (t.Col == o.Col && t.Row+1 == o.Row) || (t.Row == o.Row && t.Col-1 == o.Col)
	case "F":
		return (t.Col == o.Col && t.Row+1 == o.Row) || (t.Row == o.Row && t.Col+1 == o.Col)
	// We need to check what is adjacent to S in order to consider it 'reachable'
	case "S":
		if t.Row == o.Row {
			if o.Col < t.Col {
				return o.Label == "L" || o.Label == "-" || o.Label == "F"
			} else {
				return o.Label == "J" || o.Label == "-" || o.Label == "7"
			}
		}

		if t.Col == o.Col {
			if o.Row < t.Row {
				return o.Label == "7" || o.Label == "|" || o.Label == "F"
			} else {
				return o.Label == "J" || o.Label == "|" || o.Label == "L"
			}
		}

		fallthrough
	default:
		return false
	}
}

// Check which directions are 'open' to this tile
//
// # Except for a Start Node, this can only ever have one option
//
// However, we treat S as 'always reachable', so we if we have a choice
// between going to 'something' or S, never choose S!
func (t *Tile) NextTile(matrix *[][]string, prevTile *Tile) *Tile {
	adjacent := t.GetAdjacent(matrix)

	nextOptions := []*Tile{}
	for _, adj := range adjacent {
		if t.CanGoTo(adj) && !adj.Equals(t) && adj.Label != "." && !adj.Equals(prevTile) {
			nextOptions = append(nextOptions, adj)
		}
	}

	aochelper.PrintRefSlice(nextOptions)

	var nextChoice *Tile
	// It seems as has been found as a choice
	if len(nextOptions) > 1 {
		for _, no := range nextOptions {
			if no.Label != "S" {
				nextChoice = no
			}
		}
	} else {
		nextChoice = nextOptions[0]
	}

	return nextChoice
}

// Represents our puzzles map
type Sketch struct {
	Map           [][]string
	Start         *Tile
	Steps         []*Tile
	Farthest      *Tile
	FarthestSteps int
}

func (s Sketch) String() string {
	return fmt.Sprintf("\nStart: %v\nSteps: %v\nFarthest: %v\nFarthest Steps: %v", s.Start, s.Steps, s.Farthest, s.FarthestSteps)
}

// ---- Main ------------------------------------

func main() {
	fmt.Println("Advent of Code 2023 - Day 9")

	data, err := aochelper.NewInputData(InputFile, true)
	if err != nil {
		aochelper.PrintAndQuit("Unable to read input file: %v", err)
	}

	lines := data.GetLines()
	sketch := ParseInput(&lines)

	part1 := EvaluatePart1(&sketch)
	fmt.Printf("Part 1: %v\n", part1)
}

// ---- Helper ----------------------------------

func GetStartPoint(matrix *[][]string) *Tile {
	result := Tile{}
	result.Label = "S"

	for i := 0; i < len(*matrix); i++ {
		for j := 0; j < len((*matrix)[i]); j++ {
			if (*matrix)[i][j] == "S" {
				result.Col = j
				result.Row = i
				break
			}
		}
	}

	return &result
}

func TraversePipes(sketch *Sketch) {
	steps := (*sketch).Steps

	prevTile := Tile{Label: "Hello there"}
	currentTile := (*sketch).Start

	for {
		next := currentTile.NextTile(&sketch.Map, &prevTile)

		if next.Label == "S" {
			break
		}

		steps = append(steps, next)
		prevTile = *currentTile
		currentTile = next
	}

	sketch.FarthestSteps = (len(steps) / 2) + 1
	sketch.Farthest = steps[sketch.FarthestSteps]
}

// ---- Part 1 ----------------------------------

func ParseInput(lines *[]string) Sketch {
	result := Sketch{}

	result.Map = aochelper.CreateMatrix(lines)
	result.Steps = []*Tile{}
	result.Start = GetStartPoint(&result.Map)
	TraversePipes(&result)

	return result
}

func EvaluatePart1(sketch *Sketch) int {
	return sketch.FarthestSteps
}
