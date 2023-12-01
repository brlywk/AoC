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

