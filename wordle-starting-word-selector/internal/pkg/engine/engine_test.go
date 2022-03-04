package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	defaultTestingDictionary = Words{
		"hello",
		"no",
		"pause",
		"mo",
		"yurts",
	}
)

func TestScrabbleValue(t *testing.T) {
	tests := []struct {
		input                 string
		expectedScrabbleValue int
	}{
		{"hello", 8},
		{"no", 2},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedScrabbleValue, getScrabbleValue(test.input))
	}
}

func TestFilterWordLength(t *testing.T) {
	words = defaultTestingDictionary
	tests := []struct {
		len int
	}{
		{5},
		{2},
	}

	for _, test := range tests {
		words := words
		words = words.filter(filterLength, test.len)
		for _, word := range words {
			assert.Len(t, word, test.len)
		}
	}
}

func TestFilterScrabbleValue(t *testing.T) {
	words = defaultTestingDictionary
	tests := []struct {
		value int
	}{
		{8},
		{2},
	}

	for _, test := range tests {
		words := words
		words = words.filter(filterScrabbleValue, test.value)
		for _, word := range words {
			assert.Equal(t, test.value, getScrabbleValue(word))
		}
	}
}

func TestFilterDuplicateLetters(t *testing.T) {

	t.Run("No duplicates found", func(t *testing.T) {
		trueWords := Words{
			"no",
			"pause",
		}

		for _, trueWord := range trueWords {
			assert.True(t, filterDuplicateLetters(trueWord))
		}
	})

	t.Run("Duplicates found", func(t *testing.T) {
		falseWords := Words{
			"hello",
			"testi",
		}

		for _, falseWord := range falseWords {
			assert.False(t, filterDuplicateLetters(falseWord))
		}
	})

}
