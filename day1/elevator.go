package day1


func FirstProblem(input string) int {
	count := 0
	for _, c := range input {
		if c == '(' {
			count++
		} else if c == ')' {
			count--
		}
	}
	return count
}


func SecondProblem(input string) int {
	count := 0
	var i int
	for _, c := range input {
		if c == '(' {
			count++
		} else if c == ')' {
			count--
		}
		if count == -1 {
			break
		}
		i++
	}
	return i+1
}