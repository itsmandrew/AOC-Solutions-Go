package main

import (
	"os"

	"github.com/itsmandrew/aoc-go/utils"
)

func main() {

	ok := utils.CheckOSArguments()

	if !ok {
		os.Exit(1)
	}

	lines, err := utils.ReadLines(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	matrix := [][]rune{}

	for _, line := range lines {
		newLine := []rune(line)
		matrix = append(matrix, newLine)
	}

	// mark occupied
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			// TODO: check if spot valid, if so mark as occupied
		}
	}

	// mark occupied -> empty
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			// TODO: check if spot valid, if so mark as occupied
		}
	}

}

func markOccupied(matrix *[][]rune, r, c int) bool {

	directions := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	for _, d := range directions {
		newRow, newCol := r+d[0], c+d[1]

		if newRow < 0 || newCol < 0 || newRow >= len(*matrix) || newCol >= len(*matrix) {
			continue
		}

	}

	return true

}
