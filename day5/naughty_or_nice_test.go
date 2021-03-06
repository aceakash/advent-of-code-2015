package day5

import (
	. "github.com/aceakash/advent_of_code_2015"
	"testing"
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
		{MustReadFile(), 258},
	}

	for _, td := range testData {
		got := FirstProblem(td.input)
		if got != td.want {
			t.Errorf("Wanted %d for %v, got %d", td.want, td.input, got)
		}
	}
}

func TestSecondProblem(t *testing.T) {
	testData := []struct {
		input string
		want  int
	}{
		{`qjhvhtzxzqqjkmpb`, 1},
		{`xxyxx`, 1},
		{`uurcxstgmygtbstg`, 0},
		{`ieodomkazucvgmuy`, 0},
		{MustReadFile(), 53},
	}

	for _, td := range testData {
		got := SecondProblem(td.input)
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
		{`letter occurs twice but not together 😔`, `aeobcbc`, false},
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

func TestIsNice_ShouldNotContainSomeStrings(t *testing.T) {
	testData := []struct {
		scenario string
		input    string
		want     bool
	}{
		{`contains ab 😔`, `abaeobb`, false},
		{`contains cd 😔`, `cdaeobb`, false},
		{`contains pq 😔`, `pqaeobb`, false},
		{`contains xy 😔`, `xyaeobb`, false},
		{`contains ab although satisfies other reqs 😔`, `abbeo`, false},
		{`no ab, cd, pq, or xy`, `aeogg`, true},
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

func TestIsNice2_PairOfLettersAppearingTwice(t *testing.T) {
	testData := []struct {
		scenario string
		input    string
		want     int
	}{
		{`pair of letters appearing twice 👍`, `ababxyx`, 1},
		{`pair of letters appearing twice 👍`, `xyxyaba`, 1},
		{`pair of letters appearing twice (but not consecutive) 👍`, `xyoooooxyaba`, 1},
		{`pair of letters appearing twice (but not consecutive) 👍`, `aabcdefgaaxyx`, 1},
		{`pair of letters appearing twice but with overlap 😔`, `aaaxyx`, 0},
		{`no pair of letters appearing twice 😔`, `ieodomkazucvgmuyaba`, 0},
	}

	for _, td := range testData {
		t.Run(td.scenario, func(t *testing.T) {
			got := SecondProblem(td.input)
			if got != td.want {
				t.Errorf("Wanted %d for %v, got %d", td.want, td.input, got)
			}
		})
	}
}

func TestIsNice2_AtLeastOneLetterRepeatedIn_aba_Pattern(t *testing.T) {
	testData := []struct {
		scenario string
		input    string
		want     int
	}{
		{`repeated letter in aba patter 👍`, `abab`, 1},
		{`repeated letter in aba patter 👍`, `ababxyx`, 1},
		{`repeated letter in aba patter 👍`, `aaaxyxy`, 1},
		{`no repeated letter in aba patter 😔`, `xysdgxy`, 0},
	}

	for _, td := range testData {
		t.Run(td.scenario, func(t *testing.T) {
			got := SecondProblem(td.input)
			if got != td.want {
				t.Errorf("Wanted %d for %v, got %d", td.want, td.input, got)
			}
		})
	}
}
