package main

import (
	"fmt"
	"os"
	"testing"

	"brlywk/AoC/helper"
)

const testFile = "input_test.txt"

var (
	testData  *aochelper.InputData
	testLines []string
)

// ---- Test Setup ------------------------------

func testSetup() error {
	var err error
	testData, err = aochelper.NewInputData(testFile, true)
	if err != nil {
		fmt.Printf("Error creating testData: %v", err)
		os.Exit(1)
	}

	testLines = testData.GetLines()

	return nil
}

func TestMain(m *testing.M) {
	if err := testSetup(); err != nil {
		fmt.Printf("Test setup encountered an error: %v", err)
		os.Exit(1)
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}

// ---- Tests -----------------------------------

func TestGetStartPoint(t *testing.T) {
	matrix := aochelper.CreateMatrix(&testLines)
	actual := GetStartPoint(&matrix)
	expected := Tile{Col: 0, Row: 2, Label: "S"}

	if expected != *actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestTilesEqual(t *testing.T) {
	t1 := Tile{Col: 42, Row: 23}
	t2 := Tile{Col: 42, Row: 23}
	t3 := Tile{Col: 0, Row: 0}

	a1 := t1.Equals(&t2)
	e1 := true

	a2 := t2.Equals(&t3)
	e2 := false

	if e1 != a1 {
		t.Errorf("Expected: %v, Actual: %v", e1, a1)
	}
	if e2 != a2 {
		t.Errorf("Expected: %v, Actual: %v", e2, a2)
	}
}

func TestInBounds(t *testing.T) {
	matrix := aochelper.CreateMatrix(&testLines)
	start := GetStartPoint(&matrix)
	actual := start.InBounds(&matrix)

	expected := true

	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestGetAdjacent(t *testing.T) {
	fmt.Println("===== TestGetAdjacent =====")
	matrix := aochelper.CreateMatrix(&testLines)
	start := GetStartPoint(&matrix)
	actual := start.GetAdjacent(&matrix)

	fmt.Println("Adjacent:")
	for _, t := range actual {
		fmt.Println(*t)
	}
}

func TestCanReach(t *testing.T) {
	fmt.Println("===== TestCanReach =====")
	matrix := aochelper.CreateMatrix(&testLines)
	start := GetStartPoint(&matrix)
	adjacent := start.GetAdjacent(&matrix)

	actual := []*Tile{}

	for _, a := range adjacent {
		if start.CanGoTo(a) {
			actual = append(actual, a)
		}
	}

	fmt.Println("Reachable from start")
	for _, a := range actual {
		fmt.Println(*a)
	}
}

func TestNextTile(t *testing.T) {
	fmt.Println("===== TestNextTile =====")
	matrix := aochelper.CreateMatrix(&testLines)
	start := GetStartPoint(&matrix)

	fmt.Println("Next chosen (randomized 10 times)")
	for i := 1; i <= 10; i++ {
		next := start.NextTile(&matrix, &Tile{Label: ""})
		fmt.Printf("\t#%v:\t%v\n", i, next)
	}
}

func TestTraversePipes(t *testing.T) {
	fmt.Println("===== TestTraversePipes =====")
	sketch := ParseInput(&testLines)
	TraversePipes(&sketch)

	fmt.Println(sketch)
}

func TestEvaluatePart1 (t *testing.T) {
	sketch := ParseInput(&testLines)
	actual := EvaluatePart1(&sketch)
	expected := 8

	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
