package example

import "fmt"

func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {

		total += number
	}
	return total
}

func VariadicRun() {
	fmt.Println(sumNumbers(1, 2, 3))
	fmt.Println(sumNumbers(1, 2, 3, 4, 5, 6))
	fmt.Println(sumNumbers(1, 1, 1, 1, 1, 1, 1, 1, 1, 12, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 34, 4, 4, 4, 4, 9))
}

func init() {
	GetMyExamples().Add("variadic", VariadicRun)
}
