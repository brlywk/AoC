package main

import (
	aochelper "brlywk/AoC/helper"
	"log"
	"testing"
)

const TestFile = "input_test.txt"

var (
	testContent string
	testData    []string
)

func testSetup() error {
	var err error
	testContent, err = aochelper.ReadFile(TestFile)
	if err != nil {
		return err
	}

	testData = aochelper.ConvertContentToSlice(&testContent)

	return nil
}

func TestMain(m *testing.M) {
	if err := testSetup(); err != nil {
		log.Fatalf("Test setup encountered an error: %v", err)
	}
}

func TestParseInput(t *testing.T) {
	actual := ParseInput(&testData)
	expected := []Race{{Time: 7, BestDistance: 9}, {Time: 15, BestDistance: 40}, {Time: 30, BestDistance: 200}}
	expectedLen := len(expected)

	if len(actual) != expectedLen {
		t.Errorf("Length mismatc. Expected: %v\tActual: %v", expectedLen, len(actual))
	}

	for i, a := range actual {
		if a != expected[i] {
			t.Errorf("Element mismatch. Expected: %v\tActual: %v", expected[i], a)
		}
	}
}

// Once again, let's only test some cases all at once, for brevities sake
func TestTimePressedToDistance(t *testing.T) {
	a1 := TimePressedToDistance(7, 2)
	e1 := 10

	a2 := TimePressedToDistance(7, 4)
	e2 := 12

	a3 := TimePressedToDistance(7, 0)
	e3 := 0

	if a1 != e1 {
		t.Errorf("Expected: %v\tActual: %v", e1, a1)
	}

	if a2 != e2 {
		t.Errorf("Expected: %v\tActual: %v", e2, a2)
	}

	if a3 != e3 {
		t.Errorf("Expected: %v\tActual: %v", e3, a3)
	}
}

func TestCalculateWaysToBeatRecord(t *testing.T) {
	testRace := Race{Time: 7, BestDistance: 9}
	actual := CalculateWaysToBeatRecord(&testRace)
	expected := 4

	if actual != expected {
		t.Errorf("Expected: %v\tActual: %v", expected, actual)
	}
}

func TestEvaluatePart1(t *testing.T) {
	races := ParseInput(&testData)
	actual := EvaluatePart1(&races)
	expected := 288

	if actual != expected {
		t.Errorf("Expected: %v\tActual: %v", expected, actual)
	}
}
