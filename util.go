package advent_of_code_2015

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func MustReadFile() string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	return strings.TrimSpace(string(b))
}

func MustParseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}