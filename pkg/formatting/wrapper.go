package formatting

import "fmt"

func ExampleWrapper(name string, f func()) {
	fmt.Printf("## %s\n\n", name)
	defer fmt.Printf("\n-----\n\n")

	f()
}
