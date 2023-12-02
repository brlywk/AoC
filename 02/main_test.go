package main

import (
	"log"
	"os"
	"testing"
)

var testFileContent string

func testSetup() error {
	var err error
	testFileContent, err = ReadFile(TEST_FILE)
	return err
}

func TestMain(m *testing.M) {
	if setupErr := testSetup(); setupErr != nil {
		log.Fatalf("Test setup has failed: %v", setupErr)
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}

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
	lines := ParseInput(testFileContent)

	if len(lines) < 5 {
		t.Fatalf("Results should contain at least 5 structs")
	}
}

func TestEvaluateGame(t *testing.T) {
	parsedGames := ParseInput(testFileContent)
	result := EvaluateGames(parsedGames)

	expected := 8

	if result != expected {
		t.Fatalf("Invalid result. Expected: %v\tGot: %v", expected, result)
	}
}

func TestEvaluateGamePart2(t *testing.T) {
	parsedGames := ParseInput(testFileContent)
	result := EvaluateGamesPart2(parsedGames)

	expected := 2286

	if result != expected {
		t.Fatalf("Invalid result. Expected: %v\tGot: %v", expected, result)
	}
}
