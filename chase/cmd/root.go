package cmd

import (
	"fmt"
	"os"

	"github.com/marjamis/kata-go/chase/internal/pkg/example"

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

func subCommands() {
	for subCommandName, subCommandFunction := range example.GetMyExamples() {
		// Required to ensure the function for each subcommand isn't overridden and always points to it's own function specifically
		subcommandFunction := subCommandFunction
		rootCmd.AddCommand(&cobra.Command{
			Use: subCommandName,
			Run: func(cmd *cobra.Command, args []string) {
				subcommandFunction()
			},
		})
	}

	// Add a special subCommand "all" which will run through all the examples at the one time
	rootCmd.AddCommand(&cobra.Command{
		Use: "all",
		Run: func(cmd *cobra.Command, args []string) {
			// Note: order isn't consistent as it loops through a map. For now this is fine but may want to sort the keys first
			for _, subCommandFunction := range example.GetMyExamples() {
				subCommandFunction()
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
	subCommands()
}
