package aochelper

import (
	"testing"
)

const testFile = "input_test.txt"

func TestReadFile(t *testing.T) {
	actual, err := ReadFile(testFile)
	if err != nil {
		t.Fatalf("Error reading input file: %v", err)
	}

	expected := "Line 1\nLine 2\nLine 3"

	if actual != actual {
		t.Errorf("Expected: %v\nActual %v", expected, actual)
	}
}

func TestConvertContentToSlice(t *testing.T) {
	content, _ := ReadFile(testFile)

	actual := ConvertContentToSlice(&content)
	if len(actual) == 0 {
		t.Error("Length is 0")
	}

	expected := []string{"Line 1", "Line 2", "Line 3"}

	for i, a := range actual {
		if a != expected[i] {
			t.Errorf("Expected: %v\nActual: %v", expected[i], a)
		}
	}
}
