package aochelper

import (
	"reflect"
	"strconv"
	"testing"
)

const testFile = "input_test.txt"

// Test that empty lines are stripped correctly
func TestInputDataNoEmptyLines(t *testing.T) {
	data, err := NewInputData(testFile, true)
	if err != nil {
		t.Fatalf("NoEmptyLines\tInputData creation encountered an error: %v", err)
	}

	fileName := data.GetFileName()

	content := data.GetContent()
	expectedContent := "Line 1\nLine 2\nLine 3\n"

	lines := data.GetLines()
	expectedLines := []string{"Line 1", "Line 2", "Line 3"}

	if fileName != testFile {
		t.Errorf("Incorrect file name.\tExpected: %v\tActual: %v", testFile, fileName)
	}

	if content != expectedContent {
		t.Errorf("Incorrect content.\nExpected:\n%v\nActual:\n%v", expectedContent, content)
	}

	for i, line := range lines {
		if line != expectedLines[i] {
			t.Errorf("Incorrect line.\tExpected: %v\tActual: %v", expectedLines[i], line)
		}
	}
}

// Test that empty lines are left in
func TestInputDataWithEmptyLines(t *testing.T) {
	data, err := NewInputData(testFile, false)
	if err != nil {
		t.Fatalf("NoEmptyLines\tInputData creation encountered an error: %v", err)
	}

	fileName := data.GetFileName()

	content := data.GetContent()
	expectedContent := "Line 1\nLine 2\n\nLine 3\n"

	lines := data.GetLines()
	expectedLines := []string{"Line 1", "Line 2", "", "Line 3"}

	if fileName != testFile {
		t.Errorf("Incorrect file name.\tExpected: %v\tActual: %v", testFile, fileName)
	}

	if content != expectedContent {
		t.Errorf("Incorrect content.\nExpected:\n%v\nActual:\n%v", expectedContent, content)
	}

	for i, line := range lines {
		if line != expectedLines[i] {
			t.Errorf("Incorrect line.\tExpected: %v\tActual: %v", expectedLines[i], line)
		}
	}
}

func TestMapStringToInt(t *testing.T) {
	input := []string{"1", "2", "3"}
	actual := MapSlice(input, func(e string) int {
		n, _ := strconv.Atoi(e)
		return n
	})
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, Actual %v", expected, actual)
	}
}

func TestCreateMatrix(t *testing.T) {
	input := []string{"AB", "CD"}
	actual := CreateMatrix(&input)
	expected := [][]string{{"A", "B"}, {"C", "D"}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, Actual %v", expected, actual)
	}
}

func TestPrintRefSlice(t *testing.T) {
	s1 := "Hello there"
	s2 := "General Kenobi"
	s3 := "May the force be with you"

	input := []*string{&s1, &s2, &s3}

	PrintRefSlice(input)
}
