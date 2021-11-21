package advent2019

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6(t *testing.T) {
	assert := assert.New(t)

	fdata, _ := os.Open("../../../test/advent2019/day6Data.txt")
	defer fdata.Close()
	scanner := bufio.NewScanner(fdata)
	scanner.Split(bufio.ScanLines)
	var strings []string
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	var day6Part1Tests = []struct {
		orbits        []string
		countOfOrbits int
	}{
		//Provided Tests
		{
			[]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"},
			42,
		},
		//Personal Tests
		{
			[]string{"E)J", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "J)K", "COM)B", "B)C", "K)L"},
			42,
		},
	}

	for _, test := range day6Part1Tests {
		assert.Equal(test.countOfOrbits, Day6(test.orbits))
	}

	//Verified Solution
	// helpers.AdventWrapperInt("6", "1", day6(strings))

	var day6Part2Tests = []struct {
		orbits        []string
		countOfOrbits int
	}{
		//Provided Tests
		{
			[]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"},
			4,
		},
		//Personal Tests
		{
			[]string{"E)J", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "J)K", "COM)B", "B)C", "K)L", "K)YOU", "I)SAN"},
			4,
		},
	}

	for _, test := range day6Part2Tests {
		assert.Equal(test.countOfOrbits, Day6Part2(test.orbits))
	}

	//Verified Solution
	// helpers.AdventWrapperInt("6", "2", day6Part2(strings))
}
