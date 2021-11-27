package day2

import (
	. "github.com/aceakash/advent_of_code_2015"
	"testing"
)

func TestFirstProblem(t *testing.T) {
	input := MustReadFile()
	got := FirstProblem(input)
	want := 1598415

	if got != want {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestCalAreaForWrapping(t *testing.T) {

	present1 := PresentDimensions{
		l: 2,
		w: 3,
		h: 4,
	}

	got1 := PresentArea(present1)
	if got1 != 58 {
		t.Errorf("Expected 58 for %v, got %d", present1, got1)
	}

	present2 := PresentDimensions{
		l: 1,
		w: 1,
		h: 10,
	}

	got2 := PresentArea(present2)
	if got2 != 43 {
		t.Errorf("Expected 43 for %v, got %d", present2, got2)
	}
}


