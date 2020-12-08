package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/marjamis/advent/internal/pkg/advent2020"
	"github.com/marjamis/advent/pkg/helpers"

	"github.com/spf13/cobra"
)

type day struct {
	Day      string
	Function func(cmd *cobra.Command, args []string)
}

const dataDirectory = "./test/advent2020/"

var days = map[string]day{
	"Day1": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadIntDataFromFile(dataDirectory + "day1.txt")
			fmt.Println(advent2020.Day1(data))
			fmt.Println(advent2020.Day1Part2(data))
		}},
	"Day2": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadStringDataFromFile(dataDirectory + "day2.txt")
			fmt.Println(advent2020.Day2(data, advent2020.Day2CheckOptionGeneral))
			fmt.Println(advent2020.Day2(data, advent2020.Day2CheckOptionSpecial))
		}},
	"Day3": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadArrayDataFromFile(dataDirectory + "day3.txt")
			fmt.Println(advent2020.Day3(data, []advent2020.ToboganMovement{advent2020.ToboganMovement{3, 1}}))
			fmt.Println(advent2020.Day3(data, []advent2020.ToboganMovement{
				advent2020.ToboganMovement{1, 1},
				advent2020.ToboganMovement{3, 1},
				advent2020.ToboganMovement{5, 1},
				advent2020.ToboganMovement{7, 1},
				advent2020.ToboganMovement{1, 2},
			}))
		},
	},
	"Day4": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadDataFromFile(dataDirectory + "day4.txt")
			fmt.Println(advent2020.Day4(data, false))
			fmt.Println(advent2020.Day4(data, true))
		},
	},
	"Day5": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadStringDataFromFile(dataDirectory + "day5.txt")
			fmt.Println(advent2020.Day5(data))
			fmt.Println(advent2020.Day5Part2(data))
		},
	},
	"Day6": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadDataFromFile(dataDirectory + "day6.txt")
			fmt.Println(advent2020.Day6(data, false))
			fmt.Println(advent2020.Day6(data, true))
		},
	},
	"Day7": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadDataFromFile(dataDirectory + "day7.txt")
			fmt.Println(advent2020.Day7(data, "shiny gold", advent2020.Day7SearchOptionIsIn))
			fmt.Println(advent2020.Day7(data, "shiny gold", advent2020.Day7SearchOptionContains))
		},
	},
	"Day8": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadStringDataFromFile(dataDirectory + "day8.txt")
			fmt.Println(advent2020.Day8(data, false))
			fmt.Println(advent2020.Day8(data, true))
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
	for key, day := range days {
		year2020Cmd.AddCommand(&cobra.Command{
			Use: key,
			Aliases: []string{
				strings.ToLower(key),
			},
			Run: day.Function,
		})
	}
}

func init() {
	rootCmd.AddCommand(year2020Cmd)
	subCommands()
}
