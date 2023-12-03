package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	INPUT_FILE = "input.txt"
	TEST_FILE  = "input_test.txt"
)

type Number struct {
	Value      int
	Line       int
	StartIndex int
	EndIndex   int
}

func (num Number) String() string {
	return fmt.Sprintf("{ Value: %v, Line: %v, StartIndex: %v, EndIndex: %v}", num.Value, num.Line, num.StartIndex, num.EndIndex)
}

// ---- helper functions ----

// Naive function, so no checks of slice length
func AppendFront[T any](slice []T, elem ...T) []T {
	result := make([]T, len(slice)+len(elem)-2)
	copy(result, elem)
	result = append(result, slice...)
	return result
}

// Check if string is a number
func IsNumber(mysteryValue string) bool {
	_, err := strconv.Atoi(mysteryValue)
	return err == nil
}

// Check if string is a 'symbol' (i.e. not a . or a number)
func IsSymbol(mysteryValue string) bool {
	return mysteryValue != "." && !IsNumber(mysteryValue)
}

func main() {
	log.Println("Advent of Code 2023 - Day 3")

	content, err := ReadFile(INPUT_FILE)
	if err != nil {
		log.Fatalf("Error reading input file")
	}

	matrix := ConvertToStringMatrix(content)
	validNums := FindValidNumbers(&matrix)
	result := EvaluateGamePart1(&validNums)

	log.Printf("Part 1 - Sum: %v", result)
}

func ReadFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// Should be easier to work on an array with indices
func ConvertToStringMatrix(data string) [][]string {
	lines := strings.Split(strings.ReplaceAll(data, "\r\n", "\n"), "\n")

	var matrix [][]string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if len(line) > 0 {
			matrix = append(matrix, strings.Split(line, ""))
		}
	}

	return matrix
}

// Check for adjacent symbol
func HasAdjacentSymbols(matrix *[][]string, num *Number) bool {
	// Range for symbols are line - 1, start - 1 to line + 1, end + 1

	// Define the range of indices we need to cover
	var startIndex int
	if num.StartIndex-1 >= 0 {
		startIndex = num.StartIndex - 1
	} else {
		startIndex = num.StartIndex
	}

	var endIndex int
	if num.EndIndex+1 < len((*matrix)[0]) {
		endIndex = num.EndIndex + 1
	} else {
		endIndex = num.EndIndex
	}

	var lineStart int
	if num.Line-1 >= 0 {
		lineStart = num.Line - 1
	} else {
		lineStart = num.Line
	}

	var lineEnd int
	if num.Line+1 < len(*matrix) {
		lineEnd = num.Line + 1
	} else {
		lineEnd = num.Line
	}

	// log.Printf("Number: %v", num)
	// log.Printf("Matrix: %v\n", *matrix)
	// log.Printf("LineStart: %v\tLineEnd: %v\tIndexStart: %v\tIndexEnd: %v", lineStart, lineEnd, startIndex, endIndex)

	for i := lineStart; i <= lineEnd; i++ {
		for j := startIndex; j <= endIndex; j++ {
			if IsSymbol((*matrix)[i][j]) {
				// log.Printf("Num: %v\tChar at [%v][%v]: %v\n", num, i, j, (*matrix)[i][j])
				return true
			}
		}
	}

	return false
}

// Find a number and return start and end index
func FindValidNumbers(matrix *[][]string) []Number {
	var result []Number
	currentNumber := ""
	currentStart := -1
	saveNum := false

	for i := range *matrix {
		for j, v := range (*matrix)[i] {
			if IsNumber(v) {
				currentNumber += v
				if currentStart == -1 {
					currentStart = j
				}

				// edge case: if we reach the end and have a number we need to save it
				if j+1 == len((*matrix)[i]) {
					saveNum = true
				}
			} else {
				if currentNumber != "" {
					saveNum = true
				}
			}

			if saveNum {
				asNum, _ := strconv.Atoi(currentNumber)

				newNum := Number{
					Value:      asNum,
					Line:       i,
					StartIndex: currentStart,
					EndIndex:   j - 1,
				}

				if HasAdjacentSymbols(matrix, &newNum) {
					result = append(result, newNum)
				}

				currentNumber = ""
				currentStart = -1

				saveNum = false
			}

		}
	}

	return result
}

func EvaluateGamePart1(validNums *[]Number) int {
	sum := 0

	for _, num := range *validNums {
		sum += num.Value
	}

	return sum
}
