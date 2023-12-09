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

func TestParseInput(t *testing.T) {
	actual := ParseInput(&testLines)[0]
	expected := History{Raw: []int{0, 3, 6, 9, 12, 15}, Extras: [][]int{{3, 3, 3, 3, 3}, {0, 0, 0, 0}}, Extrapolated: []int{0, 3}, Next: 18}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestPart1(t *testing.T) {
	data := ParseInput(&testLines)
	actual := EvaluatePart1(&data)
	expected := 114

	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
