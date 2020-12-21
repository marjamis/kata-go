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
			data := helpers.ReadIntArray(dataDirectory + "day1.txt")
			fmt.Println(advent2020.Day1(data))
			fmt.Println(advent2020.Day1Part2(data))
		}},
	"Day2": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadStringArray(dataDirectory + "day2.txt")
			fmt.Println(advent2020.Day2(data, advent2020.Day2CheckOptionGeneral))
			fmt.Println(advent2020.Day2(data, advent2020.Day2CheckOptionSpecial))
		}},
	"Day3": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadStringArray2d(dataDirectory + "day3.txt")
			fmt.Println(advent2020.Day3(data, []advent2020.ToboganMovement{advent2020.ToboganMovement{X: 3, Y: 1}}))
			fmt.Println(advent2020.Day3(data, []advent2020.ToboganMovement{
				advent2020.ToboganMovement{X: 1, Y: 1},
				advent2020.ToboganMovement{X: 3, Y: 1},
				advent2020.ToboganMovement{X: 5, Y: 1},
				advent2020.ToboganMovement{X: 7, Y: 1},
				advent2020.ToboganMovement{X: 1, Y: 2},
			}))
		},
	},
	"Day4": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadString(dataDirectory + "day4.txt")
			fmt.Println(advent2020.Day4(data, false))
			fmt.Println(advent2020.Day4(data, true))
		},
	},
	"Day5": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadStringArray(dataDirectory + "day5.txt")
			fmt.Println(advent2020.Day5(data))
			fmt.Println(advent2020.Day5Part2(data))
		},
	},
	"Day6": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadString(dataDirectory + "day6.txt")
			fmt.Println(advent2020.Day6(data, false))
			fmt.Println(advent2020.Day6(data, true))
		},
	},
	"Day7": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadString(dataDirectory + "day7.txt")
			fmt.Println(advent2020.Day7(data, "shiny gold", advent2020.Day7SearchOptionIsIn))
			fmt.Println(advent2020.Day7(data, "shiny gold", advent2020.Day7SearchOptionContains))
		},
	},
	"Day8": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadStringArray(dataDirectory + "day8.txt")
			fmt.Println(advent2020.Day8(data, false))
			fmt.Println(advent2020.Day8(data, true))
		},
	},
	"Day9": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadIntArray(dataDirectory + "day9.txt")
			fmt.Println(advent2020.Day9(data, 25))
			fmt.Println(advent2020.Day9Part2(data, 25))
		},
	},
	"Day10": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadIntArray(dataDirectory + "day10.txt")
			fmt.Println(advent2020.Day10(data))
			fmt.Println(advent2020.Day10Part2(data))
		},
	},
	"Day11": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadRuneArray2d(dataDirectory + "day11.txt")
			fmt.Println(advent2020.Day11(data, advent2020.Day11RuleOptionBasic))
			fmt.Println(advent2020.Day11(data, advent2020.Day11RuleOptionAdvanced))
		},
	},
	"Day12": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadStringArray(dataDirectory + "day12.txt")
			fmt.Println(advent2020.Day12(data, advent2020.Day12MovementTypeShip))
			fmt.Println(advent2020.Day12(data, advent2020.Day12MovementTypeWaypoint))
		},
	},
	"Day13": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadStringArray(dataDirectory + "day13.txt")
			fmt.Println(advent2020.Day13(data))
			fmt.Println(advent2020.Day13Part2(data, 122012400000))
		},
	},
	"Day14": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadString(dataDirectory + "day14.txt")
			fmt.Println(advent2020.Day14(data))
			fmt.Println(advent2020.Day14Part2(data))
		},
	},
	"Day15": {
		Function: func(cmd *cobra.Command, args []string) {
			fmt.Println(advent2020.Day15([]int{2, 0, 1, 7, 4, 14, 18}, advent2020.Day15PositionOption2020))
			fmt.Println(advent2020.Day15([]int{2, 0, 1, 7, 4, 14, 18}, advent2020.Day15PositionOption30mil))
		},
	},
	"Day16": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadString(dataDirectory + "day16.txt")
			fmt.Println(advent2020.Day16(data))
			fmt.Println(advent2020.Day16Part2(data))
		},
	},

	"Day18": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadStringArray(dataDirectory + "day18.txt")
			fmt.Println(advent2020.Day18Wrapper(data))
			fmt.Println(advent2020.Day18WrapperPart2(data))
		},
	},
	"Day19": {
		Function: func(cmd *cobra.Command, args []string) {
			data := helpers.ReadString(dataDirectory + "day19.txt")
			fmt.Println(advent2020.Day19(data))
		},
	},
	"Day20": {
		Function: func(cmd *cobra.Command, args []string) {
			setLogLevel()
			data := helpers.ReadString(dataDirectory + "day20.txt")
			fmt.Println(advent2020.Day20(data))
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
