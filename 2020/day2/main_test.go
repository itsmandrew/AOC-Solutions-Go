package main

import (
	"testing"
)

func TestBasics(t *testing.T) {

	// fileName := "test_input/test1.txt"
	t.Run("testing get first and second", func(t *testing.T) {

		s := "1-3"

		got_lowest, got_highest := getFirstAndSecond(s)
		want_lowest, want_highest := 1, 3

		if got_lowest != want_lowest && got_highest != want_highest {
			t.Errorf("failed, dont want to print the rest")
		}

	})

	t.Run("testing get byte", func(t *testing.T) {
		s := "a:"

		got := getChar(s)
		want := byte('a')

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

	})

	t.Run("check valid", func(t *testing.T) {

		got := matchChecker(byte('a'), 1, 3, "aaa")
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("check valid part 2", func(t *testing.T) {
		got := matchChecker_v2(byte('c'), 2, 9, "ccccccccc")
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
