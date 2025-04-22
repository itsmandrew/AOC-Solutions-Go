package main

import (
	"testing"
)

func TestBasic(t *testing.T) {

	t.Run("testing basic example", func(t *testing.T) {

		input := []string{"1969"}

		got := calculateFuelSum(input)
		want := int64(654)

		if got != want {
			t.Errorf("Got %d want %d\n", got, want)
		}
	})

	t.Run("testing multiple example", func(t *testing.T) {

		input := []string{"12", "1969", "14"}

		got := calculateFuelSum(input)
		want := int64(658)

		if got != want {
			t.Errorf("Got %d want %d\n", got, want)
		}
	})

	t.Run("testing recursive add", func(t *testing.T) {
		input := 1969

		got := calculateRecursive(input)
		want := 966

		if got != want {
			t.Errorf("Got %d want %d\n", got, want)
		}
	})
}
