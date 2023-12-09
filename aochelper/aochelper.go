package aochelper

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// ---- Input Data ------------------------------

// Struct holding input data from a file
//
// Important:
// Use NewInputData to create this struct
type InputData struct {
	fileName string
	content  string
	lines    []string
}

// Creates a new InputData struct by reading a file and splitting the content into lines
//
// stripEmptyLines specifies if empty lines should be omitted
func NewInputData(fileName string, stripEmptyLines bool) (inputData *InputData, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var builder strings.Builder
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		if stripEmptyLines && len(line) == 0 {
			continue
		} else {
			lines = append(lines, line)
			_, err := builder.WriteString(line + "\n")
			if err != nil {
				return nil, err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	content := builder.String()
	data := InputData{
		fileName,
		content,
		lines,
	}

	return &data, nil
}

// Returns the lines of the InputData struct
func (id *InputData) GetLines() []string {
	if id.lines == nil {
		return []string{}
	}
	return id.lines
}

// Returns the content of the file as a single string
func (id *InputData) GetContent() string {
	return id.content
}

// Return the file name of the file originally read
func (id *InputData) GetFileName() string {
	return id.fileName
}

// ---- Helper Functions ------------------------

// Measure the time something takes to execute;
// Usage: defer Measure("Hello there took")()
func Measure(msg string) func() {
	start := time.Now()

	return func() {
		fmt.Printf("Measure\t\t%v\t%v\n", msg, time.Since(start))
	}
}

// Prints a message and quits the program;
// 
// Pretty much the same as log.Fatalf, just way cooler
func PrintAndQuit(msg string, m ...any) {
	fmt.Printf(msg, m...)
	os.Exit(1)
}

// Good old Map function... applies a function to each element of 'slice'
// and returns a slice with the results
func MapSlice[T any, R any](slice []T, mapFunc func(T) R) []R {
	result := make([]R, cap(slice))

	for i, v := range slice {
		result[i] = mapFunc(v)
	}

	return result
}
