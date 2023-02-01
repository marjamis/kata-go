package example

import "fmt"

func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {

		total += number
	}
	return total
}

func basicVariadic() {
	fmt.Println(sumNumbers(1, 2, 3))
	fmt.Println(sumNumbers(1, 2, 3, 4, 5, 6))
	fmt.Println(sumNumbers(1, 1, 1, 1, 1, 1, 1, 1, 1, 12, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 34, 4, 4, 4, 4, 9))
}

func init() {
	category := GetCategories().AddCategory("variadic")

	category.AddExample("basic",
		CategoryExample{
			Description: "Basic Variadic",
			Function:    basicVariadic,
		})
}
