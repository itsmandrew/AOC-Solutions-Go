package main

import (
	"bufio"
	"os"
	"testing"
)

func TestBasics(t *testing.T) {

	fileName := "test_input/test1.txt"
	t.Run("part 1 sample", func(t *testing.T) {

		arr := readInputFile(t, fileName)

		got := bruteForce(arr)
		want := 514579

		assertIntegers(t, got, want)
	})

	t.Run("part 2 sample", func(t *testing.T) {

		arr := readInputFile(t, fileName)

		got := bruteForce_v2(arr)
		want := 241861950

		assertIntegers(t, got, want)
	})
}

func assertIntegers(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func readInputFile(t testing.TB, filename string) []int {
	t.Helper()

	file, err := os.Open(filename)

	if err != nil {
		t.Fatalf("Failed to open file %s: %v", filename, err)
	}

	defer file.Close()

	scan := bufio.NewScanner(file)

	return getSliceFromFile(scan)
}
