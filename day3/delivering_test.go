package day3

import "testing"
import . "github.com/aceakash/advent_of_code_2015"

func TestFirstProblem(t *testing.T) {
	input := MustReadFile()
	got := FirstProblem(input)
	want := 123

	if got != want {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestFirstProblem_Examples(t *testing.T) {
	input := `>`
	got := FirstProblem(input)
	want := 2

	if got != want {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}


