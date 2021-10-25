package main

// Adapted from: https://towardsdatascience.com/answering-monty-hall-problem-with-monte-carlo-6c6069e39bfe
/*
Note: This example could be improved and extended in a few ways, if required, but this is just designed as a quick way to showcase the power of simulations.
*/

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const (
	numberOfDoors = 3

	// Key/Value's of the prizes
	goat      = 0
	car       = 1
	shownGoat = -1
)

func getRevealDoor(doors []int, doorSelectedByContestant int) (doorToOpen int) {
	// Find all the doors (could be one or two door) that don't have the car and hasn't been selected by the contestant
	var doorsThatCanBeOpened []int
	for door, prize := range doors {
		// Appends to an array when it's not the door with the car or the door the contestant selected
		if prize != car && door != doorSelectedByContestant {
			doorsThatCanBeOpened = append(doorsThatCanBeOpened, door)
		}
	}

	// Picks a random item from the array of doors that could be revealed and returns the door number will be openned to show a goat
	return doorsThatCanBeOpened[rand.Intn(len(doorsThatCanBeOpened))]
}

func getDoorToSwitchTo(doors []int, doorSelectedByContestant int) (doorToSwitchTo int) {
	for i := 0; i < len(doors); i++ {
		// Finds the door which hasn't already been show to have a goat and not already selected by the contestant
		if doors[i] != shownGoat && i != doorSelectedByContestant {
			doorToSwitchTo = i
		}
	}
	return
}

func montyhall() (noSwitchSuccess bool, yesSwitchSuccess bool) {
	// All doors start with a goat
	doors := []int{goat, goat, goat}

	// Door with car
	doorWithCar := rand.Intn(numberOfDoors)
	doors[doorWithCar] = car

	// Door selected by the contestant
	doorSelectedByContestant := rand.Intn(numberOfDoors)

	// Find a door that can be openned to reveal a goat, i.e. one that doesn't have the car and isn't already selected by the contestant
	doorToOpen := getRevealDoor(doors, doorSelectedByContestant)
	doors[doorToOpen] = shownGoat
	// log.Printf("Doors (seeing behind them): %+v, Contestant selected door: %d\n", doors, doorSelectedByContestant)

	// Find the available door, i.e. doesn't have the goat and isn't the one already selected by the constentant
	doorToSwitchTo := getDoorToSwitchTo(doors, doorSelectedByContestant)

	// Checks if switching to this new door will have the car and returns the required bools
	if doorToSwitchTo == doorWithCar {
		return false, true
	}

	// If switching wouldn't win the car returns the bools saying this
	return true, false
}

func montecarlo(numOfSimulations int) (nonSwitchSuccesses int, switchSuccesses int) {
	for i := 0; i < numOfSimulations; i++ {
		rand.Seed(time.Now().UnixNano())
		notSwitched, switched := montyhall()
		if notSwitched {
			nonSwitchSuccesses++
		}
		if switched {
			switchSuccesses++
		}
	}

	return
}

func main() {
	// Note: The default number of simulations is high to get as close to accurate as possible, i.e. more executions the more accurate it is, but can be overridden with the flag
	simulations := flag.Int("s", 500000, "The number of simulations to run")
	flag.Parse()

	counfOfNonSwitchSuccesses, countOfSwitchSuccesses := montecarlo(*simulations)

	fmt.Printf("Number of simulations: %d\n", *simulations)
	fmt.Printf("Wins with a switch: %d (%.2f%%)\n", countOfSwitchSuccesses, float64(countOfSwitchSuccesses)/float64(*simulations)*100)
	fmt.Printf("Wins without a switch: %d (%.2f%%)\n", counfOfNonSwitchSuccesses, float64(counfOfNonSwitchSuccesses)/float64(*simulations)*100)
}
