package advent2019

import (
	"fmt"
	"strconv"
)

func day5OpcodeBreak(data int) (int, int, int, int) {
	opcode := (data % 10) + ((data / 10 % 10) * 10)
	p1 := data / 100 % 10
	p2 := data / 1000 % 10
	p3 := data / 10000 % 10

	return opcode, p1, p2, p3
}

func day5(systemId string, v ...int) ([]int, int) {
	position := 0
	output := 0

	modeSelector := func(position int, p1 int, p2 int) (x int, y int) {
		if p1 == 0 {
			x = v[v[position+1]]
		} else {
			x = v[position+1]
		}

		if p2 != -1 {
			if p2 == 0 {
				y = v[v[position+2]]
			} else {
				y = v[position+2]
			}
		}

		return
	}

	for true {
		opcode, p1, p2, _ := day5OpcodeBreak(v[position])
		switch opcode {
		case 1:
			//Addition
			x, y := modeSelector(position, p1, p2)
			v[v[position+3]] = x + y
			position += 4
		case 2:
			//Multiplication
			x, y := modeSelector(position, p1, p2)
			v[v[position+3]] = x * y
			position += 4
		case 3:
			//Input for a parameter
			systemIdI, _ := strconv.Atoi(systemId)
			v[v[position+1]] = systemIdI
			position += 2
		case 4:
			//Output value of parameter
			x, _ := modeSelector(position, p1, -1)
			fmt.Printf("%d, ", x)
			output = x
			position += 2
		case 5:
			//jump-if-true
			x, y := modeSelector(position, p1, p2)
			if x != 0 {
				position = y
			} else {
				position += 3
			}
		case 6:
			//jump-if-false
			x, y := modeSelector(position, p1, p2)
			if x == 0 {
				position = y
			} else {
				position += 3
			}
		case 7:
			//less than
			x, y := modeSelector(position, p1, p2)
			if x < y {
				v[v[position+3]] = 1
			} else {
				v[v[position+3]] = 0
			}
			position += 4
		case 8:
			//equals
			x, y := modeSelector(position, p1, p2)
			if x == y {
				v[v[position+3]] = 1
			} else {
				v[v[position+3]] = 0
			}
			position += 4
		case 99:
			//End of app
			return v, output
		default:
			return nil, 0
		}
	}

	return nil, 0
}
