package main

import "testing"

func TestGetFileContents(t *testing.T) {
	fileName := "test_input.txt"
	data, err := GetFileContents(&fileName)
	if err != nil {
		t.Errorf("GetFileContents(test_input.txt) has thrown an error: %v", err)
	}

	if len(data) == 0 {
		t.Fatalf("GetFileContents is expected to return a string with length > 0")
	}
}

func TestEvaluateContent(t *testing.T) {
	fileName := "test_input.txt"
	data, err := GetFileContents(&fileName)
	if err != nil {
		t.Errorf("GetFileContents(test_input.txt) has thrown an error: %v", err)
	}

	result := EvaluateContent(data)

	if result != 142 {
		t.Fatalf("Expected result for test_input.txt is 142")
	}
}

func TestEvaluateContentPart2(t *testing.T) {
	fileName := "test_input_p2.txt"
	data, err := GetFileContents(&fileName)
	if err != nil {
		t.Errorf("GetFileContents(test_input.txt) has thrown an error: %v", err)
	}

	data = PrepareData(data)

	result := EvaluateContent(data)

	expected := 281

	if result != expected {
		t.Fatalf("Expected result for %v is %v, got %v", fileName, expected, result)
	}
}
