package example

import (
	"fmt"

	"github.com/marjamis/kata-go/chase/internal/pkg/example/exampleSubPackage"
	"github.com/marjamis/kata-go/pkg/formatting"
)

func structValues() {
	s := exampleSubPackage.External()

	fmt.Printf("Used struct: %+v\n", s)
	fmt.Printf("Public value in struct: %d\n", s.Public)
	// Uncomment the below to receive the compile time error of not being able to access the private value
	// fmt.Printf("Private value in struct s: %d\n", s.private)
}

func copyReference() {
	a := exampleSubPackage.Drink{
		Name: "Lemonade",
		Ice:  true,
	}

	b := a
	c := &a

	fmt.Println("Before changes to data")
	fmt.Printf("a - %+v - %p\n", a, &a)
	fmt.Printf("b - %+v - %p\n", b, &b)
	fmt.Printf("c - %+v - %p\n", *c, c)

	// As b is a copy any changes to b won't be reflected in a
	b.Ice = false
	// As c is a pointer to the data of a the change will be reflected in both
	// a and c as the underlying data is changed for both
	a.Name = "Orange Juice"
	fmt.Println("\nAfter change to data")
	// The %#v provides additional information about the struct itself
	fmt.Printf("a - %#v - %p\n", a, &a)
	fmt.Printf("b - %#v - %p\n", b, &b)
	fmt.Printf("c - %#v - %p\n", *c, c)
}

func StructsRun() {
	// Adapted from: Go in 24 Hours, Sams Teach Yourself: Next Generation Systems Programming with Golang, First Edition
	formatting.ExampleWrapper("Accessing values of a struct", structValues)
	formatting.ExampleWrapper("To make a copy or point to the same data", copyReference)
}

func init() {
	GetMyExamples().Add("structs", StructsRun)
}
