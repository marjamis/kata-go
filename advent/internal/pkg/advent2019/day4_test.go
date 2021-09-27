package advent2019

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
