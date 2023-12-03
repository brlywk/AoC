package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var (
	fileContent  string
	contentArray [][]string
)

func testSetup() error {
	var err error
	fileContent, err = ReadFile(TEST_FILE)
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	contentArray = ConvertToStringMatrix(fileContent)
	expectedLength := 10
	if len(contentArray) != expectedLength {
		return fmt.Errorf("Expected byte array length of %v, got %d",
			expectedLength, len(contentArray))
	}

	return nil
}

func TestMain(m *testing.M) {
	if err := testSetup(); err != nil {
		log.Fatalf("Test setup failed: %v", err)
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}

// ---- helper ----
func TestAppendFrontString(t *testing.T) {
	front := []string{"a", "b"}
	end := []string{"c", "d"}

	result := AppendFront(end, front...)

	expectedLen := 4
	expectedArr := []string{"a", "b", "c", "d"}

	if len(result) != expectedLen {
		t.Errorf("Length mismatch for %v. Expected %v, got %v", result, expectedLen, len(result))
	}

	for i, v := range result {
		if v != expectedArr[i] {
			t.Errorf("Elements mismatch. Expected %v, got %v", expectedArr[i], v)
		}
	}
}

func TestIsNumber(t *testing.T) {
	str1 := "a"
	str2 := "4"

	if IsNumber(str1) {
		t.Errorf("Failed. %v is not a number", str1)
	}

	if !IsNumber(str2) {
		t.Errorf("Failed. %v is a number", str2)
	}
}

func TestIsSymbol(t *testing.T) {
	str1 := "4"
	str2 := "."
	str3 := "$"

	if IsSymbol(str1) {
		t.Errorf("%v is not a symbol", str1)
	}

	if IsSymbol(str2) {
		t.Errorf("%v is not a symbol", str2)
	}

	if !IsSymbol(str3) {
		t.Errorf("%v is a symbol", str3)
	}
}

// ---- Main functions ----
func TestReadFile(t *testing.T) {
	if len(fileContent) < 1 {
		log.Fatalf("Expected some file content to be found")
	}
}

func TestHasAdjacentSymbols(t *testing.T) {
	testNum1 := Number{
		Value:      42,
		StartIndex: 1,
		EndIndex:   2,
		Line:       0,
	}
	testNum2 := Number{
		Value:      42,
		StartIndex: 1,
		EndIndex:   2,
		Line:       1,
	}

	noSymbol := [][]string{{".", "4", "2", "."}}
	hasSymbol := [][]string{{".", ".", ".", "."}, {".", "4", "2", "."}, {".", ".", ".", "$"}}

	if HasAdjacentSymbols(&noSymbol, &testNum1) {
		t.Errorf("Failed. %v has no adjacent symbol", noSymbol)
	}

	if !HasAdjacentSymbols(&hasSymbol, &testNum2) {
		t.Errorf("Failed. %v has one adjacent symbol", hasSymbol)
	}

}

func TestFindValidNumbers(t *testing.T) {
	result := FindValidNumbers(&contentArray)

	expectedLen := 8
	if len(result) != expectedLen {
		t.Errorf("Failed. Expected %v, got %v -> %v", 8, len(result), result)
	}
}

func TestEvaluateGamePart1(t *testing.T) {
	nums := FindValidNumbers(&contentArray)
	result := EvaluateGamePart1(&nums)

	expected := 4361

	if result != expected {
		t.Errorf("Failed. Expected %v, got %v", expected, result)
	}
}
