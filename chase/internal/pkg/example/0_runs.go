package example

// This file is named as it is to ensure it's init is run first and allows the myExamples map to be available within the package

import "github.com/marjamis/kata-go/pkg/formatting"

var (
	// MyExamples is used by the chase binary to auto-discover what subcommands to make
	myExamples ExamplesMap
)

// ExamplesMap is a custom type used for tracking an examples description and function
type ExamplesMap map[string]func()

// runDetails contains the requirements of an example for it to be automatically registered to the chase CLI
type runDetails struct {
	Description string
	Function    func()
}

// runs is a custom type for an array of ExampleDetails which contain related data
type runs []runDetails

// GetMyExamples returns the map of examples for use by the chase cli
func GetMyExamples() ExamplesMap {
	return myExamples
}

// Add will add additional keys and functions to the myExamples map
func (ex ExamplesMap) Add(name string, funct func()) {
	ex[name] = funct
}

func init() {
	// Initialise the variable with some space to be used
	myExamples = make(ExamplesMap)
}

// runExamples will loop through the map it's used against and execute ExampleWrapper() to have a pretty output of the example run
func (er runs) runExamples() {
	for _, v := range er {
		formatting.ExampleWrapper(v.Description, v.Function)
	}
}
