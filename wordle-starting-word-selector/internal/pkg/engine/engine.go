package engine

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
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

	dict = dictionary{}
)

type dictionary []string

func openDictionaryFile() {
	fdata, _ := os.Open("configs/dictionary.txt")
	defer fdata.Close()

	strings := []string{}
	scanner := bufio.NewScanner(fdata)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	dict = strings
}

// Engine is the workflow controller for finding a good starting word
func Engine(length int, scrabbleValue int, fullList bool) {
	openDictionaryFile()

	if fullList {
		words := getWords(length, scrabbleValue)
		fmt.Printf("There are %d words available with these filters...\n\n", len(words))
		fmt.Println(words)
	} else {
		word, err := getWord(length, scrabbleValue)
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

func (wl dictionary) filterLen(length int) (newWordList dictionary) {
	for _, word := range wl {
		if len(word) == length {
			newWordList = append(newWordList, word)
		}
	}

	return
}

func (wl dictionary) filterScrabbleValue(value int) (newWordList dictionary) {
	for _, word := range wl {
		if getScrabbleValue(word) == value {
			newWordList = append(newWordList, word)
		}
	}

	return
}

func hasDuplicateLetters(word string) bool {
	for i, testingChar := range word {
		for j, char := range word {
			if (i != j) && (char == testingChar) {
				return true
			}
		}
	}

	return false
}

func (wl dictionary) filterDuplicateLetters() (newWordList dictionary) {
	for _, word := range wl {
		if !hasDuplicateLetters(word) {
			newWordList = append(newWordList, word)
		}
	}

	return
}

func getWords(length int, scrabbleValue int) (words dictionary) {
	wordList := dict

	return wordList.
		filterLen(length).
		filterScrabbleValue(scrabbleValue).
		filterDuplicateLetters()
}

func getWord(length int, scrabbleValue int) (word string, err error) {
	rand.Seed(time.Now().UnixNano())

	wordList := getWords(length, scrabbleValue)
	if len(wordList) == 0 {
		return "", errors.New("No available word")
	}

	return wordList[rand.Intn(len(wordList))], nil

}
