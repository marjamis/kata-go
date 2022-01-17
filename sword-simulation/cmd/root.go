/*
Copyright © 2022 marjamis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"html/template"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

var (
	cfgFile             string
	numberOfPeopleToUse int
	introText           = "{{.Number}} persons standing in a circle in an order 1 to {{.Number}}. No. 1 has a sword. He kills the next person (i.e. No. 2) and gives the sword to the next (i.e. No. 3). All persons do the same until only 1 survives. Who will survive at the last?"
	outroText           = "And there you have it. Now you know where to stand..."
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sword-simulation",
	Short: "Solves the sword problem. What is the sword problem? Run the application to see...",
	Run: func(cmd *cobra.Command, args []string) {
		intro()
		fmt.Printf("\n%d\n", run(numberOfPeopleToUse).Number)
		outro()
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
	rootCmd.Flags().IntVarP(&numberOfPeopleToUse, "Number of People to use", "n", 293, "Sets the number of persons in the circle")
}

// Person contains the details of a person in the circle
type Person struct {
	Number int
}

// Persons is an array of all Person objects that make up the circle
type Persons []Person

func setup(numberOfPeople int) Persons {
	// Creates the array with the required number of Persons
	persons := make([]Person, numberOfPeople)

	// Loop through each Person and set their appropriate number based off the index
	for index := range persons {
		persons[index].Number = index + 1
	}

	return persons
}

func intro() {
	tmpl, err := template.New("Intro").Parse(introText)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, struct {
		Number int
	}{
		Number: numberOfPeopleToUse,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println()
}

func simulate(persons Persons) (survivors Persons) {
	var toRemove []int

	for i := 0; i < len(persons); i = i + 2 {
		// Killing the person next to them
		if i+2 <= len(persons) {
			toRemove = append(toRemove, i+1)
		} else {
			toRemove = append(toRemove, 0)
		}
	}

	// Sorts the indexes and puts them in reverse to allow deletion backwards through the array. This ensure the order isn't upset
	sort.Sort(sort.Reverse(sort.IntSlice(toRemove)))

	// At the end of each run it removes the "dead" persons
	for _, indexToRemove := range toRemove {
		persons = append(persons[:indexToRemove], persons[indexToRemove+1:]...)
	}

	return persons
}

func run(numberOfPeople int) Person {
	// Setup
	persons := setup(numberOfPeople)

	// Simulate - Loops until there is only one person left
	for len(persons) > 1 {
		persons = simulate(persons)
	}

	return persons[0]
}

func outro() {
	fmt.Printf("\n\n%s\n", outroText)
}
