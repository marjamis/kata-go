package helpers

import "fmt"

func AdventWrapperInt(day string, part string, result int) {
	fmt.Printf("## Advent Day %s - Part %s\n", day, part)
	defer fmt.Printf("\n-----\n")

	fmt.Printf("The solution is: %d", result)
}

func AdventWrapper(day string, part string) {
	fmt.Printf("## Advent Day %s - Part %s\n", day, part)
	defer fmt.Printf("\n-----\n")

	fmt.Printf("The result is:")
}
