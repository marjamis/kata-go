package advent

import (
	"fmt"
	"sort"
	"strconv"
)

//Find the fuel required for a set of mass and divide by three, round down, and subtract 2.
func day1CalculateFuel(mass ...int) int {
	fuel := 0
	for i := 0; i < len(mass); i++ {
		fuel += (mass[i] / 3) - 2
	}

	return fuel
}

//Find the fuel required for a mass plus the mass of the additional fuel until the required fuel is 0 or blow.
func day1CalculateFuel2(modules ...int) int {
	totalFuel := day1CalculateFuel(modules...)
	// fmt.Printf("Initial Fuel: %d\n", totalFuel)

	NeedMoreFuel := true
	additionalFuel := totalFuel
	for NeedMoreFuel {
		additionalFuel = day1CalculateFuel(additionalFuel)
		// fmt.Printf("Total: %d - Additional: %d\n", totalFuel, additionalFuel)
		if additionalFuel <= 0 {
			NeedMoreFuel = false
		} else {
			totalFuel += additionalFuel
		}
	}

	return totalFuel
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
			//TODO break out of loop properly
			return v
		default:
			return nil
		}
	}

	return nil
}

const (
	ConstCentralPointX = 15000
	ConstCentralPointY = 15000
)

func day3UpdatePoint() {

}

type coordinates struct {
	X int
	Y int
}

var track []*coordinates

func day3Direction(flag int, instruction string, grid *[30000][30000]int, x int, y int) (int, int) {
	direction := string(instruction[0])
	steps, err := strconv.Atoi(instruction[1:])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Steps to move: %d\n", steps)
	switch direction {
	case "U":
		fmt.Println("Up")
		for i := 0; i < steps; i++ {
			//TODO note about this being -1
			//TODO dedup this as it's essentially the same just with different x or y calculations
			y = y - 1
			if grid[x][y] == 1 && flag == 2 {
				grid[x][y] = 3
				track = append(track, &coordinates{x, y})
			} else {
				grid[x][y] = flag
			}
		}
	case "D":
		fmt.Println("Down")
		for i := 0; i < steps; i++ {
			y = y + 1
			if grid[x][y] == 1 && flag == 2 {
				grid[x][y] = 3
			} else {
				grid[x][y] = flag
			}
		}
	case "L":
		fmt.Println("Left")
		for i := 0; i < steps; i++ {
			x = x - 1
			if grid[x][y] == 1 && flag == 2 {
				grid[x][y] = 3
			} else {
				grid[x][y] = flag
			}
		}
	case "R":
		fmt.Println("Right")
		for i := 0; i < steps; i++ {
			x = x + 1
			if grid[x][y] == 1 && flag == 2 {
				grid[x][y] = 3
			} else {
				grid[x][y] = flag
			}
		}
	}

	return x, y
}

func day3(wire1 []string, wire2 []string) int {
	grid := [30000][30000]int{}
	//Central Point set
	grid[ConstCentralPointX][ConstCentralPointY] = -1

	//TODO add reasoning as to why this was used
	mapWires := func(label int, wire []string) {
		x := ConstCentralPointX
		y := ConstCentralPointY
		for i := range wire {
			fmt.Printf("Wire %d - X: %d Y: %d\n", label, x, y)
			x, y = day3Direction(label, wire[i], &grid, x, y)
		}
	}

	mapWires(1, wire1)
	mapWires(2, wire2)

	var num []int
	for _, coord := range track {
		num = append(num, abs(coord.X-ConstCentralPointX)+abs(coord.Y-ConstCentralPointY))
	}
	sort.Ints(num)

	return num[0]
}

//abs is simple function to return the absolute value of an integer. Absolute value being essentially an always positive number.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
