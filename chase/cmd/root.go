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
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Check the help page to see what examples are available. Exiting...")
	},
}

func subCommands() {
	for k, v := range example.GetMyExamples() {
		// Required to ensure the function isn't overridden in later internal processes and is it's own variable
		funct := v
		rootCmd.AddCommand(&cobra.Command{
			Use: k,
			Run: func(cmd *cobra.Command, args []string) {
				funct()
			},
		})
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	subCommands()
}
