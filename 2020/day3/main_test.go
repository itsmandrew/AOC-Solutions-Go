package main

import (
	"fmt"
	"os"
	"testing"
)

func TestBasics(t *testing.T) {

	fname := "test_input/test1.txt"
	t.Run("test 1 pass", func(t *testing.T) {

		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		matrix := createMatrix(file)

		got := countTreesInMatrix(matrix, 3, 1)
		want := 7

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})

	t.Run("test 2 pass", func(t *testing.T) {

		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

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

		got := getProductOfSlice(trees)
		want := int64(336)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})
}
