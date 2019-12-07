package formatting

import "fmt"

func ExampleWrapper(name string, f func()) {
	fmt.Printf("## %s\n\n", name)
	defer fmt.Printf("\n-----\n\n")

	f()
}

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
