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

		got := countTreesInMatrix(matrix)
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

		got := countTreesInMatrix_v2(matrix)
		want := 336

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})
}
