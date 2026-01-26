package main

import "testing"

func TestNewJoltage(t *testing.T) {
	a := "987654321111111"
	awant := 987654321111
	atest := findBestDigits(a, 12)

	if (atest != awant) {
		t.Errorf("isComposedOfSequences(a) = %d, wanted %d.", atest, awant)
	}

	b := "811111111111119"
	bwant := 811111111119
	btest := findBestDigits(b, 12)

	if (atest != awant) {
		t.Errorf("isComposedOfSequences(a) = %d, wanted %d.", btest, bwant)
	}

	c := "234234234234278"
	cwant := 434234234278
	ctest := findBestDigits(c, 12)

	if (atest != awant) {
		t.Errorf("isComposedOfSequences(a) = %d, wanted %d.", ctest, cwant)
	}

	d := "818181911112111"
	dwant := 888911112111
	dtest := findBestDigits(d, 12)

	if (atest != awant) {
		t.Errorf("isComposedOfSequences(a) = %d, wanted %d.", dtest, dwant)
	}

	e := "5373475263753258336423442254746263332334232217334431337464342726873125223932312363675175435324343745"
	ewant := 888911112111
	etest := findBestDigits(e, 12)

	if (atest != awant) {
		t.Errorf("isComposedOfSequences(a) = %d, wanted %d.", etest, ewant)
	}
}