package day3

import "testing"
import . "github.com/aceakash/advent_of_code_2015"

func TestFirstProblem(t *testing.T) {
	input := MustReadFile()
	got := FirstProblem(input)
	want := 2572

	if got != want {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestFirstProblem_Examples(t *testing.T) {
	testData := []struct{
		input string
		want int
	}{
		{`>`, 2,},
		{`^>v<`, 4,},
		{`^v^v^v^v^v`, 2,},
	}

	for _, td := range testData {
		got := FirstProblem(td.input)
		if got != td.want {
			t.Errorf("Wanted %d for %v, got %d", td.want, td.input, got)
		}
	}



}


