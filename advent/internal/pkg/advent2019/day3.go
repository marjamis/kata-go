package advent2019

import (
	"fmt"
	"sort"
	"strconv"
)

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

//abs is simple function to return the absolute value of an integer. Absolute value being essentially an always positive number.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func day3Direction(flag int, instruction string, grid *[25000][25000]int, x int, y int) (int, int) {
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
	//TODO devise a better strat than the 25000
	grid := [25000][25000]int{}
	//Central Point set
	grid[constCentralPoint.X][constCentralPoint.Y] = -1
	//Resetting this TODO maybe a better way
	clashes = nil

	//Note: Used a closure as didnt need to be it's own function and looks cool
	mapWires := func(flag int, wire []string) {
		x := constCentralPoint.X
		y := constCentralPoint.Y
		for i := range wire {
			fmt.Printf("Wire %s - X: %d Y: %d\n", wire[i], x, y)
			x, y = day3Direction(flag, wire[i], &grid, x, y)
		}
	}

	mapWires(1, wire1)
	mapWires(2, wire2)

	num := f()

	return num
}
