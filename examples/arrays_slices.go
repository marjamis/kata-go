package main

import (
	"fmt"

	"github.com/marjamis/kata-go/pkg/formatting"
)

func printSlice(action string, s []int) {
	// len() is the number of elements in the slice while cap() is the max size of the underlying array
	fmt.Printf("len=%d cap=%d %v - %s\n", len(s), cap(s), s, action)
}

func ranges() {
	// Dropping the elements of an array from the start of the array via the slice will mean those values are garbage collected from the array but if you reduce a slice from the end while they won't be in that current slice they will still be in the underlying array for later use. Meaning the below is shrinking the overall capacity of the underlying array by "dropping" elements from the beginning of the array via a slice but not from the end of the array/slice.

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice("All values", s)

	s = s[:0]
	printSlice("Slice the slice to give it zero length", s)

	s = s[:4]
	printSlice("Extend its length", s)

	s = s[2:]
	printSlice("Drop its first two values reducing the capacity of the array", s)

	s = s[:]
	printSlice("Simply means all elements in the current slice", s)

	s = s[0:4]
	printSlice("Extends it length", s)

	s = s[:2]
	printSlice("Reduces the length by 2", s)

	s = s[1 : len(s)-1]
	printSlice("Drops the first element and reduces the last element", s)

	s = s[:3]
	printSlice("Extends it's length to all remaining items in the array", s)
}

func arrays() {
	// This is an array literal. The type of this array is [3]bool meaning the size is apart of this arrays type.
	array := [3]bool{true, true, false}
	fmt.Printf("type=%T value=%v\n", array, array[2])

	// And this creates the same array as above and then builds a slice that references it.
	// The underlying array will be like the above, the same type, but the slice on top means the size doesn't have to be defined.
	slice := []bool{true, true, false}
	fmt.Printf("type=%T value=%v\n", slice, slice[2])
}

func appending() {
	a := [6]struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(a)

	// This will append the elements 0-1 and 3-5 together into a slice, i.e. skipping the second element.
	s := append(a[:2], a[2+1:]...)
	fmt.Println(s)
}

func structs() {
	// Creates a struct and sets it at the same time. This is good for when the struct only needs to be used in the one place.
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)
}

// Adapted from: tour.golang.org
func main() {
	// Arrays are fixed sizes and types, slices are a dynamicly sized, flexible view of the elements in an array. This allows for slices to be created with subsets of the array it's referencing. Does not store data but describes data of an array. Modifying slices data will modify the underlying array data meaning two slices of the same array will cause changes in each others view.
	formatting.ExampleWrapper("Arrays", arrays)
	formatting.ExampleWrapper("Ranges with slices ", ranges)
	formatting.ExampleWrapper("Appending with slices", appending)
	formatting.ExampleWrapper("Arrays & slices with custom types", structs)
}
