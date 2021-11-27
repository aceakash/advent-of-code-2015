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

func TestSecondProblem(t *testing.T) {
	input := MustReadFile()
	got := SecondProblem(input)
	want := 99

	if got != want {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestCalAreaForWrapping(t *testing.T) {
	testData := []struct{
		present PresentDimensions
		want int
	}{
		{PresentDimensions{2, 3, 4}, 58,},
		{PresentDimensions{1, 1, 10,}, 43,},
	}

	for _, td := range testData {
		got := PresentArea(td.present)
		if got != td.want {
			t.Errorf("Wanted %d for %v, got %d", td.want, td.present, got)
		}
	}
}


