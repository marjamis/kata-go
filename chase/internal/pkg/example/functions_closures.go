package example

import (
	"fmt"

	"github.com/marjamis/kata-go/pkg/formatting"
)

func anonymous() {
	// Note: the () at the end of the anonymous function will mean it's executed and the returned
	// value will be provided to the Printf command. Without the () it will instead return the function
	// to Printf.
	fmt.Printf("1. The output from the anonymous function is: %s\n", (func() string {
		return "anonymous function output"
	}()))

	fmt.Println("2. Beginning of fmt.Println(). " + func(mm string) string {
		mm += " Plus some additional data from within the function."
		return mm
	}("Value that is passed into the anonymous function."))
}

// Closures section was adapted from: https://www.calhoun.io/what-is-a-closure/
func newCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func closures() {
	// Creates an instantiation of the newCounter() function which will mean counter will be a function with access to
	// the n variable thats declared in newCounter() even once newCounter() has exited. The closure aspect means counter
	// will still have access to the value of n to increment when called.
	counter := newCounter()
	// Creates a new instantiation of the newCounter() with it's own n value that it will increment when counter 2 is called.
	counter2 := newCounter()

	// The benefit of closures is only when calling the variable counter or counter2 will their respective n values be incremented as newCounter() has now closed allowing restricted access to the variables assigned with the newCounter() function calls.
	fmt.Println(counter())
	fmt.Println(counter())

	fmt.Println(counter2())
	fmt.Println(counter2())

	fmt.Println(counter())
	fmt.Println(counter2())
}

func nestedFunction() {
	// First func() is to say it's an anonymous function and the second is to say it's going to return a function
	f1 := func(value string) func(string) {
		// Within the function is will return the required function by creating a function, with the returned
		// functions execution printing the input
		return func(value2 string) {
			fmt.Println(value2)
		}
	}

	var outputfunc = f1("")
	outputfunc("data")
}

func recurse(c int) {
	fmt.Printf("Forwards: %d\n", c)
	defer fmt.Printf("Backwards: %d\n", c)

	if c == 5 {
		return
	}

	recurse(c + 1)
}

func recursiveFunction() {
	count := 0
	recurse(count)
}

func ClosuresRun() {
	formatting.ExampleWrapper("Anonymous Function", anonymous)
	formatting.ExampleWrapper("Clousure", closures)
	formatting.ExampleWrapper("Nested Function", nestedFunction)
	formatting.ExampleWrapper("Recursive Function", recursiveFunction)
}

func init() {
	GetMyExamples().Add("closures", ClosuresRun)
}
