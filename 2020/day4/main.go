package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args[1:]) > 1 {
		fmt.Println("Too many arguments, exiting...")
		return
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Custom scanner for our use case, splits by empty NEW LINE
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// Check if we are at the end of the file
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		// Trim spaces and check if the current line is empty or just whitespace
		line := string(data)
		if strings.TrimSpace(line) == "" {
			return len(data), nil, nil // Found an empty line, treat as delimiter
		}

		// Otherwise, return the whole line as a token
		return len(data), data, nil
	})

	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) > 0 {
			fmt.Println("Found non-empty line:", line)
		}
	}
}
