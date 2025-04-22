package utils

import (
	"bufio"
	"fmt"
	"os"
)

func CheckOSArguments() bool {

	if len(os.Args[1:]) == 0 || len(os.Args[1:]) > 1 {
		fmt.Printf("Got too many arguments, expected 1 got %d\n", len(os.Args[1:]))
		return false
	}
	return true
}

// Given the path to a file, opens the file and creates a scanner
// and iterates through the file line by line returns a slice of strings (lines)
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
