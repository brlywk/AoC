package main

import (
	"log"
	"testing"
)

// Test file reading ability...
func TestReadFile(t *testing.T) {
	content, err := ReadFile(TEST_FILE)
	if err != nil {
		t.Errorf("Unable to read test file %v", TEST_FILE)
	}

	if content == "" {
		t.Fatalf("No content found.")
	}
}

func TestParseLine(t *testing.T) {
	testLine := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	testResult := ParseLine(testLine)

	if testResult.Id != 1 {
		log.Fatalf("ID not found. Expected: 1\tGot: %v", testResult.Id)
	}

	if testResult.MaxRed != 4 {
		log.Fatalf("Red incorrect. Expected 4\tGot: %v", testResult.MaxRed)
	}

	if testResult.MaxBlue != 6 {
		log.Fatalf("Blue incorrect. Expected 6\tGot: %v", testResult.MaxBlue)
	}

	if testResult.MaxGreen != 2 {
		log.Fatalf("Green incorrect. Expected 2\tGot: %v", testResult.MaxGreen)
	}

	if !testResult.Possible {
		log.Fatalf("Result incorrectly evaluated to impossible")
	}
}

func TestParseInput(t *testing.T) {
	content, err := ReadFile(TEST_FILE)
	if err != nil {
		t.Errorf("Unable to read test file %v", TEST_FILE)
	}

	lines:= ParseInput(content)

	if len(lines) < 5 {
		t.Fatalf("Results should contain at least 5 structs")
	}
}

func TestEvaluateGame(t *testing.T) {
	const fileToParse string = TEST_FILE

	content, err := ReadFile(fileToParse)
	if err != nil {
		log.Fatalf("Unable to read file %v. Exiting", fileToParse)
	}

	parsedGames := ParseInput(content)
	result := EvaluateGames(parsedGames)

	expected := 8

	if result != expected {
		t.Fatalf("Invalid result. Expected: %v\tGot: %v", expected, result)
	}
}

