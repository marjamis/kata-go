package advent2019

import (
	"encoding/csv"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		responseArray, _ := Day5(test.code, test.input...)
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
		_, code := Day5(test.code, test.input...)
		assert.Equal(test.resultCode, code)
	}

	//Verified Solution
	// fmt.Printf("\n")
	// _, output = day5("5", p2Ints...)
	// fmt.Printf("\n")
	// helpers.AdventWrapperInt("5", "2", output)
}
