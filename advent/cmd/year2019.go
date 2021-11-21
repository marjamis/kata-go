package cmd

import (
	"fmt"

	"github.com/marjamis/advent/internal/pkg/advent2019"
	"github.com/marjamis/advent/pkg/helpers"

	"github.com/spf13/cobra"
)

const dataDirectory2019 = "./test/advent2019/"

var days2019 = map[string]day{
	"01": {
		Function: func() {
			data := helpers.ReadIntArray(dataDirectory2019 + "day1.txt")
			fmt.Println(advent2019.Day1(false, data...))
			fmt.Println(advent2019.Day1(true, data...))
		}},
	"02": {
		Function: func() {
			// TODO this needs redoing to ensure correct result and testing
			data := helpers.ReadIntArray(dataDirectory2019 + "day2.csv")
			fmt.Println(advent2019.Day2(data...))
			fmt.Println(advent2019.Day2(data...))
		}},
	"03": {
		Function: func() {
			// TODO this needs redoing to ensure correct result and testing
			// data := helpers.ReadStringArray(dataDirectory2019 + "day3.csv")
			// fmt.Println(advent2019.Day3(data[0], data[1], advent2019.Day3Manhattan))
		}},
	"04": {
		Function: func() {
			// TODO this needs redoing to ensure correct result and testing
			fmt.Println(advent2019.Day4("353096-843212", advent2019.Day4Rules1))
			fmt.Println(advent2019.Day4("353096-843212", advent2019.Day4Rules2))
		}},
	"05": {
		Function: func() {
			// TODO this needs redoing to ensure correct result and testing
			data := helpers.ReadIntArray(dataDirectory2019 + "day5.csv")
			fmt.Println(advent2019.Day5("1", data...))
			fmt.Println(advent2019.Day5("5", data...))
		}},
	"06": {
		Function: func() {
			// TODO this needs redoing to ensure correct result and testing
		}},
}

// year2019Cmd represents the year2019 command
var year2019Cmd = &cobra.Command{
	Use:   "year2019",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		setLogLevel()
		fmt.Println("2019 called")
		printAllDaysOutput(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(year2019Cmd)
	addDaySubCommandToYearCommand(year2019Cmd, days2019)
}
