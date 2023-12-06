package aochelper

import (
	"os"
	"strings"
)

// Read a file and convert it to a string
//
// Returns the content of a file or an empty string and an error
func ReadFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// Convert a string representing the contents of a file to a string array
// for easier processing
//
// Returns an empty array if the string was empty
func ConvertContentToSlice(content *string) []string {
	if len(*content) == 0 {
		return []string{}
	}

	var result []string
	lines := strings.Split(strings.ReplaceAll(*content, "\r\n", "\n"), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if len(line) > 0 {
			result = append(result, line)
		}
	}

	return result
}
