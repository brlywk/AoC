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

type Symbol struct {
	X     int
	Y     int
	Value string
}

func (sym Symbol) String() string {
	return fmt.Sprintf("{ Value: %v, X: %v, Y: %v}", sym.Value, sym.X, sym.Y)
}

func (sym Symbol) Equals(otherSym Symbol) bool {
	return sym.Value == otherSym.Value && sym.X == otherSym.X && sym.Y == otherSym.Y
}

type Number struct {
	Value      int
	Line       int
	StartIndex int
	EndIndex   int
	Symbol
}

func (num Number) String() string {
	return fmt.Sprintf("{ Value: %v, Line: %v, StartIndex: %v, EndIndex: %v, Symbol: %v}", num.Value, num.Line, num.StartIndex, num.EndIndex, num.Symbol)
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

// Check if the symbol of two numbers are the same
func SameSymbol(num1 Number, num2 Number) bool {
	return num1.Symbol.Equals(num2.Symbol)
}

func main() {
	log.Println("Advent of Code 2023 - Day 3")

	content, err := ReadFile(INPUT_FILE)
	if err != nil {
		log.Fatalf("Error reading input file")
	}

	matrix := ConvertToStringMatrix(content)
	validNums := FindValidNumbers(&matrix)
	game1 := EvaluateGamePart1(&validNums)
	game2 := EvaluateGamePart2(&validNums)

	log.Printf("Part 1 - Sum: %v", game1)
	log.Printf("Part 2 - Ratio: %v", game2)
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
func HasAdjacentSymbols(matrix *[][]string, num *Number) (bool, Symbol) {
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
				return true, Symbol{Value: (*matrix)[i][j], X: i, Y: j}
			}
		}
	}

	return false, Symbol{}
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

				hasAdjacent, symbol := HasAdjacentSymbols(matrix, &newNum)

				if hasAdjacent {
					newNum.Symbol = symbol
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

func EvaluateGamePart2(nums *[]Number) int {
	sum := 0
	symbolMap := map[Symbol][]Number{}

	for _, num := range *nums {
		if num.Symbol.Value == "*" {
			// note to self: Go adds a key automatically if it doesn't exist and
			// append also works on nil slices
			symbolMap[num.Symbol] = append(symbolMap[num.Symbol], num)
		}
	}

	for _, v := range symbolMap {
		if len(v) == 2 {
			sum += v[0].Value * v[1].Value
		}
	}

	return sum
}

