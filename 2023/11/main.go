package main

import (
	aochelper "brlywk/AoC/helper"
	"fmt"
	"math"
	"slices"
)

const InputFile = "input.txt"

// ---- Main ------------------------------------

func main() {
	fmt.Println("Advent of Code 2023 - Day 11")

	data, _ := aochelper.NewInputData(InputFile, false)
    lines := data.GetLines()
    chaos := aochelper.CreateMatrix(&lines)
    universe := BigBang(&chaos)

    part1 := EvaluatePart1(universe)
    fmt.Printf("Part 1: %v\n", part1)
}

// ---- Structs ---------------------------------

type Universe struct {
	StarChart *[][]string
	Galaxies  []*Galaxy
	EmptySpace
}

type EmptySpace struct {
	Rows []int
	Cols []int
}

type Galaxy struct {
	Id  int
	Row int
	Col int
}

// Calculate shortest distance between two galaxies
func (g1 *Galaxy) DistanceTo(g2 *Galaxy) int {
	col := math.Abs(float64(g2.Col) - float64(g1.Col))
	row := math.Abs(float64(g2.Row) - float64(g1.Row))

	return int(col + row)
}

// ---- Creation of the known universe ----------

// Takes the unformed chaos of a newly born universe and expands the empty spaces
// in that universe
func (universe *Universe) FindEmptySpace() {
	chaos := universe.StarChart

	emptyCols := []int{}
	for i := 0; i < len((*chaos)[0]); i++ {
		emptyCols = append(emptyCols, i)
	}
	emptyRows := []int{}

	for i := 0; i < len(*chaos); i++ {
		// galaxyCol := slices.Index((*chaos)[i], "#")
		galaxyCols := []int{}
		for j := range (*chaos)[i] {
			if (*chaos)[i][j] == "#" {
				galaxyCols = append(galaxyCols, j)
			}
		}

		// fmt.Printf("%v\tGalaxy Cols: %v\n", i, galaxyCols)

		if len(galaxyCols) > 0 {
			// Remove column with galaxy in it from emptyCols
			// fmt.Printf("Galaxy found in row %v, col %v\n", i, galaxyCol)
			for _, gc := range galaxyCols {
				slices.Replace(emptyCols, gc, gc+1, -1)
			}
		} else {
			// this row has no galaxy in it, add to empty rows
			emptyRows = append(emptyRows, i)
		}
	}

	//copy remaining empty cols here
	ec := []int{}
	for _, c := range emptyCols {
		if c != -1 {
			ec = append(ec, c)
		}
	}

	universe.EmptySpace = EmptySpace{
		Rows: emptyRows,
		Cols: ec,
	}
}

// Doubles all empty rows and cols
func (universe *Universe) ExpandEmptySpace() {
	// We could transpose this as a matrix... but nah, too much work,
	// so let's just go rows first, cols second
	eRows := universe.EmptySpace.Rows
	eCols := universe.EmptySpace.Cols

	newStartChart := [][]string{}

	// double empty rows
	for i, row := range *(*universe).StarChart {
		// insert empty cols in each row before adding the row itself
		adjustedRow := []string{}
		for j, col := range row {
			adjustedRow = append(adjustedRow, string(col))

			if slices.Index(eCols, j) > -1 {
				adjustedRow = append(adjustedRow, ".")
			}
		}

		newStartChart = append(newStartChart, adjustedRow)

		if slices.Index(eRows, i) > -1 {
			newStartChart = append(newStartChart, adjustedRow)
		}
	}

	universe.StarChart = &newStartChart
}

// We need this as a separate step b/c with ever expanding space galaxies
// forever more will be drifting apart...
// (i.e. before expanding empty spaces all indices would be incorrect)
func (universe *Universe) FindGalaxies() {
	chaos := *universe.StarChart
	count := 1

	for i := range chaos {
		galaxyCols := []int{}
		for j := range chaos[i] {
			if chaos[i][j] == "#" {
				galaxyCols = append(galaxyCols, j)
			}
		}

		for _, g := range galaxyCols {
			universe.Galaxies = append(universe.Galaxies, &Galaxy{Id: count, Row: i, Col: g})
			count++
		}
	}
}

// Time for the creation of the universe
func BigBang(chaos *[][]string) *Universe {
	universe := Universe{
		StarChart: chaos,
	}

	universe.FindEmptySpace()
	universe.ExpandEmptySpace()
	universe.FindGalaxies()

	return &universe
}

// ---- Part 1 ----------------------------------

func EvaluatePart1(universe *Universe) int {
	sum := 0

	for i := range universe.Galaxies {
		for j := range universe.Galaxies {
			if i == j {
				continue
			}

			g1 := universe.Galaxies[i]
			g2 := universe.Galaxies[j]

			sum += g1.DistanceTo(g2)
		}
	}

	return sum / 2
}
