package advent2019

import (
	"fmt"
	"strconv"
	"strings"
)

func Day4Rules1(num int) bool {
	snum := strconv.Itoa(num)

	hasSameAdjacent := false
	isAscending := true
	for i := 0; i < len(snum)-1; i++ {
		if snum[i] > snum[i+1] {
			isAscending = false
			break
		}

		if snum[i] == snum[i+1] {
			hasSameAdjacent = true
		}
	}

	return hasSameAdjacent && isAscending
}

func Day4Rules2(num int) bool {
	snum := strconv.Itoa(num)

	hasSameAdjacent := false
	isAscending := true
	for i := 0; i < len(snum)-1; i++ {
		if snum[i] > snum[i+1] {
			isAscending = false
			break
		}

		if snum[i] == snum[i+1] {
			if i+2 < len(snum) {
				if snum[i+2] != snum[i] {
					hasSameAdjacent = true
				} else {
					hasSameAdjacent = false
				}
			}
		}

	}

	if hasSameAdjacent && isAscending {
		fmt.Println(num)
	}

	return hasSameAdjacent && isAscending
}

// Day4 function
func Day4(rng string, f func(int) bool) int {
	r := strings.Split(rng, "-")
	start, _ := strconv.Atoi(r[0])
	end, _ := strconv.Atoi(r[1])

	matches := 0
	for i := start; i <= end; i++ {
		if f(i) {
			matches++
		}
	}

	return matches
}
