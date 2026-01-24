package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	f := "test.txt"
	want := 3
	pass := partOne(f)

	if want != pass {
		t.Errorf("partOne(f) = %d, wanted %d.", pass, want)
	}
}

func TestPartTwo(t *testing.T) {
	f := "test.txt"
	want := 6
	pass := partTwo(f)

	if want != pass {
		t.Errorf("partTwo(f) = %d, wanted %d.", pass, want)
	}
}