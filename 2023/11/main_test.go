package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"brlywk/AoC/helper"
)

const testFile = "input_test.txt"

var (
	testData   *aochelper.InputData
	testLines  []string
	testMatrix [][]string
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
	testMatrix = aochelper.CreateMatrix(&testLines)

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

func TestDistance(t *testing.T) {
	g1 := Galaxy{Id: 1, Row: 0, Col: 4}
	g2 := Galaxy{Id: 2, Row: 10, Col: 9}

	actual1 := g1.DistanceTo(&g2)
	actual2 := g2.DistanceTo(&g1)
	expected := 15

	if actual1 != actual2 {
		t.Errorf("Not transitive. Got %v and %v", actual1, actual2)
	}

	if actual1 != expected {
		t.Errorf("Expected %v, Actual %v", expected, actual1)
	}
}

func TestFindEmptySpace(t *testing.T) {
	universe := Universe{
		StarChart: &testMatrix,
	}
	universe.FindEmptySpace()
	actualRows := universe.EmptySpace.Rows
	actualCols := universe.EmptySpace.Cols

	expectedCols := []int{2, 5, 8}
	expectedRows := []int{3, 7}

	if !reflect.DeepEqual(expectedRows, actualRows) {
		t.Errorf("Row mismatch. Expected %v, Actual %v", expectedRows, actualRows)
	}

	if !reflect.DeepEqual(expectedCols, actualCols) {
		t.Errorf("Row mismatch. Expected %v, Actual %v", expectedCols, actualCols)
	}
}

func TestExpandEmptySpace(t *testing.T) {
	matrixCopy := make([][]string, len(testMatrix))
	copy(matrixCopy, testMatrix)

	universe := Universe{
		StarChart: &testMatrix,
	}

	universe.FindEmptySpace()
	universe.ExpandEmptySpace()

	fmt.Printf("Universe:\n%v", universe.StarChart)
}

func TestFindGalaxies(t *testing.T) {
	matrixCopy := make([][]string, len(testMatrix))
	copy(matrixCopy, testMatrix)

	universe := Universe{
		StarChart: &testMatrix,
	}

	universe.FindEmptySpace()
	universe.ExpandEmptySpace()
	universe.FindGalaxies()

	for _, g := range universe.Galaxies {
		fmt.Printf("Galaxy: %v\n", g)
	}
}

func TestBigBang(t *testing.T) {
	matrixCopy := make([][]string, len(testMatrix))
	copy(matrixCopy, testMatrix)

    universe := *BigBang(&matrixCopy)


    if len(universe.Galaxies) != 9 {
        t.Errorf("Galaxies\tExpected: %v, Actual: %v", 9, len(universe.Galaxies))
    }
}

func TestEvaluatePart1(t *testing.T) {
	matrixCopy := make([][]string, len(testMatrix))
	copy(matrixCopy, testMatrix)

    universe := *BigBang(&matrixCopy)
    actual := EvaluatePart1(&universe)
    expected := 374

    if actual != expected {
        t.Errorf("Expected: %v, Actual: %v", expected, actual)
    }
}
