package cmd

import (
	"fmt"
	"os"

	"github.com/marjamis/kata-go/chase/internal/pkg/example"
	"github.com/marjamis/kata-go/chase/internal/pkg/formatting"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chase",
	Short: "chase is a quick CLI to run some basic golang tests via a single CLI command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func createSubCommands() {
	for category, categoryExamples := range example.GetCategories() {
		specificCategory := category
		// Adds a subcommand to root for each category
		categoryCommand := &cobra.Command{
			Use:   category,
			Short: "Category specific examples",
			Long:  "Category specific examples",
		}
		rootCmd.AddCommand(categoryCommand)

		// Loops through each example in a category to make subcommands of the category subcommand for specific examples
		for categoryExampleKey, categoryExample := range categoryExamples {
			categoryCommand.AddCommand(&cobra.Command{
				Use:   categoryExampleKey,
				Short: categoryExample.Description,
				Long:  categoryExample.Description,
				Run: func(cmd *cobra.Command, args []string) {
					formatting.PrintCategory(specificCategory)
					formatting.PrintExampleOutput(categoryExample.Description, categoryExample.Function)
				},
			})
		}

		categoryCommand.AddCommand(&cobra.Command{
			Use:   "all",
			Short: "Run through each example in the category",
			Long:  "Run through each example in the category",
			Run: func(cmd *cobra.Command, args []string) {
				formatting.PrintCategory(specificCategory)
				for _, functions := range example.GetCategories()[specificCategory] {
					formatting.PrintExampleOutput(functions.Description, functions.Function)
				}
			},
		})
	}

	// Add a special subCommand "all" which will run through all the examples one at a time
	rootCmd.AddCommand(&cobra.Command{
		Use:   "all",
		Short: "Run through each example in each category",
		Long:  "Run through each example in each category",
		Run: func(cmd *cobra.Command, args []string) {
			// Note: order isn't consistent as it loops through a map. For now this is fine but may want to sort the keys first
			for category, categories := range example.GetCategories() {
				formatting.PrintCategory(category)
				for _, functions := range categories {
					formatting.PrintExampleOutput(functions.Description, functions.Function)
				}
			}
		},
	})
}

// Execute runs the command, as per cobra
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	createSubCommands()
}
