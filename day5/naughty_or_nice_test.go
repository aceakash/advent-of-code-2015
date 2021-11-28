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
		scenario string
		input    string
		want     bool
	}{
		{`only one vowel 😔`, `abb`, false},
		{`three diff vowels 👍`, `aeobb`, true},
		{`three same vowels 👍`, `aaa`, true},
		{`three same vowels 👍`, `eee`, true},
		{`three vowels in order 👍`, `brehijubb`, true},
		{`three vowels out of order 👍`, `xozegavbb`, true},
		{`three vowels (two unique) 👍`, `bbreuehyu`, true},
		{`only two vowels 😔`, `aa`, false},
		{`only two vowels 😔`, `eabb`, false},
	}

	for _, td := range testData {
		t.Run(td.scenario, func(t *testing.T) {
			got := IsNice(td.input)
			if got != td.want {
				t.Errorf("Wanted %v for %v, got %v", td.want, td.input, got)
			}
		})
	}
}

func TestIsNice_ShouldContainAtLeastOneDoubleLetter(t *testing.T) {
	testData := []struct {
		scenario string
		input    string
		want     bool
	}{
		{`contains double letter 👍`, `aeobb`, true},
		{`no double letter 😔`, `aeobcdef`, false},
		{`letter occurs twice but not attached 😔`, `aeobcbc`, false},
		{`double vowel 👍`, `aeoo`, true},
	}

	for _, td := range testData {
		t.Run(td.scenario, func(t *testing.T) {
			got := IsNice(td.input)
			if got != td.want {
				t.Errorf("Wanted %v for %v, got %v", td.want, td.input, got)
			}
		})

	}
}
