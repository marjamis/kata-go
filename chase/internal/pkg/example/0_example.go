package example

// This file is named to ensure it's init is run first and allows the MyExamples variable to be available within the package.
// TODO Fix the above to something better but works for now.

var (
	// MyExamples is used by the chase binary to auto-discover what subcommands to make
	MyExamples ExamplesType
)

// ExamplesType is a custom type use for tracking functions to run and their names
type ExamplesType map[string]func()

// GetMyExamples returns the map of examples for use by the chase cli
func GetMyExamples() ExamplesType {
	return MyExamples
}

// Add will add additional keys and functions to the map
func (ex ExamplesType) Add(name string, funct func()) {
	ex[name] = funct
}

func init() {
	MyExamples = make(ExamplesType)
}
