package day5

import (
	. "github.com/aceakash/advent_of_code_2015"
	"testing"
)

type NaughtyOrNice bool

const (
	NAUGHTY NaughtyOrNice = true
	NICE    NaughtyOrNice = false
)

func TestFirstProblem(t *testing.T) {
	testData := []struct {
		input string
		want  int
	}{
		{`ugknbfddgicrmopn`, 1},
		{`aaa`, 1},
		{`jchzalrnumimnmhp`, 0},
		{`haegwjzuvuyypxyu`, 0},
		{`dvszwmarrgswjxmb`, 0},
		{MustReadFile(), 999999965},
	}

	for _, td := range testData {
		got := FirstProblem(td.input)
		if got != td.want {
			t.Errorf("Wanted %d for %v, got %d", td.want, td.input, got)
		}
	}
}

func TestIsNice_ShouldContainMinThreeVowels(t *testing.T) {
	testData := []struct {
		input string
		want  bool
	}{
		{`abb`, false},
		{`aeobb`, true},
		{`aaa`, true},
		{`eee`, true},
		{`brehijubb`, true},
		{`xazegovbb`, true},
	}

	for _, td := range testData {
		got := IsNice(td.input)
		if got != td.want {
			t.Errorf("Wanted %v for %v, got %v", td.want, td.input, got)
		}
	}
}
