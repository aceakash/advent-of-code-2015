package day4

import (
	. "github.com/aceakash/advent_of_code_2015"
	"testing"
)

func TestFirstProblem(t *testing.T) {
	testData := []struct{
		input string
		want int
	}{
		{`abcdef`, 609043,},
		{`pqrstuv`, 1048970,},
		{ MustReadFile(), 282749},
	}

	for _, td := range testData {
		got := FirstProblem(td.input)
		if got != td.want {
			t.Errorf("Wanted %d for %v, got %d", td.want, td.input, got)
		}
	}
}

func TestSecondProblem(t *testing.T) {
	testData := []struct{
		input string
		want int
	}{
		{ MustReadFile(), 9962624},
	}

	for _, td := range testData {
		got := SecondProblem(td.input)
		if got != td.want {
			t.Errorf("Wanted %d for %v, got %d", td.want, td.input, got)
		}
	}
}




