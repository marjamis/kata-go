package cmd

import (
	"fmt"

	"github.com/marjamis/advent/internal/pkg/advent2021"
	"github.com/marjamis/advent/pkg/helpers"

	"github.com/spf13/cobra"
)

const dataDirectory2021 = "./test/advent2021/"

var days2021 = map[string]day{
	"01": {
		Function: func() {
			input := helpers.ReadIntArray(dataDirectory2021 + "day1.txt")
			fmt.Println(advent2021.Day1Part1(input...))
			fmt.Println(advent2021.Day1Part2(input...))
		}},
	"02": {
		Function: func() {
			input := helpers.ReadStringArray(dataDirectory2021 + "day2.txt")
			fmt.Println(advent2021.Day2Part1(input...))
			fmt.Println(advent2021.Day2Part2(input...))
		}},
	"03": {
		Function: func() {
			input := helpers.ReadStringArray(dataDirectory2021 + "day3.txt")
			fmt.Println(advent2021.Day3Part1(input))
			fmt.Println(advent2021.Day3Part2(input))
		}},
}

// year2021Cmd represents the year2021 command
var year2021Cmd = &cobra.Command{
	Use:   "year2021",
	Short: "Runs through the each advent day for the year",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("year2021 called")
		printAllDaysOutput(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(year2021Cmd)
	addDaySubCommandToYearCommand(year2021Cmd, days2021)
}
