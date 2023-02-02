package engine

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	// Credit: These values were obtained from: https://www.thewordfinder.com/scrabble-point-values.php
	scrabbleMapping = map[rune]int{
		'a': 1,
		'b': 3,
		'c': 3,
		'd': 2,
		'e': 1,
		'f': 4,
		'g': 2,
		'h': 4,
		'i': 1,
		'j': 8,
		'k': 5,
		'l': 1,
		'm': 3,
		'n': 1,
		'o': 1,
		'p': 3,
		'q': 10,
		'r': 1,
		's': 1,
		't': 1,
		'u': 1,
		'v': 4,
		'w': 4,
		'x': 8,
		'y': 4,
		'z': 10,
	}

	words = Words{}
)

// Words a list of strings that make up the available list of words. Starting from the dictionary
// all the way down to a subset after all the filters are run
type Words []string

func openDictionaryFile(dictionary string) {
	fdata, _ := os.Open(dictionary)
	defer fdata.Close()

	strings := []string{}
	scanner := bufio.NewScanner(fdata)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	words = strings
}

// Engine is the workflow controller for finding a good starting word
func Engine(length int, scrabbleValue int, fullList bool, filterDuplicates bool, dictionary string, filterEndingInS bool) {
	openDictionaryFile(dictionary)
	words := words

	// -1 is the cobra flag default to indicate any length
	if length != -1 {
		words = words.filter(filterLength, length)
	}

	// -1 is the cobra flag default to indicate any scrabble value
	if scrabbleValue != -1 {
		words = words.filter(filterScrabbleValue, scrabbleValue)
	}

	if filterDuplicates {
		words = words.filter(filterDuplicateLetters, 0)
	}

	if filterEndingInS {
		words = words.filter(filterWordEndingInS, 0)
	}

	if fullList {
		fmt.Printf("There are %d words available with these filters... Shall these be displayed (y/n)? ", len(words))

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured", err)
			return
		}
		input = strings.TrimSuffix(input, "\n")

		if strings.Compare(input, "y") == 0 {
			fmt.Println(words)
		}
	} else {
		word, err := getWord(words)
		if err != nil {
			fmt.Printf("There is an error of: %s\n", err.Error())
		}

		fmt.Println(word)
	}
}

func getScrabbleValue(word string) (scrabbleValue int) {
	for _, char := range word {
		scrabbleValue += scrabbleMapping[char]
	}

	return
}

func getWord(words Words) (word string, err error) {
	rand.Seed(time.Now().UnixNano())

	if len(words) == 0 {
		return "", errors.New("No available word")
	}

	return words[rand.Intn(len(words))], nil
}

func (w Words) filter(check func(word string, value int) bool, value int) (filteredWords Words) {
	for _, word := range w {
		if check(word, value) {
			filteredWords = append(filteredWords, word)
		}
	}

	return
}

func filterLength(word string, length int) bool {
	return len(word) == length
}

func filterScrabbleValue(word string, scrabbleValue int) bool {
	return getScrabbleValue(word) == scrabbleValue
}

func filterDuplicateLetters(word string, none int) bool {
	for i, testingChar := range word {
		for j, char := range word {
			if (i != j) && (char == testingChar) {
				return false
			}
		}
	}

	return true
}

func filterWordEndingInS(word string, none int) bool {
	return word[len(word)-1] != 's'
}
