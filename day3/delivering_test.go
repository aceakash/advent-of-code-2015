package day3

import (
	"testing"
)
import . "github.com/aceakash/advent_of_code_2015"

func TestFirstProblem(t *testing.T) {
	testData := []struct{
		input string
		want int
	}{
		{`>`, 2,},
		{`^>v<`, 4,},
		{`^v^v^v^v^v`, 2,},
		{ MustReadFile(), 2572},
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
		{`^v`, 3,},
		{`^>v<`, 3,},
		{`^v^v^v^v^v`, 11,},
		{ MustReadFile(), 2631},
	}

	for _, td := range testData {
		got := SecondProblem(td.input)
		if got != td.want {
			t.Errorf("Wanted %d for %v, got %d", td.want, td.input, got)
		}
	}
}


