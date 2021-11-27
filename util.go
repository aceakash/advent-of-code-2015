package advent_of_code_2015

import (
	"log"
	"os"
	"strings"
)

func MustReadFile() string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	return strings.TrimSpace(string(b))
}
