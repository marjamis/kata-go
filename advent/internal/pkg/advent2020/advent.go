package advent2020

import (
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func day1CheckListAddsTo2020(pointers []int, expenseReport []int) bool {
	plus := expenseReport[pointers[0]]
	for _, pointer := range pointers[1:] {
		plus += expenseReport[pointer]
	}

	if plus == 2020 {
		return true
	}

	return false
}

func day1MultiplyListValues(pointers []int, expenseReport []int) int {
	multi := expenseReport[pointers[0]]
	for _, pointer := range pointers[1:] {
		multi *= expenseReport[pointer]
	}

	return multi
}

func Day1(expenseReport []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(expenseReport)))
	for index := range expenseReport {
		for index2 := range expenseReport[index:] {
			pointers := []int{index, index + index2}
			if day1CheckListAddsTo2020(pointers, expenseReport) {
				return day1MultiplyListValues(pointers, expenseReport)
			}
		}
	}

	return -1
}

func Day1Part2(expenseReport []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(expenseReport)))
	for index := range expenseReport {
		for index2 := range expenseReport[index:] {
			for index3 := range expenseReport[index2:] {
				pointers := []int{index, index + index2, index2 + index3}
				if day1CheckListAddsTo2020(pointers, expenseReport) {
					return day1MultiplyListValues(pointers, expenseReport)
				}
			}
		}
	}

	return -1
}

func day2GenerateConfiguration(policy string) (rangeLower int, rangeUpper int, character string, password string) {
	split := strings.Split(policy, " ")
	rant := strings.Split(split[0], "-")
	// TODO error checking and change Atoi to ParseInt
	rangeLower, _ = strconv.Atoi(rant[0])
	rangeUpper, _ = strconv.Atoi(rant[1])
	character = strings.Split(split[1], ":")[0]
	password = split[2]

	return
}

var day2checks = map[string](func(int, int, string, string) bool){
	"general": (func(rangeLower int, rangeUpper int, character string, password string) bool {
		numberOfInstances := strings.Count(password, character)
		if numberOfInstances >= rangeLower && numberOfInstances <= rangeUpper {
			return true
		}

		return false
	}),
	"special": (func(rangeLower int, rangeUpper int, character string, password string) bool {
		var X, Y bool

		// The -1's are due to the task as the customer starts from 1 rather than 0.
		if string(password[rangeLower-1]) == character {
			X = true
		}
		if string(password[rangeUpper-1]) == character {
			Y = true
		}

		if (X || Y) && !(X && Y) {
			return true
		}

		return false
	}),
}

func Day2(passwordPolicies []string, check string) (count int) {
	//TODO add a proper error check here
	f := day2checks[check]
	for _, policy := range passwordPolicies {
		if f(day2GenerateConfiguration(policy)) {
			count++
		}
	}

	return
}

func day3Counter(mapp [][]string, stepsDown int, stepsRight int) (count int) {
	// TODO investigate other types here
	// tobboganPosition = [2]int{0, 0}
	x := 0
	y := 0
	for y < (len(mapp) - 1) {
		// Perform the step
		// TODO hardcoded values
		if x+stepsRight >= len(mapp[0]) {
			x = (x + stepsRight) - len(mapp[0])
		} else {
			x += stepsRight
		}

		y += stepsDown
		if y >= len(mapp) {
			return
		}

		// Count
		// fmt.Printf("X: %d, Y: %d\n", x, y)
		if mapp[y][x] == "." {
			// fmt.Printf("0")
		} else {
			// fmt.Printf("X")
			count++
		}
	}

	return
}

func Day3(mapp [][]string, runs [][]int) (multi int) {
	// COMBAK the runs array is ambigious for inputs but first of pair is stepsDown and second is stepsRight. This should probably be fixed to a struct.
	multi = 1
	for _, run := range runs {
		multi *= day3Counter(mapp, run[1], run[0])
	}

	return
}

