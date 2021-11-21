package advent2019

import (
	"encoding/csv"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2(t *testing.T) {
	assert := assert.New(t)

	fn := "../../../test/advent2019/Day2Data.csv"
	// t.FileExists(fn)
	data, _ := os.Open(fn)
	r := csv.NewReader(data)
	var ints []int
	record, _ := r.Read()
	for _, value := range record {
		d, _ := strconv.Atoi(value)
		ints = append(ints, d)
	}

	//Part 1 - hpkill ttps://adventofcode.com/2019/day/2
	//My tests to build it up
	assert.ElementsMatch([]int{2, 5, 6, 28, 99, 4, 7}, Day2(2, 5, 6, 3, 99, 4, 7))
	assert.ElementsMatch([]int{1, 5, 6, 11, 99, 4, 7}, Day2(1, 5, 6, 3, 99, 4, 7))
	assert.ElementsMatch([]int{1, 9, 10, 28, 2, 9, 10, 3, 99, 4, 7}, Day2(1, 9, 10, 3, 2, 9, 10, 3, 99, 4, 7))
	//Provided Tests
	assert.ElementsMatch([]int{2, 0, 0, 0, 99}, Day2(1, 0, 0, 0, 99))
	assert.ElementsMatch([]int{2, 3, 0, 6, 99}, Day2(2, 3, 0, 3, 99))
	assert.ElementsMatch([]int{2, 4, 4, 5, 99, 9801}, Day2(2, 4, 4, 5, 99, 0))
	assert.ElementsMatch([]int{30, 1, 1, 4, 2, 5, 6, 0, 99}, Day2(1, 1, 1, 4, 99, 5, 6, 0, 99))
	assert.ElementsMatch([]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, Day2(1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50))

	//Substitute data as required
	ints[1] = 12
	ints[2] = 2

	//Create a copy of the data to be used for future iterations
	fresh := make([]int, len(ints))
	copy(fresh, ints)

	//Verified Solution
	// helpers.AdventWrapperInt("2", "1", Day2(fresh...)[0])

	ints[1] = 0
	ints[2] = 0
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			fresh := make([]int, len(ints))
			copy(fresh, ints)
			fresh[1] = i
			fresh[2] = j
			if Day2(fresh...)[0] == 19690720 {
				//Verified Solution
				// helpers.AdventWrapperInt("2", "2", 100*i+j)
			}
		}
	}
}
