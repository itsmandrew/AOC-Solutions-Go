package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CheckOSArguments() bool {

	if len(os.Args[1:]) == 0 || len(os.Args[1:]) > 1 {
		fmt.Printf("Wrong number of arguments, expected 1 got %d\n", len(os.Args[1:]))
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

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func StrArrayToInt(arr []string) ([]int, error) {

	res := make([]int, len(arr))

	for i := range arr {
		val, err := strconv.Atoi(string(arr[i]))

		if err != nil {
			var zeroVal []int
			return zeroVal, fmt.Errorf("conversion error: %s", err)
		}

		res[i] = val
	}

	return res, nil
}
