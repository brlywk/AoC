package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fileName := flag.String("f", "input.txt", "Filename of file to parse")

	data, err := GetFileContents(fileName)
	if err != nil {
		log.Fatalf("Unable to read file %v", fileName)
	}

	// Part 2: Adjust data before calculating results
	data = PrepareData(data)

	result := EvaluateContent(data)
	fmt.Printf("Result: %v", result)
}

func GetFileContents(fileName *string) (string, error) {
	content, err := os.ReadFile(*fileName)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// Part 2: Prepare input file to find spelled out numbers
func PrepareData(content string) string {
	// Let's cheat a little bit:
	// Issues can be caused if letters of numbers overlap, e.g. eight two -> eightwo
	// So let's just keep the first and last letter of the word and mash the number inbetween
	// This way, we prevent numbers to be lost on replace but can still regex away all letters
	wordToNumber := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	adjustedData := content 

	for numAsStr, strNum := range wordToNumber {
		adjustedData = strings.ReplaceAll(adjustedData, numAsStr, strNum)
	}

	return adjustedData
}

func EvaluateContent(content string) int {
	lines := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")

	sum := 0
	re := regexp.MustCompile(`\D`)

	for _, v := range lines {
		numOnly := re.ReplaceAllString(v, "")
		strNum := numOnly

		if len(numOnly) == 0 {
			strNum = "0"
		}

		if len(numOnly) == 1 {
			strNum = numOnly + numOnly
		}

		if len(numOnly) > 2 {
			asBytes := []byte(numOnly)
			first := asBytes[0]
			last := asBytes[len(asBytes)-1]
			strNum = string(first) + string(last)
		}

		num, err := strconv.Atoi(strNum)
		if err != nil {
			continue
		}
		sum += num
	}

	return sum
}
