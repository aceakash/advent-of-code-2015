package day2

import (
	. "github.com/aceakash/advent_of_code_2015"
	"sort"
	"strings"
)

type PresentDimensions struct {
	l int
	w int
	h int
}

func FirstProblem(input string) int {
	parsedInput := parseInput(input)

	totalArea := 0
	for _, present := range parsedInput {
		totalArea += PresentArea(present)
	}

	return totalArea
}

func SecondProblem(input string) int {
	parsedInput := parseInput(input)

	totalRibbonLength := 0
	for _, present := range parsedInput {
		totalRibbonLength += RibbonLength(present)
	}

	return totalRibbonLength
}

func RibbonLength(present PresentDimensions) int {
	//l, h, w := present.l, present.h, present.w
	p := smallestPerimeter(present)
	v := cubicVolume(present)
	return p + v
}

func PresentArea(present PresentDimensions) int {
	a := areaOfSmallestSize(present)
	return 2*(present.w*present.h+present.w*present.l+present.l*present.h) + a
}


func cubicVolume(present PresentDimensions) int {
	return present.w * present.l * present.h
}

func smallestPerimeter(present PresentDimensions) int {
	s1, s2 := twoSmallestSides(present)
	return 2 * (s1 + s2)
}


func areaOfSmallestSize(present PresentDimensions) int {
	s1, s2 := twoSmallestSides(present)
	return s1 * s2
}

func twoSmallestSides(present PresentDimensions) (a, b int) {
	dims := []int{present.h, present.l, present.w}
	sort.Ints(dims)
	return dims[0], dims[1]
}

func parseInput(input string) []PresentDimensions {
	lines := MustGetInputLinesFromFile()
	dims := []PresentDimensions{}
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		dimsStr := strings.Split(line, "x")

		dims = append(dims, PresentDimensions{
			l: MustParseInt(dimsStr[0]),
			w: MustParseInt(dimsStr[1]),
			h: MustParseInt(dimsStr[2]),
		})
	}
	return dims
}
