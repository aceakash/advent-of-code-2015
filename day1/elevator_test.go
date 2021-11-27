package day1

import "testing"
import . "github.com/aceakash/advent_of_code_2015"

func TestFirstProblem(t *testing.T) {
	input := MustReadFile()
	got := FirstProblem(input)
	want := 280

	if got != want {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestSecondProblem(t *testing.T) {
	input := MustReadFile()
	got := SecondProblem(input)
	want := 1797

	if got != want {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
