package main

import "testing"

func TestPartOne(t *testing.T) {
	a := 111
	awant := true
	atest := isComposedOfSequences(a)

	if (atest != awant) {
		t.Errorf("isComposedOfSequences(a) = %t, wanted %t.", atest, awant)
	}

	b := 112
	bwant := false
	btest := isComposedOfSequences(b)

	if (btest != bwant) {
		t.Errorf("isComposedOfSequences(a) = %t, wanted %t.", btest, bwant)
	}

	c := 1212
	cwant := true
	ctest := isComposedOfSequences(c)

	if (ctest != cwant) {
		t.Errorf("isComposedOfSequences(a) = %t, wanted %t.", ctest, cwant)
	}
}