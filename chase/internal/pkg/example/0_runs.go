package example

// This file is named as it is to ensure it's init is run first and allows the categories map to be available within the package

var (
	// MyExamples is used by the chase binary to auto-discover what subcommands to make
	categories Categories
)

// CategoryExample is the details for a specific example within a category
type CategoryExample struct {
	Description string
	Function    func()
}

// Category is a custom type used for tracking an examples description and function
type Category map[string]CategoryExample

// Categories is one or more Category objects
type Categories map[string]Category

// GetCategories returns the map of examples for use by the chase cli
func GetCategories() Categories {
	return categories
}

// AddCategory will add additional category keys
func (cats Categories) AddCategory(name string) Category {
	cats[name] = Category{}

	return cats[name]
}

// AddExample will add an example to a specific category
func (cat Category) AddExample(exampleKey string, exampleDetails CategoryExample) {
	cat[exampleKey] = exampleDetails
}

func init() {
	// Initialise the variable to be used
	categories = make(Categories)
}