func Day4(passportData string, advancedValidation bool) (count int) {
	type Validation struct {
		Type         string
		MinimumValue int
		MaximumValue int
	}

	requiredFields := []struct {
		FieldName          string
		ValidationFunction (func(string) bool)
	}{
		{ // (Birth Year)
			"byr",
			func(input string) bool {
				num, _ := strconv.Atoi(input)
				return num <= 2002 && num >= 1920
			},
		},
		{ // (Issue Year)
			"iyr",
			func(input string) bool {
				num, _ := strconv.Atoi(input)
				return num <= 2020 && num >= 2010
			},
		},
		{ // (Expiration Year)
			"eyr",
			func(input string) bool {
				num, _ := strconv.Atoi(input)
				return num <= 2030 && num >= 2020
			},
		},
		{ // (Height)
			"hgt",
			func(input string) bool {
				//TODO fix
				num, _ := strconv.Atoi(input[:len(input)-2])
				if strings.Contains(input, "cm") {
					if num >= 150 && num <= 193 {
						return true
					}
				}
				if strings.Contains(input, "in") {
					if num >= 59 && num <= 76 {
						return true
					}
				}
				return false
			},
		},
		{ // (Hair Color)
			"hcl",
			func(input string) bool {
				re := regexp.MustCompile(`#[0-9a-f]{6}`)
				return re.MatchString(input)
			},
		},
		{ // (Eye Color)
			"ecl",
			func(input string) bool {
				colours := []string{
					"amb",
					"blu",
					"brn",
					"gry",
					"grn",
					"hzl",
					"oth",
				}
				for _, colour := range colours {
					if input == colour {
						return true
					}
				}
				return false
			},
		},
		{ // (Passport ID)
			"pid",
			func(input string) bool {
				return len(input) == 9
			},
		},
		// cid (Country ID)
	}

	for _, passport := range strings.Split(passportData, "\n\n") {
		// fmt.Printf("Index: %d, Passport: %s\n", index, passport)
		rfCheck := 0
	out:
		for _, rf := range requiredFields {
			re := regexp.MustCompile(`(?m)` + rf.FieldName + `:(.*?)[\s\$]`)
			t := passport + "$"
			output := re.FindStringSubmatch(t)
			if len(output) == 2 {
				// fmt.Printf("Key: %s, Value: %s\n", rf.FieldName, output[1])
			} else {
				// fmt.Printf("Key not found: %s\n", rf.FieldName)
				continue out
			}

			if advancedValidation && rf.ValidationFunction != nil && !rf.ValidationFunction(output[1]) {
				// fmt.Printf("Invalid for: %s\n", rf.FieldName)
				continue out
			}
			rfCheck++
		}
		// TODO think of a smarter way to do this.
		if rfCheck == len(requiredFields) {
			count++
		}
		// fmt.Println()
	}

	return
}

func day5Parser(lower int, upper int, direction rune) (int, int) {
	// fmt.Printf("Lower: %d, Upper: %d\n", lower, upper)
	lf := float64(lower)
	uf := float64(upper)
	val := int(lf + (math.RoundToEven(uf-lf) / 2))
	if direction == 'F' || direction == 'L' {
		return lower, val
	}
	return val, upper
}

func day5SeatId(seatLocation string) (id int) {
	rowData := seatLocation[:7]
	colData := seatLocation[7:10]

	// fmt.Printf("Row: %s - Col: %s\n", rowData, colData)

	l := 0
	u := 127
	for _, direction := range rowData {
		l, u = day5Parser(l, u, direction)
		// fmt.Printf("Index: %d, Lower: %d, Upper: %d\n", i+1, l, u)
	}

	bl := 0
	bu := 7
	for _, direction := range colData {
		bl, bu = day5Parser(bl, bu, direction)
		// fmt.Printf("Index: %d, Lower: %d, Upper: %d\n", i+1, bl, bu)
	}

	return ((u) * 8) + bu
}

func Day5(seatLocations []string) (id int) {
	for _, seatLocation := range seatLocations {
		if num := day5SeatId(seatLocation); num > id {
			id = num
		}
	}
	return
}

func Day5Part2(seatLocations []string) (id int) {
	seatIds := []int{}
	for _, seatLocation := range seatLocations {
		// fmt.Printf("Seat Id: %d", day5SeatId(seatLocation))
		seatIds = append(seatIds, day5SeatId(seatLocation))
	}
	sort.Sort(sort.IntSlice(seatIds))

	seat := seatIds[0] - 1
	for _, seatId := range seatIds {
		// fmt.Printf("Index: %d, SeatId: %d\n", index, seatId)
		if (seatId - seat) > 1 {
			return seatId - 1
		}
		seat = seatId
	}

	return
}
