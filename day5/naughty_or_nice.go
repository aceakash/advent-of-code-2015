package day5

import (
	"fmt"
	"strings"
)

func FirstProblem(input string) int {
	lines := strings.Split(input, "\n")

	count := 0
	for _, line := range lines {
		if IsNice(line) {
			count++
		}
	}

	return count
}

func SecondProblem(input string) int {
	lines := strings.Split(input, "\n")

	count := 0
	for _, line := range lines {
		if IsNice2(line) {
			count++
		}
	}

	return count
}

func IsNice(text string) bool {
	return containsMinThreeVowels(text) && containsDoubleLetter(text) && !containsForbiddenStrings(text, []string{"ab", "cd", "pq", "xy"})
}

func IsNice2(text string) bool {
	return containsPairOfLettersTwice(text) && containsAbaPattern(text)
}

func containsAbaPattern(text string) bool {
	if len(text) < 3 {
		return false
	}
	for i := 0; i <= len(text)-3; i++ {
		if text[i] == text[i+2] && text[i] != text[i+1] {
			return true
		}
	}
	return false
}

func containsPairOfLettersTwice(text string) bool {
	if len(text) < 4 {
		return false
	}
	for i := 0; i <= len(text)-4; i++ {
		pair := fmt.Sprintf("%c%c", text[i], text[i+1])
		if strings.Contains(text[i+2:], pair) {
			return true
		}
	}
	return false
}

func containsForbiddenStrings(text string, forbidden []string) bool {
	for _, fs := range forbidden {
		if strings.Contains(text, fs) {
			return true
		}
	}
	return false
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
	const vowels = "aeiou"
	for _, r := range text {
		if strings.ContainsRune(vowels, r) {
			vowelCount++
		}
		if vowelCount == 3 {
			return true
		}
	}
	return false
}
