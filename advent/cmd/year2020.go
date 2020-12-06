package cmd

import (
	"fmt"
	"sort"

	"github.com/marjamis/advent/internal/pkg/advent2020"
	"github.com/marjamis/advent/pkg/helpers"

	"github.com/spf13/cobra"
)

type day struct {
	Day      string
	Function func(cmd *cobra.Command, args []string)
}

var days = map[string]day{
	"Day1": {
		Day: "Day1",
		Function: func(cmd *cobra.Command, args []string) {
			// TODO add these to the existing wrapper for ease and subscommands to acll specifific days
			day1data := helpers.ReadIntDataFromFile("./test/advent2020/day1.txt")
			fmt.Println(advent2020.Day1(day1data))
			fmt.Println(advent2020.Day1Part2(day1data))
		}},
	"Day2": {
		Day: "Day2",
		Function: func(cmd *cobra.Command, args []string) {
			day2data := helpers.ReadStringDataFromFile("./test/advent2020/day2.txt")
			fmt.Println(advent2020.Day2(day2data, "general"))
			fmt.Println(advent2020.Day2(day2data, "special"))
		}},
	"Day3": {
		Day: "Day3",
		Function: func(cmd *cobra.Command, args []string) {
			day3data := helpers.ReadArrayDataFromFile("./test/advent2020/day3.txt")
			fmt.Println(advent2020.Day3(day3data, [][]int{{1, 3}}))
			fmt.Println(advent2020.Day3(day3data, [][]int{
				{1, 1},
				{3, 1},
				{5, 1},
				{7, 1},
				{1, 2},
			}))
		},
	},
	"Day4": {
		Day: "Day4",
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadDataFromFile("./test/advent2020/day4.txt")
			fmt.Println(advent2020.Day4(data, false))
			fmt.Println(advent2020.Day4(data, true))
		},
	},
	"Day5": {
		Day: "Day5",
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadStringDataFromFile("./test/advent2020/day5.txt")
			fmt.Println(advent2020.Day5(data))
			fmt.Println(advent2020.Day5Part2(data))
		},
	},
}

// year2020Cmd represents the year2020 command
var year2020Cmd = &cobra.Command{
	Use:   "year2020",
	Short: "Runs through the list of each days runs.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("year2020 called")

		// As maps are unsorted keys, at least currently, I've sorted them to be in order.
		var keys []string
		for key := range days {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			days[key].Function(cmd, args)
		}
	},
}

func subCommands() {
	for _, day := range days {
		year2020Cmd.AddCommand(&cobra.Command{
			Use: day.Day,
			Run: day.Function,
		})
	}
}

func init() {
	rootCmd.AddCommand(year2020Cmd)
	subCommands()
}
