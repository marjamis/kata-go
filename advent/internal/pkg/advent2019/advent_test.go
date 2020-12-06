package advent2019

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	assert := assert.New(t)

	fdata, _ := os.Open("../../../test/advent2019/day1Data.txt")
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
	assert.Equal(2, day1(false, 12))
	assert.Equal(2, day1(false, 14))
	assert.Equal(654, day1(false, 1969))
	assert.Equal(33583, day1(false, 100756))

	//Verified Solution
	// helpers.AdventWrapperInt("1", "1", day1(false, ints...))

	//Part 2 - https://adventofcode.com/2019/day/1#part2
	//Provided Tests
	assert.Equal(2, day1(true, 14))
	assert.Equal(966, day1(true, 1969))
	assert.Equal(50346, day1(true, 100756))
	//
	// //Verified Solution
	// helpers.AdventWrapperInt("1", "2", day1(true, ints...))
}

func TestDay2(t *testing.T) {
	assert := assert.New(t)

	fn := "../../../test/advent2019/day2Data.csv"
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
	// helpers.AdventWrapperInt("2", "1", day2(fresh...)[0])

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
				// helpers.AdventWrapperInt("2", "2", 100*i+j)
			}
		}
	}
}

func TestDay3(t *testing.T) {
	assert := assert.New(t)

	fn := "../../../test/advent2019/day3Data.csv"
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
	assert.Equal(6, day3([]string{"R8", "U5", "L5", "D3"}, []string{"U7", "R6", "D4", "L4"}, day3Manhattan))
	assert.Equal(159, day3([]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}, []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}, day3Manhattan))
	assert.Equal(135, day3([]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}, []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}, day3Manhattan))
	//
	// //Verified Solution
	// helpers.AdventWrapperInt("3", "1", day3(strings[0], strings[1], day3Manhattan))

	//Part 2 - https://adventofcode.com/2019/day/3
	//Provided Tests
	assert.Equal(610, day3([]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}, []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}, day3Steps))

	//Verified Solution
	// helpers.AdventWrapperInt("3", "1", day3(strings[0], strings[1], day3Manhattan))
}

func TestDay4(t *testing.T) {
	assert := assert.New(t)

	//Part 1 - https://adventofcode.com/2019/day/4
	//My tests to build it up
	assert.Equal(9, day4("10-100", day4Rules1))
	//
	// //Verified Solution
	// helpers.AdventWrapperInt("4", "1", day4("353096-843212", day4Rules1))

	//Part 2 - https://adventofcode.com/2019/day/4
	//My tests to build it up
	assert.Equal(8, day4("10-200", day4Rules2))
	//
	// //Verified Solution
	// helpers.AdventWrapperInt("4", "2", day4("353096-843212", day4Rules2))
}

func TestDay5(t *testing.T) {
	assert := assert.New(t)

	fn := "../../../test/advent2019/day5Data.csv"
	// t.FileExists(fn)
	data, _ := os.Open(fn)
	r := csv.NewReader(data)
	var ints []int
	record, _ := r.Read()
	for _, value := range record {
		d, _ := strconv.Atoi(value)
		ints = append(ints, d)
	}

	p2Ints := make([]int, len(ints))
	copy(p2Ints, ints)

	//Part 1 - https://adventofcode.com/2019/day/5
	//My tests to build it up
	opcode, p1, p2, p3 := day5OpcodeBreak(1102)
	assert.Equal(2, opcode)
	assert.Equal(1, p1)
	assert.Equal(1, p2)
	assert.Equal(0, p3)

	//conver to tdd
	// var output int

	var day5Tests = []struct {
		code        string
		input       []int
		resultArray []int
	}{
		{
			"1",
			[]int{2, 5, 6, 3, 99, 4, 7},
			[]int{2, 5, 6, 28, 99, 4, 7},
		},
		{
			"1",
			[]int{1, 5, 6, 3, 99, 4, 7},
			[]int{1, 5, 6, 11, 99, 4, 7},
		},
		{
			"1",
			[]int{1, 9, 10, 3, 2, 9, 10, 3, 99, 4, 7},
			[]int{1, 9, 10, 28, 2, 9, 10, 3, 99, 4, 7},
		},
		{
			"1",
			[]int{2, 3, 0, 3, 99},
			[]int{2, 3, 0, 6, 99},
		},
		{
			"1",
			[]int{2, 4, 4, 5, 99, 0},
			[]int{2, 4, 4, 5, 99, 9801},
		},
		{
			"1",
			[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
		{
			"1",
			[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			[]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		//Part 2 tests

	}

	for _, test := range day5Tests {
		responseArray, _ := day5(test.code, test.input...)
		assert.ElementsMatch(test.resultArray, responseArray)
	}

	//Verified Solution
	// helpers.AdventWrapper("5", "1")
	// day5("1", ints...)
	// fmt.Println()

	var day5TestsP2 = []struct {
		code       string
		input      []int
		resultCode int
	}{
		{
			"8",
			[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			1,
		},
		{
			"7",
			[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			0,
		},
		{
			"7",
			[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8, 8},
			1,
		},
		{
			"8",
			[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8, 8},
			0,
		},
		{
			"7",
			[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			999,
		},
		{
			"8",
			[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			1000,
		},
		{
			"9",
			[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			1001,
		},
	}

	for _, test := range day5TestsP2 {
		_, code := day5(test.code, test.input...)
		assert.Equal(test.resultCode, code)
	}

	//Verified Solution
	// fmt.Printf("\n")
	// _, output = day5("5", p2Ints...)
	// fmt.Printf("\n")
	// helpers.AdventWrapperInt("5", "2", output)
}

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
		assert.Equal(test.countOfOrbits, day6(test.orbits))
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
		assert.Equal(test.countOfOrbits, day6Part2(test.orbits))
	}

	//Verified Solution
	// helpers.AdventWrapperInt("6", "2", day6Part2(strings))
}
