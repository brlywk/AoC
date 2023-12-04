package main

import (
	"log"
	"os"
	"slices"
	"testing"
)

var (
	fileContent string
)

// Setup and test for ReadFile
func testSetup() error {
	var err error
	fileContent, err = ReadFile(TestFile)
	return err
}

func TestMain(m *testing.M) {
	if err := testSetup(); err != nil {
		log.Fatalf("Unable to read test file: %v", err)
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCountSameElements(t *testing.T) {
	w := []int{1, 2, 3, 4}
	p := []int{2, 4, 6, 8, 10}
	expected := 2

	c := CountSameElements(w, p)

	if c != expected {
		t.Errorf("Expected %v, got %v", expected, c)
	}
}

func TestGetNextNumbers(t *testing.T) {
	n := 1
	c := 4
	r := GetNextNumbers(n, c)
	e := []int{2, 3, 4, 5}

	if !slices.Equal(r, e) {
		t.Errorf("Expected: %v, got %v", e, r)
	}
}

func TestParseGameInput(t *testing.T) {
	lines := ParseInput(fileContent)
	expectedLen := 6

	if len(lines) != expectedLen {
		t.Errorf("Expected %v, got %v", len(lines), expectedLen)
	}
}

func TestParseGames(t *testing.T) {
	// log.Println("\n\nTestParseGames")
	lines := ParseInput(fileContent)
	games := ParseGames(lines)
	r := len(games)
	e := 6

	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestEvaluatePart1(t *testing.T) {
	// log.Println("\n\nEvaluatePart1")
	lines := ParseInput(fileContent)
	games := ParseGames(lines)
	r := EvaluatePart1(games)
	e := 13

	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestEvaluatePart2(t *testing.T) {
	// log.Println("\n\nEvaluatePart1")
	lines := ParseInput(fileContent)
	games := ParseGames(lines)
	r := EvaluatePart2(games)
	e := 30

	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}
