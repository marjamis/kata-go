package advent2019

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3(t *testing.T) {
	assert := assert.New(t)

	fn := "../../../test/advent2019/Day3Data.csv"
	// t.FileExists(fn)
	data, _ := os.Open(fn)
	r := csv.NewReader(data)
	var strings [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		var content []string
		for i := range record {
			content = append(content, record[i])
		}
		strings = append(strings, content)
	}
	fmt.Println(strings)

	//Part 1 - https://adventofcode.com/2019/day/3
	//Provided Tests
	assert.Equal(6, Day3([]string{"R8", "U5", "L5", "D3"}, []string{"U7", "R6", "D4", "L4"}, Day3Manhattan))
	assert.Equal(159, Day3([]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}, []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}, Day3Manhattan))
	assert.Equal(135, Day3([]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}, []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}, Day3Manhattan))
	// //
	// // //Verified Solution
	// // helpers.AdventWrapperInt("3", "1", Day3(strings[0], strings[1], Day3Manhattan))

	// //Part 2 - https://adventofcode.com/2019/day/3
	// //Provided Tests
	assert.Equal(610, Day3([]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}, []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}, Day3Steps))

	//Verified Solution
	// helpers.AdventWrapperInt("3", "1", Day3(strings[0], strings[1], Day3Manhattan))
}
