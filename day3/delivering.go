package day3

func FirstProblem(input string) int {
	//var dirs []Direction = directionsFromInput(input)
	//visited := map[Location]int{}

	return 0
}

func directionsFromInput(input string) []Direction {
	return nil
}

type Direction string
const (
	N Direction = "^"
	S Direction = "V"
	W Direction = "<"
	E Direction = ">"
)

type Location struct {
	X int
	Y int
}
