package day3

import (
	"strings"
)

func FirstProblem(input string) int {
	var dirs []Direction = directionsFromInput(input)
	visited := map[Location]int{}

	currLoc := Location{0, 0}
	visited[currLoc] += 1
	for _, dir := range dirs {
		newLoc := goToNextLocation(currLoc, dir)
		//fmt.Printf("Started at %v, followed %v, ended up at %v\n", currLoc, dir, newLoc)
		visited[newLoc] += 1
		currLoc = newLoc
	}
	return countAtLeastOneVisit(visited)
}

func countAtLeastOneVisit(visited map[Location]int) int {
	//fmt.Println(visited)
	return len(visited)
}

func goToNextLocation(currLoc Location, dir Direction) Location {
	switch dir {
	case N:
		return Location{currLoc.X, currLoc.Y + 1}
	case S:
		return Location{currLoc.X, currLoc.Y - 1}
	case E:
		return Location{currLoc.X + 1, currLoc.Y}
	case W:
		return Location{currLoc.X - 1, currLoc.Y}
	default:
		panic("bad input! " + dir )
	}
	return Location{}
}

func directionsFromInput(input string) []Direction {
	stringDirs := strings.Split(input, "")
	dirs := []Direction{}
	for _, strDir := range stringDirs {
		dirs = append(dirs, Direction(strDir))
	}
	return dirs
}

type Direction string

const (
	N Direction = "^"
	S Direction = "v"
	W Direction = "<"
	E Direction = ">"
)

type Location struct {
	X int
	Y int
}
