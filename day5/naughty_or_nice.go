package day5

import (
	. "github.com/aceakash/advent_of_code_2015"
)

func FirstProblem(input string) int {
	lines := MustGetInputLinesFromFile()

	count := 0
	for _, line := range lines {
		if IsNice(line) {
			count++
		}
	}

	return count
}

func IsNice(text string) bool {
	return containsMinThreeVowels(text) && containsDoubleLetter(text)
}

func containsDoubleLetter(text string) bool {
	var prev rune = 0
	for _, r := range text {
		if r == prev {
			return true
		}
		prev = r
	}
	return false
}

func containsMinThreeVowels(text string) bool {
	vowelCount := 0
	for _, r := range text {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			vowelCount++
		}
	}
	return vowelCount >= 3
}
