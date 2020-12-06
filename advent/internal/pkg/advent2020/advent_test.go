package advent2020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	// TODO think of a way to have this more configurable outside of each test
	t.Skip()
	var expenseReport = []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	t.Run("Day 1 Test", func(t *testing.T) {
		assert.Equal(t, 514579, Day1(expenseReport))
	})
	t.Run("Day 1 Part 2 Test", func(t *testing.T) {
		assert.Equal(t, 241861950, Day1Part2(expenseReport))
	})
}

func TestDay2(t *testing.T) {
	t.Skip()
	var passwordPolicies = []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	var tests = []struct {
		testName string
		check    string
		expected int
	}{
		{
			"Day 2 Test",
			"general",
			2,
		},
		{
			"Day 2 Part 2 Test",
			"special",
			1,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			assert.Equal(t, test.expected, Day2(passwordPolicies, test.check))
		})
	}
}

func TestDay3(t *testing.T) {
	t.Skip()
	var mapp = [][]string{
		{".", ".", "#", "#", ".", ".", ".", ".", ".", ".", "."},
		{"#", ".", ".", ".", "#", ".", ".", ".", "#", ".", "."},
		{".", "#", ".", ".", ".", ".", "#", ".", ".", "#", "."},
		{".", ".", "#", ".", "#", ".", ".", ".", "#", ".", "#"},
		{".", "#", ".", ".", ".", "#", "#", ".", ".", "#", "."},
		{".", ".", "#", ".", "#", "#", ".", ".", ".", ".", "."},
		{".", "#", ".", "#", ".", "#", ".", ".", ".", ".", "#"},
		{".", "#", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
		{"#", ".", "#", "#", ".", ".", ".", "#", ".", ".", "."},
		{"#", ".", ".", ".", "#", "#", ".", ".", ".", ".", "#"},
		{".", "#", ".", ".", "#", ".", ".", ".", "#", ".", "#"},
	}

	var tests = []struct {
		testName   string
		expected   int
		stepsDown  int
		stepsRight int
	}{
		{
			"Day 3 Test",
			7,
			1,
			3,
		},
		{
			"Day 3 Test Part 2 - A",
			2,
			1,
			1,
		},
		{
			"Day 3 Test Part 2 - B",
			7,
			1,
			3,
		},
		{
			"Day 3 Test Part 2 - C",
			3,
			1,
			5,
		},
		{
			"Day 3 Test Part 2 - D",
			4,
			1,
			7,
		},
		{
			"Day 3 Test Part 2 - E",
			2,
			2,
			1,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			assert.Equal(t, test.expected, day3Counter(mapp, test.stepsDown, test.stepsRight))
		})
	}

	t.Run("Wrapper test", func(t *testing.T) {
		assert.Equal(t, 336, Day3(mapp, [][]int{
			{1, 1},
			{3, 1},
			{5, 1},
			{7, 1},
			{1, 2},
		}))
	})
}

func TestDay4(t *testing.T) {
	t.Skip()
	var passportData = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

	var validatePassportData = `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007

pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`

	t.Run("No advanced validation", func(t *testing.T) {
		assert.Equal(t, 2, Day4(passportData, false))
	})
	t.Run("Advanced validation", func(t *testing.T) {
		assert.Equal(t, 4, Day4(validatePassportData, true))
	})
}

func TestDay5(t *testing.T) {
	t.Skip()
	tests := []struct {
		input    string
		expected int
	}{
		{
			"FBFBBFFRLR",
			357,
		},
		{
			"BFFFBBFRRR",
			567,
		},
		{
			"FFFBBBFRRR",
			119,
		},
		{
			"BBFFBBFRLL",
			820,
		},
	}

	t.Run("Testing return the seat Id", func(t *testing.T) {
		for _, test := range tests {
			assert.Equal(t, test.expected, day5SeatId(test.input))
		}
	})

	seatLocations := []string{}
	for _, test := range tests {
		seatLocations = append(seatLocations, test.input)
	}

	t.Run("Testing which is the highest seat Id", func(t *testing.T) {
		assert.Equal(t, 820, Day5(seatLocations))
	})

	tests = []struct {
		input    string
		expected int
	}{
		{
			"BBFFBBFRLL",
			820,
		},
		{
			"BBFFBBFRRL",
			822,
		},
		{
			"BBFFBBFRRR",
			823,
		},
	}

	seatLocations = []string{}
	for _, test := range tests {
		seatLocations = append(seatLocations, test.input)
	}

	t.Run("Testing which is missing seat Id", func(t *testing.T) {
		assert.Equal(t, 821, Day5Part2(seatLocations))
	})

}
