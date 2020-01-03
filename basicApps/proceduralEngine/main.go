package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	numberOfOptions     = 4
	symbolTree          = "\u1F33"
	symbolMountain      = "\u26F0"
	symbolWater         = "\u1F30"
	symbolGem           = "\u1F48"
	symbolSpace         = " "
	symbolHealthPotion  = "\u2764"
	symbolKey           = "\u1F51"
	symbolTreasureChest = "\u2423"
)

var (
	mapSizeX   int
	mapSizeY   int
	seed       int64
	questItems = []string{
		symbolKey,
		symbolTreasureChest,
	}
)

func main() {
	flag.Parse()

	//Minimum map size - currently based on quest items so the grid size needs to grow to at least incorporate these items.
	if !(mapSizeX*mapSizeY >= len(questItems)) {
		log.Fatalf("Map size must be greater than the specified rows/columns allow. Size of map: %d, Items needing placement: %d", mapSizeX*mapSizeY, len(questItems))
	}

	//Determines if a new seed needs creation or use the provided seed
	var s int64
	if seed == -1 {
		s = int64(time.Now().Nanosecond())
		defer fmt.Printf("New seed is: %d\n", s)
	} else {
		s = seed
		defer fmt.Printf("Existing seed is: %d\n", s)
	}
	rand.Seed(s)

	workflow(createBlankMap())
}

func workflow(a [][]string) {
	placeQuestItems(a)

	//Go through map location and tries to place a random item based on the rules.
	for x := 0; x < mapSizeX; x++ {
		for y := 0; y < mapSizeY; y++ {
			if a[x][y] == "" {
				a[x][y] = placeGeneralLocation(a, x, y)
			}
		}
	}

	//Print out the grid
	for i := range a {
		fmt.Println(a[i])
	}
}

func placeQuestItems(a [][]string) {
	for _, qi := range questItems {
		//TODO Temporary way to ensure they're placed but should be better to ensure it's impossible to be caught in an infinite loop
		var x, y int
		for true {
			x = rand.Int() % mapSizeX
			y = rand.Int() % mapSizeY

			if a[x][y] == "" {
				break
			}
		}

		a[x][y] = qi
	}
}

func placeGeneralLocation(a [][]string, x int, y int) string {
	ri := rand.Int() % numberOfOptions
	switch ri {
	case 0:
		return symbolTree
	case 1:
		if y > 0 && a[x][y-1] == symbolTree {
			return placeHelpfulItem()
		}
		return symbolWater
	case 2:
		return symbolSpace
	case 3:
		// A mountain can't be placed directly next to another mountain
		if !checkPerimeterIsClear(a, x, y, symbolMountain) {
			return symbolSpace
		}
		return symbolMountain
	}

	return "X"
}

func placeHelpfulItem() string {
	switch rand.Int() % 2 {
	case 0:
		return symbolHealthPotion
	case 1:
		return symbolGem
	}

	return "X"
}

//checkPerimeterIsClear returns true or false on if a specific symbol is left, right, up or down of the current position provided.
func checkPerimeterIsClear(a [][]string, x int, y int, symbol string) bool {
	//Check up
	if y > 0 && a[x][y-1] == symbol {
		return false
	}

	//Check down
	if y < len(a[0])-1 && a[x][y+1] == symbol {
		return false
	}

	//Check left
	if x > 0 && a[x-1][y] == symbol {
		return false
	}

	//Check right
	if x < len(a)-1 && a[x+1][y] == symbol {
		return false
	}

	return true
}

func createBlankMap() (a [][]string) {
	//Sets up the map arrays
	for i := 0; i < mapSizeX; i++ {
		ta := make([]string, mapSizeY)
		a = append(a, ta)
	}

	return
}

func init() {
	flag.IntVar(&mapSizeX, "rows", 10, "Number of rows to be used.")
	flag.IntVar(&mapSizeY, "columns", 10, "Number of columns to be used.")
	flag.Int64Var(&seed, "seed", -1, "Uses an existing seed if it exits.")
}
