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

func TestEvalPart1(t *testing.T) {
	actual := EvaluatePart1(&testLines)
	expected := 21

	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
