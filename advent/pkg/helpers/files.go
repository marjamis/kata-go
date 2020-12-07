package helpers

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//TODO better and coding error checking throughout
func ReadIntDataFromFile(file string) (ints []int) {
	fdata, _ := os.Open(file)
	defer fdata.Close()

	scanner := bufio.NewScanner(fdata)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		ints = append(ints, i)
	}

	return
}

func ReadStringDataFromFile(file string) (strings []string) {
	fdata, _ := os.Open(file)
	defer fdata.Close()

	scanner := bufio.NewScanner(fdata)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return
}

func ReadDataFromFile(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}

	return string(data)
}

func ReadArrayDataFromFile(file string) (strings [][]string) {
	fdata, _ := os.Open(file)
	defer fdata.Close()

	scanner := bufio.NewScanner(fdata)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		substrings := []string{}
		for _, c := range scanner.Text() {
			substrings = append(substrings, string(c))
		}
		strings = append(strings, substrings)
	}

	return
}

func RemoveDuplicates(data []string) (uniques []string) {
	present := map[string]bool{}

	for _, d := range data {
		_, ok := present[d]
		if !ok {
			uniques = append(uniques, d)
			present[d] = true
		}
	}

	return uniques
}
