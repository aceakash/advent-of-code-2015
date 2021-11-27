package day4

import (
	"testing"
)

func TestFirstProblem(t *testing.T) {
	testData := []struct{
		input string
		want int
	}{
		{`abcdef`, 609043,},
		{`pqrstuv`, 1048970,},
		{`yzbqklnj`, 282749,},
		//{ MustReadFile(), 2572},
	}

	for _, td := range testData {
		got := FirstProblem(td.input)
		if got != td.want {
			t.Errorf("Wanted %d for %v, got %d", td.want, td.input, got)
		}
	}
}




