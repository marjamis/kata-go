/*
Copyright © 2022 marjamis

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/marjamis/wordle-starting-word-selector/internal/pkg/engine"

	"github.com/spf13/cobra"
)

var (
	length           *int
	scrabbleValue    *int
	fullList         *bool
	filterDuplicates *bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wsws",
	Short: "Generates a starting word for wordle",
	Run: func(cmd *cobra.Command, args []string) {
		engine.Engine(*length, *scrabbleValue, *fullList, *filterDuplicates)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	length = rootCmd.Flags().IntP("length", "l", -1, "the length of words to find. -1 means all lengths")
	scrabbleValue = rootCmd.Flags().IntP("scrabble-value", "s", -1, "the scrabble value of the word to find. -1 means all values")
	fullList = rootCmd.Flags().BoolP("full-list", "f", false, "set to true to return all available words from filters rather than a random word")
	filterDuplicates = rootCmd.Flags().BoolP("filter-duplicates", "d", true, "set to true to filter any duplicates out of the words to find")
}
