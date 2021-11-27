package day3

import (
	"strings"
)

func FirstProblem(input string) int {
	var dirs []Direction = directionsFromInput(input)
	visited := deliveryJob(Location{0, 0}, dirs)
	return countAtLeastOneVisit(visited)
}

func deliveryJob(startLoc Location, dirs []Direction) map[Location]int {
	visited := map[Location]int{}
	currLoc := startLoc
	visited[currLoc] += 1
	for _, dir := range dirs {
		newLoc := goToNextLocation(currLoc, dir)
		visited[newLoc] += 1
		currLoc = newLoc
	}
	return visited
}

func SecondProblem(input string) int {
	var dirs []Direction = directionsFromInput(input)
	santaDirs, roboDirs := splitDirs(dirs)

	santaVisited := deliveryJob(Location{0, 0}, santaDirs)
	roboVisited := deliveryJob(Location{0, 0}, roboDirs)

	visited := mergeVisited(santaVisited, roboVisited)
	return countAtLeastOneVisit(visited)
}

func mergeVisited(visited1 map[Location]int, visited2 map[Location]int) map[Location]int {
	merged := map[Location]int{}

	for location, count := range visited1 {
		if _, ok := merged[location]; !ok {
			merged[location] = count
		} else {
			merged[location] += count
		}
	}
	for location, count := range visited2 {
		if _, ok := merged[location]; !ok {
			merged[location] = count
		} else {
			merged[location] += count
		}
	}
	return merged
}

func splitDirs(dirs []Direction) ([]Direction, []Direction) {
	dirs1 := []Direction{}
	dirs2 := []Direction{}

	for i, dir := range dirs {
		if i % 2 == 0 {
			dirs1 = append(dirs1, dir)
		} else {
			dirs2 = append(dirs2, dir)
		}
	}
	return dirs1, dirs2
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
