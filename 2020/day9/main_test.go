package main

import "testing"

func TestCheckValid(t *testing.T) {

	hashMap := map[int]int{
		35: 0,
		20: 1,
		15: 2,
	}

	got := checkValid(hashMap, 55)
	want := true

	if got != want {
		t.Errorf("got %v, want %v\n", got, want)
	}
}
