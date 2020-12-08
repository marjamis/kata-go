package advent2020

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/marjamis/advent/pkg/helpers"
	log "github.com/sirupsen/logrus"
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

type Day2CheckOption int

const (
	Day2CheckOptionGeneral = 0
	Day2CheckOptionSpecial = 1
)

func day2GenerateConfiguration(policy string) (rangeLower int64, rangeUpper int64, character string, password string) {
	split := strings.Split(policy, " ")
	rant := strings.Split(split[0], "-")
	rangeLower, _ = strconv.ParseInt(rant[0], 0, 64)
	rangeUpper, _ = strconv.ParseInt(rant[1], 0, 64)
	character = strings.Split(split[1], ":")[0]
	password = split[2]

	return
}

var day2checks = map[Day2CheckOption](func(int64, int64, string, string) bool){
	Day2CheckOptionGeneral: (func(rangeLower int64, rangeUpper int64, character string, password string) bool {
		numberOfInstances := int64(strings.Count(password, character))
		if numberOfInstances >= rangeLower && numberOfInstances <= rangeUpper {
			return true
		}

		return false
	}),
	Day2CheckOptionSpecial: (func(rangeLower int64, rangeUpper int64, character string, password string) bool {
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

func Day2(passwordPolicies []string, check Day2CheckOption) (count int) {
	f := day2checks[check]
	for _, policy := range passwordPolicies {
		if f(day2GenerateConfiguration(policy)) {
			count++
		}
	}

	return
}

type ToboganMovement struct {
	X int64
	Y int64
}

func day3Counter(tobMap [][]string, tobMovement ToboganMovement) (count int) {
	var x int64
	var y int64
	for y < (int64(len(tobMap)) - 1) {
		// fmt.Printf("x: %d - tobMove.X: %d - len(tobMap): %d\n", x, tobMovement.X, int64(len(tobMap)))
		if x+tobMovement.X >= int64(len(tobMap[0])) {
			x = (x + tobMovement.X) - int64(len(tobMap[0]))
		} else {
			x += tobMovement.X
		}

		y += tobMovement.Y
		if y >= int64(len(tobMap)) {
			return
		}

		// fmt.Printf("X: %d, Y: %d\n", x, y)
		if tobMap[y][x] == "#" {
			count++
		}
	}
	// fmt.Println()

	return
}

func Day3(tobMap [][]string, runs []ToboganMovement) (multi int) {
	multi = 1
	for _, run := range runs {
		multi *= day3Counter(tobMap, run)
	}

	return
}

func Day4(passportData string, advancedValidation bool) (count int) {
	type Validation struct {
		Type         string
		MinimumValue int
		MaximumValue int
	}

	re := `(.*?)[\s\$]`
	requiredFields := []struct {
		FieldName          string
		FieldFinder        *regexp.Regexp
		ValidationFunction (func(string) bool)
	}{
		{ // (Birth Year)
			"byr",
			regexp.MustCompile(`byr:` + re),
			func(input string) bool {
				num, _ := strconv.Atoi(input)
				return num <= 2002 && num >= 1920
			},
		},
		{ // (Issue Year)
			"iyr",
			regexp.MustCompile(`iyr:` + re),
			func(input string) bool {
				num, _ := strconv.Atoi(input)
				return num <= 2020 && num >= 2010
			},
		},
		{ // (Expiration Year)
			"eyr",
			regexp.MustCompile(`eyr:` + re),
			func(input string) bool {
				num, _ := strconv.Atoi(input)
				return num <= 2030 && num >= 2020
			},
		},
		{ // (Height)
			"hgt",
			regexp.MustCompile(`hgt:` + re),
			func(input string) bool {
				num, _ := strconv.ParseInt(input[:len(input)-2], 0, 64)
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
			regexp.MustCompile(`hcl:` + re),
			func(input string) bool {
				re := regexp.MustCompile(`#[0-9a-f]{6}`)
				return re.MatchString(input)
			},
		},
		{ // (Eye Color)
			"ecl",
			regexp.MustCompile(`ecl:` + re),
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
			regexp.MustCompile(`pid:` + re),
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
			//TODO fix the below requirement of $ due to the regex not working properly
			output := rf.FieldFinder.FindStringSubmatch(passport + "$")
			if len(output) != 2 {
				continue out
			}

			if advancedValidation && rf.ValidationFunction != nil && !rf.ValidationFunction(output[1]) {
				// fmt.Printf("Invalid for: %s\n", rf.FieldName)
				continue out
			}
			rfCheck++
		}
		if rfCheck == len(requiredFields) {
			count++
		}
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

func day5Wrapper(lower int, upper int, data string) int {
	for _, d := range data {
		// fmt.Printf("Index: %d, Lower: %d, Upper: %d\n", i+1, l, u)
		lower, upper = day5Parser(lower, upper, d)
	}

	return upper
}

func day5SeatId(seatLocation string) (id int) {
	rowData := seatLocation[:7]
	colData := seatLocation[7:10]

	return (day5Wrapper(0, 127, rowData) * 8) + day5Wrapper(0, 7, colData)
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

func day6FindCount(declartionForm string, isEveryone bool) int {
	questions := map[rune]int{}

	persons := strings.Split(declartionForm, "\n")
	for _, person := range persons {
		fmt.Println(person)
		for _, char := range person {
			questions[char] += 1
		}
	}

	count := len(questions)
	if isEveryone {
		count = 0
		for _, question := range questions {
			if question == len(persons) {
				count++
			}
		}
		fmt.Printf("Number of people: %d - Count: %d - Mapping: %+v\n---\n", len(persons), count, questions)
	}

	return count
}

func Day6(declartionForms string, isEveryone bool) (count int) {
	forms := strings.Split(declartionForms, "\n\n")
	for _, form := range forms {
		count += day6FindCount(form, isEveryone)
	}

	return
}

type Contains struct {
	key      string
	numberOf int
}

type Bag struct {
	Contains []Contains
	IsIn     []string
}

type Day7SearchOption int

const (
	Day7SearchOptionIsIn     = 0
	Day7SearchOptionContains = 1
)

func day7SearchIsIn(bags map[string]Bag, findBag string, fbs *[]string) {
	for _, bag := range bags[findBag].IsIn {
		// fmt.Printf("Bag: %s\n", bag)
		*fbs = append(*fbs, bag)
		day7SearchIsIn(bags, bag, fbs)
	}

	return
}

func day7SearchContains(bags map[string]Bag, findBag string) (count int) {
	for _, bag := range bags[findBag].Contains {
		// fmt.Printf("Bag: %s - Count: %d\n", bag.key, bag.numberOf)

		if bag.numberOf != 0 {
			count += bag.numberOf + (bag.numberOf * day7SearchContains(bags, bag.key))
		}
	}

	return
}

func Day7(rulesList string, findBag string, typeOfCheck Day7SearchOption) (count int) {
	// TODO this function is just shite, fix it up to be more elegant
	rules := strings.Split(rulesList, "\n")
	bags := map[string]Bag{}
	for _, rule := range rules {
		bag := strings.Split(rule, " bags contain ")[0]

		if len(bag) < 1 {
			continue
		}
		contains := strings.Split(rule, " bags contain ")[1]
		re := regexp.MustCompile(`(.*? bag)s?[,.]`)
		nc := re.FindAllStringSubmatch(contains, -1)
		currentBags := []Contains{}

		for _, n := range nc {
			trmmed := strings.Trim(n[1], " ")
			count, _ := strconv.Atoi(strings.Split(trmmed, " ")[0])
			key := strings.Split(trmmed, " ")[1] + " " + strings.Split(trmmed, " ")[2]

			existingBag, ok := bags[key]
			if !ok {
				bags[key] = Bag{
					nil,
					[]string{
						bag,
					},
				}
			} else {
				var isin = existingBag.IsIn
				isin = append(isin, bag)
				bags[key] = Bag{
					existingBag.Contains,
					isin,
				}
			}

			currentBags = append(currentBags, Contains{key, count})
		}
		existingBag, ok := bags[bag]
		if !ok {
			bags[bag] = Bag{
				currentBags,
				nil,
			}
		} else {
			contains := existingBag.Contains
			contains = append(contains, currentBags...)
			bags[bag] = Bag{
				contains,
				existingBag.IsIn,
			}
		}
	}

	switch typeOfCheck {
	case Day7SearchOptionIsIn:
		fbs := []string{}
		day7SearchIsIn(bags, findBag, &fbs)
		count = len(helpers.RemoveDuplicates(fbs))
	case Day7SearchOptionContains:
		count = day7SearchContains(bags, findBag)
	}

	return
}

type Day8OpData struct {
	opcode    string
	direction int
	visited   bool
}

func day8ExecuteProgram(operations []Day8OpData) (acc int, ok bool) {
	pos := 0
	for pos != len(operations) {
		// fmt.Printf("Opcode: %s - Direction: %d\n", changedOperations[pos].opcode, changedOperations[pos].direction)
		if operations[pos].visited {
			// fmt.Println("Exiting...")
			break
		}

		operations[pos].visited = true

		switch operations[pos].opcode {
		case "nop":
			pos++
		case "acc":
			acc += operations[pos].direction
			pos++
		case "jmp":
			pos += operations[pos].direction
		}
	}

	if pos == len(operations) {
		ok = true
		return
	}

	return
}

func Day8(programData []string, flip bool) (acc int) {
	// TODO upgrade this to a proper wrapper
	log.Debug("Testing")
	operations := []Day8OpData{}
	for _, d := range programData {
		da, _ := strconv.Atoi(strings.Split(d, " ")[1])
		operations = append(operations, Day8OpData{
			strings.Split(d, " ")[0],
			da,
			false,
		})
	}

	if !flip {
		acc, _ = day8ExecuteProgram(operations)
	} else {
		for index := range operations {
			changedOperations := make([]Day8OpData, len(operations))
			copy(changedOperations, operations)

			if changedOperations[index].opcode == "nop" || changedOperations[index].opcode == "jmp" {
				// fmt.Printf("Before: %s\n", changedOperations[index].opcode)
				if changedOperations[index].opcode == "nop" {
					changedOperations[index].opcode = "jmp"
				} else if changedOperations[index].opcode == "jmp" {
					changedOperations[index].opcode = "nop"
				}
				// fmt.Printf("After: %s\n", changedOperations[index].opcode)
			} else {
				// fmt.Println("No change skipping...")
				continue
			}

			var ok bool
			acc, ok = day8ExecuteProgram(changedOperations)
			if ok {
				return
			}
		}
	}

	return
}
