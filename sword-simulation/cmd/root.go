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
	introText           = "{{.Number}} people standing in a circle in an order 1 to {{.Number}}. No. 1 has a sword. He kills the next person (i.e. No. 2) and gives the sword to the next (i.e. No. 3). All people do the same until only 1 survives. Who will survive at the last?"
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
	rootCmd.Flags().IntVarP(&numberOfPeopleToUse, "Number of People to use", "n", 293, "Sets the number of people in the circle")
}

// Person contains the details of a person in the circle
type Person struct {
	Number int
}

// Persons is an array of all Person objects that make up the circle
type Persons []Person

func setup(numberOfPeople int) Persons {
	// Create the array with the required number of Persons
	people := make([]Person, numberOfPeople)

	// Loop through each person and set their appropriate number based off the index
	for index := range people {
		people[index].Number = index + 1
	}

	return people
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

func simulate(people Persons) (survivors Persons) {
	var toRemove []int

	for i := 0; i < len(people); i = i + 2 {
		// Killing the person next to them
		if i+2 <= len(people) {
			toRemove = append(toRemove, i+1)
		} else {
			toRemove = append(toRemove, 0)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(toRemove)))

	for _, indexToRemove := range toRemove {
		people = append(people[:indexToRemove], people[indexToRemove+1:]...)
	}

	return people
}

func run(numberOfPeople int) Person {
	// Setup
	people := setup(numberOfPeople)

	// Simulate
	for len(people) > 1 {
		people = simulate(people)
	}

	return people[0]
}

func outro() {
	fmt.Println("\n\nAnd there you have it. Now you know where to stand...")
}
