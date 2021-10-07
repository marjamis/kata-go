package example

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

type test struct {
	ID          string
	Description string
}

func countLoop() {
	// General for loop comprised of the init statement (optional), condition expression and post statement (optional). This counts 10 times.
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("Output from the loop: %d\n\n", sum)
}

func infiniteLoop() {
	//An infinite for loop but to ensure it exits we break out if the second is divisible by 5
	for {
		fmt.Printf("The time is: %d\n", time.Now().Second())
		if time.Now().Second()%5 == 0 {
			fmt.Println("The second is divisible by 5 exiting loop and function.")
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func arrayLoop() {
	// Configures the test data
	var testData []*test
	for i := 0; i < 10; i++ {
		// Note: Append is a variadic function
		testData = append(testData, &test{
			ID:          strconv.Itoa(i),
			Description: "Test data for id " + strconv.Itoa(i),
		})
	}

	// Loop over all items of an array of pointers to a custom struct and provides the key and the value for each.
	fmt.Println("Output:")
	for index, value := range testData {
		fmt.Printf("index=%d value=%s\n", index, value)
	}
	fmt.Printf("\n\n")
}

func mapLoop() {
	// Configures the test data
	testData := make(map[string]*test)
	for i := 0; i < 10; i++ {
		testData["keys"+strconv.Itoa(i)] = &test{
			ID:          strconv.Itoa(i),
			Description: "Test data for id " + strconv.Itoa(i),
		}
	}

	// Loops over a map of keys (strings) and provides the value (custom struct) for each.
	fmt.Println("Unsorted output:")
	for key, value := range testData {
		fmt.Printf("key=%s value=%s\n", key, value)
	}

	// As ranging through a map returns keys in an unsorted manner we will need another mechanism to sort the keys.
	fmt.Println("\nNow implementing sorted keys...\nSorted output:")
	// Configures an ordered list of keys via an array which can be iterated over to get the details from the map.
	var keys []string
	for k := range testData {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("key=%s value=%s\n", k, testData[k])
	}
}

func init() {
	examples := runs{
		{"Normal count loop", countLoop},
		{"Infinite Loop", infiniteLoop},
		{"Arrays", arrayLoop},
		{"Map i.e. Hash Table", mapLoop},
	}
	GetMyExamples().Add("loops", examples.runExamples)
}
