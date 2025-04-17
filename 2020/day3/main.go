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
		return
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer file.Close()

	matrix := createMatrix(file)

	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	trees := []int{}

	for _, arr := range slopes {
		trees = append(trees, countTreesInMatrix(matrix, arr[0], arr[1]))
	}

	fmt.Printf("%v trees\n", trees)

	res := getProductOfSlice(trees)
	fmt.Printf("Product of trees given slope: %d\n", res)

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

func countTreesInMatrix(matrix [][]string, x_slope, y_slope int) int {
	// Given matrix, counts how many trees
	rows := len(matrix)
	if rows == 0 {
		return 0
	}

	cols := len(matrix[0])

	trees, col := 0, 0

	for row := y_slope; row < rows; row += y_slope {
		col = (col + x_slope) % cols

		if matrix[row][col] == "#" {
			trees++
		}
	}
	return trees
}

func getProductOfSlice(arr []int) int64 {

	res := int64(1)

	for _, val := range arr {
		res = res * int64(val)
	}

	return res
}
