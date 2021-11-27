package day4

import (
	"crypto/md5"
	"fmt"
	"math"
	"strconv"
)

func FirstProblem(input string) int {
	found := 0
	for i := 1; i<=math.MaxInt64 ; i++ {
		if md5Hash(input + strconv.Itoa(i))[0:5] == "00000" {
			found = i
			break
		}
	}
	return found
}

func SecondProblem(input string) int {
	found := 0
	for i := 1; i<=math.MaxInt64 ; i++ {
		if md5Hash(input + strconv.Itoa(i))[0:6] == "000000" {
			found = i
			break
		}
	}
	return found
}

func md5Hash(s string) string{

	hash := fmt.Sprintf("%x", md5.Sum([]byte(s)))
	//fmt.Println(s, hash)
	return hash
}
