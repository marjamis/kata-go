package formatting

import "fmt"

// PrintCategory prints the category
func PrintCategory(category string) {
	fmt.Printf("# Category: %s\n", category)
}

// PrintExampleOutput prints the examples output
func PrintExampleOutput(description string, function func()) {
	fmt.Printf("## Example description: %s\n", description)

	function()

	fmt.Printf("\n---\n\n")
}
