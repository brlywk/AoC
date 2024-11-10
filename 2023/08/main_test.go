package main

import (
	"fmt"
	"os"
	"testing"

	"brlywk/AoC/helper"
)

const testFile1 = "input_test.txt"
const testFile2 = "input_test2.txt"

var (
	testData1  *aochelper.InputData
	testLines1 []string
	testData2  *aochelper.InputData
	testLines2 []string
)

// ---- Test Setup ------------------------------

func testSetup() error {
	var err, err2 error
	testData1, err = aochelper.NewInputData(testFile1, true)
	testData2, err = aochelper.NewInputData(testFile2, true)
	if err != nil || err2 != nil {
		fmt.Printf("Error creating testData: %v", err)
		os.Exit(1)
	}

	testLines1 = testData1.GetLines()
	testLines2 = testData2.GetLines()

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

func TestParseInput1(t *testing.T) {
	directions, nodes := ParseInput(&testLines1)

	expectedPattern := "RL"

	actualLength := len(nodes)
	expectedLength := 7

	if expectedPattern != directions {
		t.Errorf("Expected: %v, Actual: %v", expectedPattern, directions)
	}

	if expectedLength != actualLength {
		t.Errorf("Expected: %v, Actual: %v", expectedLength, actualLength)
	}
}

func TestParseInput2(t *testing.T) {
	directions, nodes := ParseInput(&testLines2)

	expectedPattern := "LLR"

	actualLength := len(nodes)
	expectedLength := 3

	if expectedPattern != directions {
		t.Errorf("Expected: %v, Actual: %v", expectedPattern, directions)
	}

	if expectedLength != actualLength {
		t.Errorf("Expected: %v, Actual: %v", expectedLength, actualLength)
	}
}

func TestFindSteps1(t *testing.T) {
	directions, nodes := ParseInput(&testLines1)
	actual := FindStepCount(&directions, &nodes)
	expected := 2

	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestFindSteps2(t *testing.T) {
	directions, nodes := ParseInput(&testLines2)
	actual := FindStepCount(&directions, &nodes)
	expected := 6

	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
