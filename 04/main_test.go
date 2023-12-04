package main

import (
	"log"
	"os"
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

func TestParseGameInput(t *testing.T) {
	lines := ParseInput(fileContent)
	expectedLen := 6

	if len(lines) != expectedLen {
		t.Errorf("Expected %v, got %v", len(lines), expectedLen)
	}
}

func TestParseGames(t *testing.T) {
	lines := ParseInput(fileContent)
	games := ParseGames(lines)
	r := len(games)
	e := 6

	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestEvaluatePart1(t *testing.T) {
	lines := ParseInput(fileContent)
	games := ParseGames(lines)
	r := EvaluatePart1(games)
	e := 13

	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}
