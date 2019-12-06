package advent

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func day1(requireSubFuel bool, modules ...int) (totalFuel int) {
	for _, module := range modules {
		moduleFuel := (module / 3) - 2

		additionalFuel := moduleFuel
		if requireSubFuel {
			needMoreFuel := true
			for needMoreFuel {
				additionalFuel = (additionalFuel / 3) - 2
				if additionalFuel <= 0 {
					needMoreFuel = false
				} else {
					totalFuel += additionalFuel
				}
			}
		}

		totalFuel += moduleFuel
	}
	return
}

func day2(v ...int) []int {
	position := 0
	for true {
		switch v[position] {
		case 1:
			//Addition
			v[v[position+3]] = v[v[position+1]] + v[v[position+2]]
			position = position + 4
		case 2:
			//Multiplication
			v[v[position+3]] = v[v[position+1]] * v[v[position+2]]
			position = position + 4
		case 99:
			//End of app
			return v
		default:
			return nil
		}
	}

	return nil
}

type coordinates struct {
	X int
	Y int
}

type clash struct {
	coordinates coordinates
	stepsWire1  int
	stepsWire2  int
}

var (
	constCentralPoint = coordinates{
		X: 15000,
		Y: 15000,
	}
	clashes []*coordinates
)

func day3Direction(flag int, instruction string, grid *[30000][30000]int, x int, y int) (int, int) {
	direction := string(instruction[0])
	steps, err := strconv.Atoi(instruction[1:])
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Printf("Steps to move: %d\n", steps)
	switch direction {
	case "U":
		// fmt.Println("Up")
		for i := 0; i < steps; i++ {
			//TODO dedup this as it's essentially the same just with different x or y calculations
			//Note: This is -1 as it's go up the array which is minusing of rows
			y = y - 1
			if grid[x][y] == 1 && flag == 2 {
				grid[x][y] = 3
				clashes = append(clashes, &coordinates{x, y})
			} else {
				grid[x][y] = flag
			}
		}
	case "D":
		// fmt.Println("Down")
		for i := 0; i < steps; i++ {
			//Note: This is +1 as it's go down the array which is plussing of rows
			y = y + 1
			if grid[x][y] == 1 && flag == 2 {
				grid[x][y] = 3
				clashes = append(clashes, &coordinates{x, y})
			} else {
				grid[x][y] = flag
			}
		}
	case "L":
		// fmt.Println("Left")
		for i := 0; i < steps; i++ {
			x = x - 1
			if grid[x][y] == 1 && flag == 2 {
				grid[x][y] = 3
				clashes = append(clashes, &coordinates{x, y})
			} else {
				grid[x][y] = flag
			}
		}
	case "R":
		// fmt.Println("Right")
		for i := 0; i < steps; i++ {
			x = x + 1
			if grid[x][y] == 1 && flag == 2 {
				grid[x][y] = 3
				clashes = append(clashes, &coordinates{x, y})
			} else {
				grid[x][y] = flag
			}
		}
	}

	return x, y
}

func day3Manhattan() int {
	var num []int
	for _, clash := range clashes {
		num = append(num, abs(clash.X-constCentralPoint.X)+abs(clash.Y-constCentralPoint.Y))
	}
	sort.Ints(num)

	return num[0]
}

func day3Steps() int {
	return 1
}

func day3(wire1 []string, wire2 []string, f func() int) int {
	//TODO devise a better strat than the 30000
	grid := [30000][30000]int{}
	//Central Point set
	grid[constCentralPoint.X][constCentralPoint.Y] = -1
	//Resetting this TODO maybe a better way
	clashes = nil

	//Note: Used a closure as didnt need to be it's own function and looks cool
	mapWires := func(flag int, wire []string) {
		x := constCentralPoint.X
		y := constCentralPoint.Y
		for i := range wire {
			// fmt.Printf("Wire %d - X: %d Y: %d\n", label, x, y)
			x, y = day3Direction(flag, wire[i], &grid, x, y)
		}
	}

	mapWires(1, wire1)
	mapWires(2, wire2)

	num := f()

	return num
}

//abs is simple function to return the absolute value of an integer. Absolute value being essentially an always positive number.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func day4Rules1(num int) bool {
	snum := strconv.Itoa(num)

	hasSameAdjacent := false
	isAscending := true
	for i := 0; i < len(snum)-1; i++ {
		if snum[i] > snum[i+1] {
			isAscending = false
			i = len(snum)
			break
		}

		if snum[i] == snum[i+1] {
			hasSameAdjacent = true
		}
	}

	return hasSameAdjacent && isAscending
}

func day4Rules2(num int) bool {
	snum := strconv.Itoa(num)

	hasSameAdjacent := false
	isAscending := true
	for i := 0; i < len(snum)-1; i++ {
		if snum[i] > snum[i+1] {
			isAscending = false
			i = len(snum)
			break
		}

		if snum[i] == snum[i+1] {
			count := 0
			jk := i + 2
			for j := i + 2; j < len(snum); j++ {
				if snum[i] == snum[j] {
					count++
					jk++
				} else {
					j = len(snum)
					jk = j
				}
			}

			if count%2 == 0 {
				hasSameAdjacent = true
			} else {
				i = jk
			}
		}
	}

	return hasSameAdjacent && isAscending
}

func day4(rng string, f func(int) bool) int {
	r := strings.Split(rng, "-")
	start, _ := strconv.Atoi(r[0])
	end, _ := strconv.Atoi(r[1])

	matches := 0
	for i := start; i <= end; i++ {
		if f(i) {
			matches++
		}
	}

	return matches
}
