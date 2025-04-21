package main

import "testing"

func TestBasics(t *testing.T) {

	t.Run("testing finding seat id", func(t *testing.T) {

		input := []int{1, 2, 4, 5}
		got := findSeatID(input)
		want := 3

		if got != want {
			t.Errorf("got %d want %d\n", got, want)
		}
	})

	t.Run("testing finding seat id pt 2", func(t *testing.T) {

		input := []int{5, 6, 7, 9, 10}
		got := findSeatID(input)
		want := 8

		if got != want {
			t.Errorf("got %d want %d\n", got, want)
		}
	})
}
