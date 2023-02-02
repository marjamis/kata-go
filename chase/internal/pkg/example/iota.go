package example

import (
	"fmt"
)

const (
	// Fruits/Vegetables that are assigned a number from iota. Starts from 1

	// Apple value
	Apple int = iota + 1
	// Lemon value
	Lemon
	// Pear value
	Pear
	// As with go in general _ will assign to nothing
	_
	_
	_
	_
	_
	_
	_
	// The "= iota"'s  aren't required as once defined above any untyped var's will be the iota

	// Carrot value
	Carrot = iota
	// Cauliflower value
	Cauliflower = iota
)

func iotaRun() {
	fmt.Println("Note: iota is useful but shouldn't be used for exported modules or used when data is exported out using these values, in case the const order is changed and hence the reference is changed")
	fmt.Printf("Apple is %d\n", Apple)
	fmt.Printf("Lemon is %d\n", Lemon)
	fmt.Printf("Pear is %d\n", Pear)
	fmt.Printf("Carrot is %d\n", Carrot)
	fmt.Printf("Cauliflower is %d\n", Cauliflower)
}

func init() {
	category := GetCategories().AddCategory("iota")
	category.AddExample("basic",
		CategoryExample{
			Description: "Showing off some basic iota",
			Function:    iotaRun,
		})
}
