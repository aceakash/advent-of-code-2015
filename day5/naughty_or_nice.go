package day5

import (
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

func IsNice(text string) bool {
	return containsMinThreeVowels(text) && containsDoubleLetter(text) && !containsForbiddenStrings(text, []string{"ab", "cd", "pq", "xy"})
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
