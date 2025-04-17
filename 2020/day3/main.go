package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args[1:]) > 1 {
		fmt.Printf("Got %d arguments, expected 1\n", len(os.Args[1:]))
		os.Exit(-1)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	defer file.Close()

	matrix := createMatrix(file)

	trees := countTreesInMatrix(matrix)
	fmt.Printf("Trees in map: %d\n", trees)

}

func createMatrix(file *os.File) [][]string {
	// Creates our data structure to iterate on

	scanner := bufio.NewScanner(file)
	matrix := [][]string{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		matrix = append(matrix, line)
	}

	return matrix
}

func countTreesInMatrix(matrix [][]string) int {
	// Given matrix, counts how many trees
	trees := 0

	ROWS, COLS := len(matrix), len(matrix[0])
	rc, cc := 0, 0

	for i := 0; i < ROWS-1; i++ {
		rc, cc = rc+1, cc+3
		if matrix[rc%ROWS][cc%COLS] == "#" {
			trees += 1
		}
	}
	return trees
}

func countTreesInMatrix_v2(matrix [][]string) int {
	return 0
}
