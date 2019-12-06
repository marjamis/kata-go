package advent

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/marjamis/kata-go/pkg/formatting"
	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	// t.Skip()
	assert := assert.New(t)

	fdata, _ := os.Open("./day1Data.txt")
	defer fdata.Close()
	scanner := bufio.NewScanner(fdata)
	scanner.Split(bufio.ScanLines)
	var ints []int
	for scanner.Scan() {
		d, _ := strconv.Atoi(scanner.Text())
		ints = append(ints, d)
	}

	//Part 1 - https://adventofcode.com/2019/day/1
	//Provided Tests
	assert.Equal(2, day1CalculateFuel(12))
	assert.Equal(2, day1CalculateFuel(14))
	assert.Equal(654, day1CalculateFuel(1969))
	assert.Equal(33583, day1CalculateFuel(100756))

	//Verified Solution
	formatting.AdventWrapper("1", "1", day1CalculateFuel(ints...))

	//Part 2 - https://adventofcode.com/2019/day/1#part2
	//Provided Tests
	assert.Equal(2, day1CalculateFuel2(14))
	assert.Equal(966, day1CalculateFuel2(1969))
	assert.Equal(50346, day1CalculateFuel2(100756))

	//Verified Solution
	formatting.AdventWrapper("1", "2", day1CalculateFuel2(ints...))
}

func TestDay2(t *testing.T) {
	// t.Skip()
	assert := assert.New(t)

	fn := "./day2Data.csv"
	// t.FileExists(fn)
	data, _ := os.Open(fn)
	r := csv.NewReader(data)
	var ints []int
	record, _ := r.Read()
	for _, value := range record {
		d, _ := strconv.Atoi(value)
		ints = append(ints, d)
	}

	//Part 1 - https://adventofcode.com/2019/day/2
	//My tests to build it up
	assert.ElementsMatch([]int{2, 5, 6, 28, 99, 4, 7}, day2(2, 5, 6, 3, 99, 4, 7))
	assert.ElementsMatch([]int{1, 5, 6, 11, 99, 4, 7}, day2(1, 5, 6, 3, 99, 4, 7))
	assert.ElementsMatch([]int{1, 9, 10, 28, 2, 9, 10, 3, 99, 4, 7}, day2(1, 9, 10, 3, 2, 9, 10, 3, 99, 4, 7))
	//Provided Tests
	assert.ElementsMatch([]int{2, 0, 0, 0, 99}, day2(1, 0, 0, 0, 99))
	assert.ElementsMatch([]int{2, 3, 0, 6, 99}, day2(2, 3, 0, 3, 99))
	assert.ElementsMatch([]int{2, 4, 4, 5, 99, 9801}, day2(2, 4, 4, 5, 99, 0))
	assert.ElementsMatch([]int{30, 1, 1, 4, 2, 5, 6, 0, 99}, day2(1, 1, 1, 4, 99, 5, 6, 0, 99))
	assert.ElementsMatch([]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, day2(1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50))

	//Substitute data as required
	ints[1] = 12
	ints[2] = 2

	//Create a copy of the data to be used for future iterations
	fresh := make([]int, len(ints))
	copy(fresh, ints)

	//Verified Solution
	formatting.AdventWrapper("2", "1", day2(fresh...)[0])

	ints[1] = 0
	ints[2] = 0
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			fresh := make([]int, len(ints))
			copy(fresh, ints)
			fresh[1] = i
			fresh[2] = j
			if day2(fresh...)[0] == 19690720 {
				//Verified Solution
				formatting.AdventWrapper("2", "2", 100*i+j)
			}
		}
	}
}

func TestDay3(t *testing.T) {
	// t.Skip()
	assert := assert.New(t)

	fn := "./day3Data.csv"
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

	//Part 1 - https://adventofcode.com/2019/day/3
	//Provided Tests
	assert.Equal(6, day3([]string{"R8", "U5", "L5", "D3"}, []string{"U7", "R6", "D4", "L4"}))
	assert.Equal(159, day3([]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}, []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}))
	assert.Equal(135, day3([]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}, []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}))

	//Verified Solution
	formatting.AdventWrapper("3", "1", day3(strings[0], strings[1]))
}
